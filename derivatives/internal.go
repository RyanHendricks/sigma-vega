package derivatives

import (
	"math"
)

// d1 = (ln(s/k) + (r+((sigma^2)/2))t) / sigma*sqrt(t)
func d1(s, k, r, t, q, sigma float64) float64 {
	return (math.Log(s/k) + ((r - q + 0.5*(sigma*sigma)) * t)) / (sigma * math.Sqrt(t))
}

// d2 = d1 - (sigma * sqrt(t))
func d2(s, k, r, t, q, sigma float64) float64 {
	return d1(s, k, r, t, q, sigma) - (sigma * math.Sqrt(t))
}

// ePowMinusqT returns e ^ -qt where q represents the yield of the underlying
func ePowMinusqT(q, t float64) float64 {
	return math.Exp(-q * t)
}

func withinRequiredPrecision(x, y, precision float64) bool {
	return math.Abs(x-y) <= precision
}

func withinRequiredPercent(x, y, percent float64) bool {
	if x == y {
		return true
	}
	return 1-math.Abs(x/y) <= percent
}

// GenerateIncrementalArray returns an array of float64 values given the
// median, qty of values (including the median), and the increment between
// each value. NOTE: for an even qty, the returned value will not the same
// qty of values on either side of the median (technically no longer the median)
// median=3, qty=4, increment=1 -> [1, 2, 3, 4]
// median=3, qty=4, increment=1 -> [1, 2, 3, 4]
func GenerateIncrementalArray(median, qty, increment float64) []float64 {
	p := median - (0.5 * qty * increment)

	prices := make([]float64, int(qty))

	for i := 0; i < int(qty); i++ {
		val := p + (increment * float64(i))
		prices[i] = val
	}

	return prices
}

type fn func(float64) float64

func SomeFunction(values []float64, calc fn) []float64 {
	results := make([]float64, len(values))
	for v := range values {
		r := calc(values[v])
		results[v] = r
	}

	return results
}
