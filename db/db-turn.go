package db

import (
    . "github.com/DrunkenPoney/go-tictactoe/internal"
    "strconv"
)

type DBTurn interface {
    GetID() int64
    GetNumber() int64
    GetGame() DBGame
    GetPlayer() DBPlayer
    GetGridPos() int
    GetSubGridPos() int
    Reload()
}

type dbTurn struct {
    id      int64
    db      *database
    fetched bool
    no      int64
    game    DBGame
    player  DBPlayer
    gridPos int
    subPos  int
}

// Private
func (dbt *dbTurn) fetch(force bool) {
    if !force && dbt.db.turns[dbt.id] != nil {
        obj := dbt.db.turns[dbt.id]
        dbt.no = obj.GetNumber()
        dbt.game = obj.GetGame()
        dbt.player = obj.GetPlayer()
        dbt.gridPos = obj.GetGridPos()
        dbt.subPos = obj.GetSubGridPos()
    } else {
        rows, err := db.Query("SELECT no_coup, cadrant, position, id_partie, id_joueur FROM coup where id = ?", dbt.id)
        CheckError(err)
        defer rows.Close()
        if rows.Next() {
            cols, err := rows.Columns()
            CheckError(err)
            
            dbt.no, err = strconv.ParseInt(cols[0], 10, 64)
            CheckError(err)
            
            dbt.gridPos, err = strconv.Atoi(cols[1])
            CheckError(err)
            
            dbt.subPos, err = strconv.Atoi(cols[2])
            CheckError(err)
            
            id, err := strconv.ParseInt(cols[3], 10, 64)
            CheckError(err)
            dbt.player = &dbPlayer{id: id}
            
            id, err = strconv.ParseInt(cols[4], 10, 64)
            CheckError(err)
            dbt.game = &dbGame{id: id}
            dbt.fetched = true
            dbt.db.turns[dbt.id] = dbt
        } else {
            panic("FETCH FAILED! (no data)")
        }
    }
}

func (dbt *dbTurn) GetID() int64 {
    return dbt.id
}

func (dbt *dbTurn) GetNumber() int64 {
    if !dbt.fetched {
        dbt.fetch(false)
    }
    return dbt.no
}

func (dbt *dbTurn) GetGame() DBGame {
    if !dbt.fetched {
        dbt.fetch(false)
    }
    return dbt.game
}

func (dbt *dbTurn) GetPlayer() DBPlayer {
    if !dbt.fetched {
        dbt.fetch(false)
    }
    return dbt.player
}

func (dbt *dbTurn) GetGridPos() int {
    if !dbt.fetched {
        dbt.fetch(false)
    }
    return dbt.gridPos
}

func (dbt *dbTurn) GetSubGridPos() int {
    if !dbt.fetched {
        dbt.fetch(false)
    }
    return dbt.subPos
}

func (dbt *dbTurn) Reload() {
    dbt.fetch(true)
}
