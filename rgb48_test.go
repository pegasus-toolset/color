package color

import (
	"image/color"
	"testing"
)

func TestRGB48ToRGBAAliceBlue(t *testing.T) {
	rgb48 := RGB48{0xf0f0, 0xf8f8, 0xffff}
	r, g, b, a := rgb48.RGBA()

	if r != 0xf0f0 {
		t.Errorf("Red component was incorrect, got: 0x%x, want 0x%x.", r, 0xf0f0)
	}

	if g != 0xf8f8 {
		t.Errorf("Green component was incorrect, got: 0x%x, want 0x%x.", g, 0xf8f8)
	}

	if b != 0xffff {
		t.Errorf("Blue component was incorrect, got: 0x%x, want 0x%x.", b, 0xffff)
	}

	if a != 0xffff {
		t.Errorf("Alpha component was incorrect, got: 0x%x, want 0x%x.", a, 0xffff)
	}
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
