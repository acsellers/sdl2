package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_gamecontroller.h>
import "C"
import

/*
Unimplemented:
SDL_GameControllerAddMapping
SDL_GameControllerMappingForGUID
SDL_GameControllerGetJoystick // needed?
SDL_GameControllerUpdate
SDL_GameControllerAxis // enum
SDL_GameControllerGetAxisFromString
SDL_GameControllerGetStringForAxis
SDL_GameControllerGetBindForAxis
SDL_GameControllerGetAxis
SDL_GameControllerButton // enum
SDL_GameControllerGetButtonFromString
SDL_GameControllerGetStringForButton
SDL_GameControllerGetBindForButton
SDL_GameControllerGetButton
*/"unsafe"

type GameController struct {
	Native *C.SDL_GameController
	Parent *Joystick
}

func (gc *GameController) Mapping() string {
	cstr := C.SDL_GameControllerMapping(gc.Native)
	if cstr == nil {
		return ""
	}
	defer C.SDL_free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

func (gc *GameController) String() string {
	return C.GoString(C.SDL_GameControllerName(gc.Native))
}

func (gc *GameController) Attached() bool {
	return C.SDL_GameControllerGetAttached(gc.Native) == C.SDL_TRUE
}

func (gc *GameController) EmittingEvents() bool {
	return C.SDL_GameControllerEventState(C.int(-1)) == C.SDL_TRUE
}

func (gc *GameController) SetEvents(active bool) {
	var flag C.int
	if active {
		flag = 1
	}
	C.SDL_GameControllerEventState(flag)
}

func (gc *GameController) Close() {
	C.SDL_GameControllerClose(gc.Native)
}
