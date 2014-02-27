package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_joystick.h>
// #include <SDL2/SDL_gamecontroller.h>
import "C"
import

/*
Unimplemented:
GetGUIDString/FromString // WTF Hungarian
*/"fmt"

func ConnectedJoysticks() []*Joystick {
	n := int(C.SDL_NumJoysticks())
	joys := make([]*Joystick, n)
	for i, _ := range joys {
		joys[i] = &Joystick{
			Name:           C.GoString(C.SDL_JoystickNameForIndex(C.int(i))),
			guid:           C.SDL_JoystickGetDeviceGUID(C.int(i)),
			index:          i,
			GameController: C.SDL_IsGameController(C.int(i)) == C.SDL_TRUE,
		}
		if joys[i].GameController {
			cstr := C.SDL_GameControllerNameForIndex(C.int(i))
			if cstr != nil {
				joys[i].ControllerName = C.GoString(cstr)
			}
		}
	}
	return joys
}

type Joystick struct {
	Native         *C.SDL_Joystick
	guid           C.SDL_JoystickGUID
	index          int
	Name           string
	ControllerName string
	GameController bool
	Axes           int
	Trackballs     int
	POVHats        int
	Buttons        int
}

type HatPosition uint8

const (
	HatCentered  HatPosition = 0x0
	HatUp        HatPosition = 0x1
	HatRight     HatPosition = 0x2
	HatDown      HatPosition = 0x4
	HatLeft      HatPosition = 0x8
	HatRightUp   HatPosition = HatRight | HatUp
	HatRightDown HatPosition = HatRight | HatDown
	HatLeftUp    HatPosition = HatLeft | HatUp
	HatLeftDown  HatPosition = HatLeft | HatDown
)

func (j *Joystick) Open() error {
	j.Native = C.SDL_JoystickOpen(C.int(j.index))
	if j.Native == nil {
		return fmt.Errorf("Joystick Open Error:%v", GetError())
	}

	j.Axes = int(C.SDL_JoystickNumAxes(j.Native))
	j.Trackballs = int(C.SDL_JoystickNumBalls(j.Native))
	j.POVHats = int(C.SDL_JoystickNumHats(j.Native))
	j.Buttons = int(C.SDL_JoystickNumButtons(j.Native))

	return nil
}

func (j *Joystick) OpenController() (*GameController, error) {
	gc := C.SDL_GameControllerOpen(C.int(j.index))
	if gc == nil {
		return nil, GetError()
	}

	return &GameController{Native: gc, Parent: j}, nil
}

func (j *Joystick) Close() {
	C.SDL_JoystickClose(j.Native)
}

func (j *Joystick) Attached() bool {
	return C.SDL_JoystickGetAttached(j.Native) == C.SDL_TRUE
}

func (j *Joystick) EmittingEvents() bool {
	return C.SDL_JoystickEventState(C.int(-1)) == C.SDL_TRUE
}
func (j *Joystick) SetEvents(active bool) {
	var flag C.int
	if active {
		flag = 1
	}
	C.SDL_JoystickEventState(flag)
}
func (j *Joystick) Update() {
	C.SDL_JoystickUpdate()
}

func (j *Joystick) Axis(index int) int16 {
	return int16(C.SDL_JoystickGetAxis(j.Native, C.int(index)))
}

func (j *Joystick) Hat(index int) HatPosition {
	return HatPosition(C.SDL_JoystickGetHat(j.Native, C.int(index)))
}

func (j *Joystick) Trackball(index int) (dx, dy int32) {
	C.SDL_JoystickGetBall(j.Native, C.int(index), (*C.int)(&dx), (*C.int)(&dy))
	return
}

func (j *Joystick) Button(index int) uint8 {
	return uint8(C.SDL_JoystickGetButton(j.Native, C.int(index)))
}
