package gui

import (
	"image"

	"github.com/acsellers/sdl2"
)

type Screen struct {
	Background    Drawable
	Items         []Drawable
	ActiveItems   []Clickable
	Window        *sdl2.Window
	AlternateDraw func()
	Setup         func(*sdl2.Window)
	done          chan bool
}

func (s *Screen) Start(w *sdl2.Window) {
	s.Window = w

	if s.AlternateDraw != nil {
		s.Window.Screen = func() {
			s.AlternateDraw()
			s.Window.Present()
		}
	} else {
		if s.Setup != nil {
			s.Setup(s.Window)
		}
		s.Window.Screen = func() {
			for {
				s.Background.Draw(s.Window)
				for _, d := range s.Items {
					d.Draw(s.Window)
				}
				for _, ai := range s.ActiveItems {
					ai.Draw(s.Window)
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
	Draw(*sdl2.Window)
}

type Background struct {
	S *sdl2.Surface
}

func (b *Background) Draw(w *sdl2.Window) {
	w.RenderBackground(b.S)
}

type StaticSurface struct {
	S         *sdl2.Surface
	Placement image.Rectangle
}

func (ss *StaticSurface) Draw(w *sdl2.Window) {
	w.RenderSurface(ss.S, ss.Placement)
}

type Clickable interface {
	Drawable
	ActiveArea() image.Rectangle
	Activate()
}

type MouseoverAware interface {
	Mouseover() bool
	SetMouseover(bool)
}
