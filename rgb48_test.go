package color

import (
	"image/color"
	"testing"
)

func TestRGB48ToRGBAAliceBlue(t *testing.T) {
	rgb48 := RGB48{0xf0f0, 0xf8f8, 0xffff}
	testRGBA(t, rgb48, 0xf0f0, 0xf8f8, 0xffff, 0xffff, 0)
}

func TestRGB48DistanceFromBlackToWhite(t *testing.T) {
	black := RGB48{0x0, 0x0, 0x0}
	white := RGB48{0xffff, 0xffff, 0xffff}

	testDistanceTo(t, black, white, 113509.9497, 0.0001)
}

func TestRGB48DistanceFromRedToWebGreen(t *testing.T) {
	red := RGB48{0xffff, 0x0, 0x0}
	webGreen := RGB48{0x0, 0x8080, 0x0}

	testDistanceTo(t, red, webGreen, 73327.9145, 0.0001)
}

func TestRGBAToRGB48AliceBlue(t *testing.T) {
	rgba := color.RGBA{0xf0, 0xf8, 0xff, 0xff}
	rgb48 := RGB48Model.Convert(rgba).(RGB48)

	if rgb48.R != 0xf0f0 {
		t.Errorf("Red component was incorrect, got 0x%x, want 0x%x.", rgb48.R, 0xf0f0)
	}

	if rgb48.G != 0xf8f8 {
		t.Errorf("Green component was incorrect, got: 0x%x, want 0x%x.", rgb48.G, 0xf8f8)
	}

	if rgb48.B != 0xffff {
		t.Errorf("Blue component was incorrect, got 0x%x, want 0x%x.", rgb48.B, 0xffff)
	}
}
