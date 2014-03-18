package gui

import (
	"fmt"
	"image"

	"github.com/acsellers/sdl2"
)

type Screen struct {
	Background  Drawable
	Items       []Drawable
	ActiveItems []Clickable
	Window      *sdl2.Window
	Draw        func()
	Setup       func(*Screen)
	done        chan bool
}

func (s *Screen) SetWindow(w *sdl2.Window) {
	s.Window = w
}

func (s *Screen) Start() {
	if s.Draw != nil {
		s.Window.Screen = func() {
			s.Draw()
			s.Window.Present()
		}
	} else {
		s.Setup(s)
		s.Window.Screen = func() {
			for {
				sf, r := s.Background.Draw()
				e := s.Window.RenderSurface(sf, r)
				if e != nil {
					fmt.Println(e)
				}
				for _, d := range s.Items {
					sf, r := d.Draw()
					s.Window.RenderSurface(sf, r)
				}
				s.Window.Present()
			}
		}
	}
}

func (s *Screen) Stop() {
	s.Window.Screen = nil
}

type Drawable interface {
	Draw() (*sdl2.Surface, image.Rectangle)
	ActiveArea() image.Rectangle
}

type Clickable interface {
	Drawable
	Activate()
}

type MouseoverAware interface {
	Mouseover() bool
	SetMouseover(bool)
}
