package game

import (
    "github.com/DrunkenPoney/go-tictactoe/game/player"
)

type Game interface {
    GetPlayerO() player.Player
}

