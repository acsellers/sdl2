package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_mouse.h>
import "C"
import (
	"image"
	"runtime"
)

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

func WarpMouse(w *Window, pt image.Point) {
	C.SDL_WarpMouseInWindow(w.Native, C.int(pt.X), C.int(pt.Y))
}

// RelativeMouseState retrieves the position change of the mouse,
// and the state of the mouse buttons as well. The change period
// is the time since the last call to RelativeMouseState.
func RelativeMouseState() MouseInfo {
	ms := MouseInfo{}
	btn := C.SDL_GetRelativeMouseState((*C.int)(&ms.X), (*C.int)(&ms.Y))
	ms.parseButtons(uint32(btn))
	return ms
}

// SetRelativeMode changes the mouse mode from absolute to relative based
// on whether the activate argument is false. When RelativeMode is active,
// then the cursor is hidden and the mouse position will not change, only
// motion reported will be through RelativeMouseState.
func SetRelativeMode(activate bool) error {
	var b C.SDL_bool
	if activate {
		b = C.SDL_TRUE
	}
	r := C.SDL_SetRelativeMouseMode(b)
	if r != 0 {
		return GetError()
	}
	return nil
}

func RelativeMode() bool {
	if C.SDL_GetRelativeMouseMode() == C.SDL_TRUE {
		return true
	}
	return false
}

func CursorVisibility() bool {
	r := C.SDL_ShowCursor(C.SDL_QUERY)
	if r == SDL_TRUE {
		return true
	}
	return false
}

func SetCursorVisiblility(active bool) {
	if active {
		C.SDL_ShowCursor(C.SDL_TRUE)
	} else {
		C.SDL_ShowCursor(C.SDL_FALSE)
	}
}

func NewSystemCursor(sc SystemCursor) *Cursor {
	c := &Cursor{
		Native: C.SDL_CreateSystemCursor(C.SDL_SystemCursor(sc)),
	}
	c.setupFree()
	return c
}

func NewSurfaceCursor(s *Surface, hot image.Point) *Cursor {
	c := &Cursor{
		Native: C.SDL_CreateColorCursor(
			s.Native,
			C.int(hot.X),
			C.int(hot.Y),
		),
	}
	c.setupFree()
	return c
}

func DefaultCursor() *Cursor {
	c := &Cursor{
		Native: C.SDL_GetDefaultCursor(),
	}
	c.setupFree()
	return c
}

func CurrentCursor() *Cursor {
	return &Cursor{Native: C.SDL_GetCursor()}
}

// CreateMonoCursor creates a new cursor using the complicated method.
// Img and mask make up the cursor image according to the following method:
// | Img | Mask | Result
// |  0  |  0   | Tranparent
// |  1  |  0   | Inverted color or black
// |  0  |  1   | White
// |  1  |  1   | Black
// Note that the width must be a multiple of 8, but height doesn't have such
// a restriction. The hot image.Point is the point within the data's image
// that should be regarded as the active point.
func CreateMonoCursor(img, mask []byte, width, height int, hot image.Point) (*Cursor, error) {
	c := C.SDL_CreateCursor(
		(*C.Uint8)(&img[0]),
		(*C.Uint8)(&mask[0]),
		C.int(width),
		C.int(height),
		C.int(hot.X),
		C.int(hot.Y),
	)
	if c == nil {
		return nil, GetError()
	}
	return &Cursor{Native: c}, nil
}

type Cursor struct {
	Native *C.SDL_Cursor
}

func (c *Cursor) setupFree() {
	runtime.SetFinalizer(c, (*Cursor).doFree)
}

func (c *Cursor) SetActive() {
	C.SDL_SetCursor(c.Native)
}

func (c *Cursor) Free() {
	runtime.SetFinalizer(c, nil)
	c.doFree()
}

func (c *Cursor) doFree() {
	C.SDL_FreeCursor(c.Native)
}

func (c *Cursor) DisableFree() {
	runtime.SetFinalizer(c, nil)
}
