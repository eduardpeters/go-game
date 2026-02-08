package game_test

import (
	"errors"
	"testing"

	"github.com/eduardpeters/go-game/internal/game"
)

func TestPlaceBlackStoneUpdatesBoard(t *testing.T) {
	g := game.NewGame(9)
	stone := game.BLACK
	x := 0
	y := 0
	err := g.PlaceStone(stone, x, y)
	if err != nil {
		t.Fatal("unexpected error placing stone in empty board")
	}

	want := stone
	got, err := g.GetStoneAt(x, y)
	if err != nil {
		t.Fatal("unexpected error getting stone in board")
	}

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestCannotPlaceStoneTwice(t *testing.T) {
	g := game.NewGame(9)
	stone := game.BLACK
	y := 0
	x := 0
	err := g.PlaceStone(stone, x, y)
	if err != nil {
		t.Fatal("unexpected error placing stone in empty board")
	}
	err = g.PlaceStone(stone, x, y)
	if !errors.Is(err, game.ErrNotEmpty) {
		t.Errorf("expected error, got %v", err)
	}
}

func TestCannotGetStoneOutsideBoard(t *testing.T) {
	size := 9
	g := game.NewGame(size)

	tests := []struct {
		x, y int
	}{
		{9, 0},
		{0, 9},
		{10, 10},
		{8, 10},
		{10, 8},
		{-1, -1},
		{0, -1},
		{-1, 0},
	}
	for _, tt := range tests {
		_, err := g.GetStoneAt(tt.x, tt.y)
		if !errors.Is(err, game.ErrOutOfBounds) {
			t.Errorf("expected error in board size: %d for x: %d y: %d, got %v", size, tt.x, tt.y, err)
		}
	}
}

func TestPlacingStoneSwitchesTurns(t *testing.T) {
	g := game.NewGame(9)
	stone := game.BLACK
	y := 0
	x := 0
	err := g.PlaceStone(stone, x, y)
	if err != nil {
		t.Fatal("unexpected error placing stone in empty board")
	}

	want := game.WHITE
	got := g.CurrentStone
	if got != want {
		t.Errorf("Turn was not switched after placing stone. got %d want %d", got, want)
	}
}
