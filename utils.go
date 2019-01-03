package transcolor

import (
	"image"

	"github.com/lucasb-eyer/go-colorful"
)

func forEachImage(src image.Image, f func(r, g, b uint32)) {
	for x := src.Bounds().Min.X; x < src.Bounds().Max.X; x++ {
		for y := src.Bounds().Min.Y; y < src.Bounds().Max.Y; y++ {
			point := src.At(x, y)
			r, g, b, _ := point.RGBA()
			f(r, g, b)
		}
	}
}

func forEachLAB(src *Lab, f func(l, a, b float64)) {
	var l, a, b float64

	for i, pix := range src.Pix {
		switch i % 3 {
		case 0:
			l = pix
		case 1:
			a = pix
		case 2:
			b = pix
			f(l, a, b)
		}
	}
}

func rgbToLab(R, G, B uint32) (l, a, b float64) {
	c := colorful.Color{R: float64(R) / 65535.0, G: float64(G) / 65535.0, B: float64(B) / 65535.0}
	l, a, b = c.Lab()
	return
}

func forEachLABCounter(src *Lab, f func(l, a, b float64, counter int)) {
	counter := 0
	forEachLAB(src, func(l, a, b float64) {
		f(l, a, b, counter)
		counter += 3
	})
}
