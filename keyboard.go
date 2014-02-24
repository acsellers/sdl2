package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_keyboard.h>
import "C"

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
