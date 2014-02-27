package main

import (
	"time"

	"github.com/acsellers/sdl2"
)

func main() {
	sdl2.WatchEvents()
	time.Sleep(5 * time.Second)
}
