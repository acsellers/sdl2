package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/acsellers/sdl2"
)

func main() {
	sdl2.InitVideo()
	// Create a source for a gold color to draw with
	p, e := os.Open("run.png")
	if e != nil {
		fmt.Println(e)
		fmt.Println("You need to be in the imagesurface folder to run")
	}

	o, _ := png.Decode(p)
	// The image to use for creating our surface
	img := image.NewRGBA(image.Rect(0, 0, 64, 64))
	u := image.NewUniform(color.RGBA{0x99, 0x99, 0x00, 0xff})
	draw.Draw(img, img.Rect, u, image.ZP, draw.Src)
	s := sdl2.NewSurfaceFromImage(img)
	s.SaveToBMP("test1.bmp")

	draw.Draw(img, img.Bounds(), o, o.Bounds().Min, draw.Src)

	s = sdl2.NewSurfaceFromImage(img)
	s.SaveToBMP("test2.bmp")

	f, _ := os.Create("test.png")
	png.Encode(f, img)
}
