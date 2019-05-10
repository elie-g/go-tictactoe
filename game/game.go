package game

import (
    "fmt"
    "github.com/DrunkenPoney/go-tictactoe/ai"
    "github.com/DrunkenPoney/go-tictactoe/board"
    "github.com/DrunkenPoney/go-tictactoe/events"
    "github.com/DrunkenPoney/go-tictactoe/game/player"
    "github.com/DrunkenPoney/go-tictactoe/grid"
    . "github.com/DrunkenPoney/go-tictactoe/grid/tile"
    . "github.com/DrunkenPoney/go-tictactoe/position"
    "github.com/hajimehoshi/ebiten"
    "math/rand"
)

type Game interface {
    GetPlayerO() player.Player
    GetPlayerX() player.Player
    NextTurn(pos Position) Game
    GetBoard() board.Board
    GetAIProcess() ai.AIProcess
    CheckWinnerInGrid(tiles grid.TileGrid) player.Player
    GetCurrentPlayer() player.Player
    Reset() Game
    Draw(screen *ebiten.Image) Game
}

func NewGame(playerO player.Player, playerX player.Player, board board.Board) Game {
    if !playerO.IsCurrent() && !playerX.IsCurrent() ||
        playerO.IsCurrent() && playerX.IsCurrent() {
        rdm := rand.Float64() >= 0.5
        playerO.SetCurrent(rdm)
        playerX.SetCurrent(!rdm)
    }
    var aiProcess ai.AIProcess
    if playerO.IsCurrent() {
        aiProcess = ai.NewAIProcess(O, board.Grids())
    } else {
        aiProcess = ai.NewAIProcess(X, board.Grids())
    }
    listener := events.NewClickListener()
    g := &game{playerO, playerX, aiProcess, board, listener}
    listener.Listen(g.onClick)
    listener.Resume()
    return g
}

type game struct {
    playerO       player.Player
    playerX       player.Player
    ai            ai.AIProcess
    board         board.Board
    clickListener events.ClickListener
}

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

func (g *game) GetPlayerO() player.Player {
    return g.playerO
}

func (g *game) GetPlayerX() player.Player {
    return g.playerX
}

func (g *game) GetBoard() board.Board {
    return g.board
}

func (g *game) GetAIProcess() ai.AIProcess {
    return g.ai
}

func (g *game) CheckWinnerInGrid(tiles grid.TileGrid) player.Player {
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

func (g *game) Reset() Game {
    g.GetBoard().ResetAll()
    return g
}

func (g *game) Draw(screen *ebiten.Image) Game {
    g.GetBoard().DrawBoard(screen)
    return g
}

func (g *game) GetCurrentPlayer() player.Player {
    cur := g.playerO
    if g.playerX.IsCurrent() {
        cur = g.playerX
    }
    return cur
}
