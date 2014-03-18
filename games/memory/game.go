package main

import (
	"time"

	"github.com/acsellers/sdl2"
	"github.com/acsellers/sdl2/games/memory/menu"
)

func main() {
	w, e := sdl2.NewWindow("Memory!", 0, 0, 1024, 768,
		sdl2.WindowFlags{
			Renderer: true,
		},
	)
	if e != nil {
		panic(e)
	}

	s := menu.SetupMenu()
	s.SetWindow(w)
	s.Start()
	<-time.After(5 * time.Second)
	s.Stop()
}
