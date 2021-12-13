package canvas

import (
	"math"
)

// Rect is a rectangle in 2D defined by a position and its width and height.
type Rect struct {
	X, Y, W, H float64
}

// Add returns a rect that encompasses both the current rect and the given rect.
func (r Rect) Add(q Rect) Rect {
	if q.W == 0.0 || q.H == 0 {
		return r
	} else if r.W == 0.0 || r.H == 0 {
		return q
	}
	x0 := math.Min(r.X, q.X)
	y0 := math.Min(r.Y, q.Y)
	x1 := math.Max(r.X+r.W, q.X+q.W)
	y1 := math.Max(r.Y+r.H, q.Y+q.H)
	return Rect{x0, y0, x1 - x0, y1 - y0}
}
