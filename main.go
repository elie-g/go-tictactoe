package main

import (
    "github.com/DrunkenPoney/go-tictactoe/gui"
    "github.com/DrunkenPoney/go-tictactoe/internal"
    _ "github.com/go-sql-driver/mysql"
    "github.com/hajimehoshi/ebiten"
    "log"
)

var layout gui.Layout

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
    
    layout = gui.NewLayout("Joueur 1", "Joueur 2")
    if err := ebiten.Run(layout.Update, int(internal.ScaleWidth(800)), int(internal.ScaleHeight(800)), 1, "Go Tic-Tac-Toe"); err != nil {
        log.Fatalln(err)
    }
}
