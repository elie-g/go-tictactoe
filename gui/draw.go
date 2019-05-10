package gui

import (
	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	f "golang.org/x/image/font"
	"image/color"
)

var layoutImg *ebiten.Image
var font f.Face

func (l *layout) Draw(screen *ebiten.Image) {
	w, h := screen.Size()

	if layoutImg == nil {
		layoutImg, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	}

	l.GetGame().Draw(layoutImg)
	str := fmt.Sprintf("%s          <%d>     ||     <%d>          %s",
		l.GetGame().GetPlayerO().GetName(),
		l.GetGame().GetPlayerO().GetPoints(),
		l.GetGame().GetPlayerX().GetPoints(),
		l.GetGame().GetPlayerX().GetName())
	drawText(layoutImg, str, float64(w)/2, float64(h)-32)
	_ = screen.DrawImage(layoutImg, &ebiten.DrawImageOptions{})
}

func drawText(img *ebiten.Image, txt string, x float64, y float64) {
	if font == nil {
		tt, err := truetype.Parse(fonts.ArcadeN_ttf)
		if err != nil {
			panic(err)
		}

		font = truetype.NewFace(tt, &truetype.Options{
			Size:    24,
			DPI:     72,
			Hinting: f.HintingFull})
	}

	fontImg, _ := ebiten.NewImage(len(txt)*24-12, 34, ebiten.FilterDefault)
	_ = fontImg.Fill(color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x99})

	text.Draw(fontImg, txt, font, 12, 28, color.RGBA{R: 0xFF, G: 0x99, B: 0x00, A: 0xFF})
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(x-(float64(fontImg.Bounds().Dx())/2)+6, y)
	_ = img.DrawImage(fontImg, opts)
	_ = fontImg.Dispose()
}
