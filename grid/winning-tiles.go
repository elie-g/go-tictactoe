package grid

import (
    . "github.com/DrunkenPoney/go-tictactoe/grid/tile"
)

func (g TileGrid) GetWinningTiles() [3]*Tile {
    // Vérifie diagonales
    val := g[1][1].Value
    if val != EMPTY {
        if val == g[0][0].Value && val == g[2][2].Value {
            return [3]*Tile{g[0][0], g[1][1], g[2][2]}
        } else if val == g[0][2].Value && val == g[2][0].Value {
            return [3]*Tile{g[0][2], g[1][1], g[2][0]}
        }
    }
    
    
    for col, rows := range g {
        for row, cell := range rows {
            // Skip les cases déjà vérifiées
            if col > 0 && row > 0 {
                continue
            }
            
            val = cell.Value
            if val != EMPTY {
                if len(rows) > row+2 && rows[row+1].Value == val && rows[row+2].Value == val {
                    return [3]*Tile{cell, rows[row+1], rows[row+2]}
                } else if len(g) > col+2 && g[col+1][row].Value == val && g[col+2][row].Value == val {
                    return [3]*Tile{cell, g[col+1][row], g[col+2][row]}
                }
            }
        }
    }
    
    return [3]*Tile{}
}
