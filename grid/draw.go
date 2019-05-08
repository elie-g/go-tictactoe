package grid

import (
    "github.com/DrunkenPoney/go-tictactoe/internal"
    . "github.com/DrunkenPoney/go-tictactoe/position"
    "github.com/hajimehoshi/ebiten"
    "github.com/peterhellberg/gfx"
    "image/color"
    "math"
)

var subGridImg *ebiten.Image

func (g TileGrid) Draw(screen *ebiten.Image, strokeWidth float64, strokeColor color.Color) *ebiten.Image {
    width, height := float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy())
    colWidth, rowHeight := width/float64(len(g)), height/float64(len(g[0]))
    img, _ := ebiten.NewImage(screen.Bounds().Dx(), screen.Bounds().Dy(), ebiten.FilterDefault)
    
    _ = screen.Clear()
    for i, col := range g {
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
    
    // Board lines (#)
    if subGridImg == nil {
        im := gfx.NewImage(int(width), int(height))
        for x := colWidth; x < (width - (colWidth / 2)); x += colWidth {
            gfx.DrawLine(im, gfx.V(x, 0), gfx.V(x, height), internal.ScaleWidth(strokeWidth), strokeColor)
        }
        for y := rowHeight; y < (height - (rowHeight / 2)); y += rowHeight {
            gfx.DrawLine(im, gfx.V(0, y), gfx.V(width, y), internal.ScaleHeight(strokeWidth), strokeColor)
        }
        subGridImg, _ = ebiten.NewImageFromImage(im.SubImage(im.Rect), ebiten.FilterDefault)
    }
    
    _ = img.DrawImage(subGridImg, &ebiten.DrawImageOptions{})
    _ = screen.DrawImage(img, &ebiten.DrawImageOptions{})
    return img
}


func (g TileGrid) DrawTile(screen *ebiten.Image, pos Position) {
    w, h := screen.Size()
    x, y := pos.GetXY()
    colW, rowH := float64(w)/3, float64(h)/3
    
    img := g.At(pos).GetImage()
    scale := math.Min(colW, rowH) / float64(img.Bounds().Dx())
    opts := &ebiten.DrawImageOptions{}
    opts.GeoM.Scale(scale, scale)
    opts.GeoM.Translate(float64(x) * colW, float64(y) * rowH)
    _ = screen.DrawImage(img, opts)
}