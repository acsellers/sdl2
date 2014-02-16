package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_power.h>
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

// The basic state for a system's power supply
type PowerInfo struct {
	PowerState
	Remaining time.Duration
	Percent   int
}

func (pi PowerInfo) String() string {
	switch pi.PowerState {
	case OnBattery:
		return fmt.Sprintf(
			"Discharging: %3.0f Minutes Remaining",
			pi.Remaining.Seconds()/60.0,
		)
	case NoBattery:
		return "Plugged In"
	case Charging:
		return fmt.Sprintf("Charging: %3.0f%%", pi.Percent)
	case Charged:
		return "Charged"
	}
	return "Unknown"
}

func (pi *PowerInfo) Refresh() {
	var remain int
	pi.PowerState = PowerState(C.SDL_GetPowerInfo(
		(*C.int)(unsafe.Pointer(&remain)),
		(*C.int)(unsafe.Pointer(&pi.Percent)),
	))
	pi.Remaining = time.Duration(remain) * time.Second
}

type PowerState C.SDL_PowerState

const (
	Unknown   PowerState = C.SDL_POWERSTATE_UNKNOWN
	OnBattery            = C.SDL_POWERSTATE_ON_BATTERY
	NoBattery            = C.SDL_POWERSTATE_NO_BATTERY
	Charging             = C.SDL_POWERSTATE_CHARGING
	Charged              = C.SDL_POWERSTATE_CHARGED
)
