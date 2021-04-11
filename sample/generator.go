package sample

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"hello-grpc/pb"
)

func NewKeyboard() *pb.Keyboard {
	return &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
}

func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := RandomCPUName(brand)
	cores := randomCores()
	threads := randomThreads(cores)
	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)
	return &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(cores),
		NumberThreads: uint32(threads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}
}
func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := RandomGPUName(brand)
	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 3.0)
	memory := &pb.Memory{
		Value: uint64(randomInteger(2, 6)),
		Unit:  pb.Memory_GIGABYTE,
	}
	return &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}

}

func NewRAM() *pb.Memory  {
	return &pb.Memory{
		Value: uint64(randomInteger(4, 64)),
		Unit:  pb.Memory_GIGABYTE,
	}

}
// NewHDD returns a new sample HDD Storage
func NewHDD() *pb.Storage {
	return &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInteger(128,1024)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}
}
// NewSSD returns a new sample SSD Storage
func NewSSD() *pb.Storage {
	return &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(randomInteger(1,6)),
			Unit:  pb.Memory_TERABYTE,
		},
	}
}

func NewScreen() *pb.Screen {

	return &pb.Screen{
		Size:       randomFloat32(13,17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}


}
func NewLaptop() *pb.Laptop  {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)
	return &pb.Laptop{
		Id:          randomID(),
		Brand:       brand,
		Name:        name,
		Cpu:         NewCPU(),
		Ram:         NewRAM(),
		Gpus:        []*pb.GPU{NewGPU(),NewGPU()},
		Storages:    []*pb.Storage{NewHDD(),NewSSD()},
		Screen:      NewScreen(),
		Keyboard:    NewKeyboard(),
		Weight:      &pb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0,3.0),
		},
		PriceUsd:    randomFloat64(1500,3000),
		ReleaseYear: uint32(randomInteger(2015,2021)),
		UpdatedAt:   timestamppb.Now(),
	}
}