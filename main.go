package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"
)

func main() {
	// Define flags for the input and output files
	inputFile := flag.String("i", "", "Input file")
	outputFile := flag.String("o", "output.png", "Output file")

	// Parse the flags
	flag.Parse()

	// Check if the input file flag is set
	if *inputFile == "" {
		fmt.Println("Please specify the input file with the -i flag.")
		return
	}

	// Open the input image
	in, err := os.Open(*inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer in.Close()

	// Decode the input image
	img, err := png.Decode(in)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the dimensions of the input image
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y

	// Create a new image with the desired dimensions
	cropped := image.NewRGBA(image.Rect(0, 0, width, height-200))

	// Copy the pixels from the input image to the new image
	for y := 0; y < height-200; y++ {
		for x := 0; x < width; x++ {
			cropped.Set(x, y, img.At(x, y))
		}
	}

	// Open the output image file
	out, err := os.Create(*outputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()

	// Encode and write the output image
	err = png.Encode(out, cropped)
	if err != nil {
		fmt.Println(err)
		return
	}
}

