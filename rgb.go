package color

import "image/color"

type RGB struct {
	R, G, B uint8
}

func (rgb RGB) RGBA() (r, g, b, a uint32) {
	return uint32(rgb.R) << 8, uint32(rgb.G) << 8, uint32(rgb.B) << 8, 0xffff
}

var RGBModel color.Model = color.ModelFunc(rgbModel)

func rgbModel(c color.Color) color.Color {
	if _, ok := c.(RGB); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)}
}
