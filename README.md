<h1 align="center">🎨 Color</h1>

<p align="center">
    <a href="https://travis-ci.com/pegasus-toolset/color"><img src="https://travis-ci.com/pegasus-toolset/color.svg?branch=main" /></a>
    <a href="https://app.codecov.io/gh/pegasus-toolset/color"><img src="https://codecov.io/gh/pegasus-toolset/color/branch/main/graph/badge.svg?token=3R6863F2HA" /></a>
</p>

**Color** is an extension of the standard library package
[image/color](https://golang.org/pkg/image/color/).

It provides multiple color spaces, conversions between each of them, color
distance measuring and multiple palettes.

## Features

## Color Models/Spaces

- 24-bit RGB colors
- 48-bit RGB colors
- HSI colors
- HSL colors
- HSV colors
- L\*a\*b\* colors
- XYZ colors

### RGB

A `RGB` color is made up of three components (each an unsigned 8-bit integer):

- `R`: Red channel
- `G`: Green channel
- `B`: Blue channel

### RGB48

A `RGB48` color is made up of three components (each an unsigned 16-bit
integer):

- `R`: Red channel
- `G`: Green channel
- `B`: Blue channel

### HSL

A `HSL` color is made up of three components (each a 64-bit floating-point
number):

- `H`: Hue
- `S`: Saturation
- `L`: Lightness

> In the HSL "bi-hexcone" model, _lightness_ is defined as the average of the
> largest and smallest color components, i.e. the mid-range of the RGB
> components. This definition also puts the primary and secondary colors into a
> plane, but a plane passing halfway between white and black. The resulting
> color solid is a double-cone similar to Ostwald's, shown above.
>
> \- [Wikipedia](https://en.wikipedia.org/wiki/HSL_and_HSV#Lightness)

### HSV

A `HSV` color is made up of three components (each a 64-bit floating-point
number):

- `H`: Hue
- `S`: Saturation
- `V`: Value

> In the HSV "hexcone" model, _value_ is defined as the largest component of a
> color, our _M_ above. This places all three primaries, and also all of the
> "secondary colors"—cyan, yellow, and magenta—into a plane with white, forming
> a hexagonal pyramid out of the RGB cube.
>
> \- [Wikipedia](https://en.wikipedia.org/wiki/HSL_and_HSV#Lightness)

### HSI

A `HSI` color is made up of three components (each a 64-bit floating-point
number):

- `H`: Hue
- `S`: Saturation
- `I`: Intensity

#### Intensity

> The simplest definition is just the arithmetic mean, i.e. average, of the
> three components, in the HSI model called _intensity_. This is simply the
> projection of a point onto the neutral axis—the vertical height of a point in
> our tilted cube. The advantage is that, together with Euclidean-distance
> calculations of hue and chroma, this representation preserves distances and
> angles from the geometry of the RGB cube.
>
> \- [Wikipedia](https://en.wikipedia.org/wiki/HSL_and_HSV#Lightness)

### XYZ

`XYZ` represents a color in the CIE 1931 XYZ color space and is made up of three
components (each a 64-bit floating-point number):

- `X`: A mix of the three CIE RGB curves
- `Y`: Luminance
- `Z`: Quasi-equal to blue (of CIE RGB)

### L\*a\*b\*

`Lab` represents a color in the CIELAB color space and is made up of three
components (each a 64-bit floating-point number):

- `L`: Lightness
- `A`: Position between red and green
- `B`: Position between yellow and blue

## Palettes

- X11 (as named color constants)
- Microsoft Windows default 16-color palette (`palette.Windows16`)
- Microsoft Windows default 20-color palette (`palette.Windows20`)
- Apple Macintosh default 16-color palette (`palette.Macintosh`)
- RISC OS default palette (`palette.RISC`)

## Additional Features

- Conversion between color models/spaces
- Distance calculations between colors in all color models/spaces
