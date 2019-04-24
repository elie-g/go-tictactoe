package events

import (
	"github.com/go-vgo/robotgo"
)

func MouseClick() <-chan bool {
	ch := make(chan bool)
	go func(ch chan bool) {
		ch <- robotgo.AddEvent("mleft")
	}(ch)
	return ch
}
