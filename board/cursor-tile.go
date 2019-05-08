package board

import (
    . "github.com/DrunkenPoney/go-tictactoe/grid/tile"
    . "github.com/DrunkenPoney/go-tictactoe/position"
    "github.com/hajimehoshi/ebiten"
)

func (b *board) GetTileUnderCursor() (*Tile, Position) {
    if b.screen != nil {
        ciX, ciY := ebiten.CursorPosition()
        wi, hi := b.screen.Size()
        w, h := float64(wi), float64(hi)
    
        if ciX >= 0 && ciY >= 0 && ciX <= wi && ciY <= hi {
            w, h = w/3, h/3 // 3 = nb de sous-grilles
            wi, hi = int(w), int(h)
            pos := PositionAt(ciX/wi, ciY/hi)
    
            x, y := pos.GetXYFloat()
            x, y = (w*x)/(w/3), (h*y)/(h/3)
            if subPos := PositionAt(int(x), int(y)); subPos != INVALID {
                return b.grids[pos].At(subPos), pos
            }
        }
    }
    
    return nil, INVALID
}
