package messageservice

import (
	"context"
	"testing"
	// "github.com/dlzer/gRPC-skeleton/messageservice/models"
	pb "github.com/dlzer/gRPC-skeleton/messageservice/protos/messageservice"
)

func Test_Message_Service(t *testing.T) {

	server := MessageServer{}

	// Test Case: Attempt to send a simple message
	_, err := server.Message(context.Background(), &pb.Message{Body: "Hello"})
	if err != nil {
		t.Error("1. Server responded with an error: ", err.Error())
	}

}
