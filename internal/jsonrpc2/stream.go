package jsonrpc2

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Handler represents a JSON-RPC handler.
type Handler interface {
	Handle(method string, params json.RawMessage) (json.RawMessage, *Error)
	GetShutdownChan() chan struct{}
}

// Stream represents the JSON-RPC stream.
type Stream struct {
	reader *bufio.Reader
	writer io.Writer
	closer io.Closer
}

// NewStream creates a new JSON-RPC stream.
func NewStream(reader io.Reader, writer io.Writer) *Stream {
	return &Stream{
		reader: bufio.NewReader(reader),
		writer: writer,
		closer: nil,
	}
}

// ReadMessage reads a message from the stream.
func (s *Stream) ReadMessage() ([]byte, error) {
	var err error

	var contentLength int
	var hasContentLength bool

	for {
		line, err := s.reader.ReadString('\n')

		if err != nil {
			return nil, fmt.Errorf("could not read the message: %w", err)
		}

		line = strings.TrimSpace(line)

		if line == "" {
			break
		}

		contentLenStr, hasContentLenStr := strings.CutPrefix(line, "Content-Length: ")

		if hasContentLenStr {
			contentLength, err = strconv.Atoi(contentLenStr)

			if err != nil {
				return nil, fmt.Errorf("could not parse content length: %w", err)
			}

			hasContentLength = true
		}
	}

	if !hasContentLength {
		return nil, errors.New("no Content-Length header found")
	}

	if contentLength <= 0 {
		return nil, errors.New("invalid content length")
	}

	buf := make([]byte, contentLength)

	_, err = io.ReadFull(s.reader, buf)

	if err != nil {
		return nil, fmt.Errorf("could not read the message body: %w", err)
	}

	if len(buf) != contentLength {
		return nil, errors.New("message body length does not match Content-Length")
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

	if err != nil {
		return fmt.Errorf("could not write message: %w", err)
	}

	return nil
}
