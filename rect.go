package sdl2

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL_rect.h>
import "C"
import "image"

/*
SDL2 has all sort of rectangle functions, here are the ones that
duplicate existing image.Rectangle functions

SDL_RectEmpty     => image.Rectangle.Empty
SDL_RectEquals    => image.Rectangle.Eq
SDL_IntersectRect => image.Rectangle.Intersect != image.ZR
SDL_UnionRect     => image.Rectangle.Union

*/

// EnclosePoints will determine whether any of the points in the slice pts are
// enclosed by the clipping rectangle, then it will determine the minimal clipping
// rectangle for the slice of points and return that as an image.Rectangle.
func EnclosePoints(clipping image.Rectangle, pts []image.Point) (bool, image.Rectangle) {
	var r C.SDL_Rect
	cpts := make([]C.SDL_Point, len(pts))
	for i, pt := range pts {
		cpts[i].x, cpts[i].y = C.int(pt.X), C.int(pt.Y)
	}
	clipping.Canon()
	cr := &C.SDL_Rect{
		C.int(clipping.Min.X),
		C.int(clipping.Min.Y),
		C.int(clipping.Dx()),
		C.int(clipping.Dy()),
	}
	b := C.SDL_EnclosePoints(&cpts[0], C.int(len(cpts)), cr, &r)
	er := image.Rect(int(r.x), int(r.y), int(r.x+r.w), int(r.y+r.h))
	if b == C.SDL_TRUE {
		return true, er
	}
	return false, er
}

// IntersectRectAndLine will determine whether the line formed by the points a and b
// intersects the Rectangle r. If so it will return true and a clipped version of the
// between a and b using the points ca and cb. If the line does not intersect the
// rectangle, this will return false and two zero points as ca and cb.
func IntersectRectAndLine(r image.Rectangle, a, b image.Point) (inside bool, ca, cb image.Point) {
	r.Canon()
	sr := C.SDL_Rect{
		x: C.int(r.Min.X),
		y: C.int(r.Min.Y),
		w: C.int(r.Dx()),
		h: C.int(r.Dy()),
	}
	var ax, bx, ay, by C.int
	ax, ay = C.int(a.X), C.int(a.Y)
	bx, by = C.int(b.X), C.int(b.Y)
	cbool := C.SDL_IntersectRectAndLine(&sr, &ax, &ay, &bx, &by)
	inside = cbool == C.SDL_TRUE
	if inside {
		ca = image.Point{int(ax), int(ay)}
		cb = image.Point{int(bx), int(by)}
	}
	return
}
