package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_video.h>
import "C"
import "unsafe"

type Window struct {
	Native *C.SDL_Window
}

/*
Unimplemented:
WindowPixelFormat
SetWindowIcon
SetWindowFullscreen
GetWindowSurface
UpdateWindowSurface
UpdateWindowSurfaceRects
SetWindowBrightness
GetWindowBrightness
SetWindowGammaRamp
GetWindowGammaRamp
IsScreenSaverEnabled
EnableScreenSaver
DisableScreenSaver

GL_LoadLibrary
GL_GetProcAddress
GL_UnloadLibrary
GL_ExtensionSupported
GL_SetAttribute
GL_GetAttribute
GL_CreateContext
GL_MakeCurrent
GL_GetCurrentWindow
GL_GetCurrentContext
GL_SetSwapInterval
GL_GetSwapInterval
GL_SwapWindow
GL_DeleteContext
*/
type WindowFlags struct {
	// Flags available in NewWindow
	Fullscreen bool
	Borderless bool
	Hidden     bool
	Minimized  bool
	Maximized  bool
	Resizable  bool
	Grabbed    bool
	OpenGL     bool
	// Useful with NewWindowAndRenderer
	FullscreenDesktop bool

	// Flags availble from Flags() function
	InputFocus bool
	MouseFocus bool
}

func (wf WindowFlags) toUint32() C.Uint32 {
	var flags C.Uint32
	if wf.Fullscreen {
		flags |= C.SDL_WINDOW_FULLSCREEN
	}
	if wf.Borderless {
		flags |= C.SDL_WINDOW_BORDERLESS
	}
	if wf.Hidden {
		flags |= C.SDL_WINDOW_HIDDEN
	}
	if wf.Minimized {
		flags |= C.SDL_WINDOW_MINIMIZED
	}
	if wf.Maximized {
		flags |= C.SDL_WINDOW_MAXIMIZED
	}
	if wf.Resizable {
		flags |= C.SDL_WINDOW_RESIZABLE
	}
	if wf.Grabbed {
		flags |= C.SDL_WINDOW_INPUT_GRABBED
	}
	if wf.OpenGL {
		flags |= C.SDL_WINDOW_OPENGL
	}
	if wf.FullscreenDesktop {
		flags |= C.SDL_WINDOW_FULLSCREEN_DESKTOP
	}
	return flags
}

func NewWindow(title string, x, y, w, h int, f WindowFlags) (*Window, error) {
	tstr := C.CString(title)
	defer C.free(unsafe.Pointer(tstr))

	nw := (*C.SDL_Window)(C.SDL_CreateWindow(tstr, C.int(x), C.int(y), C.int(w), C.int(h), f.toUint32()))
	if nw == nil {
		return nil, GetError()
	}

	return &Window{Native: nw}, nil
}

// Destroy terminates a window, closes the window, it becomes an ex-window
func (w *Window) Destroy() {
	C.SDL_DestroyWindow(w.Native)
}

func (w *Window) Data()    {}
func (w *Window) SetData() {}

func (w *Window) Flags() {}

// Position requests the current windows position from SDL2, then returns the X and Y position
func (w *Window) Position() (int, int) {
	var x, y int
	C.SDL_GetWindowPosition(w.Native,
		(*C.int)(unsafe.Pointer(&x)),
		(*C.int)(unsafe.Pointer(&y)),
	)
	return x, y
}

// SetPosition instructs SDL2 to move the window to a specific position
func (w *Window) SetPosition(x, y int) {
	C.SDL_SetWindowPosition(w.Native, C.int(x), C.int(y))
}

// Size returns the width and height of the current window
func (w *Window) Size() (int, int) {
	var width, height int
	C.SDL_GetWindowSize(w.Native,
		(*C.int)(unsafe.Pointer(&width)),
		(*C.int)(unsafe.Pointer(&height)),
	)
	return width, height
}

// SetSize requests SDL2 to set the size of the window to the passed width and height
func (w *Window) SetSize(width, height int) {
	C.SDL_SetWindowSize(w.Native, C.int(width), C.int(height))
}

// MinimumSize returns the width and height of the current window
func (w *Window) MinimumSize() (int, int) {
	var width, height int
	C.SDL_GetWindowMinimumSize(w.Native,
		(*C.int)(unsafe.Pointer(&width)),
		(*C.int)(unsafe.Pointer(&height)),
	)
	return width, height
}

// SetMinimumSize requests SDL2 to set the minimum size of the window to the passed width and height
func (w *Window) SetMinimumSize(width, height int) {
	C.SDL_SetWindowSize(w.Native, C.int(width), C.int(height))
}

// MaximumSize returns the width and height of the current window
func (w *Window) MaximumSize() (int, int) {
	var width, height int
	C.SDL_GetWindowMaximumSize(w.Native,
		(*C.int)(unsafe.Pointer(&width)),
		(*C.int)(unsafe.Pointer(&height)),
	)
	return width, height
}

// SetMaximumSize requests SDL2 to set the maximum size of the window to the passed width and height
func (w *Window) SetMaximumSize(width, height int) {
	C.SDL_SetWindowMaximumSize(w.Native, C.int(width), C.int(height))
}

func (w *Window) Title() string {
	cstr := C.SDL_GetWindowTitle(w.Native)
	return C.GoString(cstr)
}

func (w *Window) SetTitle(title string) {
	tstr := C.CString(title)
	defer C.free(unsafe.Pointer(tstr))
	C.SDL_SetWindowTitle(w.Native, tstr)
}

// Show makes window visible that had previously been Hidden
func (w *Window) Show() {
	C.SDL_ShowWindow(w.Native)
}

// Hide a window
func (w *Window) Hide() {
	C.SDL_HideWindow(w.Native)
}

func (w *Window) Display() DisplayInfo {
	d := C.SDL_GetWindowDisplayIndex(w.Native)
	return getDisplayInfo(int(d))
}

func (w *Window) SetWindowDisplayMode(mode DisplayMode) {

}

// Maximize sets a window's size to the maximum area available
func (w *Window) Maximize() {
	C.SDL_MaximizeWindow(w.Native)
}

// Minimize shrinks a window to an icon in the application list
// Should come back using Restore, may not, I use Hide
func (w *Window) Minimize() {
	C.SDL_MinimizeWindow(w.Native)
}

// Raise brings a window to the foreground of the desktop
func (w *Window) Raise() {
	C.SDL_RaiseWindow(w.Native)
}

// Restore changes a Maximized window to its original size or un-minimizes the window
func (w *Window) Restore() {
	C.SDL_RestoreWindow(w.Native)
}

func (w *Window) Fullscreen() {}

func (w *Window) Grab() bool {
	b := C.SDL_GetWindowGrab(w.Native)
	if b == C.SDL_TRUE {
		return true
	}
	return false
}

func (w *Window) SetGrab(active bool) {
	var b C.SDL_bool
	if active {
		b = C.SDL_TRUE
	}
	C.SDL_SetWindowGrab(w.Native, b)
}

func (w *Window) SetIcon() {}

func (w *Window) SetBordered(hasBorder bool) {
	var b C.SDL_bool
	if hasBorder {
		b = C.SDL_TRUE
	}
	C.SDL_SetWindowBordered(w.Native, b)
}
