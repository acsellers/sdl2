package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_surface.h>
// #include <SDL2/SDL_render.h>
import "C"

import (
	"fmt"
	"image"
	"image/color"
	"runtime"
	"unsafe"
)

type BlendMode uint32

/* Unimplemented
CreateRGBSurface (NewSurface)
SetSurfacePalette
GetClipRect
SetClipRect
ConvertSurface
ConvertSurfaceFormat
ConvertPixels
FillRect
FillRects
SoftStretch
*/
const (
	// Blending equations are from SDL2 header files
	//
	// No blending
	// dstRGBA = srcRGBA
	None BlendMode = C.SDL_BLENDMODE_NONE
	// Alpha Blending
	// dstRGB = (srcRGB * srcA) + (dstRGB * (1-srcA))
	// dstA = srcA + (dstA * (1-srcA))
	Blend BlendMode = C.SDL_BLENDMODE_BLEND
	// Additive Blending
	// dstRGB = (srcRGB * srcA) + dstRGB
	// dstA = dstA
	Additive BlendMode = C.SDL_BLENDMODE_ADD
	// Color Modulation
	// dstRGB = srcRGB * dstRGB
	// dstA = dstA
	ColorModulation BlendMode = C.SDL_BLENDMODE_MOD
)

type Surface struct {
	Native        *C.SDL_Surface
	Source        *image.RGBA
	textures      []texture
	Width, Height int
	Stride        int // SDL calls it pitch, but in go's image library it's stride
	Pixels        []byte
	clip          C.SDL_Rect
	rle           bool
}

type texture struct {
	texture *C.SDL_Texture
	window  *Window
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
		Source: i,
		Width:  i.Rect.Dx(),
		Height: i.Rect.Dy(),
	}
	s.SaveToBMP("native.bmp")
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

// SetColorKey sets the color that is to be regarded as transparent
// by SDL2, note that the alpha channel is discarded.
func (s *Surface) SetColorKey(key color.RGBA) {
	p := C.SDL_MapRGB(s.Native.format, C.Uint8(key.R), C.Uint8(key.G), C.Uint8(key.B))
	C.SDL_SetColorKey(s.Native, C.SDL_TRUE, p)
}

func (s *Surface) ColorKey() (color.RGBA, error) {
	var c uint32
	if C.SDL_GetColorKey(s.Native, (*C.Uint32)(unsafe.Pointer(&c))) != 0 {
		return color.RGBA{}, GetError()
	}

	var r, g, b C.Uint8
	C.SDL_GetRGB(C.Uint32(c), s.Native.format, (*C.Uint8)(&r), (*C.Uint8)(&g), (*C.Uint8)(&b))
	return color.RGBA{uint8(r), uint8(g), uint8(b), 0xff}, nil
}

func (s *Surface) SetColorModifier(mod color.RGBA) error {
	if C.SDL_SetSurfaceColorMod(s.Native, C.Uint8(mod.R), C.Uint8(mod.G), C.Uint8(mod.B)) != 0 {
		return GetError()
	}
	return nil
}

func (s *Surface) ColorModifier() (color.RGBA, error) {
	var r, g, b C.Uint8
	if C.SDL_GetSurfaceColorMod(s.Native, (*C.Uint8)(&r), (*C.Uint8)(&g), (*C.Uint8)(&b)) != 0 {
		return color.RGBA{}, GetError()
	}
	return color.RGBA{uint8(r), uint8(g), uint8(b), 0xff}, nil
}

// SetAlphaModifier is calculated as srcAlpha * ( modAlpha / 255 )
func (s *Surface) SetAlphaModifier(a uint8) error {
	if C.SDL_SetSurfaceAlphaMod(s.Native, C.Uint8(a)) != 0 {
		return GetError()
	}
	return nil
}
func (s *Surface) AlphaModifier() (uint8, error) {
	var a uint8
	if C.SDL_GetSurfaceAlphaMod(s.Native, (*C.Uint8)(&a)) != 0 {
		return 0, GetError()
	}
	return a, nil
}

func (s *Surface) BlendMode() BlendMode {
	var bm C.SDL_BlendMode
	C.SDL_GetSurfaceBlendMode(s.Native, &bm)
	return BlendMode(bm)
}
func (s *Surface) SetBlendMode(bm BlendMode) error {
	if C.SDL_SetSurfaceBlendMode(s.Native, C.SDL_BlendMode(bm)) != 0 {
		return GetError()
	}
	return nil
}

func (s *Surface) Blit(d *Surface, dst image.Point) (image.Rectangle, error) {
	dr := C.SDL_Rect{x: C.int(dst.X), y: C.int(dst.Y)}
	r := C.SDL_BlitSurface(s.Native, nil, d.Native, &dr)
	if r != 0 {
		return image.Rectangle{}, GetError()
	}
	return image.Rect(int(dr.x), int(dr.y), int(dr.x+dr.w), int(dr.y+dr.h)), nil
}

func (s *Surface) BlitSubset(d *Surface, dst image.Point, subset image.Rectangle) (image.Rectangle, error) {
	dr := C.SDL_Rect{x: C.int(dst.X), y: C.int(dst.Y)}
	r := C.SDL_BlitSurface(
		s.Native,
		RectToNative(subset),
		d.Native,
		&dr,
	)
	if r != 0 {
		return image.Rectangle{}, GetError()
	}
	return image.Rect(int(dr.x), int(dr.y), int(dr.x+dr.w), int(dr.y+dr.h)), nil
}

func (s *Surface) BlitScaled(d *Surface, dst image.Rectangle) error {
	r := C.SDL_BlitSurface(
		s.Native,
		nil,
		d.Native,
		RectToNative(dst),
	)
	if r != 0 {
		return GetError()
	}
	return nil
}

func (s *Surface) BlitScaledSubset(d *Surface, dst, subset image.Rectangle) error {
	// TODO: what is this?
	r := C.SDL_BlitSurface(
		s.Native,
		RectToNative(subset),
		d.Native,
		RectToNative(dst),
	)
	if r != 0 {
		return GetError()
	}
	return nil

}

func (s *Surface) Texture(w *Window) *C.SDL_Texture {
	for _, t := range s.textures {
		if t.window == w {
			return t.texture
		}
	}

	tex := C.SDL_CreateTextureFromSurface(w.Renderer, s.Native)
	fmt.Println(tex == nil)
	s.textures = append(s.textures, texture{tex, w})
	return tex
}

func (s *Surface) RegenTexture(w *Window) *C.SDL_Texture {
	tex := C.SDL_CreateTextureFromSurface(w.Renderer, s.Native)
	for _, t := range s.textures {
		if t.window == w {
			C.SDL_DestroyTexture(t.texture)
			t.texture = tex
			return tex
		}
	}

	s.textures = append(s.textures, texture{tex, w})
	return tex
}

func (s *Surface) NativeClipping() *C.SDL_Rect {
	return &C.SDL_Rect{
		x: 0,
		y: 0,
		w: C.int(s.Width),
		h: C.int(s.Height),
	}
}
