package grid

import (
    . "github.com/DrunkenPoney/go-tictactoe/grid/tile"
    "github.com/hajimehoshi/ebiten"
    . "image/color"
)


type Grid interface {
    Columns() int
    Rows() int
    
    GetTiles() [][]Tile
    GetTileAt(x int, y int) Tile
    
    DrawGrid(image *ebiten.Image)
    
    GetColor() Color
    SetColor(color Color) Grid
    
    GetStrokeWidth() float64
    SetStrokeWidth(w float64) Grid
    
    GetTileUnderCursor() Tile
}

func NewGrid(columns int, rows int, color Color, strokeWidth float64) Grid {
    var tiles [][]Tile
    for x := 0; x < columns; x++ {
        var col []Tile
        for y := 0; y < rows; y++ {
            col = append(col, NewTile(EMPTY))
        }
        tiles = append(tiles, col)
    }
    return &grid{tiles, columns, rows, color, strokeWidth, nil}
}

type grid struct {
    tiles       [][]Tile
    columns     int
    rows        int
    color       Color
    strokeWidth float64
    img         *ebiten.Image
}

func (g *grid) Columns() int {
    return g.columns
}

func (g *grid) Rows() int {
    return g.rows
}

func (g *grid) GetTiles() [][]Tile {
    return g.tiles
}

func (g *grid) GetTileAt(x int, y int) Tile {
    return g.tiles[x][y]
}

func (g *grid) GetColor() Color {
    return g.color
}

func (g *grid) SetColor(color Color) Grid {
    g.color = color
    return g
}

func (g *grid) GetStrokeWidth() float64 {
    return g.strokeWidth
}

func (g *grid) SetStrokeWidth(w float64) Grid {
    g.strokeWidth = w
    return g
}