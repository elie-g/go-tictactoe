package ai

import (
    . "github.com/DrunkenPoney/go-tictactoe/ai/prediction"
    "github.com/DrunkenPoney/go-tictactoe/ai/settings"
    . "github.com/DrunkenPoney/go-tictactoe/board/bgrid"
    . "github.com/DrunkenPoney/go-tictactoe/grid/tile"
    . "github.com/DrunkenPoney/go-tictactoe/position"
    "math/rand"
)

type AIProcess interface {
    PrepareNextTurn(pos Position)
    BestMoveFor(tt TileType) Position
    GetPrediction() Prediction
}

func NewAIProcess(player TileType, board BoardGrid) AIProcess {
    return &process{prediction: NewPrediction(player, board.Clone())}
}

type process struct {
    prediction Prediction
}

func (p *process) PrepareNextTurn(pos Position) {
    p.prediction.Next(pos)
}

func (p *process) BestMoveFor(tt TileType) Position {
    bestMove, pred := INVALID, p.prediction.Predict()
    bestScore, mod := 0.0, 1.0
    if settings.REFERENCE_PLAYER != tt {
        mod = -1.0
    }
    for pos, score := range pred {
        if (score * mod) > bestScore {
            bestMove, bestScore = pos, score * mod
        }
    }
    if bestMove == INVALID {
        empty := p.prediction.CurrentLayer().Grid().EmptyTiles()
        bestMove = empty[int(rand.Float64() * float64(len(empty)))].Position
    }
    return bestMove
}

func (p *process) GetPrediction() Prediction {
    return p.prediction
}



