package color

import (
	"image/color"
	"math"
)

// HSI color space
type HSI struct {
	H, S, I float64
}

// RGBA returns the red, green, blue and alpha values for the color. Each value
// ranges within [0, 0xffff], but is represented by a uint32 so that multiplying
// by a blend factor up to 0xffff will not overflow.
func (hsi HSI) RGBA() (r, g, b, a uint32) {
	h := hsi.H / 60
	z := 1 - math.Abs(math.Mod(h, 2)-1)
	chroma := 3 * hsi.I * hsi.S / (1 + z)
	x := chroma * z

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

	m := hsi.I * (1 - hsi.S)

	return uint32((rFloat + m) * 0xffff), uint32((gFloat + m) * 0xffff), uint32((bFloat + m) * 0xffff), 0xffff
}

// DistanceTo calculates the distance to another color in the HSI color space.
//
// Before calculating the distance, the other color is converted into a HSI
// color.
func (hsi HSI) DistanceTo(c color.Color) float64 {
	other := HSIModel.Convert(c).(HSI)
	return math.Sqrt(math.Pow(other.H/360-hsi.H/360, 2) + math.Pow(other.S-hsi.S, 2) + math.Pow(other.I-hsi.I, 2))
}

// HSIModel can convert any color to a HSI color.
var HSIModel color.Model = color.ModelFunc(hsiModel)

func hsiModel(c color.Color) color.Color {
	if _, ok := c.(HSI); ok {
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

	i := (r + g + b) / 3

	var s float64 = 0
	if i > 0 {
		s = 1 - (min / i)
	}

	return HSI{h, s, i}
}
