package hexcolor

import (
	"image/color"
	"testing"
)

func TestHexValidation_ShouldBeValid(t *testing.T) {
	hexColors := []string{
		"#123456",
		"#0BCDEF",
		"#FF00AA00",
		"#12345678",
		"#9ABCDEF0",
	}

	for _, hex := range hexColors {
		if !IsValid(hex) {
			t.Errorf("Color <%s> should be valid !", hex)
		}
	}
}

func TestHexValidation_ShouldNotBeValid(t *testing.T) {
	hexColors := []string{
		"#23456",
		"#0BCDEFZ",
		"#FF0GAA00",
		"#123-5678",
		"#9DEF0",
	}

	for _, hex := range hexColors {
		if IsValid(hex) {
			t.Errorf("Color <%s> should not be valid !", hex)
		}
	}
}

func TestParseToRGBA_InvalidColor(t *testing.T) {
	hex := "#FFFFFFFG"

	c, err := ParseToRGBA(hex)

	if err == nil {
		t.Error("Got no error with an invalid color value :/")
	}

	if c != nil {
		t.Errorf("Got color <%X> with an invalid color value :/", c)
	}
}

func TestParseToRGBA_ValidColor(t *testing.T) {
	values := map[string]color.Color{
		"#000000":   color.RGBA{A: 0xFF, R: 0, G: 0, B: 0},
		"#FF0000":   color.RGBA{A: 0xFF, R: 0xFF, G: 0, B: 0},
		"#00FF00":   color.RGBA{A: 0xFF, R: 0, G: 0xFF, B: 0},
		"#0000FF":   color.RGBA{A: 0xFF, R: 0, G: 0, B: 0xFF},
		"#00000000": color.RGBA{A: 0x00, R: 0, G: 0, B: 0},
		"#00FF0000": color.RGBA{A: 0x00, R: 0xFF, G: 0, B: 0},
		"#0000FF00": color.RGBA{A: 0x00, R: 0, G: 0xFF, B: 0},
		"#000000FF": color.RGBA{A: 0x00, R: 0, G: 0, B: 0xFF},
	}

	for hex, expected := range values {
		c, err := ParseToRGBA(hex)

		if err != nil {
			t.Errorf("<%s>: got error <%s>", hex, err)
		}

		if !same(expected, c) {
			t.Errorf("<%s>: expected <%v>, got <%v>", hex, expected, c)
		}
	}

}

func same(left, right color.Color) bool {
	return left == right
}
