package main

import (
	"log"
	"net"

	"github.com/myorb/greeter/games"
	"github.com/myorb/greeter/greeter"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gamesCollection := NewCollection()

	s := grpc.NewServer()
	srv := &server{repository: gamesCollection}
	greeter.RegisterGreeterServer(s, srv)
	games.RegisterGamesServer(s, srv)
	log.Println("This is log form main")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
