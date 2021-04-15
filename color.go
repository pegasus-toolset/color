// Package color is an extension of the standard library package image/color.
package color

import "image/color"

type Color interface {
	color.Color
	DistanceTo(c color.Color) float64
}
