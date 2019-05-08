package position

type Position int

/* Les nombres sont attribués selon la formule: x + 3y
 * [0][1][2]
 * [3][4][5]
 * [6][7][8]
 */
//noinspection GoUnusedConst,GoSnakeCaseUsage
const (
    TOP_LEFT   Position = 0 // 0 + 3(0)
    TOP_CENTER Position = 1 // 1 + 3(0)
    TOP_RIGHT  Position = 2 // 2 + 3(0)
    
    MIDDLE_LEFT   Position = 3 // 0 + 3(1)
    MIDDLE_CENTER Position = 4 // 1 + 3(1)
    MIDDLE_RIGHT  Position = 5 // 2 + 3(1)
    
    BOTTOM_LEFT   Position = 6 // 0 + 3(2)
    BOTTOM_CENTER Position = 7 // 1 + 3(2)
    BOTTOM_RIGHT  Position = 8 // 2 + 3(2)
    
    INVALID Position = -1
)

func (p Position) GetXY() (int, int) {
    x := p % 3
    y := (p - x) / 3
    return int(x), int(y)
}

func (p Position) Opposite() Position {
    pos := p
    if p != INVALID {
        x, y := p.GetXY()
        pos = PositionAt(2 - x, 2 - y)
    }
    return pos
}

func PositionAt(x int, y int) Position {
    pos := x + (3 * y)
    if pos < 0 || pos > 8 {
        pos = -1
    }
    return Position(pos)
}
