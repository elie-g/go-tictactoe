package board

import (
    "github.com/DrunkenPoney/go-tictactoe/grid"
    "github.com/DrunkenPoney/go-tictactoe/internal"
    . "github.com/DrunkenPoney/go-tictactoe/position"
    "github.com/hajimehoshi/ebiten"
    . "github.com/peterhellberg/gfx"
    "sync"
)

var gridImg, boardImg *ebiten.Image
var toDraw = make(map[Position]bool)

func (b *board) DrawBoard(screen *ebiten.Image) {
    b.screen = screen
    wi, hi := screen.Size()
    w, h := float64(wi), float64(hi)
    colW, rowH := w/3, h/3
    
    // Initialise la grille princial
    if gridImg == nil {
        boardImg, _ = ebiten.NewImage(wi, hi, ebiten.FilterDefault)
        b.ResetGrid(wi, hi)
    }
    
    // Initialise les sous-grilles
    if b.cellImg[TOP_LEFT] == nil {
        var mut sync.Mutex
        var wg sync.WaitGroup
        for pos, g := range b.Grids() {
            wg.Add(1)
            go func(pos Position, g grid.TileGrid, mut *sync.Mutex, wg *sync.WaitGroup) {
                defer wg.Done()
                img, _ := ebiten.NewImage(int(colW), int(rowH), ebiten.FilterDefault)
                g.Draw(img, b.UI().SubStrokeWidth, b.UI().SubGridColor)
                mut.Lock()
                b.cellImg[pos] = img
                mut.Unlock()
            }(pos, g, &mut, &wg)
        }
        wg.Wait()
    }
    
    for pos, img := range b.cellImg {
        if toDraw[pos] {
            img = b.GridAt(pos).Draw(b.cellImg[pos], b.UI().SubStrokeWidth, b.UI().SubGridColor)
            b.cellImg[pos] = img
            toDraw[pos] = false
        }
        x, y := pos.GetXY()
        opts := &ebiten.DrawImageOptions{}
        opts.GeoM.Translate(float64(x)*colW, float64(y)*rowH)
        _ = boardImg.DrawImage(img, opts)
    }
    
    _ = boardImg.DrawImage(gridImg, &ebiten.DrawImageOptions{})
    _ = screen.DrawImage(boardImg, &ebiten.DrawImageOptions{})
}

func (b *board) ResetGrid(wi int, hi int) {
    img := NewImage(wi, hi)
    w, h := float64(wi), float64(hi)
    colW, rowH := w/3, h/3
    for x := colW; x < (w - (colW / 2)); x += colW {
        DrawLine(img, V(x, 0), V(x, h), internal.ScaleWidth(b.UI().StrokeWidth), b.UI().GridColor)
    }
    for y := rowH; y < (h - (rowH / 2)); y += rowH {
        DrawLine(img, V(0, y), V(w, y), internal.ScaleHeight(b.UI().StrokeWidth), b.UI().GridColor)
    }
    gridImg, _ = ebiten.NewImageFromImage(img.SubImage(img.Rect), ebiten.FilterDefault)
}

func (b *board) DrawTile(pos Position, subPos Position) {
    b.GridAt(pos).DrawTile(b.cellImg[pos], subPos)
    b.SetGridToDraw(pos)
}

func (b *board) DrawTileUnderCursor() {
    tile, pos := b.GetTileUnderCursor()
    b.DrawTile(pos, tile.Position)
    b.SetGridToDraw(pos)
}

func (b *board) SetGridToDraw(pos Position) {
    toDraw[pos] = true
}
