package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"petshop/proto/dog"
)

func main() {
	conn, err := grpc.Dial("localhost:9191", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}

	client := dog.NewDogServiceClient(conn)

	dogExample := &dog.Dog{
		Name:  "Scott",
		Breed: dog.Breed_POODLE,
		Age:   8,
	}

	dog1, err := client.CreateDog(context.Background(), &dog.CreateDogRequest{Dog: dogExample})

	log.Printf("dog: %v", dog1)
}
