package ai

import (
    . "github.com/DrunkenPoney/go-tictactoe/ai/prediction"
    . "github.com/DrunkenPoney/go-tictactoe/board/bgrid"
    . "github.com/DrunkenPoney/go-tictactoe/grid/tile"
    . "github.com/DrunkenPoney/go-tictactoe/position"
    "github.com/DrunkenPoney/go-tictactoe/settings"
    "math"
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
    bestMove := p.GetPrediction().CurrentLayer().GetWinPos()
    if bestMove == INVALID {
        pred := p.GetPrediction().Predict()
        bestScore, mod := -math.MaxFloat64, 1.0
        if settings.REFERENCE_PLAYER != tt {
            mod = -1.0
        }
        for pos, score := range pred {
            nextLayer := p.GetPrediction().CurrentLayer().Next(pos)
            if nextLayer.GetWinPos() == INVALID {
                score *= mod
                if score > bestScore {
                    bestMove, bestScore = pos, score
                }
            }
        }
        if bestMove == INVALID {
            empty := p.prediction.CurrentLayer().Grid().EmptyTiles()
            bestMove = empty[int(rand.Float64()*float64(len(empty)))].Position
        }
    }
    return bestMove
}

func (p *process) GetPrediction() Prediction {
    return p.prediction
}
