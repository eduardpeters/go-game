package game

import (
	"time"

	"github.com/google/uuid"
)

type Game struct {
	ID        string
	Size      int
	CreatedAt time.Time
	UpdatedAt time.Time
	Board     []int
	Moves     []string
}

func NewGame(size int) *Game {
	now := time.Now()
	g := &Game{
		ID:        uuid.NewString(),
		Size:      size,
		CreatedAt: now,
		UpdatedAt: now,
		Board:     make([]int, size*size),
		Moves:     []string{},
	}
	return g
}
