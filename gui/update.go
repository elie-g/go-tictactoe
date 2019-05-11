package gui

import (
    "fmt"
    "github.com/hajimehoshi/ebiten"
    "github.com/hajimehoshi/ebiten/ebitenutil"
)

func (l *layout) Update(screen *ebiten.Image) error {
    if ebiten.IsDrawingSkipped() {
        return nil
    }
    
    if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
        l.GetGame().OnClick()
    }
    
    if l.activeTile != nil {
        l.activeTile.Active = false
        l.GetGame().GetBoard().DrawTile(l.activeTile,l.activePos)
    }
    
    if !l.GetMenu().IsShown() {
        if l.activeTile, l.activePos = l.GetGame().GetBoard().GetTileUnderCursor(); l.activeTile != nil {
            l.activeTile.Active = true
            l.GetGame().GetBoard().DrawTile(l.activeTile, l.activePos)
        }
    }
    
    l.Draw(screen)
    _ = ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %f, TPS: %f", ebiten.CurrentFPS(), ebiten.CurrentTPS()))
    return nil
}
