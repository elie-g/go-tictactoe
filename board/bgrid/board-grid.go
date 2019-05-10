package bgrid

import (
    "fmt"
    . "github.com/DrunkenPoney/go-tictactoe/grid"
    "github.com/DrunkenPoney/go-tictactoe/grid/tile"
    . "github.com/DrunkenPoney/go-tictactoe/position"
)

type BoardGrid map[Position]TileGrid

func (bg BoardGrid) Clone() BoardGrid {
    bgrid := make(BoardGrid, len(bg))
    for pos, grid := range bg {
        bgrid[pos] = grid.Clone()
    }
    return bgrid
}

func (bg BoardGrid) Get(pos Position) TileGrid {
    return bg[pos]
}

func NewBoardGrid() BoardGrid {
    grid := make(BoardGrid)
    for i := 1; i <= 9; i++ {
        grid[Position(i)] = NewGrid(3, 3, &tile.Tile{Value: tile.EMPTY})
    }
    return grid
}

func (bg BoardGrid) String() string {
    var arr []interface{}
    appendLine := func(g TileGrid, line int) {
        arr = append(arr, g[0][line].Value, g[1][line].Value, g[2][line].Value)
    }
    
    appendLine(bg.Get(TOP_LEFT), 0)
    appendLine(bg.Get(TOP_CENTER), 0)
    appendLine(bg.Get(TOP_RIGHT), 0)
    appendLine(bg.Get(TOP_LEFT), 1)
    appendLine(bg.Get(TOP_CENTER), 1)
    appendLine(bg.Get(TOP_RIGHT), 1)
    appendLine(bg.Get(TOP_LEFT), 2)
    appendLine(bg.Get(TOP_CENTER), 2)
    appendLine(bg.Get(TOP_RIGHT), 2)
    
    appendLine(bg.Get(MIDDLE_LEFT), 0)
    appendLine(bg.Get(MIDDLE_CENTER), 0)
    appendLine(bg.Get(MIDDLE_RIGHT), 0)
    appendLine(bg.Get(MIDDLE_LEFT), 1)
    appendLine(bg.Get(MIDDLE_CENTER), 1)
    appendLine(bg.Get(MIDDLE_RIGHT), 1)
    appendLine(bg.Get(MIDDLE_LEFT), 2)
    appendLine(bg.Get(MIDDLE_CENTER), 2)
    appendLine(bg.Get(MIDDLE_RIGHT), 2)
    
    appendLine(bg.Get(BOTTOM_LEFT), 0)
    appendLine(bg.Get(BOTTOM_CENTER), 0)
    appendLine(bg.Get(BOTTOM_RIGHT), 0)
    appendLine(bg.Get(BOTTOM_LEFT), 1)
    appendLine(bg.Get(BOTTOM_CENTER), 1)
    appendLine(bg.Get(BOTTOM_RIGHT), 1)
    appendLine(bg.Get(BOTTOM_LEFT), 2)
    appendLine(bg.Get(BOTTOM_CENTER), 2)
    appendLine(bg.Get(BOTTOM_RIGHT), 2)
    
    return fmt.Sprintf("\n===============================\n" +
        "[%d][%d][%d]  [%d][%d][%d]  [%d][%d][%d]\n" +
        "[%d][%d][%d]  [%d][%d][%d]  [%d][%d][%d]\n" +
        "[%d][%d][%d]  [%d][%d][%d]  [%d][%d][%d]\n" +
        "                                  \n" +
        "[%d][%d][%d]  [%d][%d][%d]  [%d][%d][%d]\n" +
        "[%d][%d][%d]  [%d][%d][%d]  [%d][%d][%d]\n" +
        "[%d][%d][%d]  [%d][%d][%d]  [%d][%d][%d]\n" +
        "                                  \n" +
        "[%d][%d][%d]  [%d][%d][%d]  [%d][%d][%d]\n" +
        "[%d][%d][%d]  [%d][%d][%d]  [%d][%d][%d]\n" +
        "[%d][%d][%d]  [%d][%d][%d]  [%d][%d][%d]\n" +
        "===============================\n\n", arr...)
}
