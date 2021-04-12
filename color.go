package color

import "image/color"

type Color interface {
	color.Color
	DistanceTo(c color.Color) float64
}
