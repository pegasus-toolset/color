package palette

import (
	"image"
	"testing"

	"github.com/pegasus-toolset/color"
)

func TestFromImage0x0(t *testing.T) {
	image := image.NewRGBA(image.Rect(0, 0, 0, 0))

	palette := FromImage(image)

	if len(palette) != 0 {
		t.Errorf("Palette length was incorrect, got: %d, want %d.", len(palette), 0)
	}
}

func TestFromImage2x2(t *testing.T) {
	image := image.NewRGBA(image.Rect(0, 0, 2, 2))
	image.Set(0, 0, color.RGB{R: 1, G: 2, B: 3})
	image.Set(1, 0, color.RGB{R: 2, G: 3, B: 4})
	image.Set(0, 1, color.RGB{R: 3, G: 4, B: 5})
	image.Set(1, 1, color.RGB{R: 1, G: 2, B: 3})

	palette := FromImage(image)

	if len(palette) != 3 {
		t.Errorf("Palette length was incorrect, got: %d, want %d.", len(palette), 3)
	}
	if !palette.Contains(color.RGB{R: 1, G: 2, B: 3}) {
		t.Errorf("Palette should contain color rgb(1, 2, 3) but doesn't.")
	}
	if !palette.Contains(color.RGB{R: 2, G: 3, B: 4}) {
		t.Errorf("Palette should contain color rgb(2, 3, 4) but doesn't.")
	}
	if !palette.Contains(color.RGB{R: 3, G: 4, B: 5}) {
		t.Errorf("Palette should contain color rgb(3, 4, 5) but doesn't.")
	}
}
