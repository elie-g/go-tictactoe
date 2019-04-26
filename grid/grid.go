package grid

import (
	"fmt"
	. "github.com/DrunkenPoney/go-tictactoe/grid/tile"
	"github.com/hajimehoshi/ebiten"
	. "image/color"
	"math"
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

	GetGridNumber() int
	SetGridNumber(position []int) Grid

	Reset() Grid
}

func NewGrid(columns int, rows int, color Color, strokeWidth float64) Grid {
	tiles := make([][]Tile, columns)
	for x := range tiles {
		tiles[x] = make([]Tile, rows)
		for y := range tiles[x] {
			tiles[x][y] = NewTile(EMPTY, []int{x, y})

		}
	}
	return &grid{tiles, columns, rows, color, strokeWidth, nil, nil}
}

type grid struct {
	tiles       [][]Tile
	columns     int
	rows        int
	color       Color
	strokeWidth float64
	img         *ebiten.Image
	gridNumber  int
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

func (g *grid) GetGridNumber() int {
	return g.gridNumber
}

func (g *grid) SetGridNumber(position []int) Grid {
	var number int
	var col int
	col = int(math.Ceil(float64(position[0]) / 3))
	var row int
	row = int(math.Ceil(float64(position[1]) / 3))

	switch col {
	case 1:
		switch row {
		case 1:
			number = 1
		case 2:
			number = 4
		case 3:
			number = 7
		}
	case 2:
		switch row {
		case 1:
			number = 2
		case 2:
			number = 5
		case 3:
			number = 8
		}
	case 3:
		switch row {
		case 1:
			number = 3
		case 2:
			number = 6
		case 3:
			number = 9
		}
	}

	g.gridNumber = number
	return g
}

func (g *grid) Reset() Grid {
	for x, col := range g.tiles {
		for y := range col {
			g.tiles[x][y].SetActive(false)
			g.tiles[x][y].SetWinning(false)
			g.tiles[x][y].SetValue(EMPTY)
		}
	}
	return g
}
