package main

import (
    "github.com/DrunkenPoney/go-tictactoe/db"
    "github.com/DrunkenPoney/go-tictactoe/gui"
    "github.com/DrunkenPoney/go-tictactoe/internal"
    _ "github.com/go-sql-driver/mysql"
    "github.com/hajimehoshi/ebiten"
    "log"
)

var layout gui.Layout

func main() {
    defer db.CloseConnection()
    
    layout = gui.NewLayout("Joueur 1", "Joueur 2")
    if err := ebiten.Run(layout.Update, int(internal.ScaleWidth(800)), int(internal.ScaleHeight(800)), 1, "Go Tic-Tac-Toe"); err != nil {
        log.Fatalln(err)
    }
}
