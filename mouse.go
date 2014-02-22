package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_mouse.h>
import "C"

type SystemCursor int

const (
	Arrow SystemCursor = iota
	IBeam
	Wait
	Crosshair
	WaitArrow
	SizeNWSE
	SizeNESW
	SizeNS
	SizeAll
	No
	Hand
)

func MouseFocus() *Window {
	return &Window{Native: C.SDL_GetMouseFocus()}
}

type MouseInfo struct {
	X, Y                int32
	Left, Middle, Right bool
	X1, X2              bool
}

func (ms *MouseInfo) parseButtons(b uint32) {
	ms.Left = b&0x01 > 0
	ms.Middle = b&0x02 > 0
	ms.Right = b&0x04 > 0
	ms.X1 = b&0x08 > 0
	ms.X2 = b&0x10 > 0
}

func MouseState() MouseInfo {
	ms := MouseInfo{}
	btn := C.SDL_GetMouseState((*C.int)(&ms.X), (*C.int)(&ms.Y))
	ms.parseButtons(uint32(btn))
	return ms
}
