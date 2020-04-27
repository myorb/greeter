package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/myorb/greeter/greeter"
)

func main() {
	c, err := grpc.Dial(":9001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	gc := greeter.NewGreeterClient(c)

	res, err := gc.SayHello(context.TODO(), &greeter.HelloRequest{Name: "Alex"})
	if err != nil {
		panic(err)
	}
	log.Println(res.Message)
}
