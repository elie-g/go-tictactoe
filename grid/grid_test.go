package grid

import (
	. "github.com/DrunkenPoney/go-tictactoe/grid/tile"
	"image/color"
	"testing"
)

func TestStrokeWidth(t *testing.T) {
	var grid Grid = NewGrid(3, 3, color.White, 10)
	grid.SetStrokeWidth(100)
	if grid.GetStrokeWidth() != 100 {
		t.Errorf("strokeWidth incorrect")
	}
}

func TestColor(t *testing.T) {
	var grid Grid = NewGrid(3, 3, color.White, 10)
	grid.SetColor(color.White)
	if grid.GetColor() != color.White {
		t.Errorf("color incorrect")
	}
}

func TestGetTilesCountTotal(t *testing.T) {
	var tiles [][]Tile
	var grid Grid = NewGrid(3, 3, color.White, 10)
	tiles = grid.GetTiles()
	if (len(tiles)) != 3 {
		t.Errorf("total incorrect")
	}
}

func TestGetTilesAtNotNull(t *testing.T) {
	var tile Tile
	var grid Grid = NewGrid(3, 3, color.White, 10)
	tile = grid.GetTileAt(2, 2)
	if tile.IsActive() != false {
		t.Errorf("tile null")
	}
}

func TestGetRows(t *testing.T) {
	var grid Grid = NewGrid(3, 3, color.White, 10)
	if grid.Rows() != 3 {
		t.Errorf("Rows incorrect")
	}
}

func TestGetColumn(t *testing.T) {
	var grid Grid = NewGrid(3, 3, color.White, 10)
	if grid.Columns() != 3 {
		t.Errorf("Columns incorrect")
	}
}
