package main

import (
	"fmt"
	"image/color"

	"github.com/acsellers/sdl2"
)

func main() {
	sdl2.NewSimpleMessageBox(sdl2.Warning, "Simple Title", "Simple Message")

	mb := sdl2.MessageBox{
		MessageBoxType: sdl2.Error,
		Title:          "Complex Title",
		Message:        "Complex Message",
		Parent:         nil,
		Buttons: []sdl2.MessageButton{
			sdl2.MessageButton{
				Submit: true,
				Text:   "Yes",
				Click: func() {
					fmt.Println("You clicked Yes")
				},
			},
			sdl2.MessageButton{
				Quit: true,
				Text: "No",
				Click: func() {
					fmt.Println("You clicked No")
				},
			},
		},
		Colors: &sdl2.MessageBoxColorScheme{
			Background:       color.RGBA{0xff, 0x0, 0xff, 0xff},
			Text:             color.RGBA{0xff, 0x0, 0x0, 0xff},
			ButtonBorder:     color.RGBA{0xff, 0xff, 0x0, 0xff},
			ButtonBackground: color.RGBA{0x0, 0x0, 0xff, 0xff},
			ButtonSelected:   color.RGBA{0xff, 0xff, 0xff, 0xff},
		},
	}
	mb.Show()
}
