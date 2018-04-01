package main

import (
	"flag"
	"fmt"
	"github.com/dancing-koala/goradient/pkg/generator"
	"image"
	"image/png"
	"os"
	"strings"
	"time"
)

var (
	w, h         int
	outPath      string
	gradientType string
)

func main() {

	flag.IntVar(&w, "w", 800, "Width of the generated gradient image (default is 800)")
	flag.IntVar(&h, "h", 200, "Height of the generated gradient image (default is 200)")
	flag.StringVar(&outPath, "out", "./gradient-gen.png", "Path of the gradient image to be generated (default is './gradient-gen.png')")
	flag.StringVar(&gradientType, "type", generator.TYPE_POLYLINEAR, "Type of the gradient to generate (polylinear, quadratic or cubic")

	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("You must pass colors as a comma separated list, for example '#000000,#FFFFFF'")
		return
	}

	hexColors := strings.Split(os.Args[len(os.Args)-1], ",")

	fmt.Println("Generating "+gradientType+" gradient with colors:", hexColors)

	img := image.NewRGBA(image.Rect(0, 0, w, h))

	g, err := generator.GetGenerator(gradientType)

	if err != nil {
		fmt.Println("Error while getting generator: " + err.Error())
	}

	start := time.Now()

	err = g(hexColors, img, w, h)

	if err != nil {
		fmt.Println("Error while generating gradient: ", err)
		return
	}

	fmt.Println("Processing took ", time.Since(start))

	fmt.Println("Saving image to output file...")
	f, _ := os.OpenFile(outPath, os.O_WRONLY|os.O_CREATE, 0777)
	defer f.Close()

	err = png.Encode(f, img)

	if err != nil {
		fmt.Println("Error while writing generated image: ", err)
		return
	}

	fmt.Println("Done !!!")
}
