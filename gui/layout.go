package gui

import (
    . "github.com/DrunkenPoney/go-tictactoe/game"
    . "github.com/DrunkenPoney/go-tictactoe/gui/menu"
    "github.com/hajimehoshi/ebiten"
)

type Layout interface {
    Update(screen *ebiten.Image) error
    ShowMenu(show bool) // TODO
    ShowMessage() // TODO
    GetGame() Game
    NewGame() // TODO Add possibility to pause, resume, stop and restart the game
}

type layout struct {
    game Game
    menu Menu
}

