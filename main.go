package main

import (
	"fmt"
	. "github.com/DrunkenPoney/go-tictactoe/grid"
	"github.com/DrunkenPoney/go-tictactoe/grid/tile"
	"github.com/hajimehoshi/ebiten/ebitenutil"

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
		activeTile.SetActive(false)
	}

	if activeTile = grid.GetTileUnderCursor(); activeTile != nil {
		activeTile.SetActive(true)
	}

	// if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
	//
	// }

	// _ = screen.Clear()
	grid.DrawGrid(screen)

	_ = ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %f, TPS: %f", ebiten.CurrentFPS(), ebiten.CurrentTPS()))

	return nil
}

func main() {
	grid = NewGrid(3, 3, White, 10)

	grid.GetTileAt(1, 1).SetValue(tile.X)
	grid.GetTileAt(0, 0).SetValue(tile.O)
	grid.GetTileAt(1, 0).SetValue(tile.X)
	// grid.GetTileAt(2, 2).SetValue(tile.O)
	grid.GetTileAt(2, 0).SetValue(tile.X)
	// grid.GetTileAt(2, 1).SetValue(tile.O)

	if err := ebiten.Run(update, 800*2, 800*2, 1, "Go Tic-Tac-Toe"); err != nil {
		log.Fatalln(err)
	}
}
