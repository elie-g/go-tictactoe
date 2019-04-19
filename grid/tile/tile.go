package tile

import (
    "github.com/hajimehoshi/ebiten"
)

type TileType uint8

const (
    EMPTY TileType = 0
    X     TileType = 1
    O     TileType = 2
)

type Tile interface {
    GetValue() TileType
    SetValue(value TileType) Tile
    
    IsActive() bool
    SetActive(active bool) Tile
    
    IsWinning() bool
    SetWinning(win bool) Tile
    
    GetImage() *ebiten.Image
}

func NewTile(value TileType) Tile {
    return &tile{value, false, false}
}

type tile struct {
    value   TileType
    active  bool
    winning bool
}

func (t *tile) GetValue() TileType {
    return t.value
}

func (t *tile) SetValue(value TileType) Tile {
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