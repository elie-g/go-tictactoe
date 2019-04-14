package tile

import (
    "github.com/hajimehoshi/ebiten"
    "github.com/peterhellberg/gfx"
    "image/color"
    "log"
)

const (
    _DEFAULT_ uint8 = 00000000
    _X_       uint8 = 00000001
    _ACTIVE_  uint8 = 00000010
    _WINNING_ uint8 = 00000100
)

//noinspection GoSnakeCaseUsage
const (
    BASE_WIDTH  int     = 1000
    BASE_HEIGHT int     = 1000
    THICKNESS   float64 = 25
)

var tileImages = make(map[uint8]*ebiten.Image)

func (t *tile) GetImage() *ebiten.Image {
    id := _DEFAULT_
    if t.GetValue() == X { id |= _X_ }
    if t.IsActive() { id |= _ACTIVE_ }
    if t.IsWinning() { id |= _WINNING_ }
    
    if tileImages[id] == nil && id != _DEFAULT_ {
        img := gfx.NewImage(BASE_WIDTH, BASE_HEIGHT)
        rect := gfx.Polygon{
            gfx.V(0, 0),
            gfx.V(float64(BASE_WIDTH), 0),
            gfx.V(float64(BASE_WIDTH), float64(BASE_HEIGHT)),
            gfx.V(0, float64(BASE_HEIGHT)),
        }
        
        var c, stroke color.Color
        if t.IsActive() && t.IsWinning() {
            c = color.RGBA{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF}
            stroke = color.RGBA{R: 0x00, G: 0x33, B: 0x00, A: 0xEE }
        } else if t.IsActive() {
            c = color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
            stroke = color.Black
        } else if t.IsWinning() {
            c = color.RGBA{R: 0x00, G: 0x44, B: 0x00, A: 0x99}
            stroke = color.RGBA{R: 0x66, G: 0xFF, B: 0x33, A: 0x99}
        } else {
            stroke = color.White
        }
        if c != nil {
            gfx.DrawPolygon(img, rect, 0, c)
        }
        
        // println(t.GetValue())
        switch t.GetValue() {
        case X:
            w, h := float64(img.Rect.Dx()), float64(img.Rect.Dy())
            r := 0.2
            tl := gfx.V(r*w, r*h)
            tr := gfx.V((1-r)*w, r*h)
            bl := gfx.V(r*w, (1-r)*h)
            br := gfx.V((1-r)*w, (1-r)*h)
            gfx.DrawLine(img, tl, br, THICKNESS, stroke)
            gfx.DrawLine(img, tr, bl, THICKNESS, stroke)
        case O:
            center := gfx.V(float64(img.Rect.Dx())/2, float64(img.Rect.Dy())/2)
            gfx.DrawCircle(img, center, 0.7*(float64(img.Rect.Dx())/2), THICKNESS*1.6, stroke)
        case EMPTY:
            // Do nothing
        }
        
        newImg, err := ebiten.NewImageFromImage(img.SubImage(img.Rect), ebiten.FilterDefault)
        if err != nil { log.Fatal(err) }
        tileImages[id] = newImg
    }
    
    return tileImages[id]
}