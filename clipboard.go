package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_clipboard.h>
import "C"
import "unsafe"

func ClipboardText() string {
	if C.SDL_HasClipboardText() == C.SDL_TRUE {
		cstr := C.SDL_GetClipboardText()
		return C.GoString(cstr)
	}
	return ""
}

func SetClipboardText(text string) error {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	r := C.SDL_SetClipboardText(cstr)
	if int(r) == -1 {
		return GetError()
	}
	return nil
}
