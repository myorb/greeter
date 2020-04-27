package main

import (
	"google.golang.org/grpc"

	"github.com/myorb/greeter/greeter"
)

func main() {
	c, err := grpc.Dial(":9001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	gc := greeter.NewGreeterClient(c)

}
