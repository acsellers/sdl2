package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_scancode.h>
// #include <SDL2/SDL_keyboard.h>
import "C"
import "unsafe"

type Scancode int

func NewScancodeFromName(name string) Scancode {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	return Scancode(C.SDL_GetScancodeFromName(cstr))
}

func (sc Scancode) Key() Keycode {
	return Keycode(C.SDL_GetKeyFromScancode(C.SDL_Scancode(sc)))
}

func (sc Scancode) String() string {
	return C.GoString(C.SDL_GetScancodeName(C.SDL_Scancode(sc)))
}

// Scan codes from SDL_scancode, some comments reformatted
// to match standard Go comments.
const (
	ScanUnknown Scancode = C.SDL_SCANCODE_UNKNOWN
	ScanA       Scancode = C.SDL_SCANCODE_A
	ScanB       Scancode = C.SDL_SCANCODE_B
	ScanC       Scancode = C.SDL_SCANCODE_C
	ScanD       Scancode = C.SDL_SCANCODE_D
	ScanE       Scancode = C.SDL_SCANCODE_E
	ScanF       Scancode = C.SDL_SCANCODE_F
	ScanG       Scancode = C.SDL_SCANCODE_G
	ScanH       Scancode = C.SDL_SCANCODE_H
	ScanI       Scancode = C.SDL_SCANCODE_I
	ScanJ       Scancode = C.SDL_SCANCODE_J
	ScanK       Scancode = C.SDL_SCANCODE_K
	ScanL       Scancode = C.SDL_SCANCODE_L
	ScanM       Scancode = C.SDL_SCANCODE_M
	ScanN       Scancode = C.SDL_SCANCODE_N
	ScanO       Scancode = C.SDL_SCANCODE_O
	ScanP       Scancode = C.SDL_SCANCODE_P
	ScanQ       Scancode = C.SDL_SCANCODE_Q
	ScanR       Scancode = C.SDL_SCANCODE_R
	ScanS       Scancode = C.SDL_SCANCODE_S
	ScanT       Scancode = C.SDL_SCANCODE_T
	ScanU       Scancode = C.SDL_SCANCODE_U
	ScanV       Scancode = C.SDL_SCANCODE_V
	ScanW       Scancode = C.SDL_SCANCODE_W
	ScanX       Scancode = C.SDL_SCANCODE_X
	ScanY       Scancode = C.SDL_SCANCODE_Y
	ScanZ       Scancode = C.SDL_SCANCODE_Z

	Scan1 Scancode = C.SDL_SCANCODE_1
	Scan2 Scancode = C.SDL_SCANCODE_2
	Scan3 Scancode = C.SDL_SCANCODE_3
	Scan4 Scancode = C.SDL_SCANCODE_4
	Scan5 Scancode = C.SDL_SCANCODE_5
	Scan6 Scancode = C.SDL_SCANCODE_6
	Scan7 Scancode = C.SDL_SCANCODE_7
	Scan8 Scancode = C.SDL_SCANCODE_8
	Scan9 Scancode = C.SDL_SCANCODE_9
	Scan0 Scancode = C.SDL_SCANCODE_0

	ScanReturn    Scancode = C.SDL_SCANCODE_RETURN
	ScanEscape    Scancode = C.SDL_SCANCODE_ESCAPE
	ScanBackspace Scancode = C.SDL_SCANCODE_BACKSPACE
	ScanTab       Scancode = C.SDL_SCANCODE_TAB
	ScanSpace     Scancode = C.SDL_SCANCODE_SPACE

	ScanMinus        Scancode = C.SDL_SCANCODE_MINUS
	ScanEquals       Scancode = C.SDL_SCANCODE_EQUALS
	ScanLeftBracket  Scancode = C.SDL_SCANCODE_LEFTBRACKET
	ScanRightBracket Scancode = C.SDL_SCANCODE_RIGHTBRACKET

	/*
	   Located at the lower left of the return
	   key on ISO keyboards and at the right end
	   of the QWERTY row on ANSI keyboards.
	   Produces REVERSE SOLIDUS (backslash) and
	   VERTICAL LINE in a US layout, REVERSE
	   SOLIDUS and VERTICAL LINE in a UK Mac
	   layout, NUMBER SIGN and TILDE in a UK
	   Windows layout, DOLLAR SIGN and POUND SIGN
	   in a Swiss German layout, NUMBER SIGN and
	   APOSTROPHE in a German layout, GRAVE
	   ACCENT and POUND SIGN in a French Mac
	   layout, and ASTERISK and MICRO SIGN in a
	   French Windows layout.
	*/
	ScanBackslash Scancode = C.SDL_SCANCODE_BACKSLASH
	/*
	   ISO USB keyboards actually use this code
	   instead of 49 for the same key, but all
	   OSes I've seen treat the two codes
	   identically. So, as an implementor, unless
	   your keyboard generates both of those
	   codes and your OS treats them differently,
	   you should generate C.SDL_SCANCODE_BACKSLASH
	   instead of this code. As a user, you
	   should not rely on this code because SDL
	   will never generate it with most (all?)
	   keyboards.
	*/
	ScanNonUSHash  Scancode = C.SDL_SCANCODE_NONUSHASH
	ScanSemicolon  Scancode = C.SDL_SCANCODE_SEMICOLON
	ScanApostrophe Scancode = C.SDL_SCANCODE_APOSTROPHE
	/*
	   Located in the top left corner (on both ANSI
	   and ISO keyboards). Produces GRAVE ACCENT and
	   TILDE in a US Windows layout and in US and UK
	   Mac layouts on ANSI keyboards, GRAVE ACCENT
	   and NOT SIGN in a UK Windows layout, SECTION
	   SIGN and PLUS-MINUS SIGN in US and UK Mac
	   layouts on ISO keyboards, SECTION SIGN and
	   DEGREE SIGN in a Swiss German layout (Mac:
	   only on ISO keyboards), CIRCUMFLEX ACCENT and
	   DEGREE SIGN in a German layout (Mac: only on
	   ISO keyboards), SUPERSCRIPT TWO and TILDE in a
	   French Windows layout, COMMERCIAL AT and
	   NUMBER SIGN in a French Mac layout on ISO
	   keyboards, and LESS-THAN SIGN and GREATER-THAN
	   SIGN in a Swiss German, German, or French Mac
	   layout on ANSI keyboards.
	*/
	ScanGrave  Scancode = C.SDL_SCANCODE_GRAVE
	ScanComma  Scancode = C.SDL_SCANCODE_COMMA
	ScanPeriod Scancode = C.SDL_SCANCODE_PERIOD
	ScanSlash  Scancode = C.SDL_SCANCODE_SLASH

	ScanCapsLock Scancode = C.SDL_SCANCODE_CAPSLOCK

	ScanF1  Scancode = C.SDL_SCANCODE_F1
	ScanF2  Scancode = C.SDL_SCANCODE_F2
	ScanF3  Scancode = C.SDL_SCANCODE_F3
	ScanF4  Scancode = C.SDL_SCANCODE_F4
	ScanF5  Scancode = C.SDL_SCANCODE_F5
	ScanF6  Scancode = C.SDL_SCANCODE_F6
	ScanF7  Scancode = C.SDL_SCANCODE_F7
	ScanF8  Scancode = C.SDL_SCANCODE_F8
	ScanF9  Scancode = C.SDL_SCANCODE_F9
	ScanF10 Scancode = C.SDL_SCANCODE_F10
	ScanF11 Scancode = C.SDL_SCANCODE_F11
	ScanF12 Scancode = C.SDL_SCANCODE_F12

	ScanPrintScreen Scancode = C.SDL_SCANCODE_PRINTSCREEN
	ScanScrollLock  Scancode = C.SDL_SCANCODE_SCROLLLOCK
	ScanPause       Scancode = C.SDL_SCANCODE_PAUSE
	// insert on PC, help on some Mac keyboards (but does send code 73, not 117)
	ScanInsert   Scancode = C.SDL_SCANCODE_INSERT
	ScanHome     Scancode = C.SDL_SCANCODE_HOME
	ScanPageUp   Scancode = C.SDL_SCANCODE_PAGEUP
	ScanDelete   Scancode = C.SDL_SCANCODE_DELETE
	ScanEnd      Scancode = C.SDL_SCANCODE_END
	ScanPageDown Scancode = C.SDL_SCANCODE_PAGEDOWN
	ScanRight    Scancode = C.SDL_SCANCODE_RIGHT
	ScanLeft     Scancode = C.SDL_SCANCODE_LEFT
	ScanDown     Scancode = C.SDL_SCANCODE_DOWN
	ScanUp       Scancode = C.SDL_SCANCODE_UP

	// num lock on PC, clear on Mac keyboards
	ScanNumLockClear Scancode = C.SDL_SCANCODE_NUMLOCKCLEAR
	ScanNumDivide    Scancode = C.SDL_SCANCODE_KP_DIVIDE
	ScanNumMultiply  Scancode = C.SDL_SCANCODE_KP_MULTIPLY
	ScanNumMinus     Scancode = C.SDL_SCANCODE_KP_MINUS
	ScanNumPlus      Scancode = C.SDL_SCANCODE_KP_PLUS
	ScanNumEnter     Scancode = C.SDL_SCANCODE_KP_ENTER
	ScanNum1         Scancode = C.SDL_SCANCODE_KP_1
	ScanNum2         Scancode = C.SDL_SCANCODE_KP_2
	ScanNum3         Scancode = C.SDL_SCANCODE_KP_3
	ScanNum4         Scancode = C.SDL_SCANCODE_KP_4
	ScanNum5         Scancode = C.SDL_SCANCODE_KP_5
	ScanNum6         Scancode = C.SDL_SCANCODE_KP_6
	ScanNum7         Scancode = C.SDL_SCANCODE_KP_7
	ScanNum8         Scancode = C.SDL_SCANCODE_KP_8
	ScanNum9         Scancode = C.SDL_SCANCODE_KP_9
	ScanNum0         Scancode = C.SDL_SCANCODE_KP_0
	ScanNumPeriod    Scancode = C.SDL_SCANCODE_KP_PERIOD

	/*
	   This is the additional key that ISO
	   keyboards have over ANSI ones,
	   located between left shift and Y.
	   Produces GRAVE ACCENT and TILDE in a
	   US or UK Mac layout, REVERSE SOLIDUS
	   (backslash) and VERTICAL LINE in a
	   US or UK Windows layout, and
	   LESS-THAN SIGN and GREATER-THAN SIGN
	   in a Swiss German, German, or French
	   layout.
	*/
	ScanNonUSBackSlash Scancode = C.SDL_SCANCODE_NONUSBACKSLASH
	// windows contextual menu, compose
	ScanApplication Scancode = C.SDL_SCANCODE_APPLICATION
	/*
	   The USB document says this is a status flag,
	   not a physical key - but some Mac keyboards
	   do have a power key.
	*/
	ScanPower       Scancode = C.SDL_SCANCODE_POWER
	ScanNumEquals   Scancode = C.SDL_SCANCODE_KP_EQUALS
	ScanF13         Scancode = C.SDL_SCANCODE_F13
	ScanF14         Scancode = C.SDL_SCANCODE_F14
	ScanF15         Scancode = C.SDL_SCANCODE_F15
	ScanF16         Scancode = C.SDL_SCANCODE_F16
	ScanF17         Scancode = C.SDL_SCANCODE_F17
	ScanF18         Scancode = C.SDL_SCANCODE_F18
	ScanF19         Scancode = C.SDL_SCANCODE_F19
	ScanF20         Scancode = C.SDL_SCANCODE_F20
	ScanF21         Scancode = C.SDL_SCANCODE_F21
	ScanF22         Scancode = C.SDL_SCANCODE_F22
	ScanF23         Scancode = C.SDL_SCANCODE_F23
	ScanF24         Scancode = C.SDL_SCANCODE_F24
	ScanExecute     Scancode = C.SDL_SCANCODE_EXECUTE
	ScanHelp        Scancode = C.SDL_SCANCODE_HELP
	ScanMenu        Scancode = C.SDL_SCANCODE_MENU
	ScanSelect      Scancode = C.SDL_SCANCODE_SELECT
	ScanStop        Scancode = C.SDL_SCANCODE_STOP
	ScanAgain       Scancode = C.SDL_SCANCODE_AGAIN
	ScanUndo        Scancode = C.SDL_SCANCODE_UNDO
	ScanCut         Scancode = C.SDL_SCANCODE_CUT
	ScanCopy        Scancode = C.SDL_SCANCODE_COPY
	ScanPaste       Scancode = C.SDL_SCANCODE_PASTE
	ScanFind        Scancode = C.SDL_SCANCODE_FIND
	ScanMute        Scancode = C.SDL_SCANCODE_MUTE
	ScanVolumeUp    Scancode = C.SDL_SCANCODE_VOLUMEUP
	ScanVolumeDown  Scancode = C.SDL_SCANCODE_VOLUMEDOWN
	ScanNumComma    Scancode = C.SDL_SCANCODE_KP_COMMA
	ScanEqualsAS400 Scancode = C.SDL_SCANCODE_KP_EQUALSAS400

	// used on Asian keyboards, see footnotes in USB doc
	ScanInternational1 Scancode = C.SDL_SCANCODE_INTERNATIONAL1
	ScanInternational2 Scancode = C.SDL_SCANCODE_INTERNATIONAL2
	// Yen
	ScanInternational3 Scancode = C.SDL_SCANCODE_INTERNATIONAL3
	ScanInternational4 Scancode = C.SDL_SCANCODE_INTERNATIONAL4
	ScanInternational5 Scancode = C.SDL_SCANCODE_INTERNATIONAL5
	ScanInternational6 Scancode = C.SDL_SCANCODE_INTERNATIONAL6
	ScanInternational7 Scancode = C.SDL_SCANCODE_INTERNATIONAL7
	ScanInternational8 Scancode = C.SDL_SCANCODE_INTERNATIONAL8
	ScanInternational9 Scancode = C.SDL_SCANCODE_INTERNATIONAL9
	// Hangul/English toggle
	ScanLang1 Scancode = C.SDL_SCANCODE_LANG1
	// Hanja conversion
	ScanLang2 Scancode = C.SDL_SCANCODE_LANG2
	// Katakana
	ScanLang3 Scancode = C.SDL_SCANCODE_LANG3
	// Hiragana
	ScanLang4 Scancode = C.SDL_SCANCODE_LANG4
	// Zenkaku/Hankaku
	ScanLang5 Scancode = C.SDL_SCANCODE_LANG5
	// Erase-Eaze
	ScanAlterase   Scancode = C.SDL_SCANCODE_ALTERASE
	ScanSysreq     Scancode = C.SDL_SCANCODE_SYSREQ
	ScanCancel     Scancode = C.SDL_SCANCODE_CANCEL
	ScanClear      Scancode = C.SDL_SCANCODE_CLEAR
	ScanPrior      Scancode = C.SDL_SCANCODE_PRIOR
	ScanReturn2    Scancode = C.SDL_SCANCODE_RETURN2
	ScanSeparator  Scancode = C.SDL_SCANCODE_SEPARATOR
	ScanOut        Scancode = C.SDL_SCANCODE_OUT
	ScanOper       Scancode = C.SDL_SCANCODE_OPER
	ScanClearAgain Scancode = C.SDL_SCANCODE_CLEARAGAIN
	ScanCrsel      Scancode = C.SDL_SCANCODE_CRSEL
	ScanExsel      Scancode = C.SDL_SCANCODE_EXSEL

	Scan00                 Scancode = C.SDL_SCANCODE_KP_00
	Scan000                Scancode = C.SDL_SCANCODE_KP_000
	ScanThousandsSeparator Scancode = C.SDL_SCANCODE_THOUSANDSSEPARATOR
	ScanDecimalSeparator   Scancode = C.SDL_SCANCODE_DECIMALSEPARATOR
	ScanCurrencyUnit       Scancode = C.SDL_SCANCODE_CURRENCYUNIT
	ScanCurrencySubunit    Scancode = C.SDL_SCANCODE_CURRENCYSUBUNIT
	ScanLeftParen          Scancode = C.SDL_SCANCODE_KP_LEFTPAREN
	ScanRightParen         Scancode = C.SDL_SCANCODE_KP_RIGHTPAREN
	ScanLeftBrace          Scancode = C.SDL_SCANCODE_KP_LEFTBRACE
	ScanRightBrace         Scancode = C.SDL_SCANCODE_KP_RIGHTBRACE
	ScanNumTab             Scancode = C.SDL_SCANCODE_KP_TAB
	ScanNumBackspace       Scancode = C.SDL_SCANCODE_KP_BACKSPACE
	ScanNumA               Scancode = C.SDL_SCANCODE_KP_A
	ScanNumB               Scancode = C.SDL_SCANCODE_KP_B
	ScanNumC               Scancode = C.SDL_SCANCODE_KP_C
	ScanNumD               Scancode = C.SDL_SCANCODE_KP_D
	ScanNumE               Scancode = C.SDL_SCANCODE_KP_E
	ScanNumF               Scancode = C.SDL_SCANCODE_KP_F
	ScanXor                Scancode = C.SDL_SCANCODE_KP_XOR
	ScanNumPower           Scancode = C.SDL_SCANCODE_KP_POWER
	ScanPercent            Scancode = C.SDL_SCANCODE_KP_PERCENT
	ScanLess               Scancode = C.SDL_SCANCODE_KP_LESS
	ScanGreater            Scancode = C.SDL_SCANCODE_KP_GREATER
	ScanAmpersand          Scancode = C.SDL_SCANCODE_KP_AMPERSAND
	ScanDblAmpersand       Scancode = C.SDL_SCANCODE_KP_DBLAMPERSAND
	ScanVerticalBar        Scancode = C.SDL_SCANCODE_KP_VERTICALBAR
	ScanDblVerticalBar     Scancode = C.SDL_SCANCODE_KP_DBLVERTICALBAR
	ScanColon              Scancode = C.SDL_SCANCODE_KP_COLON
	ScanHash               Scancode = C.SDL_SCANCODE_KP_HASH
	ScanNumSpace           Scancode = C.SDL_SCANCODE_KP_SPACE
	ScanAt                 Scancode = C.SDL_SCANCODE_KP_AT
	ScanExclam             Scancode = C.SDL_SCANCODE_KP_EXCLAM
	ScanMemStore           Scancode = C.SDL_SCANCODE_KP_MEMSTORE
	ScanMemRecall          Scancode = C.SDL_SCANCODE_KP_MEMRECALL
	ScanMemClear           Scancode = C.SDL_SCANCODE_KP_MEMCLEAR
	ScanMemAdd             Scancode = C.SDL_SCANCODE_KP_MEMADD
	ScanMemSubtract        Scancode = C.SDL_SCANCODE_KP_MEMSUBTRACT
	ScanMemMultiply        Scancode = C.SDL_SCANCODE_KP_MEMMULTIPLY
	ScanMemDivide          Scancode = C.SDL_SCANCODE_KP_MEMDIVIDE
	ScanPlusMinus          Scancode = C.SDL_SCANCODE_KP_PLUSMINUS
	ScanNumClear           Scancode = C.SDL_SCANCODE_KP_CLEAR
	ScanClearEntry         Scancode = C.SDL_SCANCODE_KP_CLEARENTRY
	ScanBinary             Scancode = C.SDL_SCANCODE_KP_BINARY
	ScanOctal              Scancode = C.SDL_SCANCODE_KP_OCTAL
	ScanDecimal            Scancode = C.SDL_SCANCODE_KP_DECIMAL
	ScanHexadecimal        Scancode = C.SDL_SCANCODE_KP_HEXADECIMAL

	ScanLCtrl  Scancode = C.SDL_SCANCODE_LCTRL
	ScanLShift Scancode = C.SDL_SCANCODE_LSHIFT
	ScanLAlt   Scancode = C.SDL_SCANCODE_LALT
	ScanLGui   Scancode = C.SDL_SCANCODE_LGUI
	ScanRCtrl  Scancode = C.SDL_SCANCODE_RCTRL
	ScanRShift Scancode = C.SDL_SCANCODE_RSHIFT
	ScanRAlt   Scancode = C.SDL_SCANCODE_RALT
	ScanRGui   Scancode = C.SDL_SCANCODE_RGUI

	// These values are mapped from usage page 0x0C (USB consumer page).
	ScanAudioNext   Scancode = C.SDL_SCANCODE_AUDIONEXT
	ScanAudioPrev   Scancode = C.SDL_SCANCODE_AUDIOPREV
	ScanAudioStop   Scancode = C.SDL_SCANCODE_AUDIOSTOP
	ScanAudioPlay   Scancode = C.SDL_SCANCODE_AUDIOPLAY
	ScanAudioMute   Scancode = C.SDL_SCANCODE_AUDIOMUTE
	ScanMediaSelect Scancode = C.SDL_SCANCODE_MEDIASELECT
	ScanWww         Scancode = C.SDL_SCANCODE_WWW
	ScanMail        Scancode = C.SDL_SCANCODE_MAIL
	ScanCalculator  Scancode = C.SDL_SCANCODE_CALCULATOR
	ScanComputer    Scancode = C.SDL_SCANCODE_COMPUTER
	ScanSearch      Scancode = C.SDL_SCANCODE_AC_SEARCH
	ScanACHome      Scancode = C.SDL_SCANCODE_AC_HOME
	ScanBack        Scancode = C.SDL_SCANCODE_AC_BACK
	ScanForward     Scancode = C.SDL_SCANCODE_AC_FORWARD
	ScanACStop      Scancode = C.SDL_SCANCODE_AC_STOP
	ScanRefresh     Scancode = C.SDL_SCANCODE_AC_REFRESH
	ScanBookmarks   Scancode = C.SDL_SCANCODE_AC_BOOKMARKS

	// These are values that Christian Walther added (for mac keyboard?).
	ScanBrightnessDown Scancode = C.SDL_SCANCODE_BRIGHTNESSDOWN
	ScanBrightnessUp   Scancode = C.SDL_SCANCODE_BRIGHTNESSUP
	// display mirroring/dual display switch, video mode switch
	ScanDisplaySwitch Scancode = C.SDL_SCANCODE_DISPLAYSWITCH
	ScanKbIllumToggle Scancode = C.SDL_SCANCODE_KBDILLUMTOGGLE
	ScanKbIllumDown   Scancode = C.SDL_SCANCODE_KBDILLUMDOWN
	ScanKbIllumUp     Scancode = C.SDL_SCANCODE_KBDILLUMUP
	ScanEject         Scancode = C.SDL_SCANCODE_EJECT
	ScanSleep         Scancode = C.SDL_SCANCODE_SLEEP

	ScanApp1 Scancode = C.SDL_SCANCODE_APP1
	ScanApp2 Scancode = C.SDL_SCANCODE_APP2
)
