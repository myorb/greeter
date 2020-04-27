package main

import (
	"github.com/myorb/greeter/greeter"
)

type server struct {
	greeter.UnimplementedGreeterServer
}
