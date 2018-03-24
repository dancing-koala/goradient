package main

import (
	"fmt"
	"github.com/dancing-koala/gradient/pkg/hexcolor"
	"github.com/dancing-koala/gradient/pkg/interpolator"
	"image"
	"image/color"
	"image/png"
	"os"
)

type step struct {
	start int
	end   int
}

func main() {

	left, top := 0, 0
	width, height := 800, 200

	hexColors := []string{
		"#00FF00",
		"#FF0000",
		// "#00FF00",
		// "#FFFF00",
		// "#FF0000",
		// "#FFFF00",
		// "#00FF00",
		// "#00FFFF",
		// "#0000FF",
	}

	steps := make([]step, len(hexColors)-1)

	createSteps(steps, width)

	offset := 0

	img := image.NewRGBA(image.Rect(left, top, width, height))

	var startColor, endColor color.Color
	var err error

	for _, currStep := range steps {
		startColor, err = hexcolor.ParseToRGBA(hexColors[offset])

		if err != nil {
			fmt.Printf("Could not parse start color<%s>: %s\n", hexColors[offset], err)
			return
		}

		endColor, err = hexcolor.ParseToRGBA(hexColors[offset+1])

		if err != nil {
			fmt.Printf("Could not parse end color<%s>: %s\n", hexColors[offset+1], err)
			return
		}

		r, g, b, a := startColor.RGBA()
		rr, gg, bb, aa := endColor.RGBA()

		rInterpol := interpolator.NewLinearInterpolator(float64(r&0xFF), float64(rr&0xFF))
		gInterpol := interpolator.NewLinearInterpolator(float64(g&0xFF), float64(gg&0xFF))
		bInterpol := interpolator.NewLinearInterpolator(float64(b&0xFF), float64(bb&0xFF))
		aInterpol := interpolator.NewLinearInterpolator(float64(a&0xFF), float64(aa&0xFF))

		var newa, newr, newg, newb uint8
		var progress float64

		for i := currStep.start; i < currStep.end; i++ {

			progress = float64(i-currStep.start) / float64(currStep.end-currStep.start)

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

		offset++
	}

	fmt.Println("Saving image to output file...")
	f, _ := os.OpenFile("./gen/gradient.png", os.O_WRONLY|os.O_CREATE, 0777)
	defer f.Close()

	err = png.Encode(f, img)

	if err != nil {
		fmt.Println("Error while writing generated image: ", err)
		return
	}

	fmt.Println("Done !!!")
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
