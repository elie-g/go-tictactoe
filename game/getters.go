package game

import (
    "github.com/DrunkenPoney/go-tictactoe/ai"
    "github.com/DrunkenPoney/go-tictactoe/board"
    "github.com/DrunkenPoney/go-tictactoe/events"
    "github.com/DrunkenPoney/go-tictactoe/game/player"
)

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

func (g *game) GetCurrentPlayer() player.Player {
    cur := g.playerO
    if g.playerX.IsCurrent() {
        cur = g.playerX
    }
    return cur
}

func (g *game) GetClickListener() events.ClickListener {
    return g.clickListener
}