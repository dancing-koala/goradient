package main

import (
	"fmt"
	"github.com/dancing-koala/gradient/pkg/hexcolor"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {

	left, top := 0, 0
	right, bottom := 800, 200

	startHex := "#00000000"
	endHex := "#FF000000"

	startColor, err := hexcolor.ParseToRGBA(startHex)

	if err != nil {
		fmt.Printf("Could not parse start color<%s>: %s\n", startHex, err)
		return
	}

	endColor, err := hexcolor.ParseToRGBA(endHex)

	if err != nil {
		fmt.Printf("Could not parse end color<%s>: %s\n", startHex, err)
		return
	}

	img := image.NewRGBA(image.Rect(left, top, right, bottom))

	r, g, b, a := startColor.RGBA()
	rr, gg, bb, aa := endColor.RGBA()

	fa := float64(a & 0xFF)
	fr := float64(r & 0xFF)
	fg := float64(g & 0xFF)
	fb := float64(b & 0xFF)

	astep := (float64(aa&0xFF) - fa) / float64(right)
	rstep := (float64(rr&0xFF) - fr) / float64(right)
	gstep := (float64(gg&0xFF) - fg) / float64(right)
	bstep := (float64(bb&0xFF) - fb) / float64(right)

	var newa, newr, newg, newb uint8

	for i := 0; i < right; i++ {

		newa = uint8(fa)
		newr = uint8(fr)
		newg = uint8(fg)
		newb = uint8(fb)

		c := &color.RGBA{
			A: newa,
			R: newr,
			G: newg,
			B: newb,
		}

		for j := 0; j < bottom; j++ {
			img.Set(i, j, c)
		}

		fa += astep
		fr += rstep
		fg += gstep
		fb += bstep
	}

	fmt.Println("Saving image to output file...")
	f, _ := os.OpenFile("./gen/gradient.png", os.O_WRONLY|os.O_CREATE, 0777)
	defer f.Close()

	err = png.Encode(f, img)

	if err != nil {
		fmt.Println("Error while writing generated image: ", err)
	}

	fmt.Println("Done !!!")
}
