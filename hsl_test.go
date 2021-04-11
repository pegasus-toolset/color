package color

import (
	"image/color"
	"testing"
)

func TestHSLToRGBAWhite(t *testing.T) {
	hsl := HSL{0, 0, 1}
	testRGBA(t, hsl, 0xffff, 0xffff, 0xffff, 0xffff, 0)
}

func TestHSLToRGBAWebGray(t *testing.T) {
	hsl := HSL{0, 0, 0.5}
	testRGBA(t, hsl, 0x8080, 0x8080, 0x8080, 0xffff, 0x120)
}

func TestHSLToRGBABlack(t *testing.T) {
	hsl := HSL{0, 0, 0}
	testRGBA(t, hsl, 0x0, 0x0, 0x0, 0xffff, 0)
}

func TestHSLToRGBARed(t *testing.T) {
	hsl := HSL{0, 1, 0.5}
	testRGBA(t, hsl, 0xffff, 0x0, 0x0, 0xffff, 0)
}

func TestHSLToRGBAWebGreen(t *testing.T) {
	hsl := HSL{120, 1, 0.25}
	testRGBA(t, hsl, 0x0, 0x8080, 0x0, 0xffff, 0x120)
}

func TestHSLToRGBAAliceBlue(t *testing.T) {
	hsl := HSL{208, 1, 0.97}
	testRGBA(t, hsl, 0xf0f0, 0xf8f8, 0xffff, 0xffff, 0x120)
}

func TestHSLToRGBABlueViolet(t *testing.T) {
	hsl := HSL{271, 0.76, 0.53}
	testRGBA(t, hsl, 0x8a8a, 0x2b2b, 0xe2e2, 0xffff, 0x120)
}

func TestHSLToRGBACrimson(t *testing.T) {
	hsl := HSL{348, 0.83, 0.47}
	testRGBA(t, hsl, 0xdcdc, 0x1414, 0x3c3c, 0xffff, 0x120)
}

func TestHSLToRGBABeige(t *testing.T) {
	hsl := HSL{60, 0.56, 0.91}
	testRGBA(t, hsl, 0xf5f5, 0xf5f5, 0xdcdc, 0xffff, 0x120)
}

func TestHSLDistanceFromBlackToWhite(t *testing.T) {
	black := HSL{0, 0, 0}
	white := HSL{0, 0, 1}

	testDistanceTo(t, black, white, 1, 0)
}

func TestHSLDistanceFromRedToWebGreen(t *testing.T) {
	red := HSL{0, 1, 0.5}
	webGreen := HSL{120, 1, 0.25}

	testDistanceTo(t, red, webGreen, 0.4167, 0.0001)
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

func testRGBAToHSL(t *testing.T, hsl HSL, hExp, sExp, lExp, tolerance float64) {
	if hsl.H < hExp-tolerance*100 || hsl.H > hExp+tolerance*100 {
		t.Errorf("Hue component was incorrect, got: %f, want %f±%f.", hsl.H, hExp, tolerance*100)
	}

	if hsl.S < sExp-tolerance || hsl.S > sExp+tolerance {
		t.Errorf("Saturation component was incorrect, got: %f, want %f±%f.", hsl.S, sExp, tolerance)
	}

	if hsl.L < lExp-tolerance || hsl.L > lExp+tolerance {
		t.Errorf("Lightness component was incorrect, got: %f, want %f±%f.", hsl.L, lExp, tolerance)
	}
}
