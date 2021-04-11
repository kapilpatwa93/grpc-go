package service

import (
	"context"
	"hello-grpc/pb"
	"hello-grpc/sample"
	"hello-grpc/serializer"
	store2 "hello-grpc/store"
	"net"
	"testing"

	"github.com/stretchr/testify/require"

	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()
	laptopServer, serverAddress := startTestLaptopServer(t)
	laptopClient := newTestLaptopClient(t, serverAddress)

	laptop := sample.NewLaptop()
	expectedId := laptop.Id

	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}
	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, res.Id, expectedId)

	laptop2, err := laptopServer.Store.Find(laptop.Id)
	require.NoError(t, err)
	require.NotNil(t, laptop2)
	requireSameLaptop(t, laptop2, laptop)

}

func startTestLaptopServer(t *testing.T) (*LaptopServer, string) {
	store := store2.NewInMemoryLaptopStore()
	laptopServer := NewLaptopServer(store)

	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go grpcServer.Serve(listener)
	return laptopServer, listener.Addr().String()
}

func newTestLaptopClient(t *testing.T, serverAddress string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.WithBlock())
	require.NoError(t, err)

	return pb.NewLaptopServiceClient(conn)

}

func requireSameLaptop(t *testing.T, laptop1, laptop2 *pb.Laptop) {
	json1, err := serializer.ProtobufToJSON(laptop1)
	json2, err := serializer.ProtobufToJSON(laptop2)
	require.NoError(t, err)
	require.Equal(t, json2, json1)

}
