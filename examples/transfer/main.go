package main

import (
	"flag"
	"image"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/paulvasilenko/go-transcolor"
)

var (
	sourceFile = flag.String("source", "", "Source file which colors would be used as a source")
	targetFile = flag.String("target", "", "Target file is color to which we apply source color palette")
	outputPath = flag.String("out", "", "File where to save the result")
)

func main() {
	flag.Parse()
	if *sourceFile == "" {
		log.Fatalf("missing source file")
	}

	if *targetFile == "" {
		log.Fatalf("missing target file")
	}

	if *outputPath == "" {
		log.Fatalf("missing output path")
	}

	src, err := openImage(*sourceFile)
	if err != nil {
		log.Fatalf("failed to open source: %v", err)
	}

	target, err := openImage(*targetFile)
	if err != nil {
		log.Fatalf("failed to open target: %v", err)
	}

	res := transcolor.Transfer(src, target)

	if err := saveImage(res, *outputPath); err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}

func saveImage(img image.Image, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}

func openImage(path string) (image.Image, error) {
	src, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer src.Close()

	srcImg, _, err := image.Decode(src)
	if err != nil {
		return nil, err
	}

	return srcImg, nil
}
