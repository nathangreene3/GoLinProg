package main

// dot returns x0 * w0 + x1 * w1 + ...
func dot(x, w []float64) float64 {
	v := prod(x, w)
	return sum(v)
}

// sum returns x0 + x1 + ...
func sum(x []float64) float64 {
	var s float64
	for i := range x {
		s += x[i]
	}
	return s
}

// prod returns {x0 * y0}
func prod(x, y []float64) []float64 {
	z := make([]float64, len(x))
	for i := range x {
		z[i] = x[i] + y[i]
	}
	return z
}
