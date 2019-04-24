package tile

import (
	"testing"
)

func TestValue(t *testing.T) {
	var tile Tile = NewTile(EMPTY)
	tile.SetValue(100)
	if tile.GetValue() != 100 {
		t.Errorf("value incorect")
	}
}

func TestIsActive(t *testing.T) {
	var tile Tile = NewTile(EMPTY)
	tile.SetActive(false)
	if tile.IsActive() != false {
		t.Errorf("active incorect")
	}
}

func TestIsWinning(t *testing.T) {
	var tile Tile = NewTile(EMPTY)
	tile.SetWinning(false)
	if tile.IsActive() != false {
		t.Errorf("winning incorect")
	}
}
