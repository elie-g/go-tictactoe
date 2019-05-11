package settings

import (
    "github.com/DrunkenPoney/go-tictactoe/grid/tile"
    "time"
)

//noinspection GoUnusedConst,GoSnakeCaseUsage
const (
    // AI Settings
    REFERENCE_PLAYER = tile.X
    DEFAULT_PREDICTION_DEPTH = 5
    
    // Other Settings
    MIN_KEYPRESS_DELAY = 220 * time.Millisecond
    MAX_PLAYER_NAME_LENGTH = 15
)
