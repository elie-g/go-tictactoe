package db

import (
    . "github.com/DrunkenPoney/go-tictactoe/internal"
    "strconv"
    "time"
)

type DBGame interface {
    GetID() int64
    GetDate() *time.Time
    GetPlayer1() DBPlayer
    GetPlayer2() DBPlayer
    GetWinner() DBPlayer
    SetWinner(winner DBPlayer)
}

type dbGame struct {
    id      int64
    fetched bool
    date    *time.Time
    player1 DBPlayer
    player2 DBPlayer
    winner  DBPlayer
}

func (dbg *dbGame) fetch() {
    rows, err := db.Query("SELECT id_joueur1, id_joueur2, id_gagnant FROM partie WHERE id = ?", dbg.id)
    CheckError(err)
    defer rows.Close()
    if rows.Next() {
        cols, err := rows.Columns()
        CheckError(err)
        
        id, err := strconv.ParseInt(cols[0], 10, 64)
        CheckError(err)
        dbg.player1 = &dbPlayer{id: id}
        
        id, err = strconv.ParseInt(cols[1], 10, 64)
        CheckError(err)
        dbg.player2 = &dbPlayer{id: id}
        
        if len(cols[2]) > 0 {
            id, err = strconv.ParseInt(cols[2], 10, 64)
            CheckError(err)
            dbg.winner = &dbPlayer{id: id}
        }
        
        dbg.fetched = true
    } else {
        panic("FETCH FAILED! (no data)")
    }
}

func (dbg *dbGame) GetID() int64 {
    return dbg.id
}

func (dbg *dbGame) GetDate() *time.Time {
    // time.Parse() TODO
    return nil
}

func (dbg *dbGame) GetPlayer1() DBPlayer {
    if !dbg.fetched {
        dbg.fetch()
    }
    return dbg.player1
}

func (dbg *dbGame) GetPlayer2() DBPlayer {
    if !dbg.fetched {
        dbg.fetch()
    }
    return dbg.player2
}

func (dbg *dbGame) GetWinner() DBPlayer {
    if !dbg.fetched {
        dbg.fetch()
    }
    return dbg.winner
}

func (dbg *dbGame) SetWinner(winner DBPlayer) {
    _, err := db.Exec("UPDATE partie SET id_gagnant = ? WHERE id = ?", winner.GetID(), dbg.id)
    CheckError(err)
}
