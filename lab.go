package transcolor

import (
	"image"
	"math"
)

// Lab is a representation of L*a*b* image
type Lab struct {
	Pix []float64
}

// ImageToLab converts image to L*a*b* image.
// That is made public for future needs to expose Lab image struct if necessary
func ImageToLab(src image.Image) *Lab {
	lab := &Lab{}
	forEachImage(src, func(R, G, B uint32) {
		l, a, b := rgbToLab(R, G, B)
		lab.Pix = append(lab.Pix, l, a, b)
	})

	return lab
}

// LabStat is a set of stats calculated for each channel
type LabStat struct {
	LStat Stat
	AStat Stat
	BStat Stat
}

// Stat is a set of channel statistics
type Stat struct {
	Mean   float64
	StdDev float64
}

// Stat calculates L*a*b* image stats
func (src *Lab) Stat() *LabStat {
	var lMean, aMean, bMean float64
	forEachLAB(src, func(l, a, b float64) {
		lMean += l
		aMean += a
		bMean += b
	})

	amount := float64(len(src.Pix)) / 3

	lMean /= amount
	aMean /= amount
	bMean /= amount

	var lStd, aStd, bStd float64
	forEachLAB(src, func(l, a, b float64) {
		lStd += math.Pow(l-lMean, 2)
		aStd += math.Pow(a-aMean, 2)
		bStd += math.Pow(b-bMean, 2)
	})

	lStd = math.Sqrt(lStd / (amount))
	aStd = math.Sqrt(aStd / (amount))
	bStd = math.Sqrt(bStd / (amount))

	return &LabStat{
		LStat: Stat{
			Mean:   lMean,
			StdDev: lStd,
		},
		AStat: Stat{
			Mean:   aMean,
			StdDev: aStd,
		},
		BStat: Stat{
			Mean:   bMean,
			StdDev: bStd,
		},
	}
}
