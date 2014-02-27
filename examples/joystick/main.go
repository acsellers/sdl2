package main

import (
	"fmt"
	"log"
	"time"

	"github.com/acsellers/sdl2"
)

func main() {
	joys := sdl2.ConnectedJoysticks()
	if len(joys) == 0 {
		log.Fatal("Could not find any joysticks")
	} else {
		fmt.Println("Joystick Num:", len(joys))
	}

	j := joys[0]
	w, e := sdl2.NewWindow("test window", 80, 80, 80, 80,
		sdl2.WindowFlags{},
	)
	sdl2.WatchEvents()

	if e != nil {
		log.Fatal(e)
	}
	defer w.Destroy()

	fmt.Println(j.Open())
	fmt.Println(j.Name)
	fmt.Println(j.ControllerName)
	fmt.Println(j.GameController)
	fmt.Println(j.Axes)
	fmt.Println(j.Trackballs)
	fmt.Println(j.POVHats)
	fmt.Println(j.Buttons)
	// we're doing things manually for this example (don't actually do this)
	for i := 0; i < 30; i++ {
		j.Update()
		fmt.Println("Still attached", j.Attached())
		blank := true
		for i := 0; i < j.Buttons; i++ {
			if v := j.Button(i); v != 0 {
				fmt.Println("Button:", i, "\nValue:", v)
				blank = false
			}
		}
		// fmt.Println("A:", gcb.ButtonStatus(sdl2.ButtonA))
		if blank {
			fmt.Println("None of the", j.Buttons, "were pressed")
		}
		time.Sleep(time.Second)
	}
	j.Close()
}
