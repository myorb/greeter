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

	gc := games.NewGamesClient(c)

	res, err := gc.Create(context.TODO(), &games.CreateRequest{Selected: 1, Rmax: &games.Rmax{X: 1, Y: 1, Length: 1}})
	if err != nil {
		panic(err)
	}
	log.Println(res.Id)

	res1, err := gc.Get(context.TODO(), &games.GetRequest{Id: res.Id})
	if err != nil {
		panic(err)
	}
	log.Println(res1)

	res2, err := gc.Update(context.TODO(), &games.UpdateRequest{Id: res.Id, Selected: 2, Rmax: &games.Rmax{X: 2, Y: 2, Length: 2}})
	if err != nil {
		panic(err)
	}
	log.Println(res2)

	res5, err := gc.Get(context.TODO(), &games.GetRequest{Id: res.Id})
	if err != nil {
		panic(err)
	}
	log.Println(res5)

	res3, err := gc.Delete(context.TODO(), &games.DeleteRequest{Id: res.Id})
	if err != nil {
		panic(err)
	}

	log.Println(res3)
}
