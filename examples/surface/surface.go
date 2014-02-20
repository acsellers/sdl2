package main

import (
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
	o := image.NewUniform(color.RGBA{0x99, 0x99, 0x00, 0xff})
	// The image to use for creating our surface
	img := image.NewRGBA(image.Rect(0, 0, 64, 64))
	// Copy our gold color to our source surface
	draw.Draw(img, img.Rect, o, image.ZP, draw.Src)

	s := sdl2.NewSurfaceFromImage(img)
	s.SaveToBMP("test.bmp")

	f, _ := os.Create("test.png")
	png.Encode(f, img)
}
