package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"games/games"
)

func main() {
	c, err := grpc.Dial(":9001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	gamesClient := games.NewGamesClient(c)

	res2, err := gamesClient.Get(context.TODO(), &games.GetRequest{Id: "1"})
	if err != nil {
		panic(err)
	}
	log.Println(res2)
}
