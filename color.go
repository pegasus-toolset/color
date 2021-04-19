// Package color is an extension of the standard library package image/color.
package color

import "image/color"

// Color uses the image/color.Color interface but also provides a function to
// measure the distance between this and another color. The color model of this
// color is used for the distance calculation.
type Color interface {
	color.Color
	DistanceTo(c color.Color) float64
}
