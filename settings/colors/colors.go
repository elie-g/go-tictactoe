package colors

import (
    . "image/color"
)

var instance GUIColors

type GUIColors interface {
    InGameTextBackground() Color
    InGameTextColor() Color
    InGamePointsColor() Color
    MenuFadeBackground() Color
    MenuBackground() Color
    MenuButtonBackground() Color
    MenuButtonColor() Color
    MenuButtonHoverBackground() Color
}

func Colors() GUIColors {
    if instance == nil {
        instance = &guiColor {
            inGameTextBg: RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x99},
            inGameTextColor: RGBA{R: 0xFF, G: 0x99, B: 0x00, A: 0xFF},
            inGamePointsColor: RGBA{R: 0x99, G: 0xFF, B: 0x00, A: 0xFF},
            menuFadeBg: RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x33},
            menuBg: RGBA{R: 0x22, G: 0x22, B: 0x22, A: 0xFF},
            menuBtnBg: RGBA{R: 0x99, G: 0x99, B: 0x99, A: 0xFF},
            menuBtnColor: RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF},
            menuBtnHoverBg: RGBA{R: 0xCC, G: 0xCC, B: 0xCC, A: 0xCC},
        }
    }
    
    return instance
}

type guiColor struct {
    inGameTextBg Color
    inGameTextColor Color
    inGamePointsColor Color
    menuFadeBg Color
    menuBg Color
    menuBtnBg Color
    menuBtnColor Color
    menuBtnHoverBg Color
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

func (gc *guiColor) MenuFadeBackground() Color {
    return gc.menuFadeBg
}

func (gc *guiColor) MenuBackground() Color {
    return gc.menuBg
}

func (gc *guiColor) MenuButtonBackground() Color {
    return gc.menuBtnBg
}

func (gc *guiColor) MenuButtonColor() Color {
    return gc.menuBtnColor
}

func (gc *guiColor) MenuButtonHoverBackground() Color {
    return gc.menuBtnHoverBg
}
