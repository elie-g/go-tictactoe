package player

type Player interface {
    GetName() string
    SetName(name string)
    
    IsCurrent() bool
    SetCurrent(current bool)
    
    GetPoints() int
    SetPoints(points int)
    IncrementPoints() int
}

type player struct {
    name    string
    current bool
    points  int
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

func NewPlayer(name string) Player {
    return &player{name, false, 0}
}
