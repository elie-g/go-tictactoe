package db

import (
    "database/sql"
    . "github.com/DrunkenPoney/go-tictactoe/internal"
    . "github.com/DrunkenPoney/go-tictactoe/position"
    _ "github.com/go-sql-driver/mysql"
    "runtime"
)

var db *sql.DB

type Database interface {
    CreateGame(player1 DBPlayer) DBGame
    CreatePlayer(name string) DBPlayer
    CreateTurn(player DBPlayer, game DBGame, gridPos Position, subPos Position) DBTurn
    FindGame(id int64, force bool) DBGame
    FindPlayer(id int64, force bool) DBPlayer
    FindTurn(id int64, force bool) DBTurn
    LastCreatedGame() DBGame
    ClearMemory()
    DB() *sql.DB
}

func CloseConnection() {
    if db != nil {
        CheckError(db.Close())
    }
}

func NewDatabase() Database {
    var err error
    if db == nil {
        db, err = sql.Open("mysql", "tp3-veille:tp3@tcp(159.203.13.220:3306)/tp3-veille")
        CheckError(err)
    }
    err = db.Ping()
    CheckError(err)
    return &database{
        games:   make(map[int64]DBGame),
        turns:   make(map[int64]DBTurn),
        players: make(map[int64]DBPlayer),
    }
}

type database struct {
    games   map[int64]DBGame
    turns   map[int64]DBTurn
    players map[int64]DBPlayer
}

func (d *database) ClearMemory() {
    d.games = make(map[int64]DBGame)
    d.turns = make(map[int64]DBTurn)
    d.players = make(map[int64]DBPlayer)
    runtime.GC()
}

func (d *database) DB() *sql.DB {
    return db
}
