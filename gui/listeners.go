package gui

import (
    . "github.com/DrunkenPoney/go-tictactoe/gui/menu/btn"
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
