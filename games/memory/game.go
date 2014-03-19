package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/acsellers/sdl2"
	"github.com/acsellers/sdl2/games/memory/menu"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	w, e := sdl2.NewWindow("Memory!", 0, 0, 1024, 768,
		sdl2.WindowFlags{
			Renderer: true,
		},
	)
	if e != nil {
		panic(e)
	}
	fmt.Println("Setting up menu")

	s := menu.SetupMenu()
	fmt.Println("Starting")
	s.Start(w)
	<-time.After(5 * time.Second)
	fmt.Println("Stopping")
	s.Stop()

}
