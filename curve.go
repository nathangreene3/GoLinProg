package main

import "math"

type polynomial []float64

func (p polynomial) evaluate(x float64) float64 {
	var y float64
	for i := range p {
		y += math.Pow(x, float64(i)) * p[i]
	}
	return y
}

func derivative(p polynomial) polynomial {
	var q polynomial
	if 0 < len(p) {
		q = make(polynomial, len(p)-1)
		for i := range q {
			q[i] = float64(i) * p[i+1]
		}
	} else {
		panic("cannot differentiate polynomial p any further")
	}
	return q
}

func findRoot(p polynomial) {

}

func newton(p polynomial, x0, e float64) float64 {
	var x1 float64
	q := derivative(p)
	x1 = x0 - q.evaluate(x0)/p.evaluate(x0)
	return x1
}
