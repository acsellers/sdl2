package main

import (
	"fmt"

	"github.com/acsellers/sdl2"
)

func main() {
	pi := sdl2.PowerInfo{}
	pi.Refresh()
	fmt.Println(pi)
	fmt.Println(pi.Percent)
	fmt.Println(pi.Remaining)
}
