package game

import (
	"errors"
)

var (
	ErrNotEmpty    = errors.New("cannot place stone: another stone is already there")
	ErrOutOfBounds = errors.New("out of bounds: requested coordinates exceed board size")
	ErrHasEnded    = errors.New("game has ended: cannot play stones in a finished game")
)
