package db

import (
    . "github.com/DrunkenPoney/go-tictactoe/internal"
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
        row := db.QueryRow("SELECT no_coup, cadrant, position, id_partie, id_joueur FROM coup where id = ?", dbt.id)
        var idGame, idPlayer int64
        CheckError(row.Scan(&dbt.no, &dbt.gridPos, &dbt.subPos, &idGame, &idPlayer))
        dbt.game = &dbGame{id: idGame, db: dbt.db}
        dbt.player = &dbPlayer{id: idPlayer, db: dbt.db}
        dbt.fetched = true
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
