package jsonrpc2

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
)

// Server represents the JSON-RPC server.
type Server struct {
	stream  *Stream
	handler Handler
}

// NewServer creates a new JSON-RPC server.
func NewServer(handler Handler, reader io.Reader, writer io.Writer) (*Server, error) {
	if handler == nil {
		return nil, fmt.Errorf("nil handler provided")
	}

	stream := NewStream(reader, writer)

	return &Server{
		stream:  stream,
		handler: handler,
	}, nil
}

// Start starts the JSON-RPC server.
func (s *Server) Start() error {
	for {
		msg, err := s.stream.ReadMessage()

		if err != nil {
			if err == io.EOF || errors.Is(err, io.EOF) {
				return nil
			}

			return err
		}

		err = s.handleMessage(msg)

		if err != nil {
			slog.Error(err.Error())
		}
	}
}

func (s *Server) handleMessage(msg []byte) error {
	var req Request

	err := json.Unmarshal(msg, &req)

	if err != nil {
		err = s.sendError(req.ID, ErrorCodeParseError, err.Error(), nil)

		return fmt.Errorf("could not unmarshal message: %s", err)
	}

	result, handleErr := s.handler.Handle(req.Method, req.Params)

	if handleErr != nil {
		err = s.sendError(req.ID, handleErr.Code, handleErr.Message, handleErr.Data)

		if err != nil {
			return fmt.Errorf("could not send error: %s", err)
		}

		return fmt.Errorf("could not handle message: %s", handleErr.Message)
	}

	if req.ID.IsNull() || result == nil {
		return nil
	}

	err = s.sendResponse(req.ID, result)

	if err != nil {
		return fmt.Errorf("could not send response: %s", err)
	}

	return nil
}

func (s *Server) sendResponse(id RequestID, result json.RawMessage) error {
	res := NewResponse(result, id)
	data, err := json.Marshal(res)

	if err != nil {
		return fmt.Errorf("could not marshal response: %w", err)
	}

	return s.stream.WriteMessage(data)
}

func (s *Server) sendError(
	id RequestID,
	code ErrorCode,
	msg string,
	data *json.RawMessage,
) error {
	errObj := NewError(code, msg, data)
	errData, err := json.Marshal(errObj)

	if err != nil {
		return fmt.Errorf("could not marshal error: %w", err)
	}

	resp := NewErrorResponse(errData[:], id)
	respData, err := json.Marshal(resp)

	if err != nil {
		return fmt.Errorf("could not marshal error response: %w", err)
	}

	return s.stream.WriteMessage(respData)
}
