package board

import (
    "github.com/DrunkenPoney/go-tictactoe/board/ui"
    . "github.com/DrunkenPoney/go-tictactoe/grid"
    "github.com/DrunkenPoney/go-tictactoe/grid/tile"
    . "github.com/DrunkenPoney/go-tictactoe/position"
    "github.com/hajimehoshi/ebiten"
)

var DefaultPosition = MIDDLE_CENTER

type Board interface {
    Grids() map[Position]TileGrid
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
    grids   map[Position]TileGrid
    pos     Position
    prevPos Position
    screen  *ebiten.Image
    bui     *ui.BoardUI
    cellImg map[Position]*ebiten.Image
}

func NewBoard() Board {
    grids := make(map[Position]TileGrid)
    for i := 1; i <= 9; i++ {
        grids[Position(i)] = NewGrid(3, 3, &tile.Tile{Value: tile.EMPTY})
    }
    return &board{
        grids:   grids,
        pos:     DefaultPosition,
        prevPos: INVALID,
        bui:     ui.DefaultBoardUI(),
        cellImg: make(map[Position]*ebiten.Image),
    }
}
