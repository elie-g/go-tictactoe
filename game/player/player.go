package player

import (
    "github.com/DrunkenPoney/go-tictactoe/db"
)

type Player interface {
    GetName() string
    SetName(name string)
    
    IsCurrent() bool
    SetCurrent(current bool)
    
    GetPoints() int
    SetPoints(points int)
    IncrementPoints() int
    
    // Online Functions
    IsRemote() bool
    SetRemote(remote bool)
    
    IsPlayer1() bool
    SetPlayer1(player1 bool)
    
    GetDBPlayer() db.DBPlayer
    SetDBPlayer(dbPlayer db.DBPlayer)
}

type player struct {
    name      string
    current   bool
    points    int
    remote    bool
    isPlayer1 bool
    dbPlayer  db.DBPlayer
}

func (p *player) IsCurrent() bool {
    return p.current
}

func (p *player) SetCurrent(current bool) {
    p.current = current
}

func (p *player) GetName() string {
    return p.name
}

func (p *player) SetName(name string) {
    p.name = name
}

func (p *player) GetPoints() int {
    return p.points
}

func (p *player) SetPoints(points int) {
    p.points = points
}

func (p *player) IncrementPoints() int {
    p.points++
    return p.points
}

func (p *player) IsRemote() bool {
    return p.remote
}

func (p *player) SetRemote(remote bool) {
    p.remote = remote
}

func (p *player) IsPlayer1() bool {
    return p.isPlayer1
}

func (p *player) SetPlayer1(player1 bool) {
    p.isPlayer1 = player1
}

func (p *player) GetDBPlayer() db.DBPlayer {
    return p.dbPlayer
}

func (p *player) SetDBPlayer(dbPlayer db.DBPlayer) {
    p.dbPlayer = dbPlayer
    if dbPlayer != nil {
        p.SetName(dbPlayer.GetName())
    }
}

func NewPlayer(name string) Player {
    return &player{name, false, 0, false, false, nil}
}
