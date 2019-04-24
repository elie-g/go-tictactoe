package grid

import (
	. "github.com/DrunkenPoney/go-tictactoe/grid/tile"
	"github.com/hajimehoshi/ebiten"
)

func (g *grid) GetTileUnderCursor() Tile {
	if g.img == nil {
		return nil
	}

	cursorX, cursorY := ebiten.CursorPosition()
	gridPos := g.img.Bounds().Min
	cursorX -= gridPos.X
	cursorY -= gridPos.Y

	if cursorX >= 0 && cursorY >= 0 {
		width, height := g.img.Bounds().Dx(), g.img.Bounds().Dy()

		if cursorX <= width && cursorY <= height {
			col, row := cursorX/(width/g.columns), cursorY/(height/g.rows)
			// fmt.Printf("col: \x1b[33m%d\x1b[m, row: \x1b[33m%d\x1b[m\n", col, row)
			return g.GetTileAt(col, row)
		}

	}

	return nil
}
