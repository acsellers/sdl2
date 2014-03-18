package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_error.h>
// #include <SDL2/SDL.h>
import "C"
import (
	"fmt"
	"image"
)

func init() {
	C.SDL_Init(C.SDL_INIT_VIDEO)
}

func GetError() error {
	if cstr := C.SDL_GetError(); cstr != nil {
		e := C.GoString(cstr)
		C.SDL_ClearError()
		return fmt.Errorf(e)
	}
	return nil
}

func Quit() {
	C.SDL_Quit()
}

func RectToNative(r image.Rectangle) *C.SDL_Rect {
	return &C.SDL_Rect{
		x: C.int(r.Min.X),
		y: C.int(r.Min.Y),
		w: C.int(r.Dx()),
		h: C.int(r.Dy()),
	}
}
