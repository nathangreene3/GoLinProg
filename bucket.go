package main

import (
	"math"
)

type bucket struct {
	topRadius    float64
	bottomRadius float64
	height       float64
}

func (b *bucket) surfaceArea() float64 {
	return math.Pi * (math.Pow(b.bottomRadius, 2.0) + (b.topRadius+b.bottomRadius)*math.Sqrt(math.Pow(b.topRadius-b.bottomRadius, 2.0)+math.Pow(b.height, 2.0)))
}

func (b *bucket) volume() float64 {
	return math.Pi * b.height / 3.0 * (math.Pow(b.topRadius, 2.0) + b.topRadius*b.bottomRadius + math.Pow(b.bottomRadius, 2.0))
}

func (b *bucket) setHeight(area float64) {
	b.height = math.Sqrt(math.Pow((area-math.Pi*math.Pow(b.bottomRadius, 2.0))/(math.Pi*(b.topRadius+b.bottomRadius)), 2.0) - math.Pow(b.topRadius-b.bottomRadius, 2.0))
}

// Slooow. Don't set e < 0.0001.
func (b *bucket) optimize(area, e float64) {
	var (
		volume          float64
		optVolume       float64
		optTopRadius    float64
		optBottomRadius float64
	)
	b.bottomRadius = e
	for b.bottomRadius < 1.0/math.SqrtPi {
		b.topRadius = e
		for b.topRadius < 1.0 {
			b.setHeight(area)
			volume = b.volume()
			if optVolume < volume {
				optVolume = volume
				optTopRadius = b.topRadius
				optBottomRadius = b.bottomRadius
			}
			b.topRadius += e
		}
		b.bottomRadius += e
	}
	b.topRadius = optTopRadius
	b.bottomRadius = optBottomRadius
	b.setHeight(area)
}

func gradient(f func(x, y float64) float64, e float64) {

}
