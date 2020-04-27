.PHONY: proto

proto:
	protoc -I . ./greeter/greeter.proto --go_out=plugins=grpc:.