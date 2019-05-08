package ui

import (
    "image/color"
)

type BoardUI struct {
    GridColor       color.Color
    SubGridColor    color.Color
    BackgroundColor color.Color
    StrokeWidth     float64
    SubStrokeWidth  float64
}

func DefaultBoardUI() *BoardUI {
    return &BoardUI{
        GridColor:       color.White,
        SubGridColor:    color.Gray{Y: 0x99},
        BackgroundColor: color.Transparent,
        StrokeWidth:     10,
        SubStrokeWidth:  5,
    }
}
