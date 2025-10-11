package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/Dobefu/DLiteScript/scriptrunner"
)

type output struct {
	Buf string `json:"buffer,omitempty"`
	Err string `json:"error,omitempty"`
}

func main() {
	js.Global().Set("runString", runString())

	<-make(chan struct{})
}

func runString() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return fmt.Sprintf("Expected 1 argument, got %d", len(args))
		}

		input := args[0].String()
		outfile := &bytes.Buffer{}

		runner := &scriptrunner.ScriptRunner{
			OutFile: outfile,
		}

		_, err := runner.RunString(input)
		var result []byte

		if err != nil {
			result, err = json.Marshal(output{
				Buf: "",
				Err: err.Error(),
			})
		} else {
			result, err = json.Marshal(output{
				Buf: outfile.String(),
				Err: "",
			})
		}

		if err != nil {
			return "{}"
		}

		return string(result)
	})
}
