package main

import (
	"google.golang.org/grpc"
	"grpc-server/pkg/adder"
	"grpc-server/pkg/api"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	api.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	s.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
