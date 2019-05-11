package btn

type ButtonType uint8

//noinspection GoSnakeCaseUsage
const (
    BTN_RESUME  ButtonType = 0x01
    BTN_RESTART ButtonType = 0x02
    BTN_EXIT    ButtonType = 0x04
)

func (b ButtonType) Is(btn ButtonType) bool {
    return b&btn > 0
}