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

func IsValid(hex string) bool {
	return validHexRegex.Match([]byte(hex))
}

func ParseToRGBA(hex string) (color.Color, error) {
	if !IsValid(hex) {
		return nil, errors.New(hex + " is not valid !")
	}

	runes := []rune(hex)

	offset := 1

	a := uint8(0xFF)

	if len(runes) > 7 {
		a = parseUint8(string(runes[offset : offset+2]))

		offset += 2
	}

	r := parseUint8(string(runes[offset : offset+2]))
	g := parseUint8(string(runes[offset+2 : offset+4]))
	b := parseUint8(string(runes[offset+4:]))

	return color.RGBA{
		A: a,
		R: r,
		G: g,
		B: b,
	}, nil
}

func parseUint8(str string) uint8 {
	result, _ := strconv.ParseUint(str, 16, 8)
	return uint8(result)
}
