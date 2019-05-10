package colors

import (
    . "image/color"
)

var instance GUIColors

type GUIColors interface {
    InGameTextBackground() Color
    InGameTextColor() Color
    InGamePointsColor() Color
}

func Colors() GUIColors {
    if instance == nil {
        instance = &guiColor {
            inGameTextBg: RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x99},
            inGameTextColor: RGBA{R: 0xFF, G: 0x99, B: 0x00, A: 0xFF},
            inGamePointsColor: RGBA{R: 0x99, G: 0xFF, B: 0x00, A: 0xFF},
        }
    }
    
    return instance
}

type guiColor struct {
    inGameTextBg Color
    inGameTextColor Color
    inGamePointsColor Color
}

func (gc *guiColor) InGameTextBackground() Color {
    return gc.inGameTextBg
}

func (gc *guiColor) InGameTextColor() Color {
    return gc.inGameTextColor
}

func (gc *guiColor) InGamePointsColor() Color {
    return gc.inGamePointsColor
}
