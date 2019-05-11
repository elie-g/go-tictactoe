package board

import (
    . "github.com/DrunkenPoney/go-tictactoe/grid/tile"
    . "github.com/DrunkenPoney/go-tictactoe/position"
    "github.com/hajimehoshi/ebiten"
    . "math"
)

func (b *board) GetTileUnderCursor() (*Tile, Position) {
    if b.screen != nil {
        ciX, ciY := ebiten.CursorPosition()
        cX, cY := float64(ciX), float64(ciY)
        wi, hi := b.screen.Size()
        w, h := float64(wi), float64(hi)
    
        if ciX >= 0 && ciY >= 0 && ciX <= wi && ciY <= hi {
            w, h = w/3, h/3 // 3 = nb de sous-grilles
            pos := PositionAt(int(cX/w), int(cY/h))
            
            x, y := int(Mod(cX, w)/(w/3)), int(Mod(cY, h)/(h/3))
            if subPos := PositionAt(x, y); subPos != INVALID && pos != INVALID {
                return b.grids[pos].At(subPos), pos
            }
        }
    }
    
    return nil, INVALID
}
