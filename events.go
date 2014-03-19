package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_events.h>
import "C"

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
)

type Event interface {
	Type() EventType
	Since() time.Duration
	String() string
}

type EventType int

const (
	EventUnknown EventType = iota
	EventController
	EventJoystick
	EventKeyboard
	EventMouse
	EventWindow
	EventTick
)

type Keysym struct {
	Scancode
	Keycode
	KeyModifiers
}

type ControllerEventType uint8

const (
	CETUnknown ControllerEventType = iota
	ButtonPressed
	ButtonReleased
	AxisMotion
)

type ControllerEvent struct {
	ControllerEventType
	Button GameControllerButton
	Axis   GameControllerAxis
	Value  uint16
}

var (
	GameControllers map[int]<-chan ControllerEvent
)

func WatchEvents() {
	go pollEvents()
}

func pollEvents() {
	fmt.Println("Polling for Events")
	var event C.SDL_Event
	var e C.int
	for {
		e = C.SDL_WaitEventTimeout(&event, C.int(5))
		if e == 1 {
			buf := bytes.NewReader(event[0:4])
			var i uint32
			binary.Read(buf, binary.LittleEndian, &i)
			fmt.Printf("Event type: %x\n", i)
		}
	}
}

var (
	EventTimeout = 1000 // Milliseconds
	EventSlop    = 100  // Milliseconds
)
var (
	EventBus       chan Event
	internalEvents chan Event
)

func EventPoller() {
	go pollEvents()
	if EventBus == nil {
		EventBus = make(chan Event)
		internalEvents = make(chan Event)
	} else {
		return
	}

	queue := make([]Event, 0, 16)
	var ev Event
	var d time.Duration
	var timeout <-chan time.Time
EventLoop:
	for {
		if len(queue) == 0 {
			select {
			case <-time.After(time.Duration(EventTimeout) * time.Millisecond):
				queue = append(queue, NewTickEvent())
			case ev = <-internalEvents:
				queue = append(queue, ev)
			}
		}

		d = time.Duration(EventTimeout+EventSlop)*time.Millisecond - queue[0].Since()
		timeout = time.After(d)
		select {
		case EventBus <- queue[0]:
			queue = queue[1:]
		case ev = <-internalEvents:
			queue = append(queue, ev)
		case <-timeout:
			for i, qe := range queue {
				if qe.Since() < time.Duration(EventTimeout)*time.Millisecond {
					queue = queue[i:]
					continue EventLoop
				}
			}
		}
	}
}

type TickEvent struct {
	generated time.Time
}

func (te TickEvent) Type() EventType {
	return EventTick
}

func (te TickEvent) String() string {
	return fmt.Sprintf("Tick Event generated %s ago", time.Since(te.generated))
}

func (te TickEvent) Since() time.Duration {
	return time.Since(te.generated)
}

func NewTickEvent() Event {
	return TickEvent{
		generated: time.Now(),
	}
}
