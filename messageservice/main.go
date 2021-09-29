package main

import (
	"log"
	"net"

	"github.com/dlzer/gRPC-skeleton/messageservice/cmd/messageservice"
	configs "github.com/dlzer/gRPC-skeleton/messageservice/config"
	pb "github.com/dlzer/gRPC-skeleton/messageservice/protos/messageservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	log.Println("Starting Server...")

	// open config file and apply it to the configuration strcut
	config, err := configs.ConfigFromFile("config.json")
	if err != nil {
		log.Fatal("Config file error: ", err)
	}

	// models.ConnectToDB(config)
	// log.Println("Connected to Database")

	server := grpc.NewServer()
	pb.RegisterMessageServiceServer(server, messageservice.NewMessageServer(config))

	// reflection - advertises gRPC services
	reflection.Register(server)

	// creates gRPC listener
	listener, err := net.Listen("tcp", config.ListenInterface)
	if err != nil {
		log.Fatal("Error creating listener: ", err.Error())
	}

	// starts gRPC server
	log.Printf("gRPC listening on %s", config.ListenInterface)
	server.Serve(listener)
}
