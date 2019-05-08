package main

import (
	"fmt"
	"github.com/DrunkenPoney/go-tictactoe/board"
	"github.com/DrunkenPoney/go-tictactoe/events"
	g "github.com/DrunkenPoney/go-tictactoe/game"
	"github.com/DrunkenPoney/go-tictactoe/game/player"
	"github.com/DrunkenPoney/go-tictactoe/grid/tile"
	"github.com/DrunkenPoney/go-tictactoe/internal"
	. "github.com/DrunkenPoney/go-tictactoe/position"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

var activeTile *tile.Tile
var activePos Position
var game g.Game

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	if activeTile != nil {
		activeTile.Active = false
		game.GetBoard().DrawTile(activeTile, activePos)
	}

	if activeTile, activePos = game.GetBoard().GetTileUnderCursor(); activeTile != nil {
		activeTile.Active = true
		game.GetBoard().DrawTile(activeTile, activePos)
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


	// Open up our database connection.
	// db, err := sql.Open("mysql", "tp3-veille:tp3@tcp(159.203.13.220:3306)/tp3-veille")

	// if there is an error opening the connection, handle it
	// if err != nil {
	// 	panic(err.Error())
	// }
	
	// defer the close till after the main function has finished
	// executing
	// defer db.Close()


	// insert, err := db.Query("INSERT INTO partie VALUES ( 1, 1)")

	// if there is an error inserting, handle it
	// if err != nil {
	// 	panic(err.Error())
	// }
	// be careful deferring Queries if you are using transactions
	// defer insert.Close()




	// go clickListener()
	game = g.NewGame(player.NewPlayer("Joueur 1"), player.NewPlayer("Joueur 2"), board.NewBoard())
	// grid = NewGrid(3, 3, White, 10)

	// grid.GetTileAt(1, 1).SetValue(tile.X)
	// grid.GetTileAt(0, 0).SetValue(tile.O)
	// grid.GetTileAt(1, 0).SetValue(tile.X)
	// grid.GetTileAt(2, 2).SetValue(tile.O)
	// grid.GetTileAt(2, 0).SetValue(tile.X)
	// grid.GetTileAt(2, 1).SetValue(tile.O)

	if err := ebiten.Run(update, int(internal.ScaleWidth(800)), int(internal.ScaleHeight(800)), 1, "Go Tic-Tac-Toe"); err != nil {
		log.Fatalln(err)
	}
}
