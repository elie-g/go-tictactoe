package events

import (
	"github.com/go-vgo/robotgo"
)

func MouseClick(cancel <- chan bool) <-chan bool {
	ch := make(chan bool)
	go func(ch chan bool) {
		for {
			select {
			case <- cancel:
				return
			default:
				ch <- robotgo.AddEvent("mleft")
			}
		}
	}(ch)
	return ch
}
