package btn

type ButtonType uint16

const (
    BTN_RESUME        ButtonType = 0x0001
    BTN_EXIT          ButtonType = 0x0002
    BTN_PLAYER_1      ButtonType = 0x0004
    BTN_PLAYER_2      ButtonType = 0x0010
    BTN_ONLINE        ButtonType = 0x0020
    BTN_JOIN_ONLINE   ButtonType = 0x0040
    BTN_CREATE_ONLINE ButtonType = 0x0080
    BTN_RETURN        ButtonType = 0x0100
    // BTN_              ButtonType = 0x0200
    // BTN_              ButtonType = 0x0400
    // BTN_              ButtonType = 0x0800
    // BTN_              ButtonType = 0x1000
    // BTN_              ButtonType = 0x2000
    // BTN_              ButtonType = 0x4000
    BTN_HOVER ButtonType = 0x8000
)

func (btn ButtonType) Is(btnType ButtonType) bool {
    return btn&btnType > 0
}
