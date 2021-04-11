package main

import (
	"context"
	"flag"
	"hello-grpc/pb"
	"hello-grpc/sample"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

func main() {
	serverAddress := flag.String("address", "", "server address")
	flag.Parse()

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connec to server:%w", err)
	}
	laptopClient := pb.NewLaptopServiceClient(conn)

	laptop1 := sample.NewLaptop()
	laptop1.Id = "10c07996-9697-425c-9b09-c245b1bdcde2"
	laptopRequest := &pb.CreateLaptopRequest{Laptop: laptop1}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	laptopResponse, err := laptopClient.CreateLaptop(ctx, laptopRequest)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Printf("laptop already exist")
		} else {
			log.Fatal("cannot create laptop", err)
		}
	} else {
		log.Printf("laptop created successfully with id:%s", laptopResponse.Id)
	}

}
