package color

import (
	"image/color"
	"testing"
)

func TestRGBToRGBAAliceBlue(t *testing.T) {
	rgb := RGB{0xf0, 0xf8, 0xff}
	r, g, b, a := rgb.RGBA()

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

func TestRGBAToRGBAliceBlue(t *testing.T) {
	rgba := color.RGBA{0xf0, 0xf8, 0xff, 0xff}
	rgb := RGBModel.Convert(rgba).(RGB)

	if rgb.R != 0xf0 {
		t.Errorf("Red component was incorrect, got 0x%x, want 0x%x.", rgb.R, 0xf0)
	}

	if rgb.G != 0xf8 {
		t.Errorf("Green component was incorrect, got: 0x%x, want 0x%x.", rgb.G, 0xf8)
	}

	if rgb.B != 0xff {
		t.Errorf("Blue component was incorrect, got 0x%x, want 0x%x.", rgb.B, 0xff)
	}
}
