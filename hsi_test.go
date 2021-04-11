package color

import (
	"image/color"
	"testing"
)

func TestHSIToRGBAWhite(t *testing.T) {
	hsi := HSI{0, 0, 1}
	testRGBA(t, hsi, 0xffff, 0xffff, 0xffff, 0xffff, 0)
}

func TestHSIToRGBAWebGray(t *testing.T) {
	hsi := HSI{0, 0, 0.502}
	testRGBA(t, hsi, 0x8080, 0x8080, 0x8080, 0xffff, 0x120)
}

func TestHSIToRGBABlack(t *testing.T) {
	hsi := HSI{0, 0, 0}
	testRGBA(t, hsi, 0x0, 0x0, 0x0, 0xffff, 0)
}

func TestHSIToRGBARed(t *testing.T) {
	hsi := HSI{0, 1, 0.3333}
	testRGBA(t, hsi, 0xffff, 0x0, 0x0, 0xffff, 0x120)
}

func TestHISToRGBAWebGreen(t *testing.T) {
	hsi := HSI{120, 1, 0.1673}
	testRGBA(t, hsi, 0x0, 0x8080, 0x0, 0xffff, 0x120)
}

func TestHSIToRGBAAliceBlue(t *testing.T) {
	hsi := HSI{208, 0.031, 0.9712}
	testRGBA(t, hsi, 0xf0f0, 0xf8f8, 0xffff, 0xffff, 0x120)
}

func TestHSIToRGBABlueViolet(t *testing.T) {
	hsi := HSI{271.15, 0.6830, 0.532}
	testRGBA(t, hsi, 0x8a8a, 0x2b2b, 0xe2e2, 0xffff, 0x120)
}

func TestHSIToRGBACrimson(t *testing.T) {
	hsi := HSI{348, 0.8, 0.3922}
	testRGBA(t, hsi, 0xdcdc, 0x1414, 0x3c3c, 0xffff, 0x120)
}

func TestHSIToRGBABeige(t *testing.T) {
	hsi := HSI{60, 0.0704, 0.9281}
	testRGBA(t, hsi, 0xf5f5, 0xf5f5, 0xdcdc, 0xffff, 0x120)
}

func TestHSIDistanceFromBlackToWhite(t *testing.T) {
	black := HSI{0, 0, 0}
	white := HSI{0, 0, 1}

	testDistanceTo(t, black, white, 1, 0)
}

func TestHSIDistanceFromRedToWebGreen(t *testing.T) {
	red := HSI{0, 1, 0.3333}
	webGreen := HSI{120, 1, 0.1673}

	testDistanceTo(t, red, webGreen, 0.3724, 0.0001)
}

func TestRGBAToHSIWhite(t *testing.T) {
	rgba := color.RGBA{0xff, 0xff, 0xff, 0xff}
	hsi := HSIModel.Convert(rgba)
	testRGBAToHSI(t, hsi.(HSI), 0, 0, 1, 0)
}

func TestRGBAToHSIWebGray(t *testing.T) {
	rgba := color.RGBA{0x80, 0x80, 0x80, 0xff}
	hsi := HSIModel.Convert(rgba)
	testRGBAToHSI(t, hsi.(HSI), 0, 0, 0.502, 0.002)
}

func TestRGBAToHSIBlack(t *testing.T) {
	rgba := color.RGBA{0x0, 0x0, 0x0, 0xff}
	hsi := HSIModel.Convert(rgba)
	testRGBAToHSI(t, hsi.(HSI), 0, 0, 0, 0)
}

func TestRGBAToHSIRed(t *testing.T) {
	rgba := color.RGBA{0xff, 0x0, 0x0, 0xff}
	hsi := HSIModel.Convert(rgba)
	testRGBAToHSI(t, hsi.(HSI), 0, 1, 0.3333, 0.002)
}

func TestRGBAToHSIWebGreen(t *testing.T) {
	rgba := color.RGBA{0x0, 0x80, 0x0, 0xff}
	hsi := HSIModel.Convert(rgba)
	testRGBAToHSI(t, hsi.(HSI), 120, 1, 0.1673, 0.002)
}

func TestRGBAToHSIAliceBlue(t *testing.T) {
	rgba := color.RGBA{0xf0, 0xf8, 0xff, 0xff}
	hsi := HSIModel.Convert(rgba)
	testRGBAToHSI(t, hsi.(HSI), 208, 0.031, 0.9712, 0.002)
}

func testRGBAToHSI(t *testing.T, hsi HSI, hExp, sExp, iExp, tolerance float64) {
	if hsi.H < hExp-tolerance*100 || hsi.H > hExp+tolerance*100 {
		t.Errorf("Hue component was incorrect; got: %f, want %f±%f.", hsi.H, hExp, tolerance*100)
	}

	if hsi.S < sExp-tolerance || hsi.S > sExp+tolerance {
		t.Errorf("Saturation component was incorrect, got: %f, want %f±%f.", hsi.S, sExp, tolerance)
	}

	if hsi.I < iExp-tolerance || hsi.I > iExp+tolerance {
		t.Errorf("Intensity component was incorrect; got: %f, want %f±%f.", hsi.I, iExp, tolerance)
	}
}
