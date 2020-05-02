package lbr

import (
	"math"
	"testing"
)

type Test struct {
	CallPut      float64
	Strike       float64
	Maturity     float64
	RiskFreeRate float64
	DividendRate float64
	Underlying   float64
	Price        float64
	IV           float64
}

var tests = []Test{
	{1, 1220, 6.0 / 365, 0.0001, 0.0127, 1240.41, 22.1, 13.46},
	{-1, 1220, 6.0 / 365, 0.0001, 0.0127, 1240.41, 3.1, 16.13},
	{1, 1240, 6.0 / 365, 0.0001, 0.0127, 1240.41, 8.1, 12.65},
	{-1, 1240, 6.0 / 365, 0.0001, 0.0127, 1240.41, 9.4, 14.94},

	{1, 1220, 4.0 / 252, 0.0001, 0.0127, 1240.41, 22.1, 13.68},
	{-1, 1220, 4.0 / 252, 0.0001, 0.0127, 1240.41, 3.1, 16.42},
	{1, 1240, 4.0 / 252, 0.0001, 0.0127, 1240.41, 8.1, 12.87},
	{-1, 1240, 4.0 / 252, 0.0001, 0.0127, 1240.41, 9.4, 15.21},
}

func TestIV(t *testing.T) {
	for _, test := range tests {
		F := test.Underlying * math.Exp((test.RiskFreeRate-test.DividendRate)*test.Maturity)
		iv := 100 * ImpliedVolatilityFromATransformedRationalGuess(
			test.Price,
			F,
			test.Strike,
			test.Maturity,
			test.CallPut,
		)
		if math.Abs(iv-test.IV) > 0.01 {
			t.Error("Expected", test.IV, "Got", iv)
		}
	}
}

func TestNormPdfCdf(t *testing.T) {
	testvars := []struct {
		name string
		x    float64
		want float64
	}{
		// { "TestNormPdfCdf", 1.96, 0.0278},
		// { "TestNormPdfCdf2", 0.0430, 0.0278},
		{"TestNormPdfCdf2", 0.2475, 0.5977393542509256},
		{"TestNormPdfCdf2", 0.1061, 0.5422484944267403},
	}
	for _, tt := range testvars {
		t.Run(tt.name, func(t *testing.T) {
			// if got := NormPdf(tt.x); got != tt.want {
			//     t.Errorf("NormPdf() = %v, want %v", got, tt.want)
			// }
			if got := NormCdf(tt.x); got != tt.want {
				t.Errorf("NormCdf() = %v, want %v", got, tt.want)
			}
		})
	}
}
