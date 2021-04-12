package color

import (
	"image/color"
	"testing"
)

func testDistanceTo(t *testing.T, a Color, b color.Color, expected, tolerance float64) {
	distance := a.DistanceTo(b)
	if distance < expected-tolerance || distance > expected+tolerance {
		t.Errorf("Distance was incorrect, got: %f, want %f±%f.", distance, expected, tolerance)
	}
}

func testRGBA(t *testing.T, c color.Color, rExp, gExp, bExp, aExp, tolerance int64) {
	r, g, b, a := c.RGBA()

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
