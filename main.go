package main

import (
    "fmt"
    . "github.com/DrunkenPoney/go-tictactoe/grid"
    "github.com/DrunkenPoney/go-tictactoe/grid/tile"
    "github.com/hajimehoshi/ebiten/ebitenutil"
    
    // "github.com/go-vgo/robotgo"
    "github.com/hajimehoshi/ebiten"
    . "image/color"
    "log"
)

var grid Grid
var activeTile tile.Tile

func update(screen *ebiten.Image) error {
    if ebiten.IsDrawingSkipped() {
        return nil
    }
    
    if activeTile != nil {
        activeTile.SetBgColor(Black).SetColor(White)
    }
    
    if activeTile = grid.GetTileUnderCursor(); activeTile != nil {
        activeTile.SetBgColor(White).SetColor(Black)
    }
    
    // _ = screen.Clear()
    grid.DrawGrid(screen)
    
    _ = ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %f, TPS: %f", ebiten.CurrentFPS(), ebiten.CurrentTPS()))
    // text.Draw(screen, fmt.Sprintf("FPS: %d, TPS: %d", ebiten.CurrentFPS(), ebiten.CurrentTPS()), nil, 50, 50, RGBA{R: 0xFF, G: 0x99, B: 0x00, A: 0xFF})
    
    return nil
}

func main() {
    grid = NewGrid(3, 3, White, 10)
    
    grid.GetTileAt(1, 1).SetValue(tile.X)
    grid.GetTileAt(0, 0).SetValue(tile.O)
    grid.GetTileAt(1, 0).SetValue(tile.X)
    grid.GetTileAt(2, 2).SetValue(tile.O)
    grid.GetTileAt(2, 0).SetValue(tile.X)
    grid.GetTileAt(2, 1).SetValue(tile.O)
    
    if err := ebiten.Run(update, 800 * 2, 800 * 2, 1, "Go Tic-Tac-Toe"); err != nil {
        log.Fatalln(err)
    }
}
