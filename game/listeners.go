package game

import (
	. "github.com/DrunkenPoney/go-tictactoe/game/state"
	. "github.com/DrunkenPoney/go-tictactoe/grid/tile"
)

// Private
func (g *game) onClick() {
	t, pos := g.board.GetTileUnderCursor()
	if t != nil {
		if g.GetBoard().GetCurrentPos() == pos && t.Value == EMPTY {
			if g.playerO.IsCurrent() {
				t.Value = O
			} else {
				t.Value = X
			}
			g.GetBoard().DrawTile(t, pos)
			g.NextTurn(t.Position)
		}
	}
}

// Private
func (g *game) onState(state State) {
	if g.state != state {
		switch state {
		case PAUSED:
			g.Pause()
		case RUNNING:
			g.Resume()
		case STOPPED:
			// TODO
		}
	}
}
