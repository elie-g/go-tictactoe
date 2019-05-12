package db

import (
    . "github.com/DrunkenPoney/go-tictactoe/internal"
    . "github.com/DrunkenPoney/go-tictactoe/position"
    "strconv"
)

func (d *database) CreateGame(player1 DBPlayer) DBGame {
    res, err := db.Exec(`INSERT INTO partie(id_gagnant, id_joueur1, id_joueur2) VALUES (NULL, ?, NULL)`, player1.GetID())
    CheckError(err)
    id, err := res.LastInsertId()
    CheckError(err)
    return &dbGame{id: id, db: d}
}

func (d *database) CreatePlayer(name string) DBPlayer {
    res, err := db.Exec(`INSERT INTO joueur(nom) VALUES (?)`, name)
    CheckError(err)
    id, err := res.LastInsertId()
    CheckError(err)
    return &dbPlayer{id: id, db: d}
}

func (d *database) CreateTurn(player DBPlayer, game DBGame, gridPos Position, subPos Position) DBTurn {
    var noGame int
    rows, err := db.Query(`SELECT no_coup FROM coup c WHERE id_partie = ? ORDER BY no_coup DESC LIMIT 1`, game.GetID())
    CheckError(err)
    defer rows.Close()
    if rows.Next() {
        cols, err := rows.Columns()
        CheckError(err)
        noGame, err := strconv.Atoi(cols[0])
        CheckError(err)
        noGame++
    }
    res, err := db.Exec(`INSERT INTO coup(no_coup, id_partie, id_joueur, cadrant, position) VALUES (?, ?, ?, ?, ?)`,
        noGame, game.GetID(), player.GetID(), gridPos, subPos)
    CheckError(err)
    id, err := res.LastInsertId()
    CheckError(err)
    return &dbTurn{id: id, db: d}
}
