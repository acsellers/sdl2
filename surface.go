package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_surface.h>
import "C"

import (
	"fmt"
	"image"
	"image/color"
	"runtime"
	"unsafe"
)

type Surface struct {
	Native        *C.SDL_Surface
	Width, Height int
	Stride        int // SDL calls it pitch, but in go's image library it's stride
	Pixels        []byte
	clip          C.SDL_Rect
	rle           bool
}

func NewSurfaceFromImage(i *image.RGBA) *Surface {
	s := &Surface{
		Native: C.SDL_CreateRGBSurfaceFrom(
			unsafe.Pointer(&i.Pix[0]),
			C.int(i.Rect.Dx()),
			C.int(i.Rect.Dy()),
			C.int(32),
			C.int(4*i.Rect.Dx()),
			0x000000ff, // Yes these seem backwards
			0x0000ff00, // They're correct
			0x00ff0000,
			0xff000000,
		),
		Width:  i.Rect.Dx(),
		Height: i.Rect.Dy(),
	}
	runtime.SetFinalizer(s, (*Surface).Free)

	return s
}

// Use SDL's native function to save the surface
// to a bmp file
func (s *Surface) SaveToBMP(filename string) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	bstr := C.CString("wb")
	defer C.free(unsafe.Pointer(bstr))
	C.SDL_SaveBMP_RW(s.Native, C.SDL_RWFromFile(cstr, bstr), C.int(1))
}

// When the Surface was created, it was set to free the associated
// SDL surface when it was garbage collected. If you are pulling out
// the SDL surfaces and using them separately from the Surface struct,
// you should DisableFree on those surface.
func (s *Surface) DisableFree() {
	runtime.SetFinalizer(s, nil)
}

func (s *Surface) Free() {
	if s.Native != nil {
		C.SDL_FreeSurface(s.Native)
	}
}

func (s *Surface) LockPixels() {
}

func (s *Surface) UnlockPixels() {
}

// SetSurfaceRLE marks a surface to get accelerated blits, but the surface
// have LockPixels called before the pixels can be edited. This is useful
// when using a color key.
func (s *Surface) SetSurfaceRLE() error {
	if int(C.SDL_SetSurfaceRLE(s.Native, C.int(1))) != 0 {
		return GetError()
	}
	s.rle = true
	return nil
}

func (s *Surface) SetColorKey(key color.RGBA) {
	p := C.SDL_MapRGB(s.Native.format, C.Uint8(key.R), C.Uint8(key.G), C.Uint8(key.B))
	C.SDL_SetColorKey(s.Native, C.SDL_TRUE, p)
}

func (s *Surface) ColorKey() (color.RGBA, error) {
	var c uint32
	if C.SDL_GetColorKey(s.Native, (*C.Uint32)(unsafe.Pointer(&c))) != 0 {
		return color.RGBA{}, GetError()
	}

	var r, g, b, a uint32
	a = c >> 24
	b = c >> 16
	g = c >> 8
	r = c
	fmt.Println(c)

	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}, nil
}
