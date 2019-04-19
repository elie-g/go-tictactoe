package player

type Player interface {
    GetName() string
    SetName(name string) Player
    
    GetScore() int
    SetScore(score int)
}

type player struct {
    key string
    name string
}
