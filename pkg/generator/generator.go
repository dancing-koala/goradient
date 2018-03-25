package generator

import (
	"errors"
	"fmt"
	"github.com/dancing-koala/gradient/pkg/hexcolor"
	"github.com/dancing-koala/gradient/pkg/interpolator"
	"image"
	"image/color"
)

const (
	TYPE_POLYLINEAR = "polylinear"
	TYPE_QUADRATIC  = "quadratic"
	TYPE_CUBIC      = "cubic"
)

type Generator func(hexColors []string, img *image.RGBA, w, h int) error

func GetGenerator(gradientType string) (Generator, error) {

	switch gradientType {
	case TYPE_POLYLINEAR:
		return polylinearGenerator, nil

	case TYPE_QUADRATIC:
		return quadraticGenerator, nil

	case TYPE_CUBIC:
		return cubicGenerator, nil
	}

	return nil, errors.New("Unknown gradient type <" + gradientType + ">")
}

func drawGradient(img *image.RGBA, height, start, end int, rInterpol, gInterpol, bInterpol, aInterpol interpolator.Interpolator) {
	var newa, newr, newg, newb uint8
	var progress float64

	for i := start; i < end; i++ {

		progress = float64(i-start) / float64(end-start)

		newr = uint8(rInterpol.Interpolate(progress))
		newg = uint8(gInterpol.Interpolate(progress))
		newb = uint8(bInterpol.Interpolate(progress))
		newa = uint8(aInterpol.Interpolate(progress))

		c := &color.RGBA{
			R: newr,
			G: newg,
			B: newb,
			A: newa,
		}

		for j := 0; j < height; j++ {
			img.Set(i, j, c)
		}
	}
}

func createSteps(steps []step, size int) {
	gapSize := size / len(steps)

	steps[0] = step{
		start: 0,
		end:   gapSize,
	}

	if gapSize != size {
		for i := 1; i < len(steps); i++ {
			steps[i] = step{
				start: steps[i-1].start + gapSize,
				end:   steps[i-1].end + gapSize,
			}
		}
	}

	if steps[len(steps)-1].end < size {
		steps[len(steps)-1].end = size
	}
}

type step struct {
	start int
	end   int
}

func polylinearGenerator(hexColors []string, img *image.RGBA, w, h int) error {
	if len(hexColors) < 2 {
		return errors.New(fmt.Sprintf("Polylinear gradient requires at least 2 colors, %d given", len(hexColors)))
	}

	steps := make([]step, len(hexColors)-1)

	createSteps(steps, w)

	var startColor, endColor color.Color
	var err error

	offset := 0

	for _, currStep := range steps {
		startColor, err = hexcolor.ParseToRGBA(hexColors[offset])

		if err != nil {
			return err
		}

		endColor, err = hexcolor.ParseToRGBA(hexColors[offset+1])

		if err != nil {
			return err
		}

		r1, g1, b1, a1 := startColor.RGBA()
		r2, g2, b2, a2 := endColor.RGBA()

		rInterpol := interpolator.NewLinear(float64(r1&0xFF), float64(r2&0xFF))
		gInterpol := interpolator.NewLinear(float64(g1&0xFF), float64(g2&0xFF))
		bInterpol := interpolator.NewLinear(float64(b1&0xFF), float64(b2&0xFF))
		aInterpol := interpolator.NewLinear(float64(a1&0xFF), float64(a2&0xFF))

		drawGradient(img, h, currStep.start, currStep.end, rInterpol, gInterpol, bInterpol, aInterpol)

		offset++
	}

	return nil
}

func quadraticGenerator(hexColors []string, img *image.RGBA, w, h int) error {
	if len(hexColors) < 3 {
		return errors.New(fmt.Sprintf("Quadratic gradient requires at least 3 colors, %d given", len(hexColors)))
	}

	var startColor, midColor, endColor color.Color
	var err error

	startColor, err = hexcolor.ParseToRGBA(hexColors[0])

	if err != nil {
		return err
	}

	midColor, err = hexcolor.ParseToRGBA(hexColors[1])

	if err != nil {
		return err
	}

	endColor, err = hexcolor.ParseToRGBA(hexColors[2])

	if err != nil {
		return err
	}

	r1, g1, b1, a1 := startColor.RGBA()
	r2, g2, b2, a2 := midColor.RGBA()
	r3, g3, b3, a3 := endColor.RGBA()

	rInterpol := interpolator.NewQuadratic(float64(r1&0xFF), float64(r2&0xFF), float64(r3&0xFF))
	gInterpol := interpolator.NewQuadratic(float64(g1&0xFF), float64(g2&0xFF), float64(g3&0xFF))
	bInterpol := interpolator.NewQuadratic(float64(b1&0xFF), float64(b2&0xFF), float64(b3&0xFF))
	aInterpol := interpolator.NewQuadratic(float64(a1&0xFF), float64(a2&0xFF), float64(a3&0xFF))

	drawGradient(img, h, 0, w, rInterpol, gInterpol, bInterpol, aInterpol)

	return nil
}

func cubicGenerator(hexColors []string, img *image.RGBA, w, h int) error {
	if len(hexColors) < 4 {
		return errors.New(fmt.Sprintf("Cubic gradient requires at least 4 colors, %d given", len(hexColors)))
	}

	var startColor, stopAColor, stopBColor, endColor color.Color
	var err error

	startColor, err = hexcolor.ParseToRGBA(hexColors[0])

	if err != nil {
		return err
	}

	stopAColor, err = hexcolor.ParseToRGBA(hexColors[1])

	if err != nil {
		return err
	}

	stopBColor, err = hexcolor.ParseToRGBA(hexColors[2])

	if err != nil {
		return err
	}

	endColor, err = hexcolor.ParseToRGBA(hexColors[3])

	if err != nil {
		return err
	}

	r1, g1, b1, a1 := startColor.RGBA()
	r2, g2, b2, _ := stopAColor.RGBA()
	r3, g3, b3, _ := stopBColor.RGBA()
	r4, g4, b4, a4 := endColor.RGBA()

	rInterpol := interpolator.NewCubic(float64(r1&0xFF), float64(r2&0xFF), float64(r3&0xFF), float64(r4&0xFF))
	gInterpol := interpolator.NewCubic(float64(g1&0xFF), float64(g2&0xFF), float64(g3&0xFF), float64(g4&0xFF))
	bInterpol := interpolator.NewCubic(float64(b1&0xFF), float64(b2&0xFF), float64(b3&0xFF), float64(b4&0xFF))
	aInterpol := interpolator.NewLinear(float64(a1&0xFF), float64(a4&0xFF))

	drawGradient(img, h, 0, w, rInterpol, gInterpol, bInterpol, aInterpol)

	return nil
}
