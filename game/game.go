package game

import (
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

func (g *game) Reset() Game {
    g.GetBoard().ResetAll()
    return g
}

func (g *game) Draw(screen *ebiten.Image) Game {
    g.GetBoard().DrawBoard(screen)
    return g
}
