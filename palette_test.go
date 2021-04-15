package color

import "testing"

func TestContainsShouldContain(t *testing.T) {
	palette := Palette{
		RGB{0xf0, 0xf8, 0xff},
		RGB{0xfa, 0xeb, 0xd7},
		RGB{0x00, 0xff, 0xff},
	}

	if !palette.Contains(RGB{0xfa, 0xeb, 0xd7}) {
		t.Errorf("Palette should contain color but doesn't.")
	}
}

func TestContainsShouldNotContain(t *testing.T) {
	palette := Palette{
		RGB{0xf0, 0xf8, 0xff},
		RGB{0xfa, 0xeb, 0xd7},
		RGB{0x00, 0xff, 0xff},
	}

	if palette.Contains(RGB{0x7f, 0xff, 0xd4}) {
		t.Errorf("Palette should not contain color but does.")
	}
}
