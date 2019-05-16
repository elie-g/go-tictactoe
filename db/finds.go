package db

import (
    . "github.com/DrunkenPoney/go-tictactoe/internal"
    "strconv"
    "time"
)

func (d *database) FindGame(id int64, force bool) DBGame {
    if !force && d.games[id] != nil {
        return d.games[id]
    }
    var game *dbGame
    rows, err := db.Query("SELECT id, id_gagnant, id_joueur1, id_joueur2 FROM partie WHERE id = ?", id)
    CheckError(err)
    defer rows.Close()
    if rows.Next() {
        game = &dbGame{}
        cols, err := rows.Columns()
        CheckError(err)
        
        game.id, err = strconv.ParseInt(cols[0], 10, 64)
        CheckError(err)
        
        id, err := strconv.ParseInt(cols[1], 10, 64)
        CheckError(err)
        game.winner = &dbPlayer{id: id}
        
        id, err = strconv.ParseInt(cols[2], 10, 64)
        CheckError(err)
        game.player1 = &dbPlayer{id: id}
        
        id, err = strconv.ParseInt(cols[3], 10, 64)
        CheckError(err)
        game.player2 = &dbPlayer{id: id}
        
        t, err := strconv.ParseInt(cols[4], 10, 64)
        CheckError(err)
        game.date = time.Unix(t, 0)
        
        game.fetched = true
        game.db = d
        d.games[id] = game
    }
    return game
}

func (d *database) FindPlayer(id int64, force bool) DBPlayer {
    if !force && d.players[id] != nil {
        return d.players[id]
    }
    var player *dbPlayer
    rows, err := db.Query("SELECT id, nom FROM joueur WHERE id = ?", id)
    CheckError(err)
    defer rows.Close()
    if rows.NextResultSet() {
        player = &dbPlayer{}
        cols, err := rows.Columns()
        CheckError(err)
        
        player.id, err = strconv.ParseInt(cols[0], 10, 64)
        CheckError(err)
        player.name = cols[1]
        player.fetched = true
        player.db = d
        d.players[id] = player
    }
    return player
}

func (d *database) FindTurn(id int64, force bool) DBTurn {
    if !force && d.turns[id] != nil {
        return d.turns[id]
    }
    var turn *dbTurn
    rows, err := db.Query("SELECT id, no_coup, id_partie, id_joueur, cadrant, position FROM coup WHERE id = ?", id)
    CheckError(err)
    defer rows.Close()
    if rows.NextResultSet() {
        turn = &dbTurn{}
        cols, err := rows.Columns()
        CheckError(err)
        
        id, err := strconv.ParseInt(cols[0], 10, 64)
        CheckError(err)
        turn.id = id
        
        id, err = strconv.ParseInt(cols[0], 10, 64)
        CheckError(err)
        turn.no = id
        
        id, err = strconv.ParseInt(cols[0], 10, 64)
        CheckError(err)
        turn.game = &dbGame{id: id}
        
        id, err = strconv.ParseInt(cols[0], 10, 64)
        CheckError(err)
        turn.player = &dbPlayer{id: id}
        
        turn.gridPos, err = strconv.Atoi(cols[0])
        CheckError(err)
        
        turn.gridPos, err = strconv.Atoi(cols[0])
        CheckError(err)
        turn.fetched = true
        turn.db = d
        d.turns[id] = turn
    }
    return turn
}
