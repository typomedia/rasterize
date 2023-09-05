package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/pflag"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

func main() {
	start := time.Now()
	input := pflag.StringP("input", "i", "", "Input filename")
	output := pflag.StringP("output", "o", "", "Output filename")
	size := pflag.IntP("size", "s", 1000, "Size in pixels")
	debug := pflag.BoolP("debug", "d", false, "Debug")
	pflag.Parse()

	svg, _ := oksvg.ReadIcon(*input, oksvg.WarnErrorMode)
	width := int(svg.ViewBox.W)
	height := int(svg.ViewBox.H)

	if *debug {
		fmt.Println("width: ", width)
		fmt.Println("height: ", height)
	}

	if *size > 10000 {
		*size = 10000
	}

	width, height = newDimensions(width, height, *size)

	if *debug {
		fmt.Println("width: ", width)
		fmt.Println("height: ", height)
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	scanner := rasterx.NewScannerGV(width, height, img, img.Bounds())
	raster := rasterx.NewDasher(width, height, scanner)

	svg.SetTarget(0, 0, float64(width), float64(height))

	opacity := 1.0
	svg.Draw(raster, opacity)

	name := basename(*input) + ".png"
	if *output != "" {
		name = *output
	}

	if *debug {
		fmt.Println("Output:", name)
	}

	// Save the PNG file
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}

	t := time.Now()
	elapsed := t.Sub(start)

	if *debug {
		fmt.Println("Elapsed:", elapsed)
	}
}

func maxDimension(width, height int) int {
	size := width
	if height > width {
		size = height
	}
	return size
}

func newDimensions(width, height, size int) (int, int) {
	max := maxDimension(width, height)
	width = (width * size) / max
	height = (height * size) / max

	return width, height
}

func basename(filePath string) string {
	fileName := filepath.Base(filePath)
	extension := filepath.Ext(fileName)
	return fileName[0 : len(fileName)-len(extension)]
}
