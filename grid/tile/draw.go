package tile

import (
    "github.com/DrunkenPoney/go-tictactoe/internal"
    "github.com/hajimehoshi/ebiten"
    "github.com/peterhellberg/gfx"
    "image/color"
)

const (
    _DEFAULT_ uint8 = 0x00 // 0000 0000
    _X_       uint8 = 0x01 // 0000 0001
    _O_       uint8 = 0x02 // 0000 0010
    _ACTIVE_  uint8 = 0x04 // 0000 0100
    _WINNING_ uint8 = 0x10 // 0000 1000
)

//noinspection GoSnakeCaseUsage
const (
    BASE_WIDTH  int     = 1000
    BASE_HEIGHT int     = 1000
    THICKNESS   float64 = 25
)

var tileImages = make(map[uint8]*ebiten.Image)

func (t *Tile) GetImage() *ebiten.Image {
    id := _DEFAULT_
    if t.Value == X {
        id |= _X_
    }
    if t.Value == O {
        id |= _O_
    }
    if t.Active {
        id |= _ACTIVE_
    }
    if t.Winning {
        id |= _WINNING_
    }
    
    
    // Si la tile a changé
    if tileImages[id] == nil {
        img := gfx.NewImage(BASE_WIDTH, BASE_HEIGHT)
        rect := gfx.Polygon{
            gfx.V(0, 0),
            gfx.V(float64(BASE_WIDTH), 0),
            gfx.V(float64(BASE_WIDTH), float64(BASE_HEIGHT)),
            gfx.V(0, float64(BASE_HEIGHT)),
        }
        
        var c, stroke color.Color
        if t.Active && t.Winning {
            c = color.RGBA{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF}
            stroke = color.RGBA{R: 0x00, G: 0x33, B: 0x00, A: 0xEE}
        } else if t.Active {
            c = color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
            stroke = color.Black
        } else if t.Winning {
            c = color.RGBA{R: 0x00, G: 0x44, B: 0x00, A: 0x99}
            stroke = color.RGBA{R: 0x66, G: 0xFF, B: 0x33, A: 0x99}
        } else {
            c = color.Black
            stroke = color.White
        }
        if c != nil {
            gfx.DrawPolygon(img, rect, 0, c)
        }
        
        switch t.Value {
        case X:
            w, h := float64(img.Rect.Dx()), float64(img.Rect.Dy())
            r := 0.2
            tl := gfx.V(r*w, r*h)
            tr := gfx.V((1-r)*w, r*h)
            bl := gfx.V(r*w, (1-r)*h)
            br := gfx.V((1-r)*w, (1-r)*h)
            gfx.DrawLine(img, tl, br, internal.Scale(THICKNESS), stroke)
            gfx.DrawLine(img, tr, bl, internal.Scale(THICKNESS), stroke)
        case O:
            center := gfx.V(float64(img.Rect.Dx())/2, float64(img.Rect.Dy())/2)
            gfx.DrawCircle(img, center, 0.7*(float64(img.Rect.Dx())/2), internal.Scale(THICKNESS*1.6), stroke)
        case EMPTY:
            // Do nothing
        }
        
        tileImages[id], _ = ebiten.NewImageFromImage(img.SubImage(img.Rect), ebiten.FilterDefault)
    }
    return tileImages[id]
}
