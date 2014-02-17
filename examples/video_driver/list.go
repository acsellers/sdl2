package main

import (
	"fmt"

	"github.com/acsellers/sdl2"
)

func main() {
	fmt.Println(sdl2.GetVideoDrivers())
	fmt.Println(sdl2.InitVideo())
	disp := sdl2.VideoDisplays()
	fmt.Println("\nAvailable Displays:")
	for _, d := range disp {
		fmt.Println(d)
		for _, m := range d.DisplayModes {
			fmt.Println(m)
		}
	}
	fmt.Println("\nCurrent DisplayModes:")
	for _, m := range sdl2.CurrentDesktopDisplayModes() {
		fmt.Println(m)
	}
}
