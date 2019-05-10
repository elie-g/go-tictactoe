package path

import (
    . "github.com/DrunkenPoney/go-tictactoe/position"
)

type Path []Position

func (p Path) IsChildOf(parent Path) bool {
    if len(parent) < len(p) {
        isChild := true
        for i, pos := range parent {
            if isChild = pos == p[i]; !isChild {
                break
            }
        }
        return isChild
    }
    return false
}

func (p Path) Next(pos Position) Path {
    return append(p, pos)
}