package service

import (
	"context"
	"hello-grpc/pb"
	"hello-grpc/sample"
	"hello-grpc/store"
	"testing"

	"google.golang.org/grpc/status"

	"github.com/stretchr/testify/require"

	"google.golang.org/grpc/codes"
)

func TestLaptopServer_CreateLaptop(t *testing.T) {
	//t.Parallel()

	laptopNoId := sample.NewLaptop()
	laptopNoId.Id = ""

	laptopInvalidId := sample.NewLaptop()
	laptopInvalidId.Id = "invalid_id"

	newStore := store.NewInMemoryLaptopStore()
	laptop1 := sample.NewLaptop()
	newStore.Save(laptop1)

	testCases := []struct {
		name   string
		laptop *pb.Laptop
		store  store.LaptopStore
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			laptop: sample.NewLaptop(),
			store:  store.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "success_with_no_id",
			laptop: laptopNoId,
			store:  store.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "failure_with_invalid_id",
			laptop: laptopInvalidId,
			store:  store.NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "failure_with_duplicate_laptop_id",
			laptop: laptop1,
			store:  newStore,
			code:   codes.AlreadyExists,
		},
	}

	for i, _ := range testCases {
		testCase := testCases[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			req := pb.CreateLaptopRequest{
				Laptop: testCase.laptop,
			}
			server := NewLaptopServer(testCase.store)
			res, err := server.CreateLaptop(context.Background(), &req)
			if testCase.code == codes.OK {
				require.NoError(t, err)
				require.NotNilf(t, res, "not nil")
				require.NotEmptyf(t, res.Id, "not empty")
				if len(testCase.laptop.Id) > 0 {
					require.Equal(t, testCase.laptop.GetId(), res.GetId())
				}

			} else {
				require.Error(t, err, "error is nil")
				require.Nil(t, res, "response is not empty")
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, testCase.code, st.Code())
			}
		})

	}
}
