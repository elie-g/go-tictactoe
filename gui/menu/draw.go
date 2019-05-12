package menu

import (
    . "github.com/DrunkenPoney/go-tictactoe/gui/menu/btn"
    . "github.com/DrunkenPoney/go-tictactoe/gui/menu/pages"
    "github.com/DrunkenPoney/go-tictactoe/internal"
    "github.com/DrunkenPoney/go-tictactoe/settings/colors"
    . "github.com/DrunkenPoney/go-tictactoe/settings/messages"
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
        btnX := (wMenu - btnWidth) / 2
        
        if font == nil {
            tt, _ := truetype.Parse(fonts.ArcadeN_ttf)
            font = truetype.NewFace(tt, &truetype.Options{
                Size:    fntSize,
                DPI:     72,
                Hinting: f.HintingFull})
        }
        
        menuImg, _ = ebiten.NewImage(wi, hi, ebiten.FilterDefault)
        _ = menuImg.Fill(colors.Colors().MenuFadeBackground())
        
        img, _ := ebiten.NewImage(int(wMenu), int(hMenu), ebiten.FilterDefault)
        _ = img.Fill(colors.Colors().MenuBackground())
        
        var printBtn = func(btnType ButtonType, order int, nbButtons float64, btnText Message) {
            clr := colors.Colors().MenuButtonBackground()
            txtClr := colors.Colors().MenuButtonColor()
            btnY := hMenu / (nbButtons+1)
            btn2Y := btnY*float64(order) - (btnHeight / 2)
            action := m.actions[btnType]
            if !m.IsFrozen() && cX > menuX+btnX && cY > menuY+btn2Y &&
                cX < menuX+btnX+btnWidth && cY < menuY+btn2Y+btnHeight {
                btnType |= BTN_HOVER
                clr = colors.Colors().MenuButtonHoverBackground()
                txtClr = colors.Colors().InGameTextColor()
                if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && action != nil {
                    action()
                }
            }
            
            if btnImgs[btnType] == nil {
                btnImgs[btnType], _ = ebiten.NewImage(int(btnWidth), int(btnHeight), ebiten.FilterDefault)
                _ = btnImgs[btnType].Fill(clr)
                x := (btnWidth / 2) - (fntSize * float64(len(btnText.Str())) / 2)
                y := (btnHeight-fntSize)/2 + btnHeight/2
                text.Draw(btnImgs[btnType], btnText.Str(), font, int(x), int(y), txtClr)
            }
            
            opts := &ebiten.DrawImageOptions{}
            opts.GeoM.Translate(btnX, btn2Y)
            _ = img.DrawImage(btnImgs[btnType], opts)
        }
        
        // AJOUTER LES BOUTONS ICI
        switch m.GetPage() {
        case MP_ONLINE:
            printBtn(BTN_JOIN_ONLINE, 1, 3, MSG_BTN_JOIN_ONLINE)
            printBtn(BTN_CREATE_ONLINE, 2, 3, MSG_BTN_CREATE_ONLINE)
            printBtn(BTN_RETURN, 3, 3, MSG_BTN_RETURN)
        default:
            printBtn(BTN_RESUME, 1, 5, MSG_BTN_RESUME_GAME)
            printBtn(BTN_PLAYER_1, 2, 5, MSG_BTN_PLAYER_1)
            printBtn(BTN_PLAYER_2, 3, 5, MSG_BTN_PLAYER_2)
            printBtn(BTN_ONLINE, 4, 5, MSG_BTN_ONLINE)
            printBtn(BTN_EXIT, 5, 5, MSG_BTN_EXIT_GAME)
        }
        
        // //////////////////////////////////////////////////////////////////////////////////
        opts := &ebiten.DrawImageOptions{}
        opts.GeoM.Translate(menuX, menuY)
        _ = menuImg.DrawImage(img, opts)
        
        _ = screen.DrawImage(menuImg, &ebiten.DrawImageOptions{})
    }
}
