package game

import (
    "fmt"
    . "github.com/DrunkenPoney/go-tictactoe/game/player"
    "github.com/DrunkenPoney/go-tictactoe/grid"
    . "github.com/DrunkenPoney/go-tictactoe/grid/tile"
    . "github.com/DrunkenPoney/go-tictactoe/gui/message"
    . "github.com/DrunkenPoney/go-tictactoe/position"
    "github.com/sqweek/dialog"
    "os"
)

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
    if pos == INVALID {
        return g
    }
    fmt.Printf("-------------------------- NEW TURN --------------------------\n")
    winner := g.CheckWinnerInGrid(g.board.CurrentGrid())
    if winner == nil {
        g.GetPlayerX().SetCurrent(!g.GetPlayerX().IsCurrent())
        g.GetPlayerO().SetCurrent(!g.GetPlayerO().IsCurrent())
        g.GetAIProcess().PrepareNextTurn(pos)
        g.GetBoard().SetCurrentPos(pos)
        g.checkAITurn()
    } else {
        go func() {
            g.Pause()
            winner.IncrementPoints()
            msg := MSG_GAME_LOST
            if winner == g.GetPlayerO() {
                msg = MSG_GAME_WIN
            }
            ok := dialog.Message(msg.Str() + "\n\n" + MSG_NEW_GAME.Str()).Title(MSG_NEW_GAME.Str()).YesNo()
            if !ok {
                os.Exit(0)
            }
            g.Reset()
            g.checkAITurn()
        }()
    }
    return g
}

func (g *game) checkAITurn() {
    if g.GetPlayerX().IsCurrent() {
        bestPos := g.GetAIProcess().BestMoveFor(X)
        tile := g.GetBoard().CurrentGrid().At(bestPos)
        tile.Value = X
        g.GetBoard().DrawTile(tile, g.GetBoard().GetCurrentPos())
        g.NextTurn(bestPos)
    }
}
