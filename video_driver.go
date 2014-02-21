package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_video.h>
import "C"
import (
	"fmt"
	"image"
	"unsafe"
)

/*
Unimplemented:
WINDOWPOS_UNDEFINED
WINDOWPOS_CENTERED

WindowEvents
SDL_GL_*

ClosestDisplayMode
WindowDis
*/
type DisplayInfo struct {
	Id           int
	Name         string
	Bounds       image.Rectangle
	DisplayModes []DisplayMode
}

func (di DisplayInfo) String() string {
	return fmt.Sprintf("%s: %dx%d", di.Name, di.Bounds.Dx(), di.Bounds.Dy())
}

type DisplayMode struct {
	PixelFormat
	Width, Height int
	RefreshRate   int
	DriverData    interface{}
}

func (dm DisplayMode) String() string {
	return fmt.Sprintf("%dx%d@%dHz", dm.Width, dm.Height, dm.RefreshRate)
}

type PixelFormat uint32

func GetVideoDrivers() []string {
	n := int(C.SDL_GetNumVideoDrivers())
	d := make([]string, n)
	for i, _ := range d {
		cstr := C.SDL_GetVideoDriver(C.int(i))
		d[i] = C.GoString(cstr)
	}
	return d
}

func CurrentVideoDriver() string {
	return C.GoString(C.SDL_GetCurrentVideoDriver())
}

func CurrentDesktopDisplayModes() []DisplayMode {
	n := int(C.SDL_GetNumVideoDisplays())
	modes := make([]DisplayMode, int(n))
	for i, _ := range modes {
		var dm C.SDL_DisplayMode
		C.SDL_GetDesktopDisplayMode(C.int(i), &dm)
		modes[i] = DisplayMode{
			PixelFormat: PixelFormat(dm.format),
			Width:       int(dm.w),
			Height:      int(dm.h),
			RefreshRate: int(dm.refresh_rate),
			DriverData:  dm.driverdata,
		}
	}
	return modes

}

func CurrentDisplayMode(desktop int) DisplayMode {
	var dm C.SDL_DisplayMode
	C.SDL_GetCurrentDisplayMode(C.int(desktop), &dm)
	return DisplayMode{
		PixelFormat: PixelFormat(dm.format),
		Width:       int(dm.w),
		Height:      int(dm.h),
		RefreshRate: int(dm.refresh_rate),
		DriverData:  dm.driverdata,
	}

}

// InitVideo initializes the Video subsystem of SDL2 with the default driver
func InitVideo() error {
	if r := C.SDL_VideoInit(nil); r == -1 {
		return GetError()
	}
	return nil
}

// InitVideoDriver initializes the Video subsystem of SDL2 with a specific
// driver from the slice returned by GetVideoDrivers
func InitVideoDriver(driver string) error {
	cstr := C.CString(driver)
	defer C.free(unsafe.Pointer(cstr))
	if r := C.SDL_VideoInit(cstr); r == -1 {
		return GetError()
	}
	return nil
}

func QuitVideo() {
	C.SDL_VideoQuit()
}

func VideoDisplays() []DisplayInfo {
	n := int(C.SDL_GetNumVideoDisplays())
	d := make([]DisplayInfo, n)
	for i, _ := range d {
		d[i] = getDisplayInfo(i)
	}
	return d
}

func getDisplayInfo(i int) DisplayInfo {
	return DisplayInfo{
		Id:           i,
		Name:         getDisplayName(i),
		Bounds:       getDisplayBounds(i),
		DisplayModes: getDisplayModes(i),
	}
}

func getDisplayName(i int) string {
	cstr := C.SDL_GetDisplayName(C.int(i))
	return C.GoString(cstr)
}

func getDisplayBounds(i int) image.Rectangle {
	var r C.SDL_Rect
	C.SDL_GetDisplayBounds(C.int(i), &r)
	return image.Rect(int(r.x), int(r.y), int(r.x+r.w), int(r.y+r.h))
}

func getDisplayModes(i int) []DisplayMode {
	mn := C.SDL_GetNumDisplayModes(C.int(i))
	modes := make([]DisplayMode, int(mn))
	for j, _ := range modes {
		var dm C.SDL_DisplayMode
		C.SDL_GetDisplayMode(C.int(i), C.int(j), &dm)
		modes[j] = DisplayMode{
			PixelFormat: PixelFormat(dm.format),
			Width:       int(dm.w),
			Height:      int(dm.h),
			RefreshRate: int(dm.refresh_rate),
			DriverData:  dm.driverdata,
		}
	}
	return modes
}

/* Do later
type GammaRamp struct {
	Red   [256]uint16
	Green [256]uint16
	Blue  [256]uint16
}

func (w *Window) SetGammaRamp(gr GammaRamp) error {

}
*/

// ScreenSaver returns whether screensavers are enabled.
func ScreenSaver() bool {
	if C.SDL_IsScreenSaverEnabled() == C.SDL_TRUE {
		return true
	}
	return false
}

// SetScreenSaver sets whether the ScreenSaver is allowed to activate,
// if the status is true, then screensavers can blank the screen, or vice versa.
func SetScreenSaver(status bool) {
	if status {
		C.SDL_EnableScreenSaver()
	} else {
		C.SDL_DisableScreenSaver()
	}
}
