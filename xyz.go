package color

import (
	"image/color"
	"math"
)

// XYZ color space (D65)
type XYZ struct {
	X, Y, Z float64
}

// RGBA returns the red, green, blue and alpha values for the color. Each value
// ranges within [0, 0xffff], but is represented by a uint32 so that multiplying
// by a blend factor up to 0xffff will not overflow.
func (xyz XYZ) RGBA() (r, g, b, a uint32) {
	rFloat := xyz.X*3.2406 + xyz.Y*-1.5372 + xyz.Z*-0.4986
	gFloat := xyz.X*-0.9689 + xyz.Y*1.8758 + xyz.Z*0.0415
	bFloat := xyz.X*0.0557 + xyz.Y*-0.2040 + xyz.Z*1.0570

	if rFloat > 0.0031308 {
		rFloat = 1.055*math.Pow(rFloat, 1/2.4) - 0.055
	} else {
		rFloat *= 12.92
	}

	if gFloat > 0.0031308 {
		gFloat = 1.055*math.Pow(gFloat, 1/2.4) - 0.055
	} else {
		gFloat *= 12.92
	}

	if bFloat > 0.0031308 {
		bFloat = 1.055*math.Pow(bFloat, 1/2.4) - 0.055
	} else {
		bFloat *= 12.92
	}

	if rFloat < 0 {
		rFloat = 0
	}
	if gFloat < 0 {
		gFloat = 0
	}
	if bFloat < 0 {
		bFloat = 0
	}

	return uint32(rFloat * 0xffff), uint32(gFloat * 0xffff), uint32(bFloat * 0xffff), 0xffff
}

func (xyz XYZ) DistanceTo(c color.Color) float64 {
	other := XYZModel.Convert(c).(XYZ)
	return math.Sqrt(math.Pow(other.X-xyz.X, 2) + math.Pow(other.Y-xyz.Y, 2) + math.Pow(other.Z-xyz.Z, 2))
}

var XYZModel color.Model = color.ModelFunc(xyzModel)

func xyzModel(c color.Color) color.Color {
	if _, ok := c.(XYZ); ok {
		return c
	}
	rInt, gInt, bInt, _ := c.RGBA()
	r := float64(rInt) / 0xffff
	g := float64(gInt) / 0xffff
	b := float64(bInt) / 0xffff

	if r > 0.04045 {
		r = math.Pow((r+0.055)/1.055, 2.4)
	} else {
		r /= 12.92
	}

	if g > 0.04045 {
		g = math.Pow((g+0.055)/1.055, 2.4)
	} else {
		g /= 12.92
	}

	if b > 0.04045 {
		b = math.Pow((b+0.055)/1.055, 2.4)
	} else {
		b /= 12.92
	}

	x := r*0.4124 + g*0.3576 + b*0.1805
	y := r*0.2126 + g*0.7152 + b*0.0722
	z := r*0.0193 + g*0.1192 + b*0.9505

	return XYZ{x, y, z}
}
