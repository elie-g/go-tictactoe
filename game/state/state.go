package state

type State uint8

const (
    STOPPED State = 0x00
    RUNNING State = 0x01
    PAUSED  State = 0x02
)
