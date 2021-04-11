package color

import (
	"image/color"
	"testing"
)

func TestRGB48ToRGBAAliceBlue(t *testing.T) {
	rgb48 := RGB48{0xf0f0, 0xf8f8, 0xffff}
	testRGBA(t, rgb48, 0xf0f0, 0xf8f8, 0xffff, 0xffff, 0)
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
