package derivatives

import (
	"math"
	"sigma-vega/lbr"
	"sigma-vega/logger"

	"go.uber.org/zap"
)

type Greeks struct {
	Delta float64
	Theta float64
	Gamma float64
	Vega  float64
}

func (g *Greeks) String() {
	logger.Logz.Info("",
		zap.Float64("delta", g.Delta),
		zap.Float64("gamma", g.Gamma),
		zap.Float64("theta", g.Theta),
		zap.Float64("vega", g.Vega))
}

func callDelta(s, k, r, t, q, sigma float64) float64 {
	// return ePowMinusqT(q, t) * lbr.NormCdf(d1(s, k, r, t, q, sigma)) // TODO: determine importance
	return lbr.NormCdf(d1(s, k, r, t, q, sigma))
}

func putDelta(s, k, r, t, q, sigma float64) float64 {
	return lbr.NormCdf(d1(s, k, r, t, q, sigma)) - 1
}

func callTheta(s, k, r, t, q, sigma float64) float64 {
	phi := lbr.NormPdf(d1(s, k, r, t, q, sigma))

	var rate = r - q
	dTwo := d2(s, k, r, t, q, sigma)

	Sert := s * math.Pow(lbr.Exp, -rate*t)

	var a = -1 * ((Sert * sigma * phi) / (2 * math.Sqrt(t)))
	var b = r * Sert * lbr.NormCdf(d1(s, k, r, t, q, sigma)) / 100
	var c = r * k * math.Pow(lbr.Exp, -rate*t) * lbr.NormCdf(dTwo)

	return (a + b - c) * ONE_DAY_IN_YEARS
}

func gamma(s, k, r, t, q, sigma float64) float64 {
	phi := lbr.NormPdf(d1(s, k, r, t, q, sigma))

	two := ePowMinusqT(q, t)
	three := s * sigma * math.Sqrt(t)

	return (phi * two) / three
}

func vega(s, k, r, t, q, sigma float64) float64 {
	return s * math.Sqrt(t) * lbr.NormPdf(d1(s, k, r, t, q, sigma)) * 0.01
}

func putTheta(s, k, r, t, q, sigma float64) float64 {
	phi := lbr.NormPdf(d1(s, k, r, t, q, sigma))

	var rate = r - q
	dTwo := d2(s, k, r, t, q, sigma)

	Sert := s * math.Pow(Exp, -rate*t)

	var a = -1 * ((Sert * sigma * phi) / (2 * math.Sqrt(t)))
	var b = r * Sert * lbr.NormCdf(-d1(s, k, r, t, q, sigma)) / 100
	var c = r * k * math.Pow(lbr.Exp, -rate*t) * lbr.NormCdf(-dTwo)

	return (a - b + c) * ONE_DAY_IN_YEARS
}
