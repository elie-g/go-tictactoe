package db

import (
    "fmt"
    . "github.com/DrunkenPoney/go-tictactoe/internal"
    "strconv"
)

// Retourne la dernière partie créée dont le joueur2 est null
func (d *database) LastCreatedGame() DBGame {
    rows, err := db.Query("SELECT id FROM partie WHERE id_joueur2 IS NULL ORDER BY date_creation DESC LIMIT 1")
    CheckError(err)
    defer rows.Close()
    
    if rows.NextResultSet() {
        cols, err := rows.Columns()
        CheckError(err)
        fmt.Println(cols[0])
        id, err := strconv.ParseInt(cols[0], 10, 64)
        CheckError(err)
        return &dbGame{id: id, db: d}
    }
    return &dbGame{id: 7, db: d}
}

func (d *database) LastTurn(game DBGame) DBTurn {
    row := db.QueryRow("SELECT id FROM coup WHERE id_partie = ? ORDER BY no_coup DESC LIMIT 1", game.GetID())
    var id int64
    CheckError(row.Scan(&id))
    return &dbTurn{id: id, db: d}
}
