package tile

import (
	"github.com/DrunkenPoney/go-tictactoe/position"
)

type TileType uint8

const (
	EMPTY TileType = 0
	X     TileType = 1
	O     TileType = 2
)

type Tile struct {
	Value    TileType
	Active   bool
	Winning  bool
	Position position.Position
}

func (t *Tile) Clone() *Tile {
	return &Tile{
		Value:    t.Value,
		Active:   t.Active,
		Winning:  t.Winning,
		Position: t.Position,
	}
}
