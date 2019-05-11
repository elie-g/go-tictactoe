package events

import (
    "sync"
)

type ClickListener interface {
    Pause()
    Resume()
    IsRunning() bool
    Destroy()
    IsDestroyed() bool
    Listen(cb ...func())
}

func NewClickListener() ClickListener {
    var mut sync.Mutex
    stopChan := make(chan bool)
    listener := &clickListener{
        running:   false,
        destroyed: false,
        stopChan:  stopChan,
        mut:       mut,
        clickChan: MouseClick(stopChan),
        callbacks: []func(){},
    }
    go listener.init()
    return listener
}

type clickListener struct {
    running   bool
    destroyed bool
    stopChan  chan bool
    mut       sync.Mutex
    clickChan <-chan bool
    callbacks []func()
}

// Private
func (listener *clickListener) init() {
    for !listener.destroyed {
        if listener.running {
            select {
            case <-listener.clickChan:
                for _, cb := range listener.callbacks {
                    cb()
                }
            case <-listener.stopChan:
            default:
            }
        } else {
            <-listener.stopChan
            listener.clickChan = MouseClick(listener.stopChan)
        }
    }
}

func (listener *clickListener) Pause() {
    listener.running = false
    listener.stopChan <- !listener.running
}

func (listener *clickListener) Resume() {
    listener.running = !listener.destroyed
    listener.stopChan <- !listener.running
}

func (listener *clickListener) IsRunning() bool {
    return listener.running
}

func (listener *clickListener) Listen(cb ...func()) {
    listener.mut.Lock()
    listener.callbacks = append(listener.callbacks, cb...)
    listener.mut.Unlock()
}

func (listener *clickListener) Destroy() {
    listener.running = false
    listener.destroyed = true
    listener.stopChan <- !listener.running
}

func (listener *clickListener) IsDestroyed() bool {
    return listener.destroyed
}
