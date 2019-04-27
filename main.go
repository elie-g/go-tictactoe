package main

import (
	"fmt"
	"github.com/DrunkenPoney/go-tictactoe/events"
	g "github.com/DrunkenPoney/go-tictactoe/game"
	"github.com/DrunkenPoney/go-tictactoe/game/player"
	. "github.com/DrunkenPoney/go-tictactoe/grids"
	"github.com/DrunkenPoney/go-tictactoe/grid/tile"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	. "image/color"
	"log"
)

// var grid Grid
var activeTile tile.Tile
var game g.Game

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	if activeTile != nil {
		activeTile.SetActive(false)
	}

	if activeTile = game.GetBoard().GetTileUnderCursor(); activeTile != nil {
		activeTile.SetActive(true)
	}

	game.Draw(screen)

	_ = ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %f, TPS: %f", ebiten.CurrentFPS(), ebiten.CurrentTPS()))

	return nil
}

func clickListener() {
	for {
		select {
		case <-events.MouseClick():
			println("Mouse clicked!")
		default:
		}
	}
}

func main() {
	// go clickListener()
	game = g.NewGame(player.NewPlayer("Joueur 1"), player.NewPlayer("Joueur 2"), NewGrids(3, 3, White, 10))
	// grid = NewGrid(3, 3, White, 10)

	// grid.GetTileAt(1, 1).SetValue(tile.X)
	// grid.GetTileAt(0, 0).SetValue(tile.O)
	// grid.GetTileAt(1, 0).SetValue(tile.X)
	// grid.GetTileAt(2, 2).SetValue(tile.O)
	// grid.GetTileAt(2, 0).SetValue(tile.X)
	// grid.GetTileAt(2, 1).SetValue(tile.O)

	if err := ebiten.Run(update, 800, 800, 1, "Go Tic-Tac-Toe"); err != nil {
		log.Fatalln(err)
	}
}
