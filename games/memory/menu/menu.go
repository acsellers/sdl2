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
		Setup: func(s *gui.Screen) {
			s.Background = MenuBackground()
			s.Items = []gui.Drawable{MenuTitle()}
		},
	}
}

func MenuBackground() gui.Drawable {
	img := ConvertImage(GetImage("background.png"))
	surf := sdl2.NewSurfaceFromImage(img)
	return &SurfaceDrawer{
		S:    surf,
		Rect: image.Rect(0, 0, 1024, 768),
	}
}

func MenuTitle() gui.Drawable {
	img := ConvertImage(GetImage("title.png"))
	surf := sdl2.NewSurfaceFromImage(img)
	surf.SaveToBMP("test.bmp")
	return &SurfaceDrawer{
		S:    surf,
		Rect: image.Rect(366, 80, 366+img.Bounds().Dx(), 80+img.Bounds().Dy()),
	}
}

type SurfaceDrawer struct {
	S      *sdl2.Surface
	Rect   image.Rectangle
	Active image.Rectangle
}

func (sd *SurfaceDrawer) Draw() (*sdl2.Surface, image.Rectangle) {
	return sd.S, sd.Rect
}

func (sd *SurfaceDrawer) ActiveArea() image.Rectangle {
	return sd.Active
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
