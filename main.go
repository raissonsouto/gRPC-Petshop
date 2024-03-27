package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"petshop/cmd/svc"
	"petshop/proto/dog"
)

var (
	listener net.Listener
	server   *grpc.Server
)

func main() {
	initListener()

	server = grpc.NewServer()
	dog.RegisterDogServiceServer(server, &svc.DogServiceServer{})

	log.Println("Starting gRPC server...")

	err := server.Serve(listener)
	if err != nil {
		log.Panicf("Failed to start gRPC server: %v", err)
	}
}

func initListener() {
	var err error
	addr := "0.0.0.0:9191"

	listener, err = net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on address %s: %v", addr, err)
	}

	log.Printf("Started listening on address %s", addr)
}
