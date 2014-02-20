package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"time"

	"github.com/acsellers/sdl2"
)

func main() {
	sdl2.InitVideo()
	w, e := sdl2.NewWindow("test window", 80, 80, 80, 80,
		sdl2.WindowFlags{},
	)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	// Open our icon png
	p, e := os.Open("run.png")
	if e != nil {
		fmt.Println(e)
		fmt.Println("You need to be in the windowicon folder to run")
	}
	o, _ := png.Decode(p)

	// The image to use for creating our surface
	img := image.NewRGBA(image.Rect(0, 0, 64, 64))
	draw.Draw(img, img.Bounds(), o, o.Bounds().Min, draw.Src)

	s := sdl2.NewSurfaceFromImage(img)
	s.SetColorKey(color.RGBA{0xff, 0x0, 0x0, 0xff})
	s.SetSurfaceRLE()
	w.SetIcon(s)
	time.Sleep(10 * time.Second)
}
