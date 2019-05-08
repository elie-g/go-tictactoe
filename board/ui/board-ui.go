package ui

import (
    "image/color"
)

type BoardUI struct {
    GridColor       color.Color
    BackgroundColor color.Color
    StrokeWidth     float64
}

func DefaultBoardUI() *BoardUI {
    return &BoardUI{
        GridColor:       color.White,
        BackgroundColor: color.Transparent,
        StrokeWidth: 5,
    }
}