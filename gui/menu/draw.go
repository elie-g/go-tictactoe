package menu

import (
    . "github.com/DrunkenPoney/go-tictactoe/gui/colors"
    "github.com/DrunkenPoney/go-tictactoe/gui/menu/btn"
    "github.com/hajimehoshi/ebiten"
)

var menuImgs = make(map[uint8]*ebiten.Image)

func (m *menu) Draw(screen *ebiten.Image) {
    wi, hi := screen.Size()
    w, h := float64(wi), float64(hi)
    
    id := m.getBtnGroupId()
    if menuImgs[id] == nil {
        menuImgs[id], _ = ebiten.NewImage(wi, hi, ebiten.FilterDefault)
        _ = menuImgs[id].Fill(Colors().MenuFadeBackground())
        
        hM := 0.9 * h
        wM := 0.6 * hM
        img, _ := ebiten.NewImage(int(wM), int(hM), ebiten.FilterDefault)
        _ = img.Fill(Colors().MenuBackground())
        
        opts := &ebiten.DrawImageOptions{}
        
        opts.GeoM.Reset()
        opts.GeoM.Translate((w-wM)/2, (h - hM)/2)
        _ = menuImgs[id].DrawImage(img, opts)
    }
    
    if m.IsShown() {
        _ = screen.DrawImage(menuImgs[id], &ebiten.DrawImageOptions{})
    }
}

func (m *menu) getBtnGroupId() uint8 {
    var id uint8 = 0
    if m.resumeCb != nil {
        id |= uint8(btn.BTN_RESUME)
    }
    if m.restartCb != nil {
        id |= uint8(btn.BTN_RESTART)
    }
    if m.exitCb != nil {
        id |= uint8(btn.BTN_EXIT)
    }
    return id
}
