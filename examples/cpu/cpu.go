package main

import (
	"fmt"

	"github.com/acsellers/sdl2"
)

func main() {
	fmt.Printf(
		"Your CPU has %d cores\n",
		sdl2.NumCPU(),
		//sdl2.NumRam(),
	)

	ci := sdl2.NewCPUInfo()
	if ci.AltiVec {
		fmt.Println("You're on PowerPC")
	}
	if ci.MMX {
		fmt.Println("You're on x86")
	}

	if ci.SSE41 {
		fmt.Println("You're on a recent x86_64")
	}
}
