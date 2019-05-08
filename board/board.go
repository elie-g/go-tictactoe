package board

import (
    "github.com/DrunkenPoney/go-tictactoe/board/ui"
    . "github.com/DrunkenPoney/go-tictactoe/grid"
    "github.com/DrunkenPoney/go-tictactoe/position"
    "github.com/hajimehoshi/ebiten"
)

func (g *board) Grids() [3][3]TileGrid {
    return g.grids
}

func (g *board) CurrentGrid() TileGrid {
    x, y := g.pos.GetXY()
    return g.grids[x][y]
}


func (g *board) GetCurrentPos() position.Position {
    return g.pos
}

func (g *board) SetCurrentPos(pos position.Position) {
    g.pos = pos
}

func (g *board) DrawBoard(img *ebiten.Image) {
    panic("implement me")
}

func (g *board) ResetAll() {
    for _, rows := range g.grids {
        for _, grid := range rows {
            grid.Reset()
        }
    }
}

func (g *board) UI() *ui.BoardUI {
    return g.bui
}
