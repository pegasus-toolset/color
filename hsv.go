package color

import (
	"image/color"
	"math"
)

// HSV color space
type HSV struct {
	H, S, V float64
}

func (hsv HSV) RGBA() (r, g, b, a uint32) {
	chroma := hsv.V * hsv.S

	h := hsv.H / 60
	x := chroma * (1 - math.Abs(math.Mod(h, 2)-1))

	var rFloat, gFloat, bFloat float64 = chroma, x, 0

	if h >= 5 {
		rFloat = chroma
		gFloat = 0
		bFloat = x
	} else if h >= 4 {
		rFloat = x
		gFloat = 0
		bFloat = chroma
	} else if h >= 3 {
		rFloat = 0
		gFloat = x
		bFloat = chroma
	} else if h >= 2 {
		rFloat = 0
		gFloat = chroma
		bFloat = x
	} else if h >= 1 {
		rFloat = x
		gFloat = chroma
		bFloat = 0
	}

	m := hsv.V - chroma

	return uint32((rFloat + m) * 0xffff), uint32((gFloat + m) * 0xffff), uint32((bFloat + m) * 0xffff), 0xffff
}

var HSVModel color.Model = color.ModelFunc(hsvModel)

func hsvModel(c color.Color) color.Color {
	if _, ok := c.(HSV); ok {
		return c
	}
	rInt, gInt, bInt, _ := c.RGBA()
	r := float64(rInt) / 0xffff
	g := float64(gInt) / 0xffff
	b := float64(bInt) / 0xffff

	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))
	chroma := max - min

	var h float64 = 0
	if max == r {
		h = math.Mod((g-b)/chroma, 6)
	} else if max == g {
		h = (b-r)/chroma + 2
	} else if max == b {
		h = (r-g)/chroma + 4
	}
	h *= 60

	var s float64 = 0
	if max > 0 {
		s = chroma / max
	}

	return HSV{h, s, max}
}
