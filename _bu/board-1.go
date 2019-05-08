package _bu

import (
	. "github.com/DrunkenPoney/go-tictactoe/grid/tile"
	"github.com/hajimehoshi/ebiten"
	. "image/color"
)

type Board interface {
	Columns() int
	Rows() int

	GetTiles() [][]Tile
	GetTileAt(x int, y int) Tile

	DrawGrid(image *ebiten.Image)

	GetColor() Color
	SetColor(color Color) Board

	GetStrokeWidth() float64
	SetStrokeWidth(w float64) Board

	GetTileUnderCursor() *Tile

	GetGridNumber() int
	SetGridNumber(position []int) Board
	SetGridNumberFromInt(position int) Board

	GetColOffset() int
	GetRowOffset() int

	GetCurrentGrid() [][]*Tile

	Reset() Board
}

func NewGrid(columns int, rows int, color Color, strokeWidth float64) Board {
	tiles := make([][]Tile, columns)
	for x := range tiles {
		tiles[x] = make([]Tile, rows)
		for y := range tiles[x] {
			tiles[x][y] = NewTile(EMPTY, []int{x, y})

		}
	}
	return &board{tiles, columns, rows, color, strokeWidth, nil, 0}
}

type board struct {
	tiles       [][]Tile
	columns     int
	rows        int
	color       Color
	strokeWidth float64
	img         *ebiten.Image
	gridNumber  int
}

func (g *board) Columns() int {
	return g.columns
}

func (g *board) Rows() int {
	return g.rows
}

func (g *board) GetTiles() [][]Tile {
	return g.tiles
}

func (g *board) GetTileAt(x int, y int) Tile {
	return g.tiles[x][y]
}

func (g *board) GetColor() Color {
	return g.color
}

func (g *board) SetColor(color Color) Board {
	g.color = color
	return g
}

func (g *board) GetStrokeWidth() float64 {
	return g.strokeWidth
}

func (g *board) SetStrokeWidth(w float64) Board {
	g.strokeWidth = w
	return g
}

func (g *board) GetGridNumber() int {
	return g.gridNumber
}


func (g *board) SetGridNumber(position []int) Board {
	var number int
	var col int
	col = int(position[0]%3) + 1
	var row int
	row = int(position[1]%3) + 1

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

func (g *board) SetGridNumberFromInt(position int) Board {
	g.gridNumber = position
	return g
}

func (g *board) GetColOffset() int {
	var offset int
	offset = 0
	if g.gridNumber%3 == 2 {
		offset = 3
	}
	if g.gridNumber%3 == 0 {
		offset = 6
	}
	return offset
}

func (g *board) GetRowOffset() int {
	var offset int
	offset = 0
	if g.gridNumber == 4 || g.gridNumber == 5 || g.gridNumber == 6 {
		offset = 3
	}
	if g.gridNumber == 7 || g.gridNumber == 8 || g.gridNumber == 9 {
		offset = 6
	}
	return offset
}

func (g *board) GetCurrentGrid() [][]Tile {
	currentGrid := make([][]Tile, 3)
	for x := range currentGrid {
		currentGrid[x] = make([]Tile, 3)
		for y := range currentGrid[x] {
			currentGrid[x][y] = g.GetTileAt(x+g.GetColOffset(), y+g.GetRowOffset())
		}
	}

	return currentGrid
}

func (g *board) Reset() Board {
	for x, col := range g.tiles {
		for y := range col {
			g.tiles[x][y].SetActive(false)
			g.tiles[x][y].SetWinning(false)
			g.tiles[x][y].SetValue(EMPTY)
		}
	}
	return g
}
