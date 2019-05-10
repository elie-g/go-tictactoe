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
    ToggleMenu(show bool) // TODO
    ShowMessage()         // TODO
    GetGame() Game
    NewGame() // TODO Add possibility to pause, resume, stop and restart the game
}

func NewLayout(playerO string, playerX string) Layout {
    return &layout{
        game: NewGame(NewPlayer(playerO), NewPlayer(playerX), board.NewBoard()),
        menu: nil} // TODO add menu
}

type layout struct {
    activeTile *tile.Tile
    activePos  Position
    game       Game
    menu       Menu
}

func (l *layout) ToggleMenu(show bool) {
    panic("implement me")
}

func (l *layout) ShowMessage() {
    panic("implement me")
}

func (l *layout) GetGame() Game {
    return l.game
}

func (l *layout) NewGame() {
    panic("implement me")
}


