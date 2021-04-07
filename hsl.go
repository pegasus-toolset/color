package color

import (
	"image/color"
	"math"
)

type HSL struct {
	H, S, L float64
}

func (hsl HSL) RGBA() (r, g, b, a uint32) {
	chroma := (1 - math.Abs(2*hsl.L-1)) * hsl.S

	h := hsl.H / 60
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

	m := hsl.L - chroma/2

	return uint32((rFloat + m) * 0xffff), uint32((gFloat + m) * 0xffff), uint32((bFloat + m) * 0xffff), 0xffff
}

var HSLModel color.Model = color.ModelFunc(hslModel)

func hslModel(c color.Color) color.Color {
	if _, ok := c.(HSL); ok {
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

	l := (max + min) / 2

	var s float64 = 0
	if l != 1 && l != 0 {
		s = chroma / (1 - math.Abs(2*l-1))
	}

	return HSL{h, s, l}
}
