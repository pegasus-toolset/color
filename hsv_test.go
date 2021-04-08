package color

import (
	"image/color"
	"testing"
)

func TestHSVToRGBAWhite(t *testing.T) {
	hsv := HSV{0, 0, 1}
	testHSVToRGBA(t, hsv, 0xffff, 0xffff, 0xffff, 0xffff, 0)
}

func TestHSVToRGBAWebGray(t *testing.T) {
	hsv := HSV{0, 0, 0.5}
	testHSVToRGBA(t, hsv, 0x8080, 0x8080, 0x8080, 0xffff, 0x120)
}

func TestHSVToRGBABlack(t *testing.T) {
	hsv := HSV{0, 0, 0}
	testHSVToRGBA(t, hsv, 0x0, 0x0, 0x0, 0xffff, 0)
}

func TestHSVToRGBARed(t *testing.T) {
	hsv := HSV{0, 1, 1}
	testHSVToRGBA(t, hsv, 0xffff, 0x0, 0x0, 0xffff, 0)
}

func TestHSVToWebGreen(t *testing.T) {
	hsv := HSV{120, 1, 0.5}
	testHSVToRGBA(t, hsv, 0x0, 0x8080, 0x0, 0xffff, 0x120)
}

func TestHSVToRGBAAliceBlue(t *testing.T) {
	hsv := HSV{208, 0.06, 1}
	testHSVToRGBA(t, hsv, 0xf0f0, 0xf8f8, 0xffff, 0xffff, 0x120)
}

func TestHSVToRGBABlueViolet(t *testing.T) {
	hsv := HSV{271, 0.81, 0.89}
	testHSVToRGBA(t, hsv, 0x8a8a, 0x2b2b, 0xe2e2, 0xffff, 0x120)
}

func TestHSVToRGBACrimson(t *testing.T) {
	hsv := HSV{348, 0.91, 0.86}
	testHSVToRGBA(t, hsv, 0xdcdc, 0x1414, 0x3c3c, 0xffff, 0x120)
}

func TestHSVToRGBABeige(t *testing.T) {
	hsv := HSV{60, 0.1, 0.96}
	testHSVToRGBA(t, hsv, 0xf5f5, 0xf5f5, 0xdcdc, 0xffff, 0x120)
}

func TestRGBAToHSVWhite(t *testing.T) {
	rgba := color.RGBA{0xff, 0xff, 0xff, 0xff}
	hsv := HSVModel.Convert(rgba)
	testRGBAToHSV(t, hsv.(HSV), 0, 0, 1, 0)
}

func TestRGBAToHSVWebGray(t *testing.T) {
	rgba := color.RGBA{0x80, 0x80, 0x80, 0xff}
	hsv := HSVModel.Convert(rgba)
	testRGBAToHSV(t, hsv.(HSV), 0, 0, 0.5, 0.002)
}

func TestRGBAToHSVBlack(t *testing.T) {
	rgba := color.RGBA{0x0, 0x0, 0x0, 0xff}
	hsv := HSVModel.Convert(rgba)
	testRGBAToHSV(t, hsv.(HSV), 0, 0, 0, 0)
}

func TestRGBAToHSVRed(t *testing.T) {
	rgba := color.RGBA{0xff, 0x0, 0x0, 0xff}
	hsv := HSVModel.Convert(rgba)
	testRGBAToHSV(t, hsv.(HSV), 0, 1, 1, 0)
}

func TestRGBAToHSVWebGreen(t *testing.T) {
	rgba := color.RGBA{0x0, 0x80, 0x0, 0xff}
	hsv := HSVModel.Convert(rgba)
	testRGBAToHSV(t, hsv.(HSV), 120, 1, 0.5, 0.002)
}

func TestRGBAToHSVAliceBlue(t *testing.T) {
	rgba := color.RGBA{0xf0, 0xf8, 0xff, 0xff}
	hsv := HSVModel.Convert(rgba)
	testRGBAToHSV(t, hsv.(HSV), 208, 0.06, 1, 0.002)
}

func testHSVToRGBA(t *testing.T, hsv HSV, rExp, gExp, bExp, aExp, tolerance int64) {
	r, g, b, a := hsv.RGBA()

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

func testRGBAToHSV(t *testing.T, hsv HSV, hExp, sExp, vExp, tolerance float64) {
	if hsv.H < hExp-tolerance || hsv.H > hExp+tolerance {
		t.Errorf("Hue component was incorrect; got: %f, want %f±%f.", hsv.H, hExp, tolerance)
	}

	if hsv.S < sExp-tolerance || hsv.S > sExp+tolerance {
		t.Errorf("Saturation component was incorrect; got: %f, want %f±%f.", hsv.S, sExp, tolerance)
	}

	if hsv.V < vExp-tolerance || hsv.V > vExp+tolerance {
		t.Errorf("Value component was incorrect; got: %f, want %f±%f.", hsv.V, vExp, tolerance)
	}
}
