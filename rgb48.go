package color

import "image/color"

type RGB48 struct {
	R, G, B uint16
}

func (rgb48 RGB48) RGBA() (r, g, b, a uint32) {
	return uint32(rgb48.R), uint32(rgb48.G), uint32(rgb48.B), 0xffff
}

var RGB48Model color.Model = color.ModelFunc(rgb48Model)

func rgb48Model(c color.Color) color.Color {
	if _, ok := c.(RGB48); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB48{uint16(r), uint16(g), uint16(b)}
}
