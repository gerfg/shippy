package main

import (
	"context"
	"local/shippy/consignment-service/proto/consignment"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// simulated DB

// interface do simulated DB
type repository interface {
	Create(*consignment.Consignment) (*consignment.Consignment, error)
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

func main() {
	repo := &Repository{}

	// setup grpc server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to lister: %v", err)
	}
	s := grpc.NewServer()

	// Registrar nosso servico com o gRPC server
	// Isso vai amarrar nossa implementacão com a interface auto-gerada do nosso protobuf
	consignment.RegisterShippingServiceServer(s, &service{repo})

	// Registro do serkvico de reflexão no servidor gRPC
	reflection.Register(s)

	log.Println("Runing on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
