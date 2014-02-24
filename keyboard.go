package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_keyboard.h>
import "C"
import (
	"image"
	"reflect"
	"unsafe"
)

func CurrentKeyboardWindow() *Window {
	return &Window{
		Native: C.SDL_GetKeyboardFocus(),
	}
}

type KeyModifiers struct {
	LShift, RShift bool
	LCtrl, RCtrl   bool
	LAlt, RAlt     bool
	LGui, RGui     bool
	Numlock        bool
	Capslock       bool
}

func (km KeyModifiers) Shift() bool {
	return km.LShift || km.RShift
}
func (km KeyModifiers) Ctrl() bool {
	return km.LCtrl || km.RCtrl
}
func (km KeyModifiers) Alt() bool {
	return km.LAlt || km.RAlt
}
func (km KeyModifiers) Gui() bool {
	return km.LGui || km.RGui
}

func PressedKeys() []Keycode {
	// C Array to slice from https://code.google.com/p/go-wiki/wiki/cgo
	var stateLength int32
	var keyStates *C.Uint8 = C.SDL_GetKeyboardState((*C.int)(&stateLength))
	var stateSlice []C.Uint8
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&stateSlice)))
	sliceHeader.Cap = int(stateLength)
	sliceHeader.Len = int(stateLength)
	sliceHeader.Data = uintptr(unsafe.Pointer(&keyStates))

	keys := make([]Keycode, 0)
	for i, v := range stateSlice {
		if v != 0x0 {
			keys = append(keys, Scancode(i).Key())
		}
	}
	return keys
}

func Modifiers() KeyModifiers {
	var km KeyModifiers
	c := C.SDL_GetModState()
	if c == 0x0 {
		return km
	}
	if c&C.KMOD_LSHIFT != 0x0 {
		km.LShift = true
	}
	if c&C.KMOD_RSHIFT != 0x0 {
		km.RShift = true
	}
	if c&C.KMOD_LCTRL != 0x0 {
		km.LCtrl = true
	}
	if c&C.KMOD_RCTRL != 0x0 {
		km.RCtrl = true
	}
	if c&C.KMOD_LALT != 0x0 {
		km.LAlt = true
	}
	if c&C.KMOD_RALT != 0x0 {
		km.RAlt = true
	}
	if c&C.KMOD_LGUI != 0x0 {
		km.LGui = true
	}
	if c&C.KMOD_RGUI != 0x0 {
		km.RGui = true
	}
	if c&C.KMOD_NUM != 0x0 {
		km.Numlock = true
	}
	if c&C.KMOD_CAPS != 0x0 {
		km.Capslock = true
	}
	return km
}

func SetModifiers(km KeyModifiers) {
	var c C.SDL_Keymod
	if km.LShift {
		c &= C.KMOD_LSHIFT
	}
	if km.RShift {
		c &= C.KMOD_RSHIFT
	}
	if km.LCtrl {
		c &= C.KMOD_LCTRL
	}
	if km.RCtrl {
		c &= C.KMOD_RCTRL
	}
	if km.LAlt {
		c &= C.KMOD_LALT
	}
	if km.RAlt {
		c &= C.KMOD_RALT
	}
	if km.LGui {
		c &= C.KMOD_LGUI
	}
	if km.RGui {
		c &= C.KMOD_RGUI
	}
	if km.Numlock {
		c &= C.KMOD_NUM
	}
	if km.Capslock {
		c &= C.KMOD_CAPS
	}

	C.SDL_SetModState(c)
}

func StartTextInput() {
	C.SDL_StartTextInput()
}
func StopTextInput() {
	C.SDL_StopTextInput()
}
func TextInputArea(r image.Rectangle) {
	r.Canon()
	C.SDL_SetTextInputRect(&C.SDL_Rect{
		x: C.int(r.Min.X),
		y: C.int(r.Min.Y),
		w: C.int(r.Dx()),
		h: C.int(r.Dy()),
	})
}
func AcceptingTextInput() bool {
	return C.SDL_IsTextInputActive() == C.SDL_TRUE
}

func OnscreenKeyboardAvailable() bool {
	return C.SDL_HasScreenKeyboardSupport() == C.SDL_TRUE
}
