package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_messagebox.h>
import "C"

import (
	"fmt"
	"image/color"
	"unsafe"
)

type MessageBoxType uint32

const (
	Error       MessageBoxType = 0x10
	Warning     MessageBoxType = 0x20
	Information MessageBoxType = 0x40
)

func NewSimpleMessageBox(t MessageBoxType, title, message string) error {
	tstr := C.CString(title)
	defer C.free(unsafe.Pointer(tstr))
	mstr := C.CString(message)
	defer C.free(unsafe.Pointer(mstr))

	e := C.SDL_ShowSimpleMessageBox(C.Uint32(t), tstr, mstr, nil)
	if e == -1 {
		return fmt.Errorf("couldn't open messagebox")
	}

	return nil
}

type MessageBox struct {
	// Warning, Error, Information
	MessageBoxType
	Title, Message string
	Buttons        []MessageButton

	// Parent window for the message box, if nil, message box will not be parented
	Parent *Window
	// Color scheme, may be nil for the default SDL2 color scheme
	Colors *MessageBoxColorScheme
}

type MessageButton struct {
	Submit bool
	Quit   bool
	Text   string
	Click  func()
}

// Color scheme to use for Messagebox
// Note that only R,G,B attributes are used
type MessageBoxColorScheme struct {
	Background       color.Color
	Text             color.Color
	ButtonBorder     color.Color
	ButtonBackground color.Color
	ButtonSelected   color.Color
}

func toMessageBoxColor(c color.Color) C.SDL_MessageBoxColor {
	r, g, b, _ := c.RGBA()
	return C.SDL_MessageBoxColor{C.Uint8(r), C.Uint8(g), C.Uint8(b)}
}

func (mbcs *MessageBoxColorScheme) toNative() *C.SDL_MessageBoxColorScheme {
	if mbcs == nil {
		return nil
	}
	return &C.SDL_MessageBoxColorScheme{
		[5]C.SDL_MessageBoxColor{
			toMessageBoxColor(mbcs.Background),
			toMessageBoxColor(mbcs.Text),
			toMessageBoxColor(mbcs.ButtonBorder),
			toMessageBoxColor(mbcs.ButtonBackground),
			toMessageBoxColor(mbcs.ButtonSelected),
		},
	}
}

func (mb MessageBox) Show() error {
	tstr := C.CString(mb.Title)
	defer C.free(unsafe.Pointer(tstr))
	mstr := C.CString(mb.Message)
	defer C.free(unsafe.Pointer(mstr))

	buttons := make([]C.SDL_MessageBoxButtonData, len(mb.Buttons))
	for i, b := range mb.Buttons {
		bstr := C.CString(b.Text)
		defer C.free(unsafe.Pointer(bstr))

		if b.Submit {
			buttons[i] = C.SDL_MessageBoxButtonData{
				flags:    C.SDL_MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT,
				buttonid: C.int(i),
				text:     bstr,
			}
		}

		if b.Quit {
			buttons[i] = C.SDL_MessageBoxButtonData{
				flags:    C.SDL_MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT,
				buttonid: C.int(i),
				text:     bstr,
			}
		}
	}

	native := C.SDL_MessageBoxData{
		flags:       C.Uint32(mb.MessageBoxType),
		window:      nil, // TODO: Parents aren't useful yet
		title:       tstr,
		message:     mstr,
		numbuttons:  C.int(len(mb.Buttons)),
		buttons:     &buttons[0],
		colorScheme: mb.Colors.toNative(),
	}

	var button int
	e := C.SDL_ShowMessageBox(&native, (*C.int)(unsafe.Pointer(&button)))

	if e == -1 {
		return fmt.Errorf("sdl messagebox error")
	}
	mb.Buttons[button].Click()

	return nil
}
