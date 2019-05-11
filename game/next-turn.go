package game

import (
	"fmt"
	. "github.com/DrunkenPoney/go-tictactoe/game/player"
	"github.com/DrunkenPoney/go-tictactoe/grid"
	. "github.com/DrunkenPoney/go-tictactoe/grid/tile"
	. "github.com/DrunkenPoney/go-tictactoe/gui/message"
	. "github.com/DrunkenPoney/go-tictactoe/position"
	"github.com/sqweek/dialog"
)

func (g *game) CheckWinnerInGrid(tiles grid.TileGrid) Player {
	cells := tiles.GetWinningTiles()
	if cells[0] != nil {
		for _, cell := range cells {
			cell.Winning = true
		}

		if cells[0].Value == X {
			return g.playerX
		} else {
			return g.playerO
		}
	}
	return nil
}

func (g *game) NextTurn(pos Position) Game {
	if pos == INVALID {
		return g
	}
	fmt.Printf("-------------------------- NEW TURN --------------------------\n")
	winner := g.CheckWinnerInGrid(g.board.CurrentGrid())
	if winner == nil {
		g.playerX.SetCurrent(!g.playerX.IsCurrent())
		g.playerO.SetCurrent(!g.playerO.IsCurrent())
		g.ai.PrepareNextTurn(pos)
		g.GetBoard().SetCurrentPos(pos)

		if g.playerX.IsCurrent() {
			bestPos := g.ai.BestMoveFor(X)
			tile := g.GetBoard().CurrentGrid().At(bestPos)
			tile.Value = X
			g.GetBoard().DrawTile(tile, g.GetBoard().GetCurrentPos())
			g.NextTurn(bestPos)
		}
	} else {
		winner.IncrementPoints()
		ok := dialog.Message(MSG_NEW_GAME.Str()).Title(MSG_NEW_GAME.Str()).YesNo()
		if ok {
			g.Reset()
		}
	}
	return g
}
