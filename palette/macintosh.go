package palette

import "github.com/pegasus-toolset/color"

// Macintosh represents the Apple Macintosh default 16-color palette.
var Macintosh = color.Palette{
	// White
	color.RGB{R: 0xff, G: 0xff, B: 0xff},
	// Yellow
	color.RGB{R: 0xfb, G: 0xf3, B: 0x05},
	// Orange
	color.RGB{R: 0xff, G: 0x64, B: 0x03},
	// Red
	color.RGB{R: 0xdd, G: 0x09, B: 0x07},
	// Magenta
	color.RGB{R: 0xf2, G: 0x08, B: 0x84},
	// Purple
	color.RGB{R: 0x47, G: 0x00, B: 0xa5},
	// Blue
	color.RGB{R: 0x00, G: 0x00, B: 0xd3},
	// Cyan
	color.RGB{R: 0x02, G: 0xab, B: 0xea},
	// Green
	color.RGB{R: 0x1d, G: 0xb7, B: 0x14},
	// Dark Green
	color.RGB{R: 0x00, G: 0x64, B: 0x12},
	// Brown
	color.RGB{R: 0x56, G: 0x2c, B: 0x05},
	// Tan
	color.RGB{R: 0x90, G: 0x71, B: 0x3a},
	// Light Grey
	color.RGB{R: 0xc0, G: 0xc0, B: 0xc0},
	// Medium Grey
	color.RGB{R: 0x80, G: 0x80, B: 0x80},
	// Dark Grey
	color.RGB{R: 0x40, G: 0x40, B: 0x40},
	// Black
	color.RGB{R: 0x0, G: 0x0, B: 0x0},
}
