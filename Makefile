all: client server

protoc:
	@echo "Generating Go files"
	protoc -I . ./games/games.proto --go_out=plugins=grpc:.

server: protoc
	@echo "Building server"
	go build -o game-server

client: protoc
	@echo "Building client"
	cd client && go build -o ../game-client 

clean:
	rm -f game-server game-client

.PHONY: client server protoc