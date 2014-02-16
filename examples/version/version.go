package main

import (
	"fmt"

	"github.com/acsellers/sdl2"
)

func main() {
	fmt.Printf("SDL Version: %d.%d.%d\n", sdl2.Major, sdl2.Minor, sdl2.Patch)
	fmt.Printf("SDL Revision: %d\n", sdl2.Revision)
}
