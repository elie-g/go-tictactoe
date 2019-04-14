package tile

import (
    "github.com/hajimehoshi/ebiten"
)

// Private type
type tval uint8

const (
    EMPTY tval = 0
    X     tval = 1
    O     tval = 2
)

type Tile interface {
    GetValue() tval
    SetValue(value tval) Tile
    
    IsActive() bool
    SetActive(active bool) Tile
    
    IsWinning() bool
    SetWinning(win bool) Tile
    
    GetImage() *ebiten.Image
}

func NewTile(value tval) Tile {
    return &tile{value, false, false}
}

type tile struct {
    value   tval
    active  bool
    winning bool
}

func (t *tile) GetValue() tval {
    return t.value
}

func (t *tile) SetValue(value tval) Tile {
    t.value = value
    return t
}

func (t *tile) IsActive() bool {
    return t.active
}

func (t *tile) SetActive(active bool) Tile {
    t.active = active
    return t
}

func (t *tile) IsWinning() bool {
    return t.winning
}

func (t *tile) SetWinning(win bool) Tile {
    t.winning = win
    return t
}