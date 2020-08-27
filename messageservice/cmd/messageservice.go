package messageservice

import (
	"context"
	pb "github.com/dlzer/gRPC-skeleton/messageservice/protos/messageservice"
	"log"
)

// MessageServer is the abstraction of the gRPC server
type MessageServer struct {
}

// Message uses the MessageServer abstraction and the protocol buffer to pass the message
func (as *MessageServer) Message(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &pb.Message{Body: "Hello From the Server!"}, nil
}
