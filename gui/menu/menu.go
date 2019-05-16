package menu

import (
    . "github.com/DrunkenPoney/go-tictactoe/gui/menu/btn"
    . "github.com/DrunkenPoney/go-tictactoe/gui/menu/pages"
    "github.com/hajimehoshi/ebiten"
)

type Menu interface {
    SetShown(show bool)
    IsShown() bool
    Draw(screen *ebiten.Image)
    
    GetPage() MenuPage
    SetPage(page MenuPage)
    
    OnButtonClick(btn ButtonType, action func())
    Freeze(freeze bool)
    IsFrozen() bool
}

func NewMenu() Menu {
    return &menu{actions: make(map[ButtonType]func())}
}

type menu struct {
    shown   bool
    actions map[ButtonType]func()
    frozen  bool
    page    MenuPage
}

func (m *menu) SetShown(show bool) {
    m.shown = show
}

func (m *menu) IsShown() bool {
    return m.shown
}

func (m *menu) OnButtonClick(btn ButtonType, action func()) {
    m.actions[btn] = action
}

func (m *menu) Freeze(freeze bool) {
    m.frozen = freeze
}

func (m *menu) IsFrozen() bool {
    return m.frozen
}

func (m *menu) GetPage() MenuPage {
    return m.page
}

func (m *menu) SetPage(page MenuPage) {
    m.page = page
}