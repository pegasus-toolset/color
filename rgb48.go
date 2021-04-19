package color

import (
	"image/color"
	"math"
)

// RGB48 represents a color in the sRGB color space (48-bit).
type RGB48 struct {
	R, G, B uint16
}

// RGBA returns the red, green, blue and alpha values for the color. Each value
// ranges within [0, 0xffff], but is represented by a uint32 so that multiplying
// by a blend factor up to 0xffff will not overflow.
func (rgb48 RGB48) RGBA() (r, g, b, a uint32) {
	return uint32(rgb48.R), uint32(rgb48.G), uint32(rgb48.B), 0xffff
}

func (rgb48 RGB48) DistanceTo(c color.Color) float64 {
	other := RGB48Model.Convert(c).(RGB48)
	return math.Sqrt(math.Pow(float64(other.R)-float64(rgb48.R), 2) + math.Pow(float64(other.G)-float64(rgb48.G), 2) + math.Pow(float64(other.B)-float64(rgb48.B), 2))
}

var RGB48Model color.Model = color.ModelFunc(rgb48Model)

func rgb48Model(c color.Color) color.Color {
	if _, ok := c.(RGB48); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB48{uint16(r), uint16(g), uint16(b)}
}
