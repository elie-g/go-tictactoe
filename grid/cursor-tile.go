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
	// gridPos := g.img.Bounds().Min
	// cursorX -= gridPos.X
	// cursorY -= gridPos.Y

	if cursorX >= 0 && cursorY >= 0 {
		width, height := g.img.Bounds().Dx(), g.img.Bounds().Dy()

		if cursorX <= width && cursorY <= height {
			col, row := cursorX/(width/g.columns), cursorY/(height/g.rows)
			return g.GetTileAt(col, row)
		}

	}

	return nil
}
