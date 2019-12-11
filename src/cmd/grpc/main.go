package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	//pb "grpc-tutorial/grpc-polygot/api"
	pb "proto"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:9090"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterLoginServiceServer(grpcServer, newServer())

	log.Printf("gRPC Server Running...")
	grpcServer.Serve(lis)
}