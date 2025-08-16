package jsonrpc2

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Handler represents a JSON-RPC handler.
type Handler interface {
	Handle(method string, params json.RawMessage) (json.RawMessage, *Error)
}

// Stream represents the JSON-RPC stream.
type Stream struct {
	reader io.Reader
	writer io.Writer
	closer io.Closer
}

// NewStream creates a new JSON-RPC stream.
func NewStream(reader io.Reader, writer io.Writer) *Stream {
	return &Stream{
		reader: reader,
		writer: writer,
		closer: nil,
	}
}

// ReadMessage reads a message from the stream.
func (s *Stream) ReadMessage() ([]byte, error) {
	var err error

	scanner := bufio.NewScanner(s.reader)
	var contentLength int

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		contentLenStr, hasContentLenStr := strings.CutPrefix(line, "Content-Length: ")

		if hasContentLenStr {
			contentLength, err = strconv.Atoi(contentLenStr)

			if err != nil {
				return nil, fmt.Errorf("could not parse content length: %w", err)
			}
		}
	}

	if contentLength == 0 {
		return nil, nil
	}

	buf := make([]byte, contentLength)

	_, err = io.ReadFull(s.reader, buf)

	if err != nil {
		return buf, fmt.Errorf("could not read message: %w", err)
	}

	return buf, nil
}

// WriteMessage writes a message to the stream.
func (s *Stream) WriteMessage(data []byte) error {
	contentLenHeader := fmt.Sprintf("Content-Length: %d\r\n\r\n", len(data))

	_, err := s.writer.Write([]byte(contentLenHeader))

	if err != nil {
		return fmt.Errorf("could not write content length header: %w", err)
	}

	_, err = s.writer.Write(data)

	return err
}
