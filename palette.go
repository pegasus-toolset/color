package color

// Palette is a palette of colors.
type Palette []Color

// Contains checks if the palette contains a specific color.
func (palette Palette) Contains(c Color) bool {
	for _, paletteColor := range palette {
		aR, aG, aB, aA := paletteColor.RGBA()
		bR, bG, bB, bA := c.RGBA()
		if aR == bR && aG == bG && aB == bB && aA == bA {
			return true
		}
	}
	return false
}
