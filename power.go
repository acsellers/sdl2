package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_power.h>
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

// PowerInfo is the basic state for a system's power supply
type PowerInfo struct {
	PowerState
	Remaining time.Duration
	Percent   int
}

// String prints a human readable representation of the current
// state, and either a remaining time or remaining percent based
// on what the PowerState is recorded as.
//
//    p.String() // "Discharging: 74 Minutes Remaining"
//    p.String() // "Charged"
//    p.String() // "Unknown"
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

// Refresh reloads the data in the PowerInfo struct.
func (pi *PowerInfo) Refresh() {
	var remain int
	pi.PowerState = PowerState(C.SDL_GetPowerInfo(
		(*C.int)(unsafe.Pointer(&remain)),
		(*C.int)(unsafe.Pointer(&pi.Percent)),
	))
	pi.Remaining = time.Duration(remain) * time.Second
}

// PowerState is the state of the computer
type PowerState C.SDL_PowerState

const (
	// Unknown indicates that SDL2 cannot determine the battery information
	Unknown PowerState = C.SDL_POWERSTATE_UNKNOWN
	// OnBattery indicates that the system is currently running on vatter power
	OnBattery = C.SDL_POWERSTATE_ON_BATTERY
	// NoBattery indicates that computer does not have a battery
	NoBattery = C.SDL_POWERSTATE_NO_BATTERY
	// Charging indicates that the computer is currently charging
	Charging = C.SDL_POWERSTATE_CHARGING
	// Charged indicates that the computer has a fully charged battery
	Charged = C.SDL_POWERSTATE_CHARGED
)
