package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/gerfg/shippy/consignment-service/proto/consignment"
)

const (
	port = ":50051"
)

// simulated DB

// interface do simulated DB
type repository interface {
	Create(*consignment.Consignment) (*consignment.Consignment, error)
	GetAll() []*consignment.Consignment
}

// Repository is a type simulates DB
// mutex to coordinate the access to create new Consignments
// and an array of consignments
type Repository struct {
	mu           sync.RWMutex
	consignments []*consignment.Consignment
}

// Create is a function to append a new consigment in array of consignments
func (repo *Repository) Create(cons *consignment.Consignment) (*consignment.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignments, cons)
	repo.consignments = updated
	repo.mu.Unlock()
	return cons, nil
}

// GetAll return all consignments registereds
func (repo *Repository) GetAll() []*consignment.Consignment {
	return repo.consignments
}

type service struct {
	repo repository
}

func (s *service) CreateConsignment(ctx context.Context, req *consignment.Consignment) (*consignment.Response, error) {
	// Armazena nosso consignment
	cons, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	// Retorna a Reponse message criada no nosso protobuf definido
	return &consignment.Response{Created: true, Consignment: cons}, nil
}

func (s *service) GetConsignments(ctx context.Context, req *consignment.GetRequest) (*consignment.Response, error) {
	consignments := s.repo.GetAll()
	return &consignment.Response{Consignments: consignments}, nil
}

func (s *service) GetConsignment(ctx context.Context, req *consignment.GetRequest) (*consignment.Response, error) {
	cons := s.repo.GetAll()
	return &consignment.Response{Consignment: cons[0]}, nil
}

func main() {
	repo := &Repository{}

	srv := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)

	srv.Init()

	// Registrar nosso servico com o gRPC server
	// Isso vai amarrar nossa implementacão com a interface auto-gerada do nosso protobuf
	consignment.RegisterShippingServiceServer(srv, &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
