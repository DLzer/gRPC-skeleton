package messageservice

import (
	"context"
	"log"

	configs "github.com/dlzer/gRPC-skeleton/messageservice/config"

	pb "github.com/dlzer/gRPC-skeleton/messageservice/protos/messageservice"
)

// MessageServer is the abstraction of the gRPC server
type MessageServer struct {
	config configs.Configuration
}

// NewMessageServer is AuthServerFactory
func NewMessageServer(config configs.Configuration) *MessageServer {
	return &MessageServer{config}
}

// Message uses the MessageServer abstraction and the protocol buffer to pass the message
func (as *MessageServer) SayHello(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &pb.Message{Body: "Hello From the Server!"}, nil
}
