package board

import (
    "github.com/DrunkenPoney/go-tictactoe/board/ui"
    . "github.com/DrunkenPoney/go-tictactoe/grid"
    "github.com/DrunkenPoney/go-tictactoe/position"
)

func (g *board) Grids() map[position.Position]TileGrid {
    return g.grids
}

func (g *board) CurrentGrid() TileGrid {
    return g.grids[g.pos]
}

func (g *board) GetCurrentPos() position.Position {
    return g.pos
}

func (g *board) SetCurrentPos(pos position.Position) {
    g.pos = pos
}

func (g *board) ResetAll() {
    for _, grid := range g.grids {
        grid.Reset()
    }
}

func (g *board) UI() *ui.BoardUI {
    return g.bui
}
