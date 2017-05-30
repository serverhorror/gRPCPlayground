package server

import (
	"fmt"
	"log"

	"github.com/serverhorror/gRPCPlayground/pb"
	"golang.org/x/net/context"
)

type EchoServer struct{}

func (e *EchoServer) Reply(ctx context.Context, m *pb.MsgRequest) (msg *pb.MsgResponse, err error) {
	log.Printf("%T: %q", m, m)
	msg = new(pb.MsgResponse)

	msg.Value = fmt.Sprintf("Received: %v", m.GetValue())
	return msg, err
}
