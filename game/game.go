package game

import (
	"fmt"
	"github.com/DrunkenPoney/go-tictactoe/events"
	"github.com/DrunkenPoney/go-tictactoe/game/player"
	"github.com/DrunkenPoney/go-tictactoe/grid"
	"github.com/DrunkenPoney/go-tictactoe/grid/tile"
	"github.com/hajimehoshi/ebiten"
	"math/rand"
)

type Game interface {
	GetPlayerO() player.Player
	GetPlayerX() player.Player
	NextTurn() Game
	GetBoard() grid.Grid
	GetWinner() player.Player
	GetCurrentPlayer() player.Player
	Reset() Game
	Draw(screen *ebiten.Image) Game
}

func NewGame(playerO player.Player, playerX player.Player, board grid.Grid) Game {
	if !playerO.IsCurrent() && !playerX.IsCurrent() ||
		playerO.IsCurrent() && playerX.IsCurrent() {
		rdm := rand.Float64() >= 0.5
		playerO.SetCurrent(rdm)
		playerX.SetCurrent(!rdm)
	}
	listener := events.NewClickListener()
	g := &game{playerO, playerX, board, listener}
	listener.Listen(g.onClick)
	listener.Resume()
	return g
}

type game struct {
	playerO       player.Player
	playerX       player.Player
	board         grid.Grid
	clickListener events.ClickListener
}

// Private
func (g *game) onClick() {
	t := g.board.GetTileUnderCursor()
	if t.GetValue() == tile.EMPTY {
		if g.playerO.IsCurrent() {
			t.SetValue(tile.O)
		} else {
			t.SetValue(tile.X)
		}

		g.GetBoard().SetGridNumber(t.GetPosition())
		g.NextTurn()
	}
}

func (g *game) GetPlayerO() player.Player {
	return g.playerO
}

func (g *game) GetPlayerX() player.Player {
	return g.playerX
}

func (g *game) GetBoard() grid.Grid {
	return g.board
}

func (g *game) GetWinner() player.Player {
	tiles := g.board.GetTiles()
	for col, columns := range tiles {
		for row, cell := range columns {
			if cell.GetValue() != tile.EMPTY {
				if len(columns) > row+2 && columns[row+1].GetValue() == cell.GetValue() &&
					columns[row+2].GetValue() == cell.GetValue() {
					columns[row+1].SetWinning(true)
					columns[row+2].SetWinning(true)
					cell.SetWinning(true)
				} else if len(tiles) > col+2 && tiles[col+1][row].GetValue() == cell.GetValue() &&
					tiles[col+2][row].GetValue() == cell.GetValue() {
					tiles[col+1][row].SetWinning(true)
					tiles[col+2][row].SetWinning(true)
					cell.SetWinning(true)
				}
				if cell.IsWinning() {
					if cell.GetValue() == tile.X {
						return g.playerX
					} else {
						return g.playerO
					}
				}
			}
		}
	}

	return nil
}

func (g *game) willWin() bool {

	return true
}

func (g *game) NextTurn() Game {
	for {
		g.GetWinner()
		g.playerX.SetCurrent(!g.playerX.IsCurrent())
		g.playerO.SetCurrent(!g.playerO.IsCurrent())

		if g.playerX.IsCurrent() {
			g.PlayAINextMove()

		} else {
			break

		}
	}

	return g
}

func (g *game) Reset() Game {
	g.GetBoard().Reset()
	return g
}

func (g *game) Draw(screen *ebiten.Image) Game {
	g.GetBoard().DrawGrid(screen)
	return g
}

func (g *game) GetCurrentPlayer() player.Player {
	cur := g.playerO
	if g.playerX.IsCurrent() {
		cur = g.playerX
	}
	return cur
}

func (g *game) PlayAINextMove() {
	var posibility = g.GetPosibility()
	var choice = g.GetNextMove(posibility)
	fmt.Println(g.GetBoard().GetColOffset())
	fmt.Println(g.GetBoard().GetRowOffset())
	fmt.Println(posibility)

	g.GetBoard().GetTileAt(choice[0], choice[1]).SetValue(tile.X)
}

func (g *game) GetPosibility() [][]int {
	var posibility [][]int

	fmt.Println(g.GetBoard().GetColOffset())
	fmt.Println(g.GetBoard().GetRowOffset())

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.GetBoard().GetTileAt(i+g.GetBoard().GetColOffset(), j+g.GetBoard().GetRowOffset()).GetValue() == tile.EMPTY {
				choice := []int{i + g.GetBoard().GetColOffset(), j + g.GetBoard().GetRowOffset()}
				posibility = append(posibility, choice)
			}
		}
	}

	return posibility
}

func (g *game) GetNextMove(choices [][]int) []int {
	var choice []int
	choice = choices[0]

	return choice
}
