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
}
