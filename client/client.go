package main

import (
	"context"
	"io"
	"log"
	"time"

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

	// test streams
	stream, err := gc.Sync(context.Background())
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}

	ctx := stream.Context()
	done := make(chan bool)

	// first goroutine sends random Update request to stream
	// and closes int after 10 iterations
	go func() {
		log.Println("start update")
		for i := 1; i <= 10; i++ {
			log.Printf("start update %d", i)

			newGame, err := gc.Get(context.TODO(), &games.GetRequest{Id: res.Id})
			if err != nil {
				panic(err)
			}
			if err := stream.Send(&games.UpdateRequest{Id: newGame.Id, Selected: newGame.Selected * 2, Rmax: &games.Rmax{X: 2, Y: 2, Length: 2}}); err != nil {
				log.Fatalf("can not send %v", err)
			}
			time.Sleep(time.Second)
		}
		if err := stream.CloseSend(); err != nil {
			log.Println(err)
		}
	}()

	// second goroutine receives data from stream
	//
	// if stream is finished it closes done channel
	go func() {
		for {
			game, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			log.Println(game)
		}
	}()

	// third goroutine closes done channel
	// if context is done
	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
	}()

	<-done
	log.Println("finished")
}
