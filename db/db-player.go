package db

import (
    . "github.com/DrunkenPoney/go-tictactoe/internal"
)

type DBPlayer interface {
    GetID() int64
    GetName() string
    SetName(name string)
    Reload()
}

type dbPlayer struct {
    id      int64
    db      *database
    fetched bool
    name    string
}

// Private
func (dbp *dbPlayer) fetch(force bool) {
    if !force && dbp.db.players[dbp.id] != nil {
        dbp.name = dbp.db.players[dbp.id].GetName()
    } else {
        rows, err := db.Query("SELECT nom FROM joueur WHERE id = ?", dbp.id)
        CheckError(err)
        defer rows.Close()
        if rows.Next() {
            cols, err := rows.Columns()
            CheckError(err)
            dbp.name = cols[0]
            dbp.fetched = true
            dbp.db.players[dbp.id] = dbp
        } else {
            panic("FETCH FAILED! (no data)")
        }
    }
}

func (dbp *dbPlayer) GetID() int64 {
    return dbp.id
}

func (dbp *dbPlayer) GetName() string {
    if !dbp.fetched {
        dbp.fetch(false)
    }
    return dbp.name
}

func (dbp *dbPlayer) SetName(name string) {
    _, err := db.Exec("UPDATE joueur SET nom = ? WHERE id = ?", name, dbp.id)
    CheckError(err)
}

func (dbp *dbPlayer) Reload() {
    dbp.fetch(true)
}
