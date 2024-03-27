package svc

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"strconv"

	"gorm.io/gorm"
	"petshop/proto/dog"
)

type DogServiceServer struct {
	dog.UnimplementedDogServiceServer
	db *gorm.DB
}

func (s *DogServiceServer) CreateDog(_ context.Context, req *dog.CreateDogRequest) (res *dog.CreateDogResponse, err error) {
	err = s.db.Create(req.Dog).Error
	if err != nil {
		log.Printf("Failed to create dog: %v", err)
		return nil, err
	}

	log.Printf("Created dog with ID: %d", req.Dog.Id)

	res = &dog.CreateDogResponse{
		Id: req.Dog.Id,
	}

	return res, nil
}

func (s *DogServiceServer) GetDog(_ context.Context, req *dog.GetDogRequest) (res *dog.GetDogResponse, err error) {
	var fetchedDog dog.Dog

	err = s.db.First(&fetchedDog, strconv.Itoa(int(req.Id))).Error
	if err != nil {
		log.Printf("Failed to get dog with ID %d: %v", req.Id, err)
		return nil, err
	}

	log.Printf("Fetched dog: %+v", fetchedDog.String())

	res = &dog.GetDogResponse{
		Dog: &fetchedDog,
	}

	return res, nil
}

func (s *DogServiceServer) ListDog(_ context.Context, _ *emptypb.Empty) (res *dog.ListDogsResponse, err error) {
	var dogs []*dog.Dog

	err = s.db.Find(&dogs).Error
	if err != nil {
		log.Printf("Failed to list dogs: %v", err)
		return nil, err
	}

	log.Printf("Listed dogs: %+v", dogs)

	res = &dog.ListDogsResponse{
		Dog: dogs,
	}

	return res, nil
}
