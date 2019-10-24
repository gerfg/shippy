package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/gerfg/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/gerfg/shippy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
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
	repo         repository
	vesselClient vesselProto.VesselServiceClient
}

func (s *service) CreateConsignment(ctx context.Context, req *consignment.Consignment, res *consignment.Response) error {

	vesselResponse, err := s.vesselClient.FindAvaliable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found Vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}

	res.VesselId = vesselResponse.Id

	// Armazena nosso consignment
	cons, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// Retorna a Reponse message criada no nosso protobuf definido
	res.Created = true
	res.Consignment = cons
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *consignment.GetRequest, res *consignment.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {
	repo := &Repository{}

	srv := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)

	srv.Init()

	vesselClient := vesselProto.NewVesselServiceClient("shippy.service.vessel", srv.Client())

	// Registrar nosso servico com o gRPC server
	// Isso vai amarrar nossa implementacão com a interface auto-gerada do nosso protobuf
	consignment.RegisterShippingServiceHandler(srv.Server(), &service{repo, vesselClient})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
