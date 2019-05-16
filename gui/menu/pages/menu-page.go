package pages

type MenuPage uint8

const (
    MP_MAIN   MenuPage = 0x00 // Default
    MP_ONLINE MenuPage = 0x01
)

func (mp MenuPage) Is(menuPage MenuPage) bool {
    return mp&menuPage > 0
}
