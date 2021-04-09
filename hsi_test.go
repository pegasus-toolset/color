package color

import (
	"image/color"
	"testing"
)

func TestHSIToRGBAWhite(t *testing.T) {
	hsi := HSI{0, 0, 1}
	testHSIToRGBA(t, hsi, 0xffff, 0xffff, 0xffff, 0xffff, 0)
}

func TestHSIToRGBAWebGray(t *testing.T) {
	hsi := HSI{0, 0, 0.502}
	testHSIToRGBA(t, hsi, 0x8080, 0x8080, 0x8080, 0xffff, 0x120)
}

func TestHSIToRGBABlack(t *testing.T) {
	hsi := HSI{0, 0, 0}
	testHSIToRGBA(t, hsi, 0x0, 0x0, 0x0, 0xffff, 0)
}

func TestHSIToRGBARed(t *testing.T) {
	hsi := HSI{0, 1, 0.3333}
	testHSIToRGBA(t, hsi, 0xffff, 0x0, 0x0, 0xffff, 0x120)
}

func TestHISToRGBAWebGreen(t *testing.T) {
	hsi := HSI{120, 1, 0.1673}
	testHSIToRGBA(t, hsi, 0x0, 0x8080, 0x0, 0xffff, 0x120)
}

func TestHSIToRGBAAliceBlue(t *testing.T) {
	hsi := HSI{207.8, 0.031, 0.9712}
	testHSIToRGBA(t, hsi, 0xf0f0, 0xf8f8, 0xffff, 0xffff, 0x120)
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
	testRGBAToHSI(t, hsi.(HSI), 207.8, 0.031, 0.9712, 0.002)
}

func testHSIToRGBA(t *testing.T, hsi HSI, rExp, gExp, bExp, aExp, tolerance int64) {
	r, g, b, a := hsi.RGBA()

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
