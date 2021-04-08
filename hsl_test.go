package color

import (
	"image/color"
	"testing"
)

func TestHSLToRGBAWhite(t *testing.T) {
	hsl := HSL{0, 0, 1}
	testHSLToRGBA(t, hsl, 0xffff, 0xffff, 0xffff, 0xffff, 0)
}

func TestHSLToRGBAWebGray(t *testing.T) {
	hsl := HSL{0, 0, 0.5}
	testHSLToRGBA(t, hsl, 0x8080, 0x8080, 0x8080, 0xffff, 0xff)
}

func TestHSLToRGBABlack(t *testing.T) {
	hsl := HSL{0, 0, 0}
	testHSLToRGBA(t, hsl, 0x0, 0x0, 0x0, 0xffff, 0)
}

func TestHSLToRGBARed(t *testing.T) {
	hsl := HSL{0, 1, 0.5}
	testHSLToRGBA(t, hsl, 0xffff, 0x0, 0x0, 0xffff, 0)
}

func TestHSLToRGBAWebGreen(t *testing.T) {
	hsl := HSL{120, 1, 0.25}
	testHSLToRGBA(t, hsl, 0x0, 0x8080, 0x0, 0xffff, 0xff)
}

func TestHSLToRGBAAliceBlue(t *testing.T) {
	hsl := HSL{208, 1, 0.97}
	testHSLToRGBA(t, hsl, 0xf0f0, 0xf8f8, 0xffff, 0xffff, 0xff)
}

func TestRGBAToHSLWhite(t *testing.T) {
	rgba := color.RGBA{0xff, 0xff, 0xff, 0xff}
	hsl := HSLModel.Convert(rgba)
	testRGBAToHSL(t, hsl.(HSL), 0, 0, 1, 0)
}

func TestRGBAToHSLWebGray(t *testing.T) {
	rgba := color.RGBA{0x80, 0x80, 0x80, 0xff}
	hsl := HSLModel.Convert(rgba)
	testRGBAToHSL(t, hsl.(HSL), 0, 0, 0.5, 0.002)
}

func TestRGBAToHSLBlack(t *testing.T) {
	rgba := color.RGBA{0x0, 0x0, 0x0, 0xff}
	hsl := HSLModel.Convert(rgba)
	testRGBAToHSL(t, hsl.(HSL), 0, 0, 0, 0)
}

func TestRGBAToHSLRed(t *testing.T) {
	rgba := color.RGBA{0xff, 0x0, 0x0, 0xff}
	hsl := HSLModel.Convert(rgba)
	testRGBAToHSL(t, hsl.(HSL), 0, 1, 0.5, 0)
}

func TestRGBAToHSLWebGreen(t *testing.T) {
	rgba := color.RGBA{0x0, 0x80, 0x0, 0xff}
	hsl := HSLModel.Convert(rgba)
	testRGBAToHSL(t, hsl.(HSL), 120, 1, 0.25, 0.002)
}

func TestRGBAToHSLAliceBlue(t *testing.T) {
	rgba := color.RGBA{0xf0, 0xf8, 0xff, 0xff}
	hsl := HSLModel.Convert(rgba)
	testRGBAToHSL(t, hsl.(HSL), 208, 1, 0.97, 0.002)
}

func testHSLToRGBA(t *testing.T, hsl HSL, rExp, gExp, bExp, aExp, tolerance int64) {
	r, g, b, a := hsl.RGBA()

	if int64(r) < rExp-tolerance || int64(r) > rExp+tolerance {
		t.Errorf("Red component was incorrect, got: 0x%x, want 0x%x±%d.", r, rExp, tolerance)
	}

	if int64(g) < gExp-tolerance || int64(g) > gExp+tolerance {
		t.Errorf("Green component was incorrect, got: 0x%x, want 0x%x±%d.", g, gExp, tolerance)
	}

	if int64(b) < bExp-tolerance || int64(b) > bExp+tolerance {
		t.Errorf("Blue component was incorrect, got: 0x%x, want 0x%x±%d.", b, bExp, tolerance)
	}

	if int64(a) != aExp {
		t.Errorf("Alpha component was incorrect, got: 0x%x, want 0x%x.", a, aExp)
	}
}

func testRGBAToHSL(t *testing.T, hsl HSL, hExp, sExp, lExp, tolerance float64) {
	if hsl.H < hExp-tolerance || hsl.H > hExp+tolerance {
		t.Errorf("Hue component was incorrect, got: %f, want %f±%f.", hsl.H, hExp, tolerance)
	}

	if hsl.S < sExp-tolerance || hsl.S > sExp+tolerance {
		t.Errorf("Saturation component was incorrect, got: %f, want %f±%f.", hsl.S, sExp, tolerance)
	}

	if hsl.L < lExp-tolerance || hsl.L > lExp+tolerance {
		t.Errorf("Lightness component was incorrect, got: %f, want %f±%f.", hsl.L, lExp, tolerance)
	}
}
