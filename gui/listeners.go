package gui

import (
    "fmt"
    "github.com/hajimehoshi/ebiten"
    "time"
)

var pressTimes = make(map[ebiten.Key]time.Time)
const MIN_KEY_DELAY = 220 * time.Millisecond

func (l *layout) checkKeypress() {
    if ebiten.IsKeyPressed(ebiten.KeyEscape) {
        if time.Since(pressTimes[ebiten.KeyEscape]) > MIN_KEY_DELAY {
            pressTimes[ebiten.KeyEscape] = time.Now()
            fmt.Println("\033[96;1mESCAPE PRESSED!!!\033[m")
            l.ToggleMenu()
        }
    }
}
