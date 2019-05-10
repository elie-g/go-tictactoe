package gui

import (
    "github.com/DrunkenPoney/go-tictactoe/internal"
    "github.com/golang/freetype/truetype"
    "github.com/hajimehoshi/ebiten"
    "github.com/hajimehoshi/ebiten/examples/resources/fonts"
    f "golang.org/x/image/font"
)

var layoutImg *ebiten.Image
var font f.Face
var fntSize = internal.Scale(12)

func (l *layout) Draw(screen *ebiten.Image) {
    w, h := screen.Size()
    
    if layoutImg == nil {
        layoutImg, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
    }
    
    if font == nil {
        tt, err := truetype.Parse(fonts.ArcadeN_ttf)
        if err != nil {
            panic(err)
        }
        font = truetype.NewFace(tt, &truetype.Options{Size: fntSize, DPI: 72, Hinting: f.HintingFull})
    }
    
    l.GetGame().Draw(layoutImg)
    DrawGameStats(l.GetGame().GetPlayerO(), l.GetGame().GetPlayerX())
    _ = screen.DrawImage(layoutImg, &ebiten.DrawImageOptions{})
}
