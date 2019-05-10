package board

import (
	. "github.com/DrunkenPoney/go-tictactoe/board/bgrid"
	"github.com/DrunkenPoney/go-tictactoe/board/ui"
	. "github.com/DrunkenPoney/go-tictactoe/grid"
	"github.com/DrunkenPoney/go-tictactoe/grid/tile"
	. "github.com/DrunkenPoney/go-tictactoe/position"
)

func (b *board) Grids() BoardGrid {
	return b.grids
}

func (b *board) GridAt(pos Position) TileGrid {
	return b.grids.Get(pos)
}

func (b *board) CurrentGrid() TileGrid {
	return b.grids.Get(b.pos)
}

func (b *board) GetCurrentPos() Position {
	return b.pos
}

func (b *board) SetCurrentPos(pos Position) {
	b.prevPos = b.pos
	b.pos = pos
}

func (b *board) GetPreviousPos() Position {
	return b.prevPos
}

func (b *board) ResetAll() {
	for _, grid := range b.grids {
		grid.Reset()
	}
}

func (b *board) UI() *ui.BoardUI {
	return b.bui
}

func (b *board) IsValidTile(tile *tile.Tile) bool {
	for i := TOP_LEFT; i < (BOTTOM_RIGHT); i++ {
		for j := TOP_LEFT; j < (BOTTOM_RIGHT); j++ {
			b.GridAt(i).At(j).Active = false
		}
	}

	for i := TOP_LEFT; i < (BOTTOM_RIGHT); i++ {
		b.CurrentGrid().At(i).Active = true
		b.DrawTile(b.CurrentGrid().At(i), b.CurrentGrid().At(i).Position)

		if b.CurrentGrid().At(i) == tile {
			return true
		}
	}
	return false
}
