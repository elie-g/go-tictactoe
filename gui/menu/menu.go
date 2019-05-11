package menu

import (
    "github.com/hajimehoshi/ebiten"
)

type Menu interface {
    SetShown(show bool)
    IsShown() bool
    Draw(screen *ebiten.Image)
    
    OnExit(cb func())
    OnRestart(cb func())
    OnResume(cb func())
}

func NewMenu() Menu {
    return &menu{}
}

type menu struct {
    shown bool
    exitCb func()
    restartCb func()
    resumeCb func()
}

func (m *menu) SetShown(show bool) {
    m.shown = show
}

func (m *menu) IsShown() bool {
    return m.shown
}

func (m *menu) OnExit(cb func()) {
    m.exitCb = cb
}

func (m *menu) OnRestart(cb func()) {
    m.restartCb = cb
}

func (m *menu) OnResume(cb func()) {
    m.resumeCb = cb
}




