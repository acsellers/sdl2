package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_keyboard.h>
import "C"
import "image"

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

func StartTextInput() {
	C.SDL_StartTextInput()
}
func StopTextInput() {
	C.SDL_StopTextInput()
}
func TextInputArea(r image.Rectangle) {
	r.Canon()
	C.SDL_SetTextInputRect(&SDL_Rect{
		x: r.Min.X,
		y: r.Min.Y,
		w: r.Dx(),
		h: r.Dy(),
	})
}
func AcceptingTextInput() bool {
	return C.SDL_IsTextInputActive() == C.SDL_TRUE
}

func OnscreenKeyboardAvailable() bool {
	return C.SDL_HasScreenKeyboardSupport() == C.SDL_TRUE
}
