package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/gerfg/shippy/consignment-service/proto/consignment"
	"github.com/micro/go-micro"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*consignment.Consignment, error) {
	var cons *consignment.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &cons)
	return cons, err
}

func main() {
	service := micro.NewService(micro.Name("shippy.cli.consignment"))
	service.Init()

	client := consignment.NewShippingServiceClient("shippy.service.consignment", service.Client())

	// contact the server and print out its response
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	cons, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), cons)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &consignment.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
