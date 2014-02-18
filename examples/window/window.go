package main

import (
	"fmt"
	"os"
	"time"

	"github.com/acsellers/sdl2"
)

func main() {
	fmt.Println("Start Window Example")
	w, e := sdl2.NewWindow("test window", 80, 80, 80, 80,
		sdl2.WindowFlags{},
	)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	time.Sleep(time.Second)
	w.SetBordered(false)
	for i := 0; i < 5; i++ {
		time.Sleep(333 * time.Millisecond)
		fmt.Println("Position")
		fmt.Println(w.Position())
		w.SetPosition(80*(i+2), 30*i)
		fmt.Println("Size")
		fmt.Println(w.Size())
		w.SetSize(40*(i+1), 60*(i+1))
	}
	time.Sleep(500 * time.Millisecond)
	w.Maximize()
	time.Sleep(500 * time.Millisecond)
	w.Restore()
	time.Sleep(500 * time.Millisecond)
	w.Hide()
	time.Sleep(time.Second)
	w.Show()
	time.Sleep(500 * time.Millisecond)
	w.Destroy()
}
