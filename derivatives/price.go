package derivatives

import (
	"math"
)

const Exp = 2.71828182845904523536028747135266249775724709369995957496696763 // https://oeis.org/A001113

// forwardPrice returns the Forward price for an asset
// s = current price of the yield-bearing asset
// t = time to maturity of forward contract in years
// r = risk free rate of interest with continuous compounding
// q = average yield per annum with continuous compounding
// F0 = S0e^(r-q)T or S0e^rT (if no dividend; q = 0)
func forwardPrice(s, t, r, q float64) float64 {
	if q != 0 {
		return forwardPriceWithYield(s, t, r, q)
	}
	return s * math.Pow(Exp, r*t)
}

//
// forwardPriceWithYield returns the Forward price for an asset with a dividend
// F0 = S0e^(r-q)T
func forwardPriceWithYield(s, t, r, q float64) float64 {
	exponent := (r - q) * t
	fPrice := s * math.Exp(exponent)
	return fPrice
}

// Value at expiration for a call option
// (s - k)+ where k = spot price k = strike price
func CallPayoff(s, k float64) float64 {
	if k >= s {
		return 0
	}
	return s - k
}

// Value at expiration for a put option
// (k - s)+ where k = spot price k = strike price
func PutPayoff(s, k float64) float64 {
	if k <= s {
		return 0
	}
	return k - s
}
