package scriptrunner

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"

	vm "github.com/Dobefu/vee-em"
)

type vmRuntime struct {
	constPool    []string
	functionPool []string
	program      []byte
}

const (
	errUnexpectedEndOfBytecode = "unexpected end of bytecode: %d > %d"
)

func (rt *vmRuntime) loadBytecode(bytecode []byte) error {
	if len(bytecode) < 4 {
		return errors.New("bytecode too short")
	}

	magicHeader := bytecode[:4]

	if string(magicHeader) != "DLS\x01" {
		return fmt.Errorf("invalid magic header: %s", string(magicHeader))
	}

	offset := 4
	constPool, newOffset, err := rt.deserializeConstPool(bytecode, offset)

	if err != nil {
		return err
	}

	rt.constPool = constPool
	offset = newOffset

	functionPool, newOffset, err := rt.deserializeFunctionPool(bytecode, offset)

	if err != nil {
		return err
	}

	rt.functionPool = functionPool
	offset = newOffset

	rt.program = bytecode[offset:]

	return nil
}

func (rt *vmRuntime) deserializeConstPool(
	bytecode []byte,
	offset int,
) ([]string, int, error) {
	if offset+8 > len(bytecode) {
		return nil, 0, fmt.Errorf(
			errUnexpectedEndOfBytecode,
			offset+8,
			len(bytecode),
		)
	}

	count := binary.BigEndian.Uint64(bytecode[offset : offset+8])
	offset += 8

	pool := make([]string, 0, count)

	for range count {
		if offset+8 > len(bytecode) {
			return nil, 0, fmt.Errorf(
				errUnexpectedEndOfBytecode,
				offset+8,
				len(bytecode),
			)
		}

		length := binary.BigEndian.Uint64(bytecode[offset : offset+8])
		offset += 8

		if length > math.MaxInt {
			return nil, 0, fmt.Errorf(
				"length exceeds maximum int value: %d > %d",
				length,
				math.MaxInt,
			)
		}

		if (offset + int(length)) > len(bytecode) {
			return nil, 0, fmt.Errorf(
				errUnexpectedEndOfBytecode,
				offset+int(length),
				len(bytecode),
			)
		}

		str := string(bytecode[offset : offset+int(length)])
		pool = append(pool, str)
		offset += int(length)
	}

	return pool, offset, nil
}

func (rt *vmRuntime) deserializeFunctionPool(
	bytecode []byte,
	offset int,
) ([]string, int, error) {
	if offset+8 > len(bytecode) {
		return nil, 0, fmt.Errorf(
			errUnexpectedEndOfBytecode,
			offset+8,
			len(bytecode),
		)
	}

	count := binary.BigEndian.Uint64(bytecode[offset : offset+8])
	offset += 8

	pool := make([]string, 0, count)

	for range count {
		if offset+8 > len(bytecode) {
			return nil, 0, fmt.Errorf(
				errUnexpectedEndOfBytecode,
				offset+8,
				len(bytecode),
			)
		}

		length := binary.BigEndian.Uint64(bytecode[offset : offset+8])
		offset += 8

		if length > math.MaxInt {
			return nil, 0, fmt.Errorf(
				"length exceeds maximum int value: %d > %d",
				length,
				math.MaxInt,
			)
		}

		if offset+int(length) > len(bytecode) {
			return nil, 0, fmt.Errorf(
				errUnexpectedEndOfBytecode,
				offset+int(length),
				len(bytecode),
			)
		}

		name := string(bytecode[offset : offset+int(length)])
		pool = append(pool, name)
		offset += int(length)
	}

	return pool, offset, nil
}

func (rt *vmRuntime) run(out io.Writer) error {
	handler := rt.createHostCallHandler(out)
	v := vm.New(rt.program, vm.WithHostCallHandler(handler))

	err := v.Run()

	if err != nil {
		return fmt.Errorf("failed to run bytecode: %s", err.Error())
	}

	return nil
}

func (rt *vmRuntime) createHostCallHandler(out io.Writer) vm.HostCallHandler {
	return func(
		functionIndex int64,
		arg1Reg uint64,
		numArgs uint64,
		registers [32]int64,
	) (int64, error) {
		if functionIndex < 0 || int(functionIndex) >= len(rt.functionPool) {
			return 0, fmt.Errorf(
				"invalid function index: %d < 0 or %d >= %d",
				functionIndex,
				functionIndex,
				len(rt.functionPool),
			)
		}

		functionName := rt.functionPool[functionIndex]

		args := make([]int64, numArgs)

		for i := range numArgs {
			registerIndex := (arg1Reg + i) & vm.NumRegistersMask

			if registerIndex >= vm.NumRegisters {
				return 0, fmt.Errorf(
					"invalid register index: %d >= %d",
					registerIndex,
					vm.NumRegisters-1,
				)
			}

			args[i] = registers[registerIndex] // #nosec: G115
		}

		return rt.callFunction(functionName, args, registers, out)
	}
}

func (rt *vmRuntime) callFunction(
	name string,
	args []int64,
	_ [32]int64,
	out io.Writer,
) (int64, error) {
	switch name {
	case "printf":
		return rt.callPrintf(args, out)

	default:
		return 0, fmt.Errorf("unknown function: %s", name)
	}
}

func (rt *vmRuntime) callPrintf(args []int64, out io.Writer) (int64, error) {
	if len(args) == 0 {
		return 0, nil
	}

	formatIndex := args[0]

	if formatIndex < 0 || int(formatIndex) >= len(rt.constPool) {
		return 0, fmt.Errorf(
			"invalid format string index: %d < 0 or %d >= %d",
			formatIndex,
			formatIndex,
			len(rt.constPool),
		)
	}

	format := rt.constPool[formatIndex]

	formatArgs := rt.parseFormatString(format, args)

	_, err := fmt.Fprintf(out, format, formatArgs...)

	if err != nil {
		return 0, fmt.Errorf("failed to format string: %s", err.Error())
	}

	return 0, nil
}

func (rt *vmRuntime) parseFormatString(format string, args []int64) []any {
	formatArgs := make([]any, 0)
	formatIdx := 0

	for i := 1; i < len(args); i++ {
		argValue := args[i]

		isStringSpec := false

	getFormatSpecifier:
		for j := formatIdx; j < len(format)-1; j++ {
			if format[j] != '%' {
				continue
			}

			switch format[j+1] {
			case 's':
				isStringSpec = true
				formatIdx = j + 2

				break getFormatSpecifier
			default:
				formatIdx = j + 2

				break getFormatSpecifier
			}
		}

		if isStringSpec && argValue >= 0 && int(argValue) < len(rt.constPool) {
			formatArgs = append(formatArgs, rt.constPool[argValue])
		} else {
			formatArgs = append(formatArgs, float64(argValue))
		}
	}

	return formatArgs
}
