package main

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

// Game struct
type Game struct {
	ID       string `json:"id"`
	Selected int64  `json:"selected"`
	Rmax     Rmax
}

// Rmax struct
type Rmax struct {
	X      int32 `json:"x"`
	Y      int32 `json:"y"`
	Length int32 `json:"length"`
}

// Collection struct
type Collection struct {
	sync.Mutex
	games map[string]*Game
}

// NewCollection func
func NewCollection() *Collection {
	return &Collection{
		games: map[string]*Game{
			// "0": &Game{ID: "0", Selected: 1, Rmax: Rmax{X: 0, Y: 0, Length: 0}},
			// "1": &Game{ID: "1", Selected: 2, Rmax: Rmax{X: 2, Y: 1, Length: 3}},
		},
	}
}

// NewGame func
func (c *Collection) NewGame(game *Game) (*Game, error) {
	c.Lock()
	defer c.Unlock()

	game.ID = uuid.New().String()
	c.games[game.ID] = game

	return game, nil
}

// GetGame func
func (c *Collection) GetGame(id string) (*Game, error) {
	c.Lock()
	defer c.Unlock()

	if g, ok := c.games[id]; ok {
		return g, nil
	}

	return nil, errors.New("game not found")
}

// UpdateGame func
func (c *Collection) UpdateGame(game *Game) error {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.games[game.ID]; ok {
		c.games[game.ID] = game
		return nil
	}

	return errors.New("game not found")
}

// DeleteGame func
func (c *Collection) DeleteGame(id string) error {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.games[id]; ok {
		delete(c.games, id)
		return nil
	}

	return errors.New("game not found")
}
