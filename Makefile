.PHONY: proto

proto:
	protoc -I . ./games/games.proto --go_out=plugins=grpc:.