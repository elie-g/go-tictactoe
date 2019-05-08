package internal

import (
    "github.com/hajimehoshi/ebiten"
)

//noinspection GoSnakeCaseUsage
const (
    REF_WIDTH  float64 = 1920
    REF_HEIGHT float64 = 1080
)

func ScaleFactors() (float64, float64) {
    x, y := ebiten.ScreenSizeInFullscreen()
    return float64(x) / REF_WIDTH, float64(y) / REF_HEIGHT
}

func ScaleWidth(w float64) float64 {
    sW, _ := ScaleFactors()
    return sW * w
}

func ScaleHeight(h float64) float64 {
    _, sH := ScaleFactors()
    return sH * h
}
