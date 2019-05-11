package prediction

import (
    "fmt"
    . "github.com/DrunkenPoney/go-tictactoe/ai/prediction/layer"
    . "github.com/DrunkenPoney/go-tictactoe/board"
    "github.com/DrunkenPoney/go-tictactoe/board/bgrid"
    "github.com/DrunkenPoney/go-tictactoe/grid/tile"
    . "github.com/DrunkenPoney/go-tictactoe/position"
    "github.com/DrunkenPoney/go-tictactoe/settings"
    "sync"
)

var mut sync.Mutex

type Prediction interface {
    CurrentLayer() PredictionLayer
    Predict() map[Position]float64
    Next(pos Position)
    SetMaxDepth(maxDepth int)
    MaxDepth() int
}

func NewPrediction(player tile.TileType, board bgrid.BoardGrid) Prediction {
    return &prediction{
        layer:    NewLayer(DefaultPosition, player, board),
        maxDepth: settings.DEFAULT_PREDICTION_DEPTH,
        mut:      &mut}
}

type prediction struct {
    layer    PredictionLayer
    maxDepth int
    mut      *sync.Mutex
}

func (pred *prediction) CurrentLayer() PredictionLayer {
    return pred.layer
}

func (pred *prediction) Next(pos Position) {
    pred.layer = pred.layer.Next(pos)
}

func (pred *prediction) MaxDepth() int {
    return pred.maxDepth
}

func (pred *prediction) SetMaxDepth(maxDepth int) {
    pred.maxDepth = maxDepth
}

func (pred *prediction) Predict() map[Position]float64 {
    pred.mut.Lock()
    println(pred.CurrentLayer().Board().String())
    predictions := make(map[Position]float64)
    var wg sync.WaitGroup
    var mut sync.Mutex
    for pos, layer := range pred.CurrentLayer().Layers() {
        wg.Add(1)
        go func(pos Position, layer PredictionLayer, wg *sync.WaitGroup, mut *sync.Mutex) {
            defer wg.Done()
            res := (layer.GetScore() * 3) + pred.calcLayer(layer)
            mut.Lock()
            predictions[pos] = res
            mut.Unlock()
        }(pos, layer, &wg, &mut)
    }
    wg.Wait()
    pred.mut.Unlock()
    str := "Prediction: "
    for pos, score := range predictions {
        x, y := pos.GetXY()
        str += fmt.Sprintf("\n\t> Pos %d (%d,%d) => %f", pos, x, y, score)
    }
    fmt.Println(str)
    return predictions
}

func (pred *prediction) calcLayer(layer PredictionLayer) float64 {
    score := 0.0
    if layer.Depth()-pred.layer.Depth() < pred.MaxDepth() {
        var wg sync.WaitGroup
        var mut sync.Mutex
        count := 0.0
        for _, subLayer := range layer.Layers() {
            wg.Add(1)
            count++
            go func(subLayer PredictionLayer, wg *sync.WaitGroup, mut *sync.Mutex) {
                defer wg.Done()
                scr := (subLayer.GetScore() * 3) + pred.calcLayer(subLayer)
                mut.Lock()
                score += scr
                mut.Unlock()
            }(subLayer, &wg, &mut)
        }
        wg.Wait()
        score /= count
    }
    return score
}
