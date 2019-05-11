package board

import (
	. "github.com/DrunkenPoney/go-tictactoe/board/bgrid"
	"github.com/DrunkenPoney/go-tictactoe/board/ui"
	. "github.com/DrunkenPoney/go-tictactoe/grid"
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
	for _, rows := range b.GridAt(b.pos) {
		for _, cell := range rows {
			cell.Enabled = false
		}
	}
	for _, rows := range b.GridAt(pos) {
		for _, cell := range rows {
			cell.Enabled = true
		}
	}
	b.SetGridToDraw(b.pos)
	b.SetGridToDraw(pos)
	b.prevPos = b.pos
	b.pos = pos
}

func (b *board) GetPreviousPos() Position {
	return b.prevPos
}

func (b *board) ResetAll() {
	for pos, grid := range b.Grids() {
		grid.Reset()
		b.SetGridToDraw(pos)
	}
	b.SetCurrentPos(DefaultPosition)
}

func (b *board) UI() *ui.BoardUI {
	return b.bui
}