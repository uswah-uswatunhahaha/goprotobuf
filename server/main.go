package main

import (
	pb "./custsomer"
)

const (
	port = ":50051"
)

// server is used to implement custsomer.CustomerServer.
type server struct {
	savedCustomers []*pb.CustomerServer
}
