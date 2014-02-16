package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_version.h>
import "C"

var (
	Major    uint8
	Minor    uint8
	Patch    uint8
	Revision int
)

func init() {
	sv := C.SDL_version{}
	C.SDL_GetVersion(&sv)
	Major = uint8(sv.major)
	Minor = uint8(sv.minor)
	Patch = uint8(sv.patch)
	Revision = int(C.SDL_GetRevisionNumber())
}
