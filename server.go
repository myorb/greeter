package main

import (
	"context"
	"io"
	"log"

	"games/games"
)

type server struct {
	repository *Collection
}

// Get func
func (s *server) Get(ctx context.Context, r *games.GetRequest) (*games.Game, error) {
	g, err := s.repository.GetGame(r.Id)
	if err != nil {
		return nil, err
	}
	return convert(g), nil
}

// Create func
func (s *server) Create(ctx context.Context, r *games.CreateRequest) (*games.Game, error) {
	g, err := s.repository.NewGame(&Game{Selected: r.Selected, Rmax: Rmax{X: r.Rmax.X, Y: r.Rmax.Y, Length: r.Rmax.Length}})
	if err != nil {
		return nil, err
	}
	return convert(g), nil
}

// Update func
func (s *server) Update(c context.Context, r *games.UpdateRequest) (*games.BoolResponce, error) {
	err := s.repository.UpdateGame(&Game{ID: r.Id, Selected: r.Selected, Rmax: Rmax{X: r.Rmax.X, Y: r.Rmax.Y, Length: r.Rmax.Length}})
	if err != nil {
		return nil, err
	}
	return &games.BoolResponce{Status: true}, nil
}

// Delete func
func (s *server) Delete(c context.Context, r *games.DeleteRequest) (*games.BoolResponce, error) {
	err := s.repository.DeleteGame(r.Id)
	if err != nil {
		return nil, err
	}
	return &games.BoolResponce{Status: true}, nil
}

// Sync func
func (s *server) Sync(srv games.Games_SyncServer) error {
	log.Println("start new server")
	ctx := srv.Context()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// receive data from stream
		req, err := srv.Recv()
		if err == io.EOF {
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}

		if err = s.repository.UpdateGame(&Game{ID: req.Id, Selected: req.Selected, Rmax: Rmax{X: req.Rmax.X, Y: req.Rmax.Y, Length: req.Rmax.Length}}); err != nil {
			return err
		}

		g, err := s.repository.GetGame(req.Id)
		if err != nil {
			return err
		}
		resp := convert(g)

		if err := srv.Send(resp); err != nil {
			return err
		}
		log.Printf("game updated %s", g.ID)
	}
}

// converting form grpc to Game resp struct
func convert(g *Game) *games.Game {
	return &games.Game{Id: g.ID, Selected: g.Selected, Rmax: &games.Rmax{X: g.Rmax.X, Y: g.Rmax.Y, Length: g.Rmax.Length}}
}
