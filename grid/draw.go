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
                scale := math.Min(colWidth, rowHeight)/float64(tileImg.Bounds().Dx())
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

// func (g *grid) DrawGrid(img *ebiten.Image) {
//     width, height := float64(img.Bounds().Dx()), float64(img.Bounds().Dy())
//     colWidth, rowHeight := width/float64(g.columns), height/float64(g.rows)
//     im, _ := ebiten.NewImage(int(width), int(height), ebiten.FilterDefault)
//     ctxt := gg.NewContextForImage(im)
//
//     for i, col := range g.tiles {
//         x := float64(i) * colWidth
//         for j, tile := range col {
//             y := float64(j) * rowHeight
//             if tile.GetBgColor() != nil {
//                 ctxt.DrawRectangle(x, y, colWidth, rowHeight)
//                 ctxt.SetColor(tile.GetBgColor())
//                 ctxt.Fill()
//             }
//
//             if tile.GetValue() != EMPTY {
//                 if tile.GetValue() == X {
//                     var tl, tr, br, bl []float64
//                     ratio := 0.8
//                     tl = []float64{x + (colWidth * (1 - ratio)), y + (rowHeight * (1 - ratio))}
//                     tr = []float64{x + (colWidth * ratio), y + (rowHeight * (1 - ratio))}
//                     br = []float64{x + (colWidth * ratio), y + (rowHeight * ratio)}
//                     bl = []float64{x + (colWidth * (1 - ratio)), y + (rowHeight * ratio)}
//                     ctxt.DrawLine(tl[0], tl[1], br[0], br[1])
//                     ctxt.DrawLine(tr[0], tr[1], bl[0], bl[1])
//
//                 } else if tile.GetValue() == O {
//                     ctxt.DrawCircle(x+(colWidth/2), y+(rowHeight/2), (0.6)*(colWidth/2))
//                 }
//
//                 ctxt.SetLineWidth(tile.GetStrokeWidth())
//                 ctxt.SetColor(tile.GetColor())
//                 ctxt.Stroke()
//             }
//         }
//     }
//
//     // Grid lines (#)
//     for x := colWidth; x < (width - (colWidth / 2)); x += colWidth {
//         ctxt.DrawLine(x, 0, x, height)
//     }
//     for y := rowHeight; y < (height - (rowHeight / 2)); y += rowHeight {
//         ctxt.DrawLine(0, y, width, y)
//     }
//
//     ctxt.SetLineWidth(g.strokeWidth)
//     ctxt.SetColor(g.color)
//     ctxt.Stroke()
//
//     newImg, err := ebiten.NewImageFromImage(ctxt.Image(), ebiten.FilterDefault)
//     if err != nil { log.Fatal(err) }
//     err = img.DrawImage(newImg, &ebiten.DrawImageOptions{})
//     if err != nil { log.Fatal(err) }
//     g.img = newImg
// }