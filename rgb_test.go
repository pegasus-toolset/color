package color

import "testing"

func TestRGB2RGBAAliceBlue(t *testing.T) {
	rgb := RGB{0xf0, 0xf8, 0xff}
	r, g, b, a := rgb.RGBA()

	if r != 0xf000 {
		t.Errorf("Red component was incorrect, got: 0x%x, want 0x%x.", r, 0xf000)
	}

	if g != 0xf800 {
		t.Errorf("Green component was incorrect, got: 0x%x, want 0x%x.", g, 0xf800)
	}

	if b != 0xff00 {
		t.Errorf("Blue component was incorrect, got: 0x%x, want 0x%x.", b, 0xff00)
	}

	if a != 0xffff {
		t.Errorf("Alpha component was incorrect, got: 0x%x, want 0x%x.", a, 0xffff)
	}
}
