build:
	protoc --proto_path=$(GOPATH)/src:. --micro_out=. --go_out=. \
		proto/consignment/consignment.proto
	sudo docker build -t consignment-service .

run:
	sudo docker run -p 50051:50051 \
	-e MICRO_SERVER_ADDRESS=:50051 \
	consignment-service