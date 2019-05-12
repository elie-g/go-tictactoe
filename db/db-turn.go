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
}

type dbTurn struct {
    id      int64
    fetched bool
    no      int64
    game    DBGame
    player  DBPlayer
    gridPos int
    subPos  int
}

// Private
func (dbt *dbTurn) fetch() {
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
    } else {
        panic("FETCH FAILED! (no data)")
    }
}

func (dbt *dbTurn) GetID() int64 {
    return dbt.id
}

func (dbt *dbTurn) GetNumber() int64 {
    if !dbt.fetched {
        dbt.fetch()
    }
    return dbt.no
}

func (dbt *dbTurn) GetGame() DBGame {
    if !dbt.fetched {
        dbt.fetch()
    }
    return dbt.game
}

func (dbt *dbTurn) GetPlayer() DBPlayer {
    if !dbt.fetched {
        dbt.fetch()
    }
    return dbt.player
}

func (dbt *dbTurn) GetGridPos() int {
    if !dbt.fetched {
        dbt.fetch()
    }
    return dbt.gridPos
}

func (dbt *dbTurn) GetSubGridPos() int {
    if !dbt.fetched {
        dbt.fetch()
    }
    return dbt.subPos
}
