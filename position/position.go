package position

type Position int

/* Les nombres sont attribués selon la formule: 1 + x + 3y
 * [1][2][3]
 * [4][5][6]
 * [7][8][9]
 */
//noinspection GoUnusedConst,GoSnakeCaseUsage
const (
    TOP_LEFT   Position = 1 // 1 + 0 + 3(0)
    TOP_CENTER Position = 2 // 1 + 1 + 3(0)
    TOP_RIGHT  Position = 3 // 1 + 2 + 3(0)
    
    MIDDLE_LEFT   Position = 4 // 1 + 0 + 3(1)
    MIDDLE_CENTER Position = 5 // 1 + 1 + 3(1)
    MIDDLE_RIGHT  Position = 6 // 1 + 2 + 3(1)
    
    BOTTOM_LEFT   Position = 7 // 1 + 0 + 3(2)
    BOTTOM_CENTER Position = 8 // 1 + 1 + 3(2)
    BOTTOM_RIGHT  Position = 9 // 1 + 2 + 3(2)
    
    INVALID Position = 0
)

func (p Position) GetXY() (int, int) {
    x := (p - 1) % 3
    y := (p - 1 - x) / 3
    return int(x), int(y)
}

func (p Position) GetXYFloat() (float64, float64) {
    x, y := p.GetXY()
    return float64(x), float64(y)
}

func (p Position) Opposite() Position {
    pos := p
    if p != INVALID {
        x, y := p.GetXY()
        pos = PositionAt(2-x, 2-y)
    }
    return pos
}

func PositionAt(x int, y int) Position {
    pos := 1 + x + (3 * y)
    if pos < 0 || pos > 9 {
        pos = 0
    }
    return Position(pos)
}
