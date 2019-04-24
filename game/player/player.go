package player

type Player interface {
	GetName() string
	SetName(name string) Player

	IsCurrent() bool
	SetCurrent(current bool) Player
}

type player struct {
	name    string
	current bool
}

func (p *player) IsCurrent() bool {
	return p.current
}

func (p *player) SetCurrent(current bool) Player {
	p.current = current
	return p
}

func (p *player) GetName() string {
	return p.name
}

func (p *player) SetName(name string) Player {
	p.name = name
	return p
}

func NewPlayer(name string) Player {
	return &player{name, false}
}
