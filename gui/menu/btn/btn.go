package btn

type ButtonType uint8

const (
    BTN_HOVER    ButtonType = 0x80
    BTN_RESUME   ButtonType = 0x01
    BTN_EXIT     ButtonType = 0x02
    BTN_PLAYER_1 ButtonType = 0x04
    BTN_PLAYER_2 ButtonType = 0x10
    // BTN_         ButtonType = 0x20
    // BTN_         ButtonType = 0x40
)

func (btn ButtonType) Is(btnType ButtonType) bool {
    return btn&btnType > 0
}
