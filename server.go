package main

import (
	"context"

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

// converting form grpc to Game resp struct
func convert(g *Game) *games.Game {
	return &games.Game{Id: g.ID, Selected: g.Selected, Rmax: &games.Rmax{X: g.Rmax.X, Y: g.Rmax.Y, Length: g.Rmax.Length}}
}
