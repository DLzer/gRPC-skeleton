package main

import (
	"log"
	"net"

	"github.com/dlzer/gRPC-skeleton/messageservice/cmd/messageservice"
	"github.com/dlzer/gRPC-skeleton/messageservice/config"
	"github.com/dlzer/gRPC-skeleton/messageservice/models"
	pb "github.com/golang-friends/slack-clone/authservice/protos/authservice"
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

	models.ConnectToDB(config)
	log.Println("Connected to Database")

	server := grpc.NewServer()
	pb.RegisterAuthServiceServer(server, &authservice.AuthServer{})

	// reflection - advertises gRPC services
	reflection.Register(server)

	// creates gRPC listener
	listener, err := net.Listen("tcp", config.ListenInterface)
	if err != nil {
		log.Fatal("Error creating listener: ", err.Error())
	}

	// starts gRPC server
	log.Printf("gRPC server hosted on %s", config.ListenInterface)
	server.Serve(listener)
}
