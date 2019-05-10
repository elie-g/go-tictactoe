package game

import (
    "fmt"
    . "github.com/DrunkenPoney/go-tictactoe/game/player"
    "github.com/DrunkenPoney/go-tictactoe/grid"
    . "github.com/DrunkenPoney/go-tictactoe/grid/tile"
    . "github.com/DrunkenPoney/go-tictactoe/position"
)

// Private
func (g *game) onClick() {
    t, pos := g.board.GetTileUnderCursor()
    if t != nil && t.Value == EMPTY {
        if g.playerO.IsCurrent() {
            t.Value = O
        } else {
            t.Value = X
        }
        g.GetBoard().DrawTile(t, pos)
        g.NextTurn(t.Position)
    }
}

func (g *game) CheckWinnerInGrid(tiles grid.TileGrid) Player {
    cells := tiles.GetWinningTiles()
    if cells[0] != nil {
        for _, cell := range cells {
            cell.Winning = true
        }
        
        if cells[0].Value == X {
            return g.playerX
        } else {
            return g.playerO
        }
    }
    return nil
}

func (g *game) NextTurn(pos Position) Game {
    if pos == INVALID { return g }
    fmt.Printf("-------------------------- NEW TURN --------------------------\n")
    g.CheckWinnerInGrid(g.board.CurrentGrid())
    g.playerX.SetCurrent(!g.playerX.IsCurrent())
    g.playerO.SetCurrent(!g.playerO.IsCurrent())
    g.ai.PrepareNextTurn(pos)
    g.GetBoard().SetCurrentPos(pos)
    
    if g.playerX.IsCurrent() {
        bestPos := g.ai.BestMoveFor(X)
        tile := g.GetBoard().CurrentGrid().At(bestPos)
        tile.Value = X
        g.GetBoard().DrawTile(tile, g.GetBoard().GetCurrentPos())
        g.NextTurn(bestPos)
    }
    return g
}