package game

import (
    "github.com/DrunkenPoney/go-tictactoe/ai"
    "github.com/DrunkenPoney/go-tictactoe/board"
    . "github.com/DrunkenPoney/go-tictactoe/game/player"
    . "github.com/DrunkenPoney/go-tictactoe/game/state"
    "github.com/DrunkenPoney/go-tictactoe/grid"
    . "github.com/DrunkenPoney/go-tictactoe/grid/tile"
    . "github.com/DrunkenPoney/go-tictactoe/position"
    "github.com/hajimehoshi/ebiten"
    "math/rand"
    "runtime"
)

type Game interface {
    GetPlayerO() Player
    GetPlayerX() Player
    NextTurn(pos Position) Game
    GetBoard() board.Board
    GetAIProcess() ai.AIProcess
    CheckWinnerInGrid(tiles grid.TileGrid) Player
    GetCurrentPlayer() Player
    
    StateChannel() chan State
    OnClick()
    
    Resume()
    Pause()
    IsPaused() bool
    Reset() Game
    
    Draw(screen *ebiten.Image) Game
}

func NewGame(playerO Player, playerX Player, board board.Board) Game {
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
    g := &game{playerO, playerX, aiProcess,
        board, make(chan State), RUNNING}
    go func() {
        for {
            g.onState(<-g.stateChan)
        }
    }()
    return g
}

type game struct {
    playerO   Player
    playerX   Player
    ai        ai.AIProcess
    board     board.Board
    stateChan chan State
    state     State
}

func (g *game) Reset() Game {
    g.Pause()
    g.GetBoard().ResetAll()
    tile := O
    if g.GetPlayerX().IsCurrent() {
        tile = X
    }
    g.ai = nil
    // Le jeu freeze si on ne force pas le garbage collection
    runtime.GC()
    g.ai = ai.NewAIProcess(tile, g.GetBoard().Grids())
    g.Resume()
    return g
}

func (g *game) Draw(screen *ebiten.Image) Game {
    g.GetBoard().DrawBoard(screen)
    return g
}

func (g *game) StateChannel() chan State {
    return g.stateChan
}

func (g *game) Pause() {
    if g.state != PAUSED && g.state != STOPPED {
        g.state = PAUSED
        g.stateChan <- PAUSED
    }
}

func (g *game) Resume() {
    if g.state != RUNNING && g.state != STOPPED {
        g.state = RUNNING
        g.stateChan <- RUNNING
    }
}

func (g *game) IsPaused() bool {
    return g.state == PAUSED || g.state == STOPPED
}
