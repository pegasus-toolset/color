package color

import (
	"image/color"
	"testing"
)

func TestRGBToRGBAAliceBlue(t *testing.T) {
	rgb := RGB{0xf0, 0xf8, 0xff}
	testRGBA(t, rgb, 0xf0f0, 0xf8f8, 0xffff, 0xffff, 0)
}

func TestRGBToRGBARed(t *testing.T) {
	rgb := RGB{0xff, 0x0, 0x0}
	testRGBA(t, rgb, 0xffff, 0x0, 0x0, 0xffff, 0)
}

func TestRGBToRGBAWebGreen(t *testing.T) {
	rgb := RGB{0x0, 0x80, 0x0}
	testRGBA(t, rgb, 0x0, 0x8080, 0x0, 0xffff, 0)
}

func TestRGBDistanceFromBlackToWhite(t *testing.T) {
	black := RGB{0, 0, 0}
	white := RGB{255, 255, 255}

	testDistanceTo(t, black, white, 441.673, 0.0001)
}

func TestRGBDistanceFromRedToWebGreen(t *testing.T) {
	red := RGB{255, 0, 0}
	webGreen := RGB{0, 128, 0}

	testDistanceTo(t, red, webGreen, 285.3226, 0.0001)
}

func TestRGBAToRGBAliceBlue(t *testing.T) {
	rgba := color.RGBA{0xf0, 0xf8, 0xff, 0xff}
	rgb := RGBModel.Convert(rgba)
	testRGBAToRGB(t, rgb.(RGB), 0xf0, 0xf8, 0xff)
}

func TestRGBAToRGBRed(t *testing.T) {
	rgba := color.RGBA{0xff, 0x0, 0x0, 0xff}
	rgb := RGBModel.Convert(rgba)
	testRGBAToRGB(t, rgb.(RGB), 0xff, 0x0, 0x0)
}

func TestRGBAToRGBWebGreen(t *testing.T) {
	rgba := color.RGBA{0x0, 0x80, 0x0, 0xff}
	rgb := RGBModel.Convert(rgba)
	testRGBAToRGB(t, rgb.(RGB), 0x0, 0x80, 0x0)
}

func testRGBAToRGB(t *testing.T, rgb RGB, rExp, gExp, bExp uint8) {
	if rgb.R != rExp {
		t.Errorf("Red component was incorrect, got 0x%x, want 0x%x.", rgb.R, rExp)
	}

	if rgb.G != gExp {
		t.Errorf("Green component was incorrect, got: 0x%x, want 0x%x.", rgb.G, gExp)
	}

	if rgb.B != bExp {
		t.Errorf("Blue component was incorrect, got 0x%x, want 0x%x.", rgb.B, bExp)
	}
}
