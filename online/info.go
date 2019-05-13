package online

import (
    . "github.com/DrunkenPoney/go-tictactoe/db"
)

type OnlineData struct {
    DB             Database
    Game           DBGame
    LocalPlayer    DBPlayer
    RemotePlayer   DBPlayer
    IsLocalPlayer1 bool
    IsPlayerAI     bool
}
