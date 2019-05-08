package board

import (
    "github.com/DrunkenPoney/go-tictactoe/board/ui"
    . "github.com/DrunkenPoney/go-tictactoe/grid"
    "github.com/DrunkenPoney/go-tictactoe/grid/tile"
    "github.com/DrunkenPoney/go-tictactoe/position"
    "github.com/hajimehoshi/ebiten"
)

var DefaultPosition = position.MIDDLE_CENTER

type Board interface {
    Grids() map[position.Position]TileGrid
    CurrentGrid() TileGrid
    GetCurrentPos() position.Position
    SetCurrentPos(pos position.Position)
    GetTileUnderCursor() *tile.Tile
    ResetAll()
    DrawBoard(screen *ebiten.Image)
    UI() *ui.BoardUI
}

type board struct {
    grids  map[position.Position]TileGrid
    pos    position.Position
    screen *ebiten.Image
    bui    *ui.BoardUI
}

func NewBoard() Board {
    grids := make(map[position.Position]TileGrid)
    for i := 0; i < 9; i++ {
        grids[position.Position(i)] = NewGrid(3, 3, &tile.Tile{Value: tile.EMPTY})
    }
    return &board{
        grids: grids,
        pos: DefaultPosition,
        bui: ui.DefaultBoardUI(),
    }
}