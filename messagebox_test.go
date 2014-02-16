package sdl2

import (
	"fmt"
	"os"
)

func ExampleNewSimpleMessageBox() {
	// Create an error message to the user that tells them their OpenGL
	// version is too old to run Awesome Game.
	NewSimpleMessageBox(
		Error,
		"Awesome Game 1.0.0",
		"You have OpenGL 2.0, this game requires 2.1. Please upgrade your GPU drivers",
	)
}

func ExampleMessageBox() {
	var mode int
	mb := MessageBox{
		Title:   "Awesome Game 1.0.1",
		Message: "Which Window Mode would you like?",
		Buttons: []MessageButton{
			MessageButton{
				Submit: true,
				Text:   "Fullscreen",
				Click: func() {
					mode = 1
				},
			},
			MessageButton{
				Text: "Borderless",
				Click: func() {
					mode = 2
				},
			},
			MessageButton{
				Text: "Windowed",
				Click: func() {
					mode = 3
				},
			},
		},
	}
	err := mb.Show()
	if err != nil {
		fmt.Println("Could not load messagebox, environment be broke")
		os.Exit(1)
	}
}
