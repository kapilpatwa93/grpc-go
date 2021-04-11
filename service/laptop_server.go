package service

import (
	"context"
	"errors"
	"hello-grpc/pb"
	"hello-grpc/store"
	"log"

	"google.golang.org/grpc/status"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)

// LaptopServer is the server that provides laptop services
type LaptopServer struct {
	pb.UnimplementedLaptopServiceServer
	Store store.LaptopStore
}

func (server LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("recieved a create-laptop request with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptopId is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop ID: %v", err)
		}
		laptop.Id = id.String()
	}
	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("dead line is exceeded")
		return nil, status.Errorf(codes.DeadlineExceeded, "deadline is exceeded")
	}

	if ctx.Err() == context.Canceled {
		log.Printf("request is cancelled")
		return nil, status.Errorf(codes.Canceled, "request is cancelled by the client")
	}
	// save the laptop to the store

	err := server.Store.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, store.ErrAlreadyExists) {
			log.Printf("laptop already exist")
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot save laptop to the store:%v", err)
	}

	log.Printf("saved laptop with id:%s", laptop.Id)

	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return res, nil
}

//// CreateLaptop is unary RPC to create a new Laptop
//func (server LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
//
//}

// NewLaptopServer returns new LaptopServer
func NewLaptopServer(store store.LaptopStore) *LaptopServer {
	return &LaptopServer{
		UnimplementedLaptopServiceServer: pb.UnimplementedLaptopServiceServer{},
		Store:                            store,
	}
}
