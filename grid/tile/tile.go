package tile

import (
    . "image/color"
)

// Private type
type tval uint8

const (
    EMPTY tval = 0
    X     tval = 1
    O     tval = 2
    // X     tval = '\uf00d'
    // O     tval = '\uf111'
)

type Tile interface {
    GetValue() tval
    SetValue(value tval) Tile
    
    GetColor() Color
    SetColor(color Color) Tile
    
    GetBgColor() Color
    SetBgColor(color Color) Tile
    
    GetStrokeWidth() float64
    SetStrokeWidth(w float64) Tile
}

func NewTile(value tval, color Color, strkWidth float64) Tile {
    return &tile{value, color, Transparent, strkWidth}
}

type tile struct {
    value tval
    color Color
    bgColor Color
    strokeWidth float64
}

func (t *tile) GetValue() tval {
    return t.value
}

func (t *tile) SetValue(value tval) Tile {
    t.value = value
    return t
}

func (t *tile) GetColor() Color {
    return t.color
}

func (t *tile) SetColor(color Color) Tile {
    t.color = color
    return t
}

func (t *tile) GetBgColor() Color {
    return t.bgColor
}

func (t *tile) SetBgColor(color Color) Tile {
    t.bgColor = color
    return t
}

func (t *tile) GetStrokeWidth() float64 {
    return t.strokeWidth
}

func (t *tile) SetStrokeWidth(w float64) Tile {
    t.strokeWidth = w
    return t
}

