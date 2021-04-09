package color

import (
	"image/color"
	"testing"
)

func TestLabToRGBAWhite(t *testing.T) {
	lab := Lab{100, 0, 0}
	testLabToRGBA(t, lab, 0xffff, 0xffff, 0xffff, 0xffff, 0x120)
}

func TestLabToRGBAWebGray(t *testing.T) {
	lab := Lab{53.59, 0, 0}
	testLabToRGBA(t, lab, 0x8080, 0x8080, 0x8080, 0xffff, 0x120)
}

func TestLabToRGBABlack(t *testing.T) {
	lab := Lab{0, 0, 0}
	testLabToRGBA(t, lab, 0x0, 0x0, 0x0, 0xffff, 0)
}

func TestLabToRGBARed(t *testing.T) {
	lab := Lab{53.24, 80.09, 67.2}
	testLabToRGBA(t, lab, 0xffff, 0x0, 0x0, 0xffff, 0x120)
}

func TestLabToRGBAWebGreen(t *testing.T) {
	lab := Lab{46.23, -51.7, 49.9}
	testLabToRGBA(t, lab, 0x0, 0x8080, 0x0, 0xffff, 0x120)
}

func TestLabToRGBAAliceBlue(t *testing.T) {
	lab := Lab{97.18, -1.35, -4.26}
	testLabToRGBA(t, lab, 0xf0f0, 0xf8f8, 0xffff, 0xffff, 0x120)
}

func TestRGBAToLabWhite(t *testing.T) {
	rgba := color.RGBA{0xff, 0xff, 0xff, 0xff}
	lab := LabModel.Convert(rgba)
	testRGBAToLab(t, lab.(Lab), 100, 0, 0, 0.2)
}

func TestRGBAToLabWebGray(t *testing.T) {
	rgba := color.RGBA{0x80, 0x80, 0x80, 0xff}
	lab := LabModel.Convert(rgba)
	testRGBAToLab(t, lab.(Lab), 53.59, 0, 0, 0.2)
}

func TestRGBAToLabBlack(t *testing.T) {
	rgba := color.RGBA{0x0, 0x0, 0x0, 0xff}
	lab := LabModel.Convert(rgba)
	testRGBAToLab(t, lab.(Lab), 0, 0, 0, 0)
}

func TestRGBAToLabRed(t *testing.T) {
	rgba := color.RGBA{0xff, 0x0, 0x0, 0xff}
	lab := LabModel.Convert(rgba)
	testRGBAToLab(t, lab.(Lab), 53.24, 80.09, 67.2, 0.2)
}

func TestRGBAToLabWebGreen(t *testing.T) {
	rgba := color.RGBA{0x0, 0x80, 0x0, 0xff}
	lab := LabModel.Convert(rgba)
	testRGBAToLab(t, lab.(Lab), 46.23, -51.7, 49.9, 0.2)
}

func TestRGBAToLabAliceBlue(t *testing.T) {
	rgba := color.RGBA{0xf0, 0xf8, 0xff, 0xff}
	lab := LabModel.Convert(rgba)
	testRGBAToLab(t, lab.(Lab), 97.18, -1.35, -4.26, 0.2)
}

func testLabToRGBA(t *testing.T, lab Lab, rExp, gExp, bExp, aExp, tolerance int64) {
	r, g, b, a := lab.RGBA()

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

func testRGBAToLab(t *testing.T, lab Lab, lExp, aExp, bExp, tolerance float64) {
	if lab.L < lExp-tolerance || lab.L > lExp+tolerance {
		t.Errorf("Lightness component was incorrect, got: %f, want %f±%f.", lab.L, lExp, tolerance)
	}

	if lab.A < aExp-tolerance || lab.A > aExp+tolerance {
		t.Errorf("a* axis component was incorrect, got: %f, want %f±%f.", lab.A, aExp, tolerance)
	}

	if lab.B < bExp-tolerance || lab.B > bExp+tolerance {
		t.Errorf("b* axis component was incorrect, got: %f, want %f±%f.", lab.B, bExp, tolerance)
	}
}
