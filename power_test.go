package sdl2

import "time"

func ExamplePowerInfo() {
	p := PowerInfo{}
	p.Refresh()
	if p.PowerState == OnBattery && p.Remaining < 5*time.Minute {
		// ShutdownGame()
	}
}
