package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_stdinc.h>
// #include <SDL2/SDL_pixels.h>
import "C"
import (
	"image/color"
	"runtime"
	"unsafe"
)

/*
Unimplemented:
SDL_Palette
*/

type PixelFormat uint32

const (
	Unknown     PixelFormat = C.SDL_PIXELFORMAT_UNKNOWN
	Index1LSB   PixelFormat = C.SDL_PIXELFORMAT_INDEX1LSB
	Index1MSB   PixelFormat = C.SDL_PIXELFORMAT_INDEX1MSB
	Index4LSB   PixelFormat = C.SDL_PIXELFORMAT_INDEX4LSB
	Index4MSB   PixelFormat = C.SDL_PIXELFORMAT_INDEX4MSB
	Index8      PixelFormat = C.SDL_PIXELFORMAT_INDEX8
	RGB332      PixelFormat = C.SDL_PIXELFORMAT_RGB332
	RGB444      PixelFormat = C.SDL_PIXELFORMAT_RGB444
	RGB555      PixelFormat = C.SDL_PIXELFORMAT_RGB555
	BGR555      PixelFormat = C.SDL_PIXELFORMAT_BGR555
	ARGB4444    PixelFormat = C.SDL_PIXELFORMAT_ARGB4444
	RGBA4444    PixelFormat = C.SDL_PIXELFORMAT_RGBA4444
	ABGR4444    PixelFormat = C.SDL_PIXELFORMAT_ABGR4444
	BGRA4444    PixelFormat = C.SDL_PIXELFORMAT_BGRA4444
	ARGB1555    PixelFormat = C.SDL_PIXELFORMAT_ARGB1555
	RGBA5551    PixelFormat = C.SDL_PIXELFORMAT_RGBA5551
	ABGR1555    PixelFormat = C.SDL_PIXELFORMAT_ABGR1555
	BGRA5551    PixelFormat = C.SDL_PIXELFORMAT_BGRA5551
	RGB565      PixelFormat = C.SDL_PIXELFORMAT_RGB565
	BGR565      PixelFormat = C.SDL_PIXELFORMAT_BGR565
	RGB24       PixelFormat = C.SDL_PIXELFORMAT_RGB24
	BGR24       PixelFormat = C.SDL_PIXELFORMAT_BGR24
	RGB888      PixelFormat = C.SDL_PIXELFORMAT_RGB888
	RGBX8888    PixelFormat = C.SDL_PIXELFORMAT_RGBX8888
	BGR888      PixelFormat = C.SDL_PIXELFORMAT_BGR888
	BGRX8888    PixelFormat = C.SDL_PIXELFORMAT_BGRX8888
	ARGB8888    PixelFormat = C.SDL_PIXELFORMAT_ARGB8888
	RGBA8888    PixelFormat = C.SDL_PIXELFORMAT_RGBA8888
	ABGR8888    PixelFormat = C.SDL_PIXELFORMAT_ABGR8888
	BGRA8888    PixelFormat = C.SDL_PIXELFORMAT_BGRA8888
	ARGB2101010 PixelFormat = C.SDL_PIXELFORMAT_ARGB2101010
	YV12        PixelFormat = C.SDL_PIXELFORMAT_YV12
	IYUV        PixelFormat = C.SDL_PIXELFORMAT_IYUV
	YUY2        PixelFormat = C.SDL_PIXELFORMAT_YUY2
	UYVY        PixelFormat = C.SDL_PIXELFORMAT_UYVY
	YVYU        PixelFormat = C.SDL_PIXELFORMAT_YVYU
)

func (pf PixelFormat) String() string {
	cstr := C.SDL_GetPixelFormatName(C.Uint32(pf))
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

type PixelMask struct {
	BPP        int32
	R, G, B, A uint32
}

func (pf PixelFormat) Mask() (PixelMask, bool) {
	var pm PixelMask
	r := C.SDL_PixelFormatEnumToMasks(C.Uint32(pf),
		(*C.int)(&pm.BPP),
		(*C.Uint32)(&pm.R),
		(*C.Uint32)(&pm.G),
		(*C.Uint32)(&pm.B),
		(*C.Uint32)(&pm.A),
	)
	if r == C.SDL_TRUE {
		return pm, true
	} else {
		return PixelMask{}, false
	}
}

func (pm PixelMask) PixelFormat() PixelFormat {
	return PixelFormat(C.SDL_MasksToPixelFormatEnum(
		C.int(pm.BPP),
		C.Uint32(pm.R),
		C.Uint32(pm.G),
		C.Uint32(pm.B),
		C.Uint32(pm.A)),
	)
}

type NativePixelFormat struct {
	Native *C.SDL_PixelFormat
}

func (pf PixelFormat) Native() NativePixelFormat {
	f := NativePixelFormat{Native: C.SDL_AllocFormat(C.Uint32(pf))}
	runtime.SetFinalizer(f, func(*C.SDL_PixelFormat) {
		C.SDL_FreeFormat(f.Native)
	})
	return f
}

func (npf *NativePixelFormat) SetPalette(p *Palette) {
	C.SDL_SetPixelFormatPalette((*C.SDL_PixelFormat)(npf.Native), p.Native)
}

func (pf PixelFormat) RGB(c color.RGBA) uint32 {
	return uint32(C.SDL_MapRGB(pf.Native().Native, C.Uint8(c.R), C.Uint8(c.G), C.Uint8(c.B)))
}
func (pf PixelFormat) RGBA(c color.RGBA) uint32 {
	return uint32(C.SDL_MapRGBA(pf.Native().Native, C.Uint8(c.R), C.Uint8(c.G), C.Uint8(c.B), C.Uint8(c.A)))
}
func (pf PixelFormat) GetRGB(pxl uint32) color.RGBA {
	var r, g, b uint8
	C.SDL_GetRGB(C.Uint32(pxl), pf.Native().Native, (*C.Uint8)(&r), (*C.Uint8)(&g), (*C.Uint8)(&b))
	return color.RGBA{R: r, G: g, B: b, A: 0xff}
}
func (pf PixelFormat) GetRGBA(pxl uint32) color.RGBA {
	var r, g, b, a uint8
	C.SDL_GetRGBA(C.Uint32(pxl), pf.Native().Native, (*C.Uint8)(&r), (*C.Uint8)(&g), (*C.Uint8)(&b), (*C.Uint8)(&a))
	return color.RGBA{R: r, G: g, B: b, A: a}
}

func GammaRamp(gamma float32) [256]uint16 {
	var r [256]uint16
	C.SDL_CalculateGammaRamp(C.float(gamma), (*C.Uint16)(&r[0]))
	return r
}

type Palette struct {
	Native *C.SDL_Palette
}

// Note that the color.Palette should be made up of color.RGBA
// for best results
func ConvertPalette(p color.Palette) *Palette {
	sp := Palette{
		Native: C.SDL_AllocPalette(C.int(len([]color.Color(p)))),
	}
	runtime.SetFinalizer(p, func(p Palette) {
		C.SDL_FreePalette(p.Native)
	})

	colors := make([]C.SDL_Color, len([]color.Color(p)))
	for i, c := range []color.Color(p) {
		r, g, b, a := c.RGBA()
		colors[i] = C.SDL_Color{
			r: C.Uint8(uint8(r)),
			g: C.Uint8(uint8(g)),
			b: C.Uint8(uint8(b)),
			a: C.Uint8(uint8(a)),
		}
	}
	C.SDL_SetPaletteColors(sp.Native, &colors[0], 0, C.int(len(colors)))
	return &sp
}
func (p *Palette) DisableFree() {
	runtime.SetFinalizer(p, nil)
}

func (p *Palette) Free() {
	p.DisableFree()
	C.SDL_FreePalette(p.Native)
}
