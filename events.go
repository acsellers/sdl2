package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_events.h>
import "C"

import (
	"bytes"
	"encoding/binary"
	"fmt"
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
