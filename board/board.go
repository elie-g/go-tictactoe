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