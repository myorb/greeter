package main

import (
	"context"
	"fmt"

	"github.com/myorb/greeter/greeter"
)

type server struct {
	greeter.UnimplementedGreeterServer
}

func (*server) SayHello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloReply, error) {
	return &greeter.HelloReply{Message: fmt.Sprintf("hi %s", req.Name)}, nil
}
