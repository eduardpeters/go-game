package game

import (
	"time"

	"github.com/google/uuid"
)

const (
	EMPTY int = 0
	BLACK int = 1
	WHITE int = 2
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

func (g Game) GetStoneAt(x, y int) (int, error) {
	if x < 0 || x >= g.Size || y < 0 || y >= g.Size {
		return 0, ErrOutOfBounds
	}
	return g.Board[y*g.Size+x], nil
}
func (g Game) setStoneAt(stone int, x, y int) {
	g.Board[y*g.Size+x] = stone
}

func (g Game) PlaceStone(stone int, x, y int) error {
	free, err := g.isIntersectionFree(x, y)
	if err != nil {
		return err
	}
	if !free {
		return ErrNotEmpty
	}
	g.setStoneAt(stone, x, y)
	return nil
}

func (g Game) isIntersectionFree(x, y int) (bool, error) {
	stone, err := g.GetStoneAt(x, y)
	if err != nil {
		return false, err
	}
	return stone == EMPTY, nil
}
