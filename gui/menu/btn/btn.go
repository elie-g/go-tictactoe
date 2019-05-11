package btn

type ButtonType uint8

const (
    BTN_RESUME ButtonType = 0x01
    BTN_EXIT   ButtonType = 0x02
    BTN_HOVER  ButtonType = 0x04
)
