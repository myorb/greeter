.PHONY: proto

proto:
	protoc -I . ./greeter/greeter.proto --go_out=plugins=grpc:.

gproto:
	protoc -I . ./games/games.proto --go_out=plugins=grpc:.