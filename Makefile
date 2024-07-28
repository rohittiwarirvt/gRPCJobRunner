.PHONY: proto

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	helloworld/helloworld.proto

server:
	go run greeter_service/main.go -port 50053

client:
	go run greeter_client/main.go -name "Rohit Tiwari" -addr "localhost:50053"