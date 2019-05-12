package db

import (
    . "github.com/DrunkenPoney/go-tictactoe/internal"
    "strconv"
)

func (d *database) LastCreatedGame() DBGame {
    var game *dbGame
    rows, err := db.Query("SELECT id FROM partie ORDER BY date_creation DESC LIMIT 1")
    CheckError(err)
    defer rows.Close()
    if rows.Next() {
        game = &dbGame{}
        cols, err := rows.Columns()
        CheckError(err)
        game.id, err = strconv.ParseInt(cols[0], 10, 64)
        CheckError(err)
    }
    return game
}
