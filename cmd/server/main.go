package main

import (
	"flag"
	"fmt"
	"hello-grpc/pb"
	"hello-grpc/service"
	"hello-grpc/store"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 0, "port of the server")
	flag.Parse()
	serverAddress := fmt.Sprintf("0.0.0.0:%d", *port)
	log.Printf("start server on %s", serverAddress)
	laptopServer := service.NewLaptopServer(store.NewInMemoryLaptopStore())
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)
	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalf("failed to start server:%w", err)
	}
	log.Printf("server started")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed to start grpc server:%w", err)
	}

}
