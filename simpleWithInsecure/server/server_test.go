package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/serverhorror/gRPCPlayground/pb"
)

func TestEchoServer_Echo(t *testing.T) {

	s := EchoServer{}
	req := &pb.MsgRequest{
		Value: "...",
	}

	resp, err := s.Reply(context.Background(), req)
	if err != nil {
		t.Error(errors.Wrap(err, "could not s.Reply"))
	}
	expected := &pb.MsgResponse{
		Value: fmt.Sprintf("received: %v", req.GetValue()),
	}
	if resp != expected {
		t.Error(fmt.Errorf("expected: %q, got: %q",
			expected,
			resp.String()))
	}
}
