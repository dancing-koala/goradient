package hexcolor

import (
	"errors"
	"image/color"
	"regexp"
	"strconv"
)

var (
	validHexRegex = regexp.MustCompile("^#([A-F0-9]{2}){3,4}$")
)

const (
	maskR = 0x00FF0000
	maskG = 0x0000FF00
	maskB = 0x000000FF
	maskA = 0xFF000000

	defaultA = uint8(0xFF)
)

func IsValid(hex string) bool {
	return validHexRegex.Match([]byte(hex))
}

func ParseToRGBA(hex string) (color.Color, error) {
	if !IsValid(hex) {
		return nil, errors.New(hex + " is not valid !")
	}

	withAlpha := len([]rune(hex)) > 7

	colorInt, err := strconv.ParseUint(hex[1:], 16, 64)

	if err != nil {
		return nil, errors.New(hex + " could not be parsed: " + err.Error())
	}

	a := defaultA

	if withAlpha {
		a = uint8(colorInt & maskA >> 24)
	}

	r := uint8(colorInt & maskR >> 16)
	g := uint8(colorInt & maskG >> 8)
	b := uint8(colorInt & maskB)

	return color.RGBA{
		R: r,
		G: g,
		B: b,
		A: a,
	}, nil
}
