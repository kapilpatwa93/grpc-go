package serializer

import (
	"hello-grpc/pb"
	"hello-grpc/sample"
	"testing"

	"github.com/golang/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	laptop1 := sample.NewLaptop()
	err := WriteProtobufToBinaryFile(laptop1, binaryFile)
	if err != nil {
		t.Fatal("failed to write binary file")
	}

	laptop2 := &pb.Laptop{}
	err = ReadProtobufFomBinaryFile(binaryFile, laptop2)
	if err != nil {
		t.Fatal("error while reading file")
	}
	if !proto.Equal(laptop1, laptop2) {
		t.Fatal("both messages are different")
	}

	jsonFile := "../tmp/laptop.json"
	err = WriteProtobufToJSONFile(laptop1, jsonFile)
	if err != nil {
		t.Fatal("failed to write json file")
	}

}
