package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_error.h>
// #include <SDL2/SDL.h>
import "C"
import "fmt"

func init() {
	C.SDL_Init(C.SDL_INIT_EVERYTHING)
}

func GetError() error {
	if cstr := C.SDL_GetError(); cstr != nil {
		C.SDL_ClearError()
		return fmt.Errorf(C.GoString(cstr))
	}
	return nil
}
