package grid

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/peterhellberg/gfx"
	"image/color"
	"math"
)

func (g *grid) DrawGrid(screen *ebiten.Image) {
	width, height := float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy())
	colWidth, rowHeight := width/float64(g.columns), height/float64(g.rows)
	img, _ := ebiten.NewImage(screen.Bounds().Dx(), screen.Bounds().Dy(), ebiten.FilterDefault)

	_ = img.Fill(color.Black)
	for i, col := range g.GetTiles() {
		x := float64(i) * colWidth
		for j, tile := range col {
			y := float64(j) * rowHeight
			tileImg := tile.GetImage()
			if tileImg != nil {
				opts := &ebiten.DrawImageOptions{}
				scale := math.Min(colWidth, rowHeight) / float64(tileImg.Bounds().Dx())
				opts.GeoM.Scale(scale, scale)
				opts.GeoM.Translate(x, y)
				_ = img.DrawImage(tileImg, opts)
			}
		}
	}

	// Grid lines (#)
	im := gfx.NewImage(int(width), int(height))
	for x := colWidth; x < (width - (colWidth / 2)); x += colWidth {
		gfx.DrawLine(im, gfx.V(x, 0), gfx.V(x, height), g.strokeWidth, g.color)
	}
	for y := rowHeight; y < (height - (rowHeight / 2)); y += rowHeight {
		gfx.DrawLine(im, gfx.V(0, y), gfx.V(width, y), g.strokeWidth, g.color)
	}
	newImg, _ := ebiten.NewImageFromImage(im.SubImage(im.Rect), ebiten.FilterDefault)
	_ = img.DrawImage(newImg, &ebiten.DrawImageOptions{})
	_ = screen.DrawImage(img, &ebiten.DrawImageOptions{})
	g.img = img
}
