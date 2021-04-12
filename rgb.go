package color

import (
	"image/color"
	"math"
)

// sRGB color space (24-bit)
type RGB struct {
	R, G, B uint8
}

func (rgb RGB) RGBA() (r, g, b, a uint32) {
	r = uint32(rgb.R)
	r |= r << 8
	g = uint32(rgb.G)
	g |= g << 8
	b = uint32(rgb.B)
	b |= b << 8
	a = 0xffff
	return
}

func (rgb RGB) DistanceTo(c color.Color) float64 {
	other := RGBModel.Convert(c).(RGB)
	return math.Sqrt(math.Pow(float64(other.R)-float64(rgb.R), 2) + math.Pow(float64(other.G)-float64(rgb.G), 2) + math.Pow(float64(other.B)-float64(rgb.B), 2))
}

var RGBModel color.Model = color.ModelFunc(rgbModel)

func rgbModel(c color.Color) color.Color {
	if _, ok := c.(RGB); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)}
}
