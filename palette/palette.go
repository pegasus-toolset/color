// Package palette contains various pre-defined color palettes.
package palette

import (
	"image"

	"github.com/pegasus-toolset/color"
)

// FromImage returns the palette used by an image.
func FromImage(image image.Image) color.Palette {
	var palette color.Palette

	bounds := image.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := color.RGBModel.Convert(image.At(x, y)).(color.RGB)
			if !palette.Contains(c) {
				palette = append(palette, c)
			}
		}
	}

	return palette
}
