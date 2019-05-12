package db

import (
    . "github.com/DrunkenPoney/go-tictactoe/internal"
)

type DBPlayer interface {
    GetID() int64
    GetName() string
    SetName(name string)
}

type dbPlayer struct {
    id      int64
    fetched bool
    name    string
}

// Private
func (dbp *dbPlayer) fetch() {
    rows, err := db.Query("SELECT nom FROM joueur WHERE id = ?", dbp.id)
    CheckError(err)
    defer rows.Close()
    if rows.Next() {
        cols, err := rows.Columns()
        CheckError(err)
        dbp.name = cols[0]
        dbp.fetched = true
    } else {
        panic("FETCH FAILED! (no data)")
    }
}

func (dbp *dbPlayer) GetID() int64 {
    return dbp.id
}

func (dbp *dbPlayer) GetName() string {
    if !dbp.fetched {
        dbp.fetch()
    }
    return dbp.name
}

func (dbp *dbPlayer) SetName(name string) {
    _, err := db.Exec("UPDATE joueur SET nom = ? WHERE id = ?", name, dbp.id)
    CheckError(err)
}
