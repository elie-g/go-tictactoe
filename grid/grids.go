package grid

import (
	. "image/color"
)

type Grids interface {
	Columns() int
	Rows() int

	GetGrids() [][]Grid
	GetGridAt(x int, y int) Grid
}

func NewGrids(columns int, rows int) Grids {
	gridSlice := make([][]Grid, columns)
	for x := range gridSlice {
		gridSlice[x] = make([]Grid, rows)
		for y := range gridSlice[x] {
			gridSlice[x][y] = NewGrid(3, 3, White, 10)

		}
	}
	return &grids{gridSlice, columns, rows}
}

type grids struct {
	grids   [][]Grid
	columns int
	rows    int
}

func (g *grids) Columns() int {
	return g.columns
}

func (g *grids) Rows() int {
	return g.rows
}

func (g *grids) GetGrids() [][]Grid {
	return g.grids
}

func (g *grids) GetGridAt(x int, y int) Grid {
	return g.grids[x][y]
}
