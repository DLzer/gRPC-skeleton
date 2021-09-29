package messageservice

import (
	"context"
	"testing"

	pb "github.com/dlzer/gRPC-skeleton/messageservice/protos/messageservice"
)

func Test_Message_Service(t *testing.T) {

	server := MessageServer{}

	// Test Case: Attempt to send a simple message
	_, err := server.SayHello(context.Background(), &pb.Message{Body: "Hello"})
	if err != nil {
		t.Error("1. Server responded with an error: ", err.Error())
	}

	// Test Case: Response from the server
	res, err := server.SayHello(context.Background(), &pb.Message{Body: "Hello"})
	if err != nil {
		t.Error("1a. Could not get response from server")
	}

	want := "Hello From the Server!"

	if res.GetBody() != "" {
		got := res.GetBody()
		if got != want {
			t.Errorf("2. Body does not match requested. Got %q wanted %q", got, want)
		}
	} else {
		t.Error("3. Message body is blank")
	}

}
