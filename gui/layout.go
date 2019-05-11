package gui

import (
    "github.com/DrunkenPoney/go-tictactoe/board"
    . "github.com/DrunkenPoney/go-tictactoe/game"
    . "github.com/DrunkenPoney/go-tictactoe/game/player"
    "github.com/DrunkenPoney/go-tictactoe/grid/tile"
    . "github.com/DrunkenPoney/go-tictactoe/gui/menu"
    . "github.com/DrunkenPoney/go-tictactoe/position"
    "github.com/hajimehoshi/ebiten"
)

type Layout interface {
    Update(screen *ebiten.Image) error
    ToggleMenu() bool // TODO
    ShowMessage()     // TODO
    GetGame() Game // TODO Add possibility to pause, resume and restart the game
    GetMenu() Menu
    Draw(screen *ebiten.Image)
}

func NewLayout(playerO string, playerX string) Layout {
    layout := &layout{
        game: NewGame(NewPlayer(playerO), NewPlayer(playerX), board.NewBoard()),
        menu: NewMenu()} // TODO add menu
    return layout
}

type layout struct {
    activeTile *tile.Tile
    activePos  Position
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

func (l *layout) ShowMessage() {
    panic("implement me")
}

func (l *layout) GetGame() Game {
    return l.game
}

func (l *layout) GetMenu() Menu {
    return l.menu
}
