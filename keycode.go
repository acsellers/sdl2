package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_keycode.h>
// #include <SDL2/SDL_keyboard.h>
import "C"
import "unsafe"

type Keycode rune

func NewKeycodeFromName(name string) Keycode {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	return Keycode(C.SDL_GetKeyFromName(cstr))
}

func (kc Keycode) Scancode() Scancode {
	return Scancode(C.SDL_GetScancodeFromKey(C.SDL_Keycode(kc)))
}

func (kc Keycode) String() string {
	return C.GoString(C.SDL_GetKeyName(C.SDL_Keycode(kc)))
}

const (
	KeyUnknown  Keycode = C.SDLK_UNKNOWN
	Return      Keycode = C.SDLK_RETURN
	Escape      Keycode = C.SDLK_ESCAPE
	Backspace   Keycode = C.SDLK_BACKSPACE
	Tab         Keycode = C.SDLK_TAB
	Space       Keycode = C.SDLK_SPACE
	Exclaim     Keycode = C.SDLK_EXCLAIM
	DoubleQuote Keycode = C.SDLK_QUOTEDBL
	Hash        Keycode = C.SDLK_HASH
	Percent     Keycode = C.SDLK_PERCENT
	Dollar      Keycode = C.SDLK_DOLLAR
	Ampersand   Keycode = C.SDLK_AMPERSAND
	Quote       Keycode = C.SDLK_QUOTE
	LeftParen   Keycode = C.SDLK_LEFTPAREN
	RightParen  Keycode = C.SDLK_RIGHTPAREN
	Asterisk    Keycode = C.SDLK_ASTERISK
	Plus        Keycode = C.SDLK_PLUS
	Comma       Keycode = C.SDLK_COMMA
	Minus       Keycode = C.SDLK_MINUS
	Period      Keycode = C.SDLK_PERIOD
	Slash       Keycode = C.SDLK_SLASH

	Key0      Keycode = C.SDLK_0
	Key1      Keycode = C.SDLK_1
	Key2      Keycode = C.SDLK_2
	Key3      Keycode = C.SDLK_3
	Key4      Keycode = C.SDLK_4
	Key5      Keycode = C.SDLK_5
	Key6      Keycode = C.SDLK_6
	Key7      Keycode = C.SDLK_7
	Key8      Keycode = C.SDLK_8
	Key9      Keycode = C.SDLK_9
	Colon     Keycode = C.SDLK_COLON
	Semicolon Keycode = C.SDLK_SEMICOLON
	Less      Keycode = C.SDLK_LESS
	Equals    Keycode = C.SDLK_EQUALS
	Greater   Keycode = C.SDLK_GREATER
	Question  Keycode = C.SDLK_QUESTION
	At        Keycode = C.SDLK_AT

	// Skip uppercase letters
	KeyA         Keycode = C.SDLK_a
	KeyB         Keycode = C.SDLK_b
	KeyC         Keycode = C.SDLK_c
	KeyD         Keycode = C.SDLK_d
	KeyE         Keycode = C.SDLK_e
	KeyF         Keycode = C.SDLK_f
	KeyG         Keycode = C.SDLK_g
	KeyH         Keycode = C.SDLK_h
	KeyI         Keycode = C.SDLK_i
	KeyJ         Keycode = C.SDLK_j
	KeyK         Keycode = C.SDLK_k
	KeyL         Keycode = C.SDLK_l
	KeyM         Keycode = C.SDLK_m
	KeyN         Keycode = C.SDLK_n
	KeyO         Keycode = C.SDLK_o
	KeyP         Keycode = C.SDLK_p
	KeyQ         Keycode = C.SDLK_q
	KeyR         Keycode = C.SDLK_r
	KeyS         Keycode = C.SDLK_s
	KeyT         Keycode = C.SDLK_t
	KeyU         Keycode = C.SDLK_u
	KeyV         Keycode = C.SDLK_v
	KeyW         Keycode = C.SDLK_w
	KeyX         Keycode = C.SDLK_x
	KeyY         Keycode = C.SDLK_y
	KeyZ         Keycode = C.SDLK_z
	LeftBracket  Keycode = C.SDLK_LEFTBRACKET
	BackSlash    Keycode = C.SDLK_BACKSLASH
	RightBracket Keycode = C.SDLK_RIGHTBRACKET
	Caret        Keycode = C.SDLK_CARET
	Underscore   Keycode = C.SDLK_UNDERSCORE
	Backquote    Keycode = C.SDLK_BACKQUOTE
	CapsLock     Keycode = C.SDLK_CAPSLOCK

	F1  Keycode = C.SDLK_F1
	F2  Keycode = C.SDLK_F2
	F3  Keycode = C.SDLK_F3
	F4  Keycode = C.SDLK_F4
	F5  Keycode = C.SDLK_F5
	F6  Keycode = C.SDLK_F6
	F7  Keycode = C.SDLK_F7
	F8  Keycode = C.SDLK_F8
	F9  Keycode = C.SDLK_F9
	F10 Keycode = C.SDLK_F10
	F11 Keycode = C.SDLK_F11
	F12 Keycode = C.SDLK_F12

	PrintScreen Keycode = C.SDLK_PRINTSCREEN
	ScrollLock  Keycode = C.SDLK_SCROLLLOCK
	Pause       Keycode = C.SDLK_PAUSE
	Insert      Keycode = C.SDLK_INSERT
	Home        Keycode = C.SDLK_HOME
	PageUp      Keycode = C.SDLK_PAGEUP
	Delete      Keycode = C.SDLK_DELETE
	End         Keycode = C.SDLK_END
	PageDown    Keycode = C.SDLK_PAGEDOWN
	Right       Keycode = C.SDLK_RIGHT
	Left        Keycode = C.SDLK_LEFT
	Down        Keycode = C.SDLK_DOWN
	Up          Keycode = C.SDLK_UP

	NumLockClear Keycode = C.SDLK_NUMLOCKCLEAR
	NumDivide    Keycode = C.SDLK_KP_DIVIDE
	NumMultiply  Keycode = C.SDLK_KP_MULTIPLY
	NumMinus     Keycode = C.SDLK_KP_MINUS
	NumPlus      Keycode = C.SDLK_KP_PLUS
	NumEnter     Keycode = C.SDLK_KP_ENTER
	Num1         Keycode = C.SDLK_KP_1
	Num2         Keycode = C.SDLK_KP_2
	Num3         Keycode = C.SDLK_KP_3
	Num4         Keycode = C.SDLK_KP_4
	Num5         Keycode = C.SDLK_KP_5
	Num6         Keycode = C.SDLK_KP_6
	Num7         Keycode = C.SDLK_KP_7
	Num8         Keycode = C.SDLK_KP_8
	Num9         Keycode = C.SDLK_KP_9
	Num0         Keycode = C.SDLK_KP_0
	NumPeriod    Keycode = C.SDLK_KP_PERIOD

	Application    Keycode = C.SDLK_APPLICATION
	Power          Keycode = C.SDLK_POWER
	NumEquals      Keycode = C.SDLK_KP_EQUALS
	F13            Keycode = C.SDLK_F13
	F14            Keycode = C.SDLK_F14
	F15            Keycode = C.SDLK_F15
	F16            Keycode = C.SDLK_F16
	F17            Keycode = C.SDLK_F17
	F18            Keycode = C.SDLK_F18
	F19            Keycode = C.SDLK_F19
	F20            Keycode = C.SDLK_F20
	F21            Keycode = C.SDLK_F21
	F22            Keycode = C.SDLK_F22
	F23            Keycode = C.SDLK_F23
	F24            Keycode = C.SDLK_F24
	Execute        Keycode = C.SDLK_EXECUTE
	Help           Keycode = C.SDLK_HELP
	Menu           Keycode = C.SDLK_MENU
	Select         Keycode = C.SDLK_SELECT
	Stop           Keycode = C.SDLK_STOP
	Again          Keycode = C.SDLK_AGAIN
	Undo           Keycode = C.SDLK_UNDO
	Cut            Keycode = C.SDLK_CUT
	Copy           Keycode = C.SDLK_COPY
	Paste          Keycode = C.SDLK_PASTE
	Find           Keycode = C.SDLK_FIND
	Mute           Keycode = C.SDLK_MUTE
	Volumeup       Keycode = C.SDLK_VOLUMEUP
	Volumedown     Keycode = C.SDLK_VOLUMEDOWN
	NumComma       Keycode = C.SDLK_KP_COMMA
	NumEqualsas400 Keycode = C.SDLK_KP_EQUALSAS400

	Alterase   Keycode = C.SDLK_ALTERASE
	Sysreq     Keycode = C.SDLK_SYSREQ
	Cancel     Keycode = C.SDLK_CANCEL
	Clear      Keycode = C.SDLK_CLEAR
	Prior      Keycode = C.SDLK_PRIOR
	Return2    Keycode = C.SDLK_RETURN2
	Separator  Keycode = C.SDLK_SEPARATOR
	Out        Keycode = C.SDLK_OUT
	Oper       Keycode = C.SDLK_OPER
	Clearagain Keycode = C.SDLK_CLEARAGAIN
	Crsel      Keycode = C.SDLK_CRSEL
	Exsel      Keycode = C.SDLK_EXSEL

	Key00              Keycode = C.SDLK_KP_00
	Key000             Keycode = C.SDLK_KP_000
	Thousandsseparator Keycode = C.SDLK_THOUSANDSSEPARATOR
	Decimalseparator   Keycode = C.SDLK_DECIMALSEPARATOR
	Currencyunit       Keycode = C.SDLK_CURRENCYUNIT
	Currencysubunit    Keycode = C.SDLK_CURRENCYSUBUNIT
	Leftparen          Keycode = C.SDLK_KP_LEFTPAREN
	Rightparen         Keycode = C.SDLK_KP_RIGHTPAREN
	Leftbrace          Keycode = C.SDLK_KP_LEFTBRACE
	Rightbrace         Keycode = C.SDLK_KP_RIGHTBRACE
	NumTab             Keycode = C.SDLK_KP_TAB
	NumBackspace       Keycode = C.SDLK_KP_BACKSPACE
	NumA               Keycode = C.SDLK_KP_A
	NumB               Keycode = C.SDLK_KP_B
	NumC               Keycode = C.SDLK_KP_C
	NumD               Keycode = C.SDLK_KP_D
	NumE               Keycode = C.SDLK_KP_E
	NumF               Keycode = C.SDLK_KP_F
	Xor                Keycode = C.SDLK_KP_XOR
	NumPower           Keycode = C.SDLK_KP_POWER
	NumPercent         Keycode = C.SDLK_KP_PERCENT
	NumLess            Keycode = C.SDLK_KP_LESS
	NumGreater         Keycode = C.SDLK_KP_GREATER
	NumAmpersand       Keycode = C.SDLK_KP_AMPERSAND
	Dblampersand       Keycode = C.SDLK_KP_DBLAMPERSAND
	Verticalbar        Keycode = C.SDLK_KP_VERTICALBAR
	Dblverticalbar     Keycode = C.SDLK_KP_DBLVERTICALBAR
	NumColon           Keycode = C.SDLK_KP_COLON
	NumHash            Keycode = C.SDLK_KP_HASH
	NumSpace           Keycode = C.SDLK_KP_SPACE
	NumAt              Keycode = C.SDLK_KP_AT
	Exclam             Keycode = C.SDLK_KP_EXCLAM
	Memstore           Keycode = C.SDLK_KP_MEMSTORE
	Memrecall          Keycode = C.SDLK_KP_MEMRECALL
	Memclear           Keycode = C.SDLK_KP_MEMCLEAR
	Memadd             Keycode = C.SDLK_KP_MEMADD
	Memsubtract        Keycode = C.SDLK_KP_MEMSUBTRACT
	Memmultiply        Keycode = C.SDLK_KP_MEMMULTIPLY
	Memdivide          Keycode = C.SDLK_KP_MEMDIVIDE
	Plusminus          Keycode = C.SDLK_KP_PLUSMINUS
	NumClear           Keycode = C.SDLK_KP_CLEAR
	Clearentry         Keycode = C.SDLK_KP_CLEARENTRY
	Binary             Keycode = C.SDLK_KP_BINARY
	Octal              Keycode = C.SDLK_KP_OCTAL
	Decimal            Keycode = C.SDLK_KP_DECIMAL
	Hexadecimal        Keycode = C.SDLK_KP_HEXADECIMAL

	Lctrl  Keycode = C.SDLK_LCTRL
	Lshift Keycode = C.SDLK_LSHIFT
	Lalt   Keycode = C.SDLK_LALT
	Lgui   Keycode = C.SDLK_LGUI
	Rctrl  Keycode = C.SDLK_RCTRL
	Rshift Keycode = C.SDLK_RSHIFT
	Ralt   Keycode = C.SDLK_RALT
	Rgui   Keycode = C.SDLK_RGUI

	Mode Keycode = C.SDLK_MODE

	Audionext   Keycode = C.SDLK_AUDIONEXT
	Audioprev   Keycode = C.SDLK_AUDIOPREV
	Audiostop   Keycode = C.SDLK_AUDIOSTOP
	Audioplay   Keycode = C.SDLK_AUDIOPLAY
	Audiomute   Keycode = C.SDLK_AUDIOMUTE
	Mediaselect Keycode = C.SDLK_MEDIASELECT
	Www         Keycode = C.SDLK_WWW
	Mail        Keycode = C.SDLK_MAIL
	Calculator  Keycode = C.SDLK_CALCULATOR
	Computer    Keycode = C.SDLK_COMPUTER
	Search      Keycode = C.SDLK_AC_SEARCH
	AltHome     Keycode = C.SDLK_AC_HOME
	Back        Keycode = C.SDLK_AC_BACK
	Forward     Keycode = C.SDLK_AC_FORWARD
	AltStop     Keycode = C.SDLK_AC_STOP
	Refresh     Keycode = C.SDLK_AC_REFRESH
	Bookmarks   Keycode = C.SDLK_AC_BOOKMARKS

	Brightnessdown Keycode = C.SDLK_BRIGHTNESSDOWN
	Brightnessup   Keycode = C.SDLK_BRIGHTNESSUP
	Displayswitch  Keycode = C.SDLK_DISPLAYSWITCH
	Kbdillumtoggle Keycode = C.SDLK_KBDILLUMTOGGLE
	Kbdillumdown   Keycode = C.SDLK_KBDILLUMDOWN
	Kbdillumup     Keycode = C.SDLK_KBDILLUMUP
	Eject          Keycode = C.SDLK_EJECT
	Sleep          Keycode = C.SDLK_SLEEP
)
