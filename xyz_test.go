package color

import (
	"image/color"
	"testing"
)

func TestXYZToRGBAWhite(t *testing.T) {
	xyz := XYZ{0.9505, 1, 1.089}
	testXYZToRGBA(t, xyz, 0xffff, 0xffff, 0xffff, 0xffff, 0x120)
}

func TestXYZToRGBAWebGray(t *testing.T) {
	xyz := XYZ{0.2052, 0.2159, 0.2351}
	testXYZToRGBA(t, xyz, 0x8080, 0x8080, 0x8080, 0xffff, 0x120)
}

func TestXYZToRGBABlack(t *testing.T) {
	xyz := XYZ{0, 0, 0}
	testXYZToRGBA(t, xyz, 0x0, 0x0, 0x0, 0xffff, 0)
}

func TestXYZToRGBARed(t *testing.T) {
	xyz := XYZ{0.4124, 0.2126, 0.0193}
	testXYZToRGBA(t, xyz, 0xffff, 0x0, 0x0, 0xffff, 0x120)
}

func TestXYZToRGBAWebGreen(t *testing.T) {
	xyz := XYZ{0.0772, 0.1544, 0.0257}
	testXYZToRGBA(t, xyz, 0x0, 0x8080, 0x0, 0xffff, 0x120)
}

func TestXYZToRGBAAliceBlue(t *testing.T) {
	xyz := XYZ{0.8755, 0.9288, 1.0792}
	testXYZToRGBA(t, xyz, 0xf0f0, 0xf8f8, 0xffff, 0xffff, 0x120)
}

func TestRGBAToXYZWhite(t *testing.T) {
	rgba := color.RGBA{0xff, 0xff, 0xff, 0xff}
	xyz := XYZModel.Convert(rgba)
	testRGBAToXYZ(t, xyz.(XYZ), 0.9505, 1, 1.089, 0)
}

func TestRGBAToXYZWebGray(t *testing.T) {
	rgba := color.RGBA{0x80, 0x80, 0x80, 0xff}
	xyz := XYZModel.Convert(rgba)
	testRGBAToXYZ(t, xyz.(XYZ), 0.2052, 0.2159, 0.2351, 0.002)
}

func TestRGBAToXYZBlack(t *testing.T) {
	rgba := color.RGBA{0x0, 0x0, 0x0, 0xff}
	xyz := XYZModel.Convert(rgba)
	testRGBAToXYZ(t, xyz.(XYZ), 0, 0, 0, 0)
}

func TestRGBAToXYZRed(t *testing.T) {
	rgba := color.RGBA{0xff, 0x0, 0x0, 0xff}
	xyz := XYZModel.Convert(rgba)
	testRGBAToXYZ(t, xyz.(XYZ), 0.4124, 0.2126, 0.0193, 0.002)
}

func TestRGBAToXYZWebGreen(t *testing.T) {
	rgba := color.RGBA{0x0, 0x80, 0x0, 0xff}
	xyz := XYZModel.Convert(rgba)
	testRGBAToXYZ(t, xyz.(XYZ), 0.0772, 0.1544, 0.0257, 0.002)
}

func TestRGBAToAliceBlue(t *testing.T) {
	rgba := color.RGBA{0xf0, 0xf8, 0xff, 0xff}
	xyz := XYZModel.Convert(rgba)
	testRGBAToXYZ(t, xyz.(XYZ), 0.8755, 0.9288, 1.0792, 0.002)
}

func testXYZToRGBA(t *testing.T, xyz XYZ, rExp, gExp, bExp, aExp, tolerance int64) {
	r, g, b, a := xyz.RGBA()

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
		t.Errorf("Alpha component was incorrect, got 0x%x, want 0x%x.", a, aExp)
	}
}

func testRGBAToXYZ(t *testing.T, xyz XYZ, xExp, yExp, zExp, tolerance float64) {
	if xyz.X < xExp-tolerance || xyz.X > xExp+tolerance {
		t.Errorf("X component was incorrect; got: %f, want %f±%f.", xyz.X, xExp, tolerance)
	}

	if xyz.Y < yExp-tolerance || xyz.Y > yExp+tolerance {
		t.Errorf("Y component was incorrect; got: %f, want %f±%f.", xyz.Y, yExp, tolerance)
	}

	if xyz.Z < zExp-tolerance || xyz.Z > zExp+tolerance {
		t.Errorf("Z component was incorrect; got: %f, want %f±%f.", xyz.Z, zExp, tolerance)
	}
}
