package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_version.h>
// #include <SDL2/SDL_platform.h>
import "C"

var (
	// Platform is the generic name of the current computer's OS,
	// Ex: Linux. Windows
	Platform string
	// Major verion of SDL, should be 2
	Major uint8
	// Minor version of SDL2
	Minor uint8
	// Patch level of SDL2
	Patch uint8
	// Revision is an incrementing number from SDL2's mercurial repository
	Revision int
)

func init() {
	sv := C.SDL_version{}
	C.SDL_GetVersion(&sv)
	Major = uint8(sv.major)
	Minor = uint8(sv.minor)
	Patch = uint8(sv.patch)
	Revision = int(C.SDL_GetRevisionNumber())

	cstr := C.SDL_GetPlatform()
	Platform = C.GoString(cstr)
}
