gen:
	protoc --proto_path=proto --go_out=. ./proto/*.proto --go-grpc_out=.

clean:
	rm ./pb/*.go

print:
	pwd

server:
	go run ./cmd/server/main.go -port 8080

client:
	go run ./cmd/client/main.go -address 0.0.0.0:8080

test:
	go test -v -race -cover ./...

#protoc --proto_path=proto \
--go_out=pb \
--go_opt=Mmemory_message.proto=example.com/project/proto \
--go_opt=Mprocessor_message.proto=example.com/project/proto \
--go_opt=paths=source_relative memory_message.proto processor_message.proto

#protoc --proto_path=proto \
--go_out=pb \
./proto/*.proto \
--go_opt=paths=source_relative memory_message.proto processor_message.proto storage_message.proto