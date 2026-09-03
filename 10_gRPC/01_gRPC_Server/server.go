package main

import (
	"context"
	"log"
	"net"

	pb "simplegrpcserver/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedCalculateServer
}

func (s server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	return &pb.AddResponse{
		Sum: req.A + req.B,
	}, nil
}

func main() {
	// ---------- Start: 007 ------------
	cert := "cert.pem"
	key := "key.pem"
	// ---------- End: 007 ------------

	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	// ---------- Start: 007 ------------
	creds, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		log.Fatalln("Failed to load credentials", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	// ---------- End: 007 ------------

	pb.RegisterCalculateServer(grpcServer, &server{})

	log.Println("Server is running on port", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("Failed to serve:", err)
	}
}
