package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_scancode.h>
import "C"

type ScanCode uint32

// Scan codes from SDL_scancode, some comments reformatted
// to match standard Go comments.
const (
	ScanUnknown ScanCode = C.SDL_SCANCODE_UNKNOWN
	ScanA       ScanCode = C.SDL_SCANCODE_A
	ScanB       ScanCode = C.SDL_SCANCODE_B
	ScanC       ScanCode = C.SDL_SCANCODE_C
	ScanD       ScanCode = C.SDL_SCANCODE_D
	ScanE       ScanCode = C.SDL_SCANCODE_E
	ScanF       ScanCode = C.SDL_SCANCODE_F
	ScanG       ScanCode = C.SDL_SCANCODE_G
	ScanH       ScanCode = C.SDL_SCANCODE_H
	ScanI       ScanCode = C.SDL_SCANCODE_I
	ScanJ       ScanCode = C.SDL_SCANCODE_J
	ScanK       ScanCode = C.SDL_SCANCODE_K
	ScanL       ScanCode = C.SDL_SCANCODE_L
	ScanM       ScanCode = C.SDL_SCANCODE_M
	ScanN       ScanCode = C.SDL_SCANCODE_N
	ScanO       ScanCode = C.SDL_SCANCODE_O
	ScanP       ScanCode = C.SDL_SCANCODE_P
	ScanQ       ScanCode = C.SDL_SCANCODE_Q
	ScanR       ScanCode = C.SDL_SCANCODE_R
	ScanS       ScanCode = C.SDL_SCANCODE_S
	ScanT       ScanCode = C.SDL_SCANCODE_T
	ScanU       ScanCode = C.SDL_SCANCODE_U
	ScanV       ScanCode = C.SDL_SCANCODE_V
	ScanW       ScanCode = C.SDL_SCANCODE_W
	ScanX       ScanCode = C.SDL_SCANCODE_X
	ScanY       ScanCode = C.SDL_SCANCODE_Y
	ScanZ       ScanCode = C.SDL_SCANCODE_Z

	Scan1 ScanCode = C.SDL_SCANCODE_1
	Scan2 ScanCode = C.SDL_SCANCODE_2
	Scan3 ScanCode = C.SDL_SCANCODE_3
	Scan4 ScanCode = C.SDL_SCANCODE_4
	Scan5 ScanCode = C.SDL_SCANCODE_5
	Scan6 ScanCode = C.SDL_SCANCODE_6
	Scan7 ScanCode = C.SDL_SCANCODE_7
	Scan8 ScanCode = C.SDL_SCANCODE_8
	Scan9 ScanCode = C.SDL_SCANCODE_9
	Scan0 ScanCode = C.SDL_SCANCODE_0

	ScanReturn    ScanCode = C.SDL_SCANCODE_RETURN
	ScanEscape    ScanCode = C.SDL_SCANCODE_ESCAPE
	ScanBackspace ScanCode = C.SDL_SCANCODE_BACKSPACE
	ScanTab       ScanCode = C.SDL_SCANCODE_TAB
	ScanSpace     ScanCode = C.SDL_SCANCODE_SPACE

	ScanMinus        ScanCode = C.SDL_SCANCODE_MINUS
	ScanEquals       ScanCode = C.SDL_SCANCODE_EQUALS
	ScanLeftBracket  ScanCode = C.SDL_SCANCODE_LEFTBRACKET
	ScanRightBracket ScanCode = C.SDL_SCANCODE_RIGHTBRACKET

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
	ScanBackslash ScanCode = C.SDL_SCANCODE_BACKSLASH
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
	ScanNonUSHash  ScanCode = C.SDL_SCANCODE_NONUSHASH
	ScanSemicolon  ScanCode = C.SDL_SCANCODE_SEMICOLON
	ScanApostrophe ScanCode = C.SDL_SCANCODE_APOSTROPHE
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
	ScanGrave  ScanCode = C.SDL_SCANCODE_GRAVE
	ScanComma  ScanCode = C.SDL_SCANCODE_COMMA
	ScanPeriod ScanCode = C.SDL_SCANCODE_PERIOD
	ScanSlash  ScanCode = C.SDL_SCANCODE_SLASH

	ScanCapsLock ScanCode = C.SDL_SCANCODE_CAPSLOCK

	ScanF1  ScanCode = C.SDL_SCANCODE_F1
	ScanF2  ScanCode = C.SDL_SCANCODE_F2
	ScanF3  ScanCode = C.SDL_SCANCODE_F3
	ScanF4  ScanCode = C.SDL_SCANCODE_F4
	ScanF5  ScanCode = C.SDL_SCANCODE_F5
	ScanF6  ScanCode = C.SDL_SCANCODE_F6
	ScanF7  ScanCode = C.SDL_SCANCODE_F7
	ScanF8  ScanCode = C.SDL_SCANCODE_F8
	ScanF9  ScanCode = C.SDL_SCANCODE_F9
	ScanF10 ScanCode = C.SDL_SCANCODE_F10
	ScanF11 ScanCode = C.SDL_SCANCODE_F11
	ScanF12 ScanCode = C.SDL_SCANCODE_F12

	ScanPrintScreen ScanCode = C.SDL_SCANCODE_PRINTSCREEN
	ScanScrollLock  ScanCode = C.SDL_SCANCODE_SCROLLLOCK
	ScanPause       ScanCode = C.SDL_SCANCODE_PAUSE
	// insert on PC, help on some Mac keyboards (but does send code 73, not 117)
	ScanInsert   ScanCode = C.SDL_SCANCODE_INSERT
	ScanHome     ScanCode = C.SDL_SCANCODE_HOME
	ScanPageUp   ScanCode = C.SDL_SCANCODE_PAGEUP
	ScanDelete   ScanCode = C.SDL_SCANCODE_DELETE
	ScanEnd      ScanCode = C.SDL_SCANCODE_END
	ScanPageDown ScanCode = C.SDL_SCANCODE_PAGEDOWN
	ScanRight    ScanCode = C.SDL_SCANCODE_RIGHT
	ScanLeft     ScanCode = C.SDL_SCANCODE_LEFT
	ScanDown     ScanCode = C.SDL_SCANCODE_DOWN
	ScanUp       ScanCode = C.SDL_SCANCODE_UP

	// num lock on PC, clear on Mac keyboards
	ScanNumLockClear ScanCode = C.SDL_SCANCODE_NUMLOCKCLEAR
	ScanNumDivide    ScanCode = C.SDL_SCANCODE_KP_DIVIDE
	ScanNumMultiply  ScanCode = C.SDL_SCANCODE_KP_MULTIPLY
	ScanNumMinus     ScanCode = C.SDL_SCANCODE_KP_MINUS
	ScanNumPlus      ScanCode = C.SDL_SCANCODE_KP_PLUS
	ScanNumEnter     ScanCode = C.SDL_SCANCODE_KP_ENTER
	ScanNum1         ScanCode = C.SDL_SCANCODE_KP_1
	ScanNum2         ScanCode = C.SDL_SCANCODE_KP_2
	ScanNum3         ScanCode = C.SDL_SCANCODE_KP_3
	ScanNum4         ScanCode = C.SDL_SCANCODE_KP_4
	ScanNum5         ScanCode = C.SDL_SCANCODE_KP_5
	ScanNum6         ScanCode = C.SDL_SCANCODE_KP_6
	ScanNum7         ScanCode = C.SDL_SCANCODE_KP_7
	ScanNum8         ScanCode = C.SDL_SCANCODE_KP_8
	ScanNum9         ScanCode = C.SDL_SCANCODE_KP_9
	ScanNum0         ScanCode = C.SDL_SCANCODE_KP_0
	ScanNumPeriod    ScanCode = C.SDL_SCANCODE_KP_PERIOD

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
	ScanNonUSBackSlash ScanCode = C.SDL_SCANCODE_NONUSBACKSLASH
	// windows contextual menu, compose
	ScanApplication ScanCode = C.SDL_SCANCODE_APPLICATION
	/*
	   The USB document says this is a status flag,
	   not a physical key - but some Mac keyboards
	   do have a power key.
	*/
	ScanPower       ScanCode = C.SDL_SCANCODE_POWER
	ScanNumEquals   ScanCode = C.SDL_SCANCODE_KP_EQUALS
	ScanF13         ScanCode = C.SDL_SCANCODE_F13
	ScanF14         ScanCode = C.SDL_SCANCODE_F14
	ScanF15         ScanCode = C.SDL_SCANCODE_F15
	ScanF16         ScanCode = C.SDL_SCANCODE_F16
	ScanF17         ScanCode = C.SDL_SCANCODE_F17
	ScanF18         ScanCode = C.SDL_SCANCODE_F18
	ScanF19         ScanCode = C.SDL_SCANCODE_F19
	ScanF20         ScanCode = C.SDL_SCANCODE_F20
	ScanF21         ScanCode = C.SDL_SCANCODE_F21
	ScanF22         ScanCode = C.SDL_SCANCODE_F22
	ScanF23         ScanCode = C.SDL_SCANCODE_F23
	ScanF24         ScanCode = C.SDL_SCANCODE_F24
	ScanExecute     ScanCode = C.SDL_SCANCODE_EXECUTE
	ScanHelp        ScanCode = C.SDL_SCANCODE_HELP
	ScanMenu        ScanCode = C.SDL_SCANCODE_MENU
	ScanSelect      ScanCode = C.SDL_SCANCODE_SELECT
	ScanStop        ScanCode = C.SDL_SCANCODE_STOP
	ScanAgain       ScanCode = C.SDL_SCANCODE_AGAIN
	ScanUndo        ScanCode = C.SDL_SCANCODE_UNDO
	ScanCut         ScanCode = C.SDL_SCANCODE_CUT
	ScanCopy        ScanCode = C.SDL_SCANCODE_COPY
	ScanPaste       ScanCode = C.SDL_SCANCODE_PASTE
	ScanFind        ScanCode = C.SDL_SCANCODE_FIND
	ScanMute        ScanCode = C.SDL_SCANCODE_MUTE
	ScanVolumeUp    ScanCode = C.SDL_SCANCODE_VOLUMEUP
	ScanVolumeDown  ScanCode = C.SDL_SCANCODE_VOLUMEDOWN
	ScanNumComma    ScanCode = C.SDL_SCANCODE_KP_COMMA
	ScanEqualsAS400 ScanCode = C.SDL_SCANCODE_KP_EQUALSAS400

	// used on Asian keyboards, see footnotes in USB doc
	ScanInternational1 ScanCode = C.SDL_SCANCODE_INTERNATIONAL1
	ScanInternational2 ScanCode = C.SDL_SCANCODE_INTERNATIONAL2
	// Yen
	ScanInternational3 ScanCode = C.SDL_SCANCODE_INTERNATIONAL3
	ScanInternational4 ScanCode = C.SDL_SCANCODE_INTERNATIONAL4
	ScanInternational5 ScanCode = C.SDL_SCANCODE_INTERNATIONAL5
	ScanInternational6 ScanCode = C.SDL_SCANCODE_INTERNATIONAL6
	ScanInternational7 ScanCode = C.SDL_SCANCODE_INTERNATIONAL7
	ScanInternational8 ScanCode = C.SDL_SCANCODE_INTERNATIONAL8
	ScanInternational9 ScanCode = C.SDL_SCANCODE_INTERNATIONAL9
	// Hangul/English toggle
	ScanLang1 ScanCode = C.SDL_SCANCODE_LANG1
	// Hanja conversion
	ScanLang2 ScanCode = C.SDL_SCANCODE_LANG2
	// Katakana
	ScanLang3 ScanCode = C.SDL_SCANCODE_LANG3
	// Hiragana
	ScanLang4 ScanCode = C.SDL_SCANCODE_LANG4
	// Zenkaku/Hankaku
	ScanLang5 ScanCode = C.SDL_SCANCODE_LANG5
	// Erase-Eaze
	ScanAlterase   ScanCode = C.SDL_SCANCODE_ALTERASE
	ScanSysreq     ScanCode = C.SDL_SCANCODE_SYSREQ
	ScanCancel     ScanCode = C.SDL_SCANCODE_CANCEL
	ScanClear      ScanCode = C.SDL_SCANCODE_CLEAR
	ScanPrior      ScanCode = C.SDL_SCANCODE_PRIOR
	ScanReturn2    ScanCode = C.SDL_SCANCODE_RETURN2
	ScanSeparator  ScanCode = C.SDL_SCANCODE_SEPARATOR
	ScanOut        ScanCode = C.SDL_SCANCODE_OUT
	ScanOper       ScanCode = C.SDL_SCANCODE_OPER
	ScanClearAgain ScanCode = C.SDL_SCANCODE_CLEARAGAIN
	ScanCrsel      ScanCode = C.SDL_SCANCODE_CRSEL
	ScanExsel      ScanCode = C.SDL_SCANCODE_EXSEL

	Scan00                 ScanCode = C.SDL_SCANCODE_KP_00
	Scan000                ScanCode = C.SDL_SCANCODE_KP_000
	ScanThousandsSeparator ScanCode = C.SDL_SCANCODE_THOUSANDSSEPARATOR
	ScanDecimalSeparator   ScanCode = C.SDL_SCANCODE_DECIMALSEPARATOR
	ScanCurrencyUnit       ScanCode = C.SDL_SCANCODE_CURRENCYUNIT
	ScanCurrencySubunit    ScanCode = C.SDL_SCANCODE_CURRENCYSUBUNIT
	ScanLeftParen          ScanCode = C.SDL_SCANCODE_KP_LEFTPAREN
	ScanRightParen         ScanCode = C.SDL_SCANCODE_KP_RIGHTPAREN
	ScanLeftBrace          ScanCode = C.SDL_SCANCODE_KP_LEFTBRACE
	ScanRightBrace         ScanCode = C.SDL_SCANCODE_KP_RIGHTBRACE
	ScanNumTab             ScanCode = C.SDL_SCANCODE_KP_TAB
	ScanNumBackspace       ScanCode = C.SDL_SCANCODE_KP_BACKSPACE
	ScanNumA               ScanCode = C.SDL_SCANCODE_KP_A
	ScanNumB               ScanCode = C.SDL_SCANCODE_KP_B
	ScanNumC               ScanCode = C.SDL_SCANCODE_KP_C
	ScanNumD               ScanCode = C.SDL_SCANCODE_KP_D
	ScanNumE               ScanCode = C.SDL_SCANCODE_KP_E
	ScanNumF               ScanCode = C.SDL_SCANCODE_KP_F
	ScanXor                ScanCode = C.SDL_SCANCODE_KP_XOR
	ScanNumPower           ScanCode = C.SDL_SCANCODE_KP_POWER
	ScanPercent            ScanCode = C.SDL_SCANCODE_KP_PERCENT
	ScanLess               ScanCode = C.SDL_SCANCODE_KP_LESS
	ScanGreater            ScanCode = C.SDL_SCANCODE_KP_GREATER
	ScanAmpersand          ScanCode = C.SDL_SCANCODE_KP_AMPERSAND
	ScanDblAmpersand       ScanCode = C.SDL_SCANCODE_KP_DBLAMPERSAND
	ScanVerticalBar        ScanCode = C.SDL_SCANCODE_KP_VERTICALBAR
	ScanDblVerticalBar     ScanCode = C.SDL_SCANCODE_KP_DBLVERTICALBAR
	ScanColon              ScanCode = C.SDL_SCANCODE_KP_COLON
	ScanHash               ScanCode = C.SDL_SCANCODE_KP_HASH
	ScanNumSpace           ScanCode = C.SDL_SCANCODE_KP_SPACE
	ScanAt                 ScanCode = C.SDL_SCANCODE_KP_AT
	ScanExclam             ScanCode = C.SDL_SCANCODE_KP_EXCLAM
	ScanMemStore           ScanCode = C.SDL_SCANCODE_KP_MEMSTORE
	ScanMemRecall          ScanCode = C.SDL_SCANCODE_KP_MEMRECALL
	ScanMemClear           ScanCode = C.SDL_SCANCODE_KP_MEMCLEAR
	ScanMemAdd             ScanCode = C.SDL_SCANCODE_KP_MEMADD
	ScanMemSubtract        ScanCode = C.SDL_SCANCODE_KP_MEMSUBTRACT
	ScanMemMultiply        ScanCode = C.SDL_SCANCODE_KP_MEMMULTIPLY
	ScanMemDivide          ScanCode = C.SDL_SCANCODE_KP_MEMDIVIDE
	ScanPlusMinus          ScanCode = C.SDL_SCANCODE_KP_PLUSMINUS
	ScanNumClear           ScanCode = C.SDL_SCANCODE_KP_CLEAR
	ScanClearEntry         ScanCode = C.SDL_SCANCODE_KP_CLEARENTRY
	ScanBinary             ScanCode = C.SDL_SCANCODE_KP_BINARY
	ScanOctal              ScanCode = C.SDL_SCANCODE_KP_OCTAL
	ScanDecimal            ScanCode = C.SDL_SCANCODE_KP_DECIMAL
	ScanHexadecimal        ScanCode = C.SDL_SCANCODE_KP_HEXADECIMAL

	ScanLCtrl  ScanCode = C.SDL_SCANCODE_LCTRL
	ScanLShift ScanCode = C.SDL_SCANCODE_LSHIFT
	ScanLAlt   ScanCode = C.SDL_SCANCODE_LALT
	ScanLGui   ScanCode = C.SDL_SCANCODE_LGUI
	ScanRCtrl  ScanCode = C.SDL_SCANCODE_RCTRL
	ScanRShift ScanCode = C.SDL_SCANCODE_RSHIFT
	ScanRAlt   ScanCode = C.SDL_SCANCODE_RALT
	ScanRGui   ScanCode = C.SDL_SCANCODE_RGUI

	// These values are mapped from usage page 0x0C (USB consumer page).
	ScanAudioNext   ScanCode = C.SDL_SCANCODE_AUDIONEXT
	ScanAudioPrev   ScanCode = C.SDL_SCANCODE_AUDIOPREV
	ScanAudioStop   ScanCode = C.SDL_SCANCODE_AUDIOSTOP
	ScanAudioPlay   ScanCode = C.SDL_SCANCODE_AUDIOPLAY
	ScanAudioMute   ScanCode = C.SDL_SCANCODE_AUDIOMUTE
	ScanMediaSelect ScanCode = C.SDL_SCANCODE_MEDIASELECT
	ScanWww         ScanCode = C.SDL_SCANCODE_WWW
	ScanMail        ScanCode = C.SDL_SCANCODE_MAIL
	ScanCalculator  ScanCode = C.SDL_SCANCODE_CALCULATOR
	ScanComputer    ScanCode = C.SDL_SCANCODE_COMPUTER
	ScanSearch      ScanCode = C.SDL_SCANCODE_AC_SEARCH
	ScanACHome      ScanCode = C.SDL_SCANCODE_AC_HOME
	ScanBack        ScanCode = C.SDL_SCANCODE_AC_BACK
	ScanForward     ScanCode = C.SDL_SCANCODE_AC_FORWARD
	ScanACStop      ScanCode = C.SDL_SCANCODE_AC_STOP
	ScanRefresh     ScanCode = C.SDL_SCANCODE_AC_REFRESH
	ScanBookmarks   ScanCode = C.SDL_SCANCODE_AC_BOOKMARKS

	// These are values that Christian Walther added (for mac keyboard?).
	ScanBrightnessDown ScanCode = C.SDL_SCANCODE_BRIGHTNESSDOWN
	ScanBrightnessUp   ScanCode = C.SDL_SCANCODE_BRIGHTNESSUP
	// display mirroring/dual display switch, video mode switch
	ScanDisplaySwitch ScanCode = C.SDL_SCANCODE_DISPLAYSWITCH
	ScanKbIllumToggle ScanCode = C.SDL_SCANCODE_KBDILLUMTOGGLE
	ScanKbIllumDown   ScanCode = C.SDL_SCANCODE_KBDILLUMDOWN
	ScanKbIllumUp     ScanCode = C.SDL_SCANCODE_KBDILLUMUP
	ScanEject         ScanCode = C.SDL_SCANCODE_EJECT
	ScanSleep         ScanCode = C.SDL_SCANCODE_SLEEP

	ScanApp1 ScanCode = C.SDL_SCANCODE_APP1
	ScanApp2 ScanCode = C.SDL_SCANCODE_APP2
)
