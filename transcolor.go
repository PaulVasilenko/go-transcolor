package transcolor

import (
	"image"
	"image/color"

	"github.com/lucasb-eyer/go-colorful"
)

// Transfer transfers color palette from one image to other
// It adjust color of source image based on target image
func Transfer(src, target image.Image) image.Image {
	srcLab := ImageToLab(src)
	targetLab := ImageToLab(target)

	srcLabStat := srcLab.Stat()
	targetLabStat := targetLab.Stat()

	forEachLABCounter(targetLab, func(l, a, b float64, counter int) {
		targetLab.Pix[counter] = calculateNewPix(l, targetLabStat.LStat, srcLabStat.LStat)
		targetLab.Pix[counter+1] = calculateNewPix(a, targetLabStat.AStat, srcLabStat.AStat)
		targetLab.Pix[counter+2] = calculateNewPix(b, targetLabStat.BStat, srcLabStat.BStat)
	})

	newTargetRGBA := image.NewRGBA(target.Bounds())
	ind := 0
	for x := target.Bounds().Min.X; x < target.Bounds().Max.X; x++ {
		for y := target.Bounds().Min.Y; y < target.Bounds().Max.Y; y++ {
			_, _, _, A := target.At(x, y).RGBA()
			c := colorful.Lab(targetLab.Pix[ind], targetLab.Pix[ind+1], targetLab.Pix[ind+2])
			R, G, B := c.Clamped().RGB255()
			newTargetRGBA.Set(x, y, color.RGBA{
				R: R,
				G: G,
				B: B,
			})
			ind += 3
		}
	}

	return newTargetRGBA
}

func calculateNewPix(src float64, targetStat, srcStat Stat) float64 {
	return (src-targetStat.Mean)*(srcStat.StdDev/targetStat.StdDev) + srcStat.Mean
}
