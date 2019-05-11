package gui

import (
    "fmt"
    "github.com/DrunkenPoney/go-tictactoe/game/player"
    "github.com/DrunkenPoney/go-tictactoe/internal"
    "github.com/DrunkenPoney/go-tictactoe/settings/colors"
    "github.com/hajimehoshi/ebiten"
    "github.com/hajimehoshi/ebiten/text"
    "math"
)

func DrawGameStats(playerLeft player.Player, playerRight player.Player) {
    wi, hi := layoutImg.Bounds().Dx(), int(fntSize+internal.ScaleHeight(2))
    w := float64(wi)
    
    img, _ := ebiten.NewImage(wi, hi, ebiten.FilterDefault)
    _ = img.Fill(colors.Colors().InGameTextBackground())
    
    yFnt := int(internal.Scale(1) + fntSize)
    // ///////////////////////////////// LEFT SIDE
    name, pts := playerLeft.GetName(), fmt.Sprintf("<%d>", playerLeft.GetPoints())
    // xJ1: Name Player 1 // xS1: Points Player 1
    xS1 := (w/9)*3 + (w / 18) - (fntSize * float64(len(pts)) / 2)
    xJ1 := xS1/2 - (fntSize * float64(len(name)) / 2)
    if xJ1+(fntSize*float64(len(name))/2) > xS1 {
        xJ1 -= ((xJ1 + (fntSize * float64(len(name)) / 2)) - xS1) - (fntSize * float64(len(name)) / 2)
    }
    text.Draw(img, name, font, int(math.Round(xJ1)), yFnt, colors.Colors().InGameTextColor())
    text.Draw(img, pts, font, int(math.Round(xS1+internal.ScaleWidth(3))), yFnt, colors.Colors().InGamePointsColor())
    
    // //////////////////////////////////// CENTER
    sep := "||"
    x := (w / 2) - (fntSize * float64(len(sep)) / 2)
    text.Draw(img, sep, font, int(math.Round(x)), yFnt, colors.Colors().InGameTextColor())
    
    // //////////////////////////////// RIGHT SIDE
    name, pts = playerRight.GetName(), fmt.Sprintf("<%d>", playerRight.GetPoints())
    // xJ2: Name Player 2 // xS2: Points Player 2
    xS2 := (w / 9 * 6) - (w / 18) - (fntSize * float64(len(pts)) / 2)
    xJ2 := w - (xS1 / 2) - (fntSize * float64(len(name)) / 2)
    if xJ2 < xS2+(fntSize*float64(len(pts))/2) {
        xJ2 += ((xJ2 + (fntSize * float64(len(name)) / 2)) - xS2) - (fntSize * float64(len(name)) / 2)
    }
    text.Draw(img, pts, font, int(xS2), yFnt, colors.Colors().InGamePointsColor())
    text.Draw(img, name, font, int(xJ2+internal.ScaleWidth(6)), yFnt, colors.Colors().InGameTextColor())
    
    opts := &ebiten.DrawImageOptions{}
    opts.GeoM.Translate(0, float64(layoutImg.Bounds().Dy()-hi))
    _ = layoutImg.DrawImage(img, opts)
}
