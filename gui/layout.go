package gui

import (
    "github.com/DrunkenPoney/go-tictactoe/board"
    "github.com/DrunkenPoney/go-tictactoe/db"
    . "github.com/DrunkenPoney/go-tictactoe/game"
    . "github.com/DrunkenPoney/go-tictactoe/game/player"
    "github.com/DrunkenPoney/go-tictactoe/grid/tile"
    . "github.com/DrunkenPoney/go-tictactoe/gui/menu"
    . "github.com/DrunkenPoney/go-tictactoe/online"
    . "github.com/DrunkenPoney/go-tictactoe/position"
    "github.com/hajimehoshi/ebiten"
)

type Layout interface {
    Update(screen *ebiten.Image) error
    ToggleMenu() bool
    GetGame() Game // TODO Add possibility to pause, resume and restart the game
    GetMenu() Menu
    GetOnlineData() *OnlineData
    IsOnline() bool
    Draw(screen *ebiten.Image)
}

func NewLayout(playerO string, playerX string) Layout {
    layout := &layout{menu: NewMenu()}
    layout.game = NewGame(NewPlayer(playerO), NewPlayer(playerX), board.NewBoard())
    layout.initListeners()
    return layout
}

type layout struct {
    activeTile *tile.Tile
    activePos  Position
    onlineData *OnlineData
    game       Game
    menu       Menu
}

func (l *layout) ToggleMenu() bool {
    l.GetMenu().SetShown(!l.GetMenu().IsShown())
    if l.GetMenu().IsShown() {
        l.GetGame().Pause()
    } else {
        l.GetGame().Resume()
    }
    return l.GetMenu().IsShown()
}

func (l *layout) GetGame() Game {
    return l.game
}

func (l *layout) GetMenu() Menu {
    return l.menu
}

func (l *layout) IsOnline() bool {
    return l.onlineData != nil
}

func (l *layout) GetOnlineData() *OnlineData {
    if l.onlineData == nil {
        l.onlineData = &OnlineData{DB: db.NewDatabase()}
    }
    return l.onlineData
}
