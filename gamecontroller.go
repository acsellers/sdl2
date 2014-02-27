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

func (gc *GameController) ButtonStatus(gcb GameControllerButton) uint8 {
	return uint8(C.SDL_GameControllerGetButton(gc.Native, gcb.Native()))
}

func (gc *GameController) AxisStatus(gca GameControllerAxis) int16 {
	return int16(C.SDL_GameControllerGetAxis(gc.Native, gca.Native()))
}

type GameControllerButton int

func (gcb GameControllerButton) Native() C.SDL_GameControllerButton {
	return C.SDL_GameControllerButton(gcb)
}

const (
	ButtonA GameControllerButton = iota
	ButtonB
	ButtonX
	ButtonY
	ButtonBack
	ButtonGuide
	ButtonStart
	ButtonLeftStick
	ButtonRightStick
	ButtonLeftShoulder
	ButtonRightShoulder
	ButtonDpadUp
	ButtonDpadDown
	ButtonDpadLeft
	ButtonDpadRight
	ButtonInvalid GameControllerButton = -1
)

type GameControllerAxis int

func GameControllerAxisFromString(axis string) GameControllerAxis {
	cstr := C.CString(axis)
	defer C.free(unsafe.Pointer(cstr))
	return GameControllerAxis(C.SDL_GameControllerGetAxisFromString(cstr))
}

const (
	Invalid      GameControllerAxis = -1
	LeftX        GameControllerAxis = C.SDL_CONTROLLER_AXIS_LEFTX
	LeftY        GameControllerAxis = C.SDL_CONTROLLER_AXIS_LEFTY
	RightX       GameControllerAxis = C.SDL_CONTROLLER_AXIS_RIGHTX
	RightY       GameControllerAxis = C.SDL_CONTROLLER_AXIS_RIGHTY
	TriggerLeft  GameControllerAxis = C.SDL_CONTROLLER_AXIS_TRIGGERLEFT
	TriggerRight GameControllerAxis = C.SDL_CONTROLLER_AXIS_TRIGGERRIGHT
)

func (gca GameControllerAxis) Native() C.SDL_GameControllerAxis {
	return C.SDL_GameControllerAxis(gca)
}
func (gca GameControllerAxis) String() string {
	return C.GoString(C.SDL_GameControllerGetStringForAxis(gca.Native()))
}
