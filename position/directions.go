package position

func (p Position) Up(i int) Position {
    if p == INVALID { return INVALID }
    x, y := p.GetXY()
    return PositionAt(x, y-i)
}

func (p Position) UpRight(i int) Position {
    if p == INVALID { return INVALID }
    x, y := p.GetXY()
    return PositionAt(x+i, y-i)
}

func (p Position) Right(i int) Position {
    if p == INVALID { return INVALID }
    x, y := p.GetXY()
    return PositionAt(x+i, y)
}

func (p Position) DownRight(i int) Position {
    if p == INVALID { return INVALID }
    x, y := p.GetXY()
    return PositionAt(x+i, y+i)
}

func (p Position) Down(i int) Position {
    if p == INVALID { return INVALID }
    x, y := p.GetXY()
    return PositionAt(x, y+i)
}

func (p Position) DownLeft(i int) Position {
    if p == INVALID { return INVALID }
    x, y := p.GetXY()
    return PositionAt(x-i, y+i)
}

func (p Position) Left(i int) Position {
    if p == INVALID { return INVALID }
    x, y := p.GetXY()
    return PositionAt(x-i, y)
}

func (p Position) UpLeft(i int) Position {
    if p == INVALID { return INVALID }
    x, y := p.GetXY()
    return PositionAt(x-i, y-i)
}