package menu

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"

	"github.com/acsellers/sdl2"
	"github.com/acsellers/sdl2/gui"
)

func SetupMenu() *gui.Screen {
	return &gui.Screen{
		Background: MenuBackground(),
		Items:      []gui.Drawable{MenuTitle()},
	}
}

func MenuBackground() gui.Drawable {
	img := ConvertImage(GetImage("background.png"))
	surf := sdl2.NewSurfaceFromImage(img)
	fmt.Println("background")
	return &gui.Background{
		S: surf,
	}
}

func MenuTitle() gui.Drawable {
	img := ConvertImage(GetImage("title.png"))
	surf := sdl2.NewSurfaceFromImage(img)
	fmt.Println("title")
	return &gui.StaticSurface{
		S:         surf,
		Placement: image.Rect(366, 80, 366+img.Bounds().Dx(), 80+img.Bounds().Dy()),
	}
}

func GetImage(name string) image.Image {
	f, e := os.Open(name)
	if e != nil {
		fmt.Print("Could not open image:", name)
		panic(e)
	}

	img, e := png.Decode(f)
	if e != nil {
		fmt.Print("Could not parse image:", name)
		panic(e)
	}

	return img
}

func ConvertImage(img image.Image) *image.RGBA {
	rimg := image.NewRGBA(img.Bounds())
	draw.Draw(rimg, img.Bounds(), img, image.ZP, draw.Src)
	f, _ := os.Create("intermediate.png")
	png.Encode(f, rimg)

	return rimg
}
