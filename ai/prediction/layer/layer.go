package layer

import (
    "fmt"
    . "github.com/DrunkenPoney/go-tictactoe/ai/path"
    . "github.com/DrunkenPoney/go-tictactoe/board/bgrid"
    . "github.com/DrunkenPoney/go-tictactoe/grid"
    . "github.com/DrunkenPoney/go-tictactoe/grid/tile"
    . "github.com/DrunkenPoney/go-tictactoe/position"
    "github.com/DrunkenPoney/go-tictactoe/settings"
)

type PredictionLayer interface {
    Layers() map[Position]PredictionLayer
    Next(pos Position) PredictionLayer
    Depth() int
    
    GetScore() float64
    GetWinPossibilities() map[Position]bool
    GetWinPos() Position
    CurrentPlayer() TileType
    
    GridPosition() Position
    Board() BoardGrid
    Grid() TileGrid
}

func NewLayer(position Position, player TileType, board BoardGrid) PredictionLayer {
    return &layer{
        position: position,
        player:   player,
        depth:    0,
        board:    board,
        path:     Path{},
    }
}

type layer struct {
    position Position
    player   TileType
    depth    int
    path     Path
    board    BoardGrid
}

func (l *layer) CurrentPlayer() TileType {
    return l.player
}

func (l *layer) GridPosition() Position {
    return l.position
}

func (l *layer) Depth() int {
    return l.depth
}

func (l *layer) Board() BoardGrid {
    return l.board
}

func (l *layer) Grid() TileGrid {
    return l.Board().Get(l.GridPosition())
}

func (l *layer) Layers() map[Position]PredictionLayer {
    layers := make(map[Position]PredictionLayer)
    for _, tile := range l.Grid().EmptyTiles() {
        layers[tile.Position] = l.Next(tile.Position)
    }
    return layers
}

func (l *layer) Next(pos Position) PredictionLayer {
    if l.board.Get(l.GridPosition()).At(pos).Value == EMPTY {
        board := l.Board().Clone()
        board.Get(l.position).At(pos).Value = l.CurrentPlayer()
        return &layer{
            position: pos,
            player:   l.player.Opponent(),
            depth:    l.depth + 1,
            path:     l.path.Next(pos),
            board:    board,
        }
    }
    return nil
}

func (l *layer) GetWinPossibilities() map[Position]bool {
    emptyTiles := l.Grid().EmptyTiles()
    possibilities := make(map[Position]bool)
    for _, tile := range emptyTiles {
        clone := l.Grid().Clone()
        clone.At(tile.Position).Value = l.CurrentPlayer()
        possibilities[tile.Position] = clone.GetWinningTiles()[0] != nil
    }
    return possibilities
}

func (l *layer) GetWinPos() Position {
    for pos, win := range l.GetWinPossibilities() {
        if win {
            fmt.Printf("\t > Pos %d => CAN WIN!!!\tPlayer: %v // Depth: %v\n", pos, l.CurrentPlayer(), l.Depth())
            return pos
        }
    }
    return INVALID
}

func (l *layer) GetScore() float64 {
    score := 0.0
    possibilities := l.GetWinPossibilities()
    for _, winning := range possibilities {
        if winning {
            if l.CurrentPlayer() == settings.REFERENCE_PLAYER {
                score += 10
            } else {
                score += -10
            }
        }
    }
    return score / float64(len(possibilities))
}
