package game

import (
    "fmt"
    . "github.com/DrunkenPoney/go-tictactoe/db"
    "github.com/DrunkenPoney/go-tictactoe/game/player"
    . "github.com/DrunkenPoney/go-tictactoe/grid/tile"
    . "github.com/DrunkenPoney/go-tictactoe/position"
    "github.com/DrunkenPoney/go-tictactoe/settings"
    . "github.com/DrunkenPoney/go-tictactoe/settings/messages"
    "github.com/sqweek/dialog"
    "os"
    "time"
)

func (g *game) IsOnline() bool {
    return g.online
}

func (g *game) InitOnline(dbGame DBGame, begin bool) {
    g.Reset()
    g.db = NewDatabase()
    db := g.db
    g.dbGame = dbGame
    g.online = true
    
    lastTurn := db.LastTurn(dbGame)
    // if !begin && lastTurn == nil {
    go g.waitTurn(g.noTurn)
    // }
    
    if lastTurn != nil {
        gridPos := Position(lastTurn.GetGridPos())
        tile := g.GetBoard().GridAt(gridPos).At(Position(lastTurn.GetSubGridPos()))
        if g.GetPlayerO().IsCurrent() {
            tile.Value = O
        } else {
            tile.Value = X
        }
        g.GetBoard().DrawTile(tile, gridPos)
        // g.NextTurn(tile.Position)
    }
}

func (g *game) ExitOnline() {
    // TODO
}

func (g *game) waitTurn(noTurn int) DBTurn {
    turn := g.db.LastTurn(g.dbGame)
    fmt.Println("WAITING")
    for turn == nil || turn.GetNumber() < int64(noTurn) {
        <-time.After(settings.REFRESH_DELAY)
        turn = g.db.LastTurn(g.dbGame)
    }
    fmt.Println("\033[36mCarotte Bleue!\033[m")
    
    gridPos := Position(turn.GetGridPos())
    tile := g.GetBoard().GridAt(gridPos).At(Position(turn.GetSubGridPos()))
    if g.GetPlayerO().IsCurrent() {
        tile.Value = O
    } else {
        tile.Value = X
    }
    g.GetBoard().DrawTile(tile, gridPos)
    g.noTurn++
    g.GetPlayerX().SetCurrent(true)
    g.GetPlayerO().SetCurrent(false)
    winner := g.CheckWinnerInGrid(g.board.CurrentGrid())
    if winner == nil {
        g.GetBoard().SetCurrentPos(tile.Position)
        g.GetAIProcess().PrepareNextTurn(tile.Position)
    
        bestPos := g.GetAIProcess().BestMoveFor(X)
        tile = g.GetBoard().CurrentGrid().At(bestPos)
        tile.Value = X
        turn = g.db.CreateTurn(g.GetPlayerX().GetDBPlayer(), g.dbGame, g.GetBoard().GetCurrentPos(), bestPos)
        g.noTurn = int(turn.GetNumber())
        g.GetPlayerX().SetCurrent(false)
        g.GetPlayerO().SetCurrent(true)
        winner = g.CheckWinnerInGrid(g.board.CurrentGrid())
        if winner == nil {
            g.GetBoard().SetCurrentPos(tile.Position)
            g.GetAIProcess().PrepareNextTurn(tile.Position)
        }
    }
    if winner != nil {
        g.endGame(winner)
    }
    // g.NextTurn(tile.Position)
    g.waitTurn(g.noTurn + 1)
    return turn
}

func (g *game) endGame(winner player.Player) {
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
    // g.checkAITurn()
}
