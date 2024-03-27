package svc

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"petshop/proto/person"

	"gorm.io/gorm"
)

type PersonServiceServer struct {
	person.UnimplementedPersonServiceServer
	db *gorm.DB
}

func (s *PersonServiceServer) GetPerson(_ context.Context, req *person.GetPersonRequest) (res *person.GetPersonResponse, err error) {
	var fetchedPerson person.Person

	err = s.db.First(&fetchedPerson, req.Id).Error
	if err != nil {
		log.Printf("Failed to get person with ID %d: %v", req.Id, err)
		return nil, err
	}

	log.Printf("Fetched person: %+v", fetchedPerson.String())

	res = &person.GetPersonResponse{
		Person: &fetchedPerson,
	}

	return res, nil
}

func (s *PersonServiceServer) ListPersons(_ context.Context, _ *emptypb.Empty) (res *person.ListPersonsResponse, err error) {
	var persons []*person.Person

	err = s.db.Find(&persons).Error
	if err != nil {
		log.Printf("Failed to list persons: %v", err)
		return nil, err
	}

	log.Printf("Listed persons: %+v", persons)

	res = &person.ListPersonsResponse{
		Person: persons,
	}

	return res, nil
}

func (s *PersonServiceServer) CreatePerson(_ context.Context, req *person.CreatePersonRequest) (res *person.CreatePersonResponse, err error) {
	err = s.db.Create(req.Person).Error

	if err != nil {
		log.Printf("Failed to create person: %v", err)
		return nil, err
	}

	log.Printf("Created person with ID: %d", req.Person.Id)

	res = &person.CreatePersonResponse{
		Id: req.Person.Id,
	}

	return res, nil
}
