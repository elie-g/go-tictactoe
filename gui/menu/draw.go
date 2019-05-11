package menu

import (
    . "github.com/DrunkenPoney/go-tictactoe/gui/colors"
    . "github.com/DrunkenPoney/go-tictactoe/gui/menu/btn"
    . "github.com/DrunkenPoney/go-tictactoe/gui/message"
    "github.com/DrunkenPoney/go-tictactoe/internal"
    "github.com/golang/freetype/truetype"
    "github.com/hajimehoshi/ebiten"
    "github.com/hajimehoshi/ebiten/examples/resources/fonts"
    "github.com/hajimehoshi/ebiten/text"
    f "golang.org/x/image/font"
)

var menuImg *ebiten.Image
var fntSize = internal.Scale(32)
var font f.Face
var btnImgs = make(map[ButtonType]*ebiten.Image)

func (m *menu) Draw(screen *ebiten.Image) {
    if m.IsShown() {
        wi, hi := screen.Size()
        w, h := float64(wi), float64(hi)
        hMenu := 0.9 * h
        wMenu := 0.6 * hMenu
        menuX, menuY := (w-wMenu)/2, (h-hMenu)/2
        btnWidth, btnHeight := 0.8*wMenu, fntSize*2
        ciX, ciY := ebiten.CursorPosition()
        cX, cY := float64(ciX), float64(ciY)
        
        if font == nil {
            tt, _ := truetype.Parse(fonts.ArcadeN_ttf)
            font = truetype.NewFace(tt, &truetype.Options{
                Size:    fntSize,
                DPI:     72,
                Hinting: f.HintingFull})
        }
        
        menuImg, _ = ebiten.NewImage(wi, hi, ebiten.FilterDefault)
        _ = menuImg.Fill(Colors().MenuFadeBackground())
        
        img, _ := ebiten.NewImage(int(wMenu), int(hMenu), ebiten.FilterDefault)
        _ = img.Fill(Colors().MenuBackground())
        
        //////////////////////////////////////////////////////////////// BTN_RESUME
        btnType := BTN_RESUME
        clr := Colors().MenuButtonBackground()
        txtClr := Colors().MenuButtonColor()
        btnX, btnY := (wMenu-btnWidth)/2, hMenu/7
        if cX > menuX+btnX && cY > menuY+btnY &&
            cX < menuX+btnX+btnWidth && cY < menuY+btnY+btnHeight {
            btnType |= BTN_HOVER
            clr = Colors().MenuButtonHoverBackground()
            txtClr = Colors().InGameTextColor()
            if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && m.resumeCb != nil {
                m.resumeCb()
            }
        }
        
        if btnImgs[btnType] == nil {
            btnImgs[btnType], _ = ebiten.NewImage(int(btnWidth), int(btnHeight), ebiten.FilterDefault)
            _ = btnImgs[btnType].Fill(clr)
            x := (btnWidth / 2) - (fntSize * float64(len(MSG_RESUME_GAME.Str())) / 2)
            y := (btnHeight-fntSize)/2 + btnHeight/2
            text.Draw(btnImgs[btnType], MSG_RESUME_GAME.Str(), font, int(x), int(y), txtClr)
        }
        
        opts := &ebiten.DrawImageOptions{}
        opts.GeoM.Translate(btnX, btnY)
        _ = img.DrawImage(btnImgs[btnType], opts)
        
        // ////////////////////////////////////////////////////////// BTN_EXIT
        btnType = BTN_EXIT
        clr = Colors().MenuButtonBackground()
        txtClr = Colors().MenuButtonColor()
        btnY = btnY * 3
        if cX > menuX+btnX && cY > menuY+btnY &&
            cX < menuX+btnX+btnWidth && cY < menuY+btnY+btnHeight {
            btnType |= BTN_HOVER
            clr = Colors().MenuButtonHoverBackground()
            txtClr = Colors().InGameTextColor()
            if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && m.exitCb != nil {
                m.exitCb()
            }
        }
        
        if btnImgs[btnType] == nil {
            btnImgs[btnType], _ = ebiten.NewImage(int(btnWidth), int(btnHeight), ebiten.FilterDefault)
            _ = btnImgs[btnType].Fill(clr)
            x := (btnWidth / 2) - (fntSize * float64(len(MSG_EXIT_GAME.Str())) / 2)
            y := (btnHeight-fntSize)/2 + btnHeight/2
            text.Draw(btnImgs[btnType], MSG_EXIT_GAME.Str(), font, int(x), int(y), txtClr)
        }
        
        opts = &ebiten.DrawImageOptions{}
        opts.GeoM.Translate(btnX, btnY)
        _ = img.DrawImage(btnImgs[btnType], opts)
        
        ////////////////////////////////////////////////////////////////////////////////////
        opts.GeoM.Reset()
        opts.GeoM.Translate(menuX, menuY)
        _ = menuImg.DrawImage(img, opts)
        
        _ = screen.DrawImage(menuImg, &ebiten.DrawImageOptions{})
    }
}
