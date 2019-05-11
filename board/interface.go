package board

import (
	. "github.com/DrunkenPoney/go-tictactoe/board/bgrid"
	"github.com/DrunkenPoney/go-tictactoe/board/ui"
	. "github.com/DrunkenPoney/go-tictactoe/grid"
	"github.com/DrunkenPoney/go-tictactoe/grid/tile"
	. "github.com/DrunkenPoney/go-tictactoe/position"
	"github.com/hajimehoshi/ebiten"
)

var DefaultPosition = MIDDLE_CENTER

type Board interface {
	Grids() BoardGrid
	CurrentGrid() TileGrid
	GridAt(pos Position) TileGrid
	GetCurrentPos() Position
	SetCurrentPos(pos Position)
	GetPreviousPos() Position
	GetTileUnderCursor() (*tile.Tile, Position)
	ResetAll()
	UI() *ui.BoardUI

	DrawBoard(screen *ebiten.Image)
	DrawTile(t *tile.Tile, pos Position)
	DrawTileUnderCursor()
	SetGridToDraw(pos Position)
}

type board struct {
	grids       BoardGrid
	pos         Position
	prevPos     Position
	screen      *ebiten.Image
	bui         *ui.BoardUI
	cellImg     map[Position]*ebiten.Image
}

func NewBoard() Board {
	b :=  &board{
		grids:   NewBoardGrid(),
		prevPos: INVALID,
		bui:     ui.DefaultBoardUI(),
		cellImg: make(map[Position]*ebiten.Image),
	}
	b.SetCurrentPos(DefaultPosition)
	return b
}
