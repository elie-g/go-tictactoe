package board

import (
    . "github.com/DrunkenPoney/go-tictactoe/grid/tile"
    "github.com/DrunkenPoney/go-tictactoe/position"
    "github.com/hajimehoshi/ebiten"
)

func (g *board) GetTileUnderCursor() *Tile {
    if g.screen == nil {
        return nil
    }
    
    cursorX, cursorY := ebiten.CursorPosition()
    
    if cursorX >= 0 && cursorY >= 0 {
        width, height := g.screen.Size()
        width = width / 3 // 3 = nb de sous-grilles
        height = height / 3
        
        pos := position.PositionAt(cursorX/width, cursorY/height)
        // Si curseur est dans la sous-grille actuelle
        if pos == g.GetCurrentPos() {
            x, y := pos.GetXY()
            x, y = (width*x)/(width/3), (height*y)/(height/3)
            
        }
        
        if cursorX <= width && cursorY <= height {
            col, row := cursorX/(width/g.columns), cursorY/(height/g.rows)
            return g.TileAt(col, row)
        }
        
    }
    
    return nil
}
