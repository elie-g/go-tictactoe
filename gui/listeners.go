package gui

import (
    "fmt"
    . "github.com/DrunkenPoney/go-tictactoe/gui/menu/btn"
    . "github.com/DrunkenPoney/go-tictactoe/gui/menu/pages"
    "github.com/DrunkenPoney/go-tictactoe/settings"
    . "github.com/DrunkenPoney/go-tictactoe/settings/messages"
    "github.com/hajimehoshi/ebiten"
    "github.com/martinlindhe/inputbox"
    "github.com/sqweek/dialog"
    "os"
    "strings"
    "time"
)

var pressTimes = make(map[ebiten.Key]time.Time)

func (l *layout) checkKeypress() {
    if ebiten.IsKeyPressed(ebiten.KeyEscape) {
        if time.Since(pressTimes[ebiten.KeyEscape]) > settings.MIN_KEYPRESS_DELAY {
            pressTimes[ebiten.KeyEscape] = time.Now()
            l.ToggleMenu()
        }
    }
}

// //////////////////////////////// Button Listeners //////////////////////////////// //
func (l *layout) initListeners() {
    l.GetMenu().OnButtonClick(BTN_EXIT, l.onBtnExit)
    l.GetMenu().OnButtonClick(BTN_RESUME, l.onBtnResume)
    l.GetMenu().OnButtonClick(BTN_PLAYER_1, l.onBtnPlayer1)
    l.GetMenu().OnButtonClick(BTN_PLAYER_2, l.onBtnPlayer2)
    l.GetMenu().OnButtonClick(BTN_ONLINE, l.onBtnOnline)
    l.GetMenu().OnButtonClick(BTN_RETURN, l.onBtnReturn)
    l.GetMenu().OnButtonClick(BTN_JOIN_ONLINE, l.onBtnJoinOnline)
    l.GetMenu().OnButtonClick(BTN_CREATE_ONLINE, l.onBtnCreateOnline)
}

func (l *layout) onBtnExit() {
    os.Exit(0)
}

func (l *layout) onBtnResume() {
    l.GetMenu().SetShown(false)
    l.GetGame().Resume()
}

func (l *layout) onBtnPlayer1() {
    l.GetMenu().Freeze(true)
    name, ok := inputbox.InputBox(MSG_PLAYER_1_BOX_TITLE.Str(), MSG_PLAYER_1_BOX.Str(), "")
    name = strings.TrimSpace(name)
    if ok && len(name) > 0 {
        if len(name) > settings.MAX_PLAYER_NAME_LENGTH {
            dialog.Message(MSG_NAME_TOO_LONG_TITLE.Str()).
                Title(MSG_NAME_TOO_LONG.Str(settings.MAX_PLAYER_NAME_LENGTH)).
                Error()
        } else {
            l.GetGame().GetPlayerO().SetName(name)
        }
    }
    l.GetMenu().Freeze(false)
}

func (l *layout) onBtnPlayer2() {
    l.GetMenu().Freeze(true)
    name, ok := inputbox.InputBox(MSG_PLAYER_2_BOX_TITLE.Str(), MSG_PLAYER_2_BOX.Str(), "")
    name = strings.TrimSpace(name)
    if ok && len(name) > 0 {
        if len(name) > settings.MAX_PLAYER_NAME_LENGTH {
            dialog.Message(MSG_NAME_TOO_LONG_TITLE.Str()).
                Title(MSG_NAME_TOO_LONG.Str(settings.MAX_PLAYER_NAME_LENGTH)).
                Error()
        } else {
            l.GetGame().GetPlayerX().SetName(name)
        }
    }
    l.GetMenu().Freeze(false)
}

func (l *layout) onBtnOnline() {
    l.GetMenu().SetPage(MP_ONLINE)
}

func (l *layout) onBtnReturn() {
    l.GetMenu().SetPage(MP_MAIN)
}

func (l *layout) onBtnJoinOnline() {
    game := l.GetOnlineData().DB.LastCreatedGame()
    fmt.Printf("%v // %v\n", game, game == nil)
    if game != nil {
        l.GetOnlineData().IsPlayerAI = dialog.Message(MSG_CHOOSE_PLAYER.Str()).Title(MSG_CHOOSE_PLAYER_TITLE.Str()).YesNo()
        l.GetOnlineData().RemotePlayer = game.GetPlayer1()
        l.GetOnlineData().IsLocalPlayer1 = false
        player := l.GetGame().GetPlayerO()
        if l.GetOnlineData().IsPlayerAI {
            player = l.GetGame().GetPlayerX()
            l.GetGame().GetPlayerO().SetName(l.GetOnlineData().RemotePlayer.GetName())
            l.GetGame().GetPlayerO().SetRemote(true)
        } else {
            l.GetGame().GetPlayerX().SetName(l.GetOnlineData().RemotePlayer.GetName())
            l.GetGame().GetPlayerX().SetRemote(false)
        }
        l.GetOnlineData().LocalPlayer = l.GetOnlineData().DB.CreatePlayer(player.GetName())
        l.GetOnlineData().Game.SetPlayer2(l.GetOnlineData().LocalPlayer)
        // TODO Who plays the next turn?
        l.GetGame().Reset()
    } else {
        dialog.Message(MSG_NO_GAME_AVAILABLE.Str()).Error()
    }
}

func (l *layout) onBtnCreateOnline() {
    l.GetOnlineData().IsPlayerAI = dialog.Message(MSG_CHOOSE_PLAYER.Str()).Title(MSG_CHOOSE_PLAYER_TITLE.Str()).YesNo()
    player := l.GetGame().GetPlayerO()
    if l.GetOnlineData().IsPlayerAI {
        player = l.GetGame().GetPlayerX()
        l.GetGame().GetPlayerO().SetRemote(true)
    } else {
        l.GetGame().GetPlayerX().SetRemote(false)
    }
    l.GetOnlineData().IsLocalPlayer1 = true
    l.GetOnlineData().LocalPlayer = l.GetOnlineData().DB.CreatePlayer(player.GetName())
    l.GetOnlineData().Game = l.GetOnlineData().DB.CreateGame(l.GetOnlineData().LocalPlayer)
    // TODO Who plays the next turn?
    // TODO Wait other player to connect
    l.GetGame().Reset()
}
