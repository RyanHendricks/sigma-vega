package lbr

import "math"

//
// The original source code resides at www.jaeckel.org/LetsBeRational.7z.
//
// Comments unchanged
//
// ======================================================================================
// Copyright © 2013-2014 Peter Jäckel.
//
// Permission to use, copy, modify, and distribute this software is freely granted,
// provided that this notice is preserved.
//
// WARRANTY DISCLAIMER
// The Software is provided "as is" without warranty of any kind, either express or implied,
// including without limitation any implied warranties of condition, uninterrupted use,
// merchantability, fitness for a particular purpose, or non-infringement.
// ======================================================================================
//
// The asymptotic expansion  Φ(z) = φ(z)/|z|·[1-1/z^2+...],  Abramowitz & Stegun (26.2.12), suffices for Φ(z) to have
// relative accuracy of 1.64E-16 for z<=-10 with 17 terms inside the square brackets (not counting the leading 1).
// This translates to a maximum of about 9 iterations below, which is competitive with a call to erfc() and never
// less accurate when z<=-10. Note that, as mentioned in section 4 (and discussion of figures 2 and 3) of George
// Marsaglia's article "Evaluating the Normal Distribution" (available at http://www.jstatsoft.org/v11/a05/paper),
// for values of x approaching -8 and below, the error of any cumulative normal function is actually dominated by
// the hardware (or compiler implementation) accuracy of exp(-x²/2) which is not reliably more than 14 digits when
// x becomes large. Still, we should switch to the asymptotic only when it is beneficial to do so.
//
// type LbrNormDist interface {
//     erfCody(float64) float64
//     erfcCody(float64) float64
//     erfcxCody(float64) float64
//     normCdf(float64) float64
//     normPdf(float64) float64
//     inverseNormCdf(float64) float64
// }

func erfcxCody(x float64) float64 {
	return math.Exp(x*x) * math.Erfc(x)
}

func NormCdf(z float64) float64 {
	if z <= NormCdfAsymptoticExpansionFirstThreshold {
		var sum = 1.0

		if z >= NormCdfAsymptoticExpansionSecondThreshold {
			var zsqr = z * z
			var i float64 = 1
			var g float64 = 1
			var x, y float64
			var a = math.MaxFloat64
			var lasta float64

			for {
				lasta = a
				x = (4*i - 3) / zsqr
				y = x * ((4*i - 1) / zsqr)
				a = g * (x - y)
				sum -= a
				g *= y
				i++

				a = math.Abs(a)

				if !(lasta > a && a >= math.Abs(sum*DblEpsilon)) {
					break
				}
			}
		}
		return -NormPdf(z) * sum / z
	}
	return 0.5 * math.Erfc(-z*OneOverSqrtTwo)
}

func NormPdf(x float64) float64 {
	return OneOverSqrtTwoPi * math.Exp(-0.5*x*x)
}

//
// ALGORITHM AS241  APPL. STATIST. (1988) VOL. 37, NO. 3
//
// Produces the normal deviate Z corresponding to a given lower
// tail area of u Z is accurate to about 1 part in 10**16.
// see http://lib.stat.cmu.edu/apstat/241
//
func inverseNormCdf(u float64) float64 {
	const split1 float64 = 0.425
	const split2 float64 = 5.0
	const const1 float64 = 0.180625
	const const2 float64 = 1.6

	// Coefficients for P close to 0.5
	const A0 float64 = 3.3871328727963666080e0
	const A1 float64 = 1.3314166789178437745e+2
	const A2 float64 = 1.9715909503065514427e+3
	const A3 float64 = 1.3731693765509461125e+4
	const A4 float64 = 4.5921953931549871457e+4
	const A5 float64 = 6.7265770927008700853e+4
	const A6 float64 = 3.3430575583588128105e+4
	const A7 float64 = 2.5090809287301226727e+3
	const B1 float64 = 4.2313330701600911252e+1
	const B2 float64 = 6.8718700749205790830e+2
	const B3 float64 = 5.3941960214247511077e+3
	const B4 float64 = 2.1213794301586595867e+4
	const B5 float64 = 3.9307895800092710610e+4
	const B6 float64 = 2.8729085735721942674e+4
	const B7 float64 = 5.2264952788528545610e+3
	// Coefficients for P not close to 0, 0.5 or 1.
	const C0 float64 = 1.42343711074968357734e0
	const C1 float64 = 4.63033784615654529590e0
	const C2 float64 = 5.76949722146069140550e0
	const C3 float64 = 3.64784832476320460504e0
	const C4 float64 = 1.27045825245236838258e0
	const C5 float64 = 2.41780725177450611770e-1
	const C6 float64 = 2.27238449892691845833e-2
	const C7 float64 = 7.74545014278341407640e-4
	const D1 float64 = 2.05319162663775882187e0
	const D2 float64 = 1.67638483018380384940e0
	const D3 float64 = 6.89767334985100004550e-1
	const D4 float64 = 1.48103976427480074590e-1
	const D5 float64 = 1.51986665636164571966e-2
	const D6 float64 = 5.47593808499534494600e-4
	const D7 float64 = 1.05075007164441684324e-9
	// Coefficients for P very close to 0 or 1
	const E0 float64 = 6.65790464350110377720e0
	const E1 float64 = 5.46378491116411436990e0
	const E2 float64 = 1.78482653991729133580e0
	const E3 float64 = 2.96560571828504891230e-1
	const E4 float64 = 2.65321895265761230930e-2
	const E5 float64 = 1.24266094738807843860e-3
	const E6 float64 = 2.71155556874348757815e-5
	const E7 float64 = 2.01033439929228813265e-7
	const F1 float64 = 5.99832206555887937690e-1
	const F2 float64 = 1.36929880922735805310e-1
	const F3 float64 = 1.48753612908506148525e-2
	const F4 float64 = 7.86869131145613259100e-4
	const F5 float64 = 1.84631831751005468180e-5
	const F6 float64 = 1.42151175831644588870e-7
	const F7 float64 = 2.04426310338993978564e-15

	if u <= 0 {
		return math.Log(u)
	}

	if u >= 1 {
		return math.Log(1 - u)
	}

	var q = u - 0.5
	if math.Abs(q) <= split1 {
		var r = const1 - q*q
		return q * (((((((A7*r+A6)*r+A5)*r+A4)*r+A3)*r+A2)*r+A1)*r + A0) /
			(((((((B7*r+B6)*r+B5)*r+B4)*r+B3)*r+B2)*r+B1)*r + 1.0)
	} else {
		var r float64
		if q < 0 {
			r = u
		} else {
			r = 1.0 - u
		}
		r = math.Sqrt(-math.Log(r))
		var ret float64
		if r < split2 {
			r = r - const2
			ret = (((((((C7*r+C6)*r+C5)*r+C4)*r+C3)*r+C2)*r+C1)*r + C0) /
				(((((((D7*r+D6)*r+D5)*r+D4)*r+D3)*r+D2)*r+D1)*r + 1.0)
		} else {
			r = r - split2
			ret = (((((((E7*r+E6)*r+E5)*r+E4)*r+E3)*r+E2)*r+E1)*r + E0) /
				(((((((F7*r+F6)*r+F5)*r+F4)*r+F3)*r+F2)*r+F1)*r + 1.0)
		}
		if q < 0 {
			return -ret
		} else {
			return ret
		}
	}
}
