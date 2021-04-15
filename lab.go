package color

import (
	"image/color"
	"math"
)

// Lab represents a color in the CIELAB color space (D65).
type Lab struct {
	L, A, B float64
}

func (lab Lab) RGBA() (r, g, b, a uint32) {
	y := (lab.L + 16) / 116
	x := lab.A/500 + y
	z := y - lab.B/200

	if math.Pow(y, 3) > 0.008856 {
		y = math.Pow(y, 3)
	} else {
		y = (y - 16) / 116 / 7.787
	}

	if math.Pow(x, 3) > 0.008856 {
		x = math.Pow(x, 3)
	} else {
		x = (x - 16) / 116 / 7.787
	}

	if math.Pow(z, 3) > 0.008856 {
		z = math.Pow(z, 3)
	} else {
		z = (z - 16) / 116 / 7.787
	}

	xyz := XYZ{x * 0.95047, y, z * 1.08883}
	return xyz.RGBA()
}

func (lab Lab) DistanceTo(c color.Color) float64 {
	other := LabModel.Convert(c).(Lab)
	return math.Sqrt(math.Pow(other.L-lab.L, 2) + math.Pow(other.A-lab.A, 2) + math.Pow(other.B-lab.B, 2))
}

var LabModel color.Model = color.ModelFunc(labModel)

func labModel(c color.Color) color.Color {
	if _, ok := c.(Lab); ok {
		return c
	}

	xyz := XYZModel.Convert(c).(XYZ)

	x := xyz.X / 0.95047
	y := xyz.Y
	z := xyz.Z / 1.08883

	if x > 0.008856 {
		x = math.Pow(x, 1.0/3)
	} else {
		x = (7.787*x + 16) / 116
	}

	if y > 0.008856 {
		y = math.Pow(y, 1.0/3)
	} else {
		y = (7.787*y + 16) / 116
	}

	if z > 0.008856 {
		z = math.Pow(z, 1.0/3)
	} else {
		z = (7.787*z + 16) / 116
	}

	l := 116*y - 16
	a := 500 * (x - y)
	b := 200 * (y - z)

	return Lab{l, a, b}
}
