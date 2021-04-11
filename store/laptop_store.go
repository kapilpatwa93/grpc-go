package store

import (
	"errors"
	"fmt"
	"hello-grpc/pb"
	"sync"

	"github.com/jinzhu/copier"
)

var ErrAlreadyExists = errors.New("record already exists")

type LaptopStore interface {
	Save(laptop *pb.Laptop) error
	Find(id string) (*pb.Laptop, error)
}

type InMemoryStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

func NewInMemoryLaptopStore() *InMemoryStore {
	return &InMemoryStore{
		mutex: sync.RWMutex{},
		data:  make(map[string]*pb.Laptop),
	}
}

func (store InMemoryStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	if _, exist := store.data[laptop.Id]; exist {
		return ErrAlreadyExists
	}
	newLaptop := &pb.Laptop{}
	err := copier.Copy(newLaptop, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data: %w", err)
	}
	store.data[laptop.Id] = newLaptop
	return nil
}

func (store InMemoryStore) Find(id string) (*pb.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()
	laptop, exist := store.data[id]
	if !exist {
		return nil, nil
	}
	laptop2 := &pb.Laptop{}
	err := copier.Copy(laptop2, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}
	return laptop2, nil
}
