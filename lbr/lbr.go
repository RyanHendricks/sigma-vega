package lbr

import (
	"math"
)

/* q=±1 */
// type letsberational interface {
//     set_implied_volatility_maximum_iterations(n float64) float64
//     set_implied_volatility_output_type(k float64) float64
//     set_implied_volatility_householder_method_order(m float64) float64
//     normalised_black_call(x, s float64) float64
//     normalised_vega(x, s float64) float64
//     normalised_black(x, s, q float64) float64
//     black(F, K, sigma, T, q float64) float64
//     normalised_implied_volatility_from_a_transformed_rational_guess_with_limited_iterations(beta, x, q float64, N int) float64
//     normalised_implied_volatility_from_a_transformed_rational_guess(beta, x, q float64) float64
//     implied_volatility_from_a_transformed_rational_guess_with_limited_iterations(price, F, K, T, q float64, N int) float64
//     implied_volatility_from_a_transformed_rational_guess(price, F, K, T float64, q int) float64
// }

func isBelowHorizon(x float64) bool {
	return math.Abs(x) < DenormalisationCutoff
}

func householderFactor(newton, halley, hh3 float64) float64 {
	return (1.0 + 0.5*halley*newton) / (1.0 + newton*(halley+hh3*newton/6.0))
}

func normalisedIntrinsic(x, q float64) float64 {
	if q*x <= 0 {
		return 0
	}

	var x2 = x * x

	var qSign float64 = 1

	if q < 0 {
		qSign = -1
	}

	if x2 < 98.0*FourthRootDbtEpsilon {
		return math.Abs(math.Max(qSign*x*(1+x2*((1.0/24.0)+x2*((1.0/1920.0)+x2*((1.0/322560.0)+(1.0/92897280.0)*x2)))), 0.0))
	}

	var bMax = math.Exp(0.5 * x)

	var oneOverBMax = 1.0 / bMax

	return math.Abs(math.Max(qSign*(bMax-oneOverBMax), 0.0))
}

func normalisedIntrinsicCall(x float64) float64 {
	return normalisedIntrinsic(x, 1.0)
}

func asymptoticExpansionOfNormalizedBlackCall(h, t float64) float64 {
	var e = (t / h) * (t / h)

	var r = (h + t) * (h - t)

	var q = (h / r) * (h / r)
	var asymptoticExpansionSum = 2.0 + q*(-6.0e0-2.0*e+3.0*q*(1.0e1+e*(2.0e1+2.0*e)+5.0*q*(-1.4e1+e*(-7.0e1+e*(-4.2e1-2.0*e))+7.0*q*(1.8e1+e*(1.68e2+e*(2.52e2+e*(7.2e1+2.0*e)))+9.0*q*(-2.2e1+e*(-3.3e2+e*(-9.24e2+e*(-6.6e2+e*(-1.1e2-2.0*e))))+1.1e1*q*(2.6e1+e*(5.72e2+e*(2.574e3+e*(3.432e3+e*(1.43e3+e*(1.56e2+2.0*e)))))+1.3e1*q*(-3.0e1+e*(-9.1e2+e*(-6.006e3+e*(-1.287e4+e*(-1.001e4+e*(-2.73e3+e*(-2.1e2-2.0*e))))))+1.5e1*q*(3.4e1+e*(1.36e3+e*(1.2376e4+e*(3.8896e4+e*(4.862e4+e*(2.4752e4+e*(4.76e3+e*(2.72e2+2.0*e)))))))+1.7e1*q*(-3.8e1+e*(-1.938e3+e*(-2.3256e4+e*(-1.00776e5+e*(-1.84756e5+e*(-1.51164e5+e*(-5.4264e4+e*(-7.752e3+e*(-3.42e2-2.0*e))))))))+1.9e1*q*(4.2e1+e*(2.66e3+e*(4.0698e4+e*(2.3256e5+e*(5.8786e5+e*(7.05432e5+e*(4.0698e5+e*(1.08528e5+e*(1.197e4+e*(4.2e2+2.0*e)))))))))+2.1e1*q*(-4.6e1+e*(-3.542e3+e*(-6.7298e4+e*(-4.90314e5+e*(-1.63438e6+e*(-2.704156e6+e*(-2.288132e6+e*(-9.80628e5+e*(-2.01894e5+e*(-1.771e4+e*(-5.06e2-2.0*e))))))))))+2.3e1*q*(5.0e1+e*(4.6e3+e*(1.0626e5+e*(9.614e5+e*(4.08595e6+e*(8.9148e6+e*(1.04006e7+e*(6.53752e6+e*(2.16315e6+e*(3.542e5+e*(2.53e4+e*(6.0e2+2.0*e)))))))))))+2.5e1*q*(-5.4e1+e*(-5.85e3+e*(-1.6146e5+e*(-1.77606e6+e*(-9.37365e6+e*(-2.607579e7+e*(-4.01166e7+e*(-3.476772e7+e*(-1.687257e7+e*(-4.44015e6+e*(-5.9202e5+e*(-3.51e4+e*(-7.02e2-2.0*e))))))))))))+2.7e1*q*(5.8e1+e*(7.308e3+e*(2.3751e5+e*(3.12156e6+e*(2.003001e7+e*(6.919458e7+e*(1.3572783e8+e*(1.5511752e8+e*(1.0379187e8+e*(4.006002e7+e*(8.58429e6+e*(9.5004e5+e*(4.7502e4+e*(8.12e2+2.0*e)))))))))))))+2.9e1*q*(-6.2e1+e*(-8.99e3+e*(-3.39822e5+e*(-5.25915e6+e*(-4.032015e7+e*(-1.6934463e8+e*(-4.1250615e8+e*(-6.0108039e8+e*(-5.3036505e8+e*(-2.8224105e8+e*(-8.870433e7+e*(-1.577745e7+e*(-1.472562e6+e*(-6.293e4+e*(-9.3e2-2.0*e))))))))))))))+3.1e1*q*(6.6e1+e*(1.0912e4+e*(4.74672e5+e*(8.544096e6+e*(7.71342e7+e*(3.8707344e8+e*(1.14633288e9+e*(2.07431664e9+e*(2.33360622e9+e*(1.6376184e9+e*(7.0963464e8+e*(1.8512208e8+e*(2.7768312e7+e*(2.215136e6+e*(8.184e4+e*(1.056e3+2.0*e)))))))))))))))+3.3e1*(-7.0e1+e*(-1.309e4+e*(-6.49264e5+e*(-1.344904e7+e*(-1.4121492e8+e*(-8.344518e8+e*(-2.9526756e9+e*(-6.49588632e9+e*(-9.0751353e9+e*(-8.1198579e9+e*(-4.6399188e9+e*(-1.6689036e9+e*(-3.67158792e8+e*(-4.707164e7+e*(-3.24632e6+e*(-1.0472e5+e*(-1.19e3-2.0*e)))))))))))))))))*q))))))))))))))))
	var b = OneOverSqrtTwoPi * math.Exp(-0.5*(h*h+t*t)) * (t / r) * asymptoticExpansionSum

	return math.Abs(math.Max(b, 0.0))
}

func normalisedBlackCallUsingErfcx(h, t float64) float64 {
	var b = 0.5 * math.Exp(-0.5*(h*h+t*t)) * (erfcxCody(-OneOverSqrtTwo*(h+t)) - erfcxCody(-OneOverSqrtTwo*(h-t)))
	return math.Abs(math.Max(b, 0.0))
}

func smallTExpansionOfNormalizedBlackCall(h, t float64) float64 {
	var a = 1 + h*(0.5*SqrtTwoPi)*erfcxCody(-OneOverSqrtTwo*h)
	var w = t * t
	var h2 = h * h
	var expansion = 2 * t * (a + w*((-1+3*a+a*h2)/6+w*((-7+15*a+h2*(-1+10*a+a*h2))/120+w*((-57+105*a+h2*(-18+105*a+h2*(-1+21*a+a*h2)))/5040+w*((-561+945*a+h2*(-285+1260*a+h2*(-33+378*a+h2*(-1+36*a+a*h2))))/362880+w*((-6555+10395*a+h2*(-4680+17325*a+h2*(-840+6930*a+h2*(-52+990*a+h2*(-1+55*a+a*h2)))))/39916800+((-89055+135135*a+h2*(-82845+270270*a+h2*(-20370+135135*a+h2*(-1926+25740*a+h2*(-75+2145*a+h2*(-1+78*a+a*h2))))))*w)/6227020800.0))))))
	var b = OneOverSqrtTwoPi * math.Exp(-0.5*(h*h+t*t)) * expansion

	return math.Abs(math.Max(b, 0.0))
}

func normalizedBlackCallUsingNormCdf(x, s float64) float64 {
	var h = x / s
	var t = 0.5 * s
	var bMax = math.Exp(0.5 * x)
	var b = NormCdf(h+t)*bMax - NormCdf(h-t)/bMax

	return math.Abs(math.Max(b, 0.0))
}

func normalisedBlackCallWithOptimalUseOfCodysFunctions(x, s float64) float64 {
	var codysThreshold = 0.46875
	var h = x / s
	var t = 0.5 * s
	var q1 = -OneOverSqrtTwo * (h + t)
	var q2 = -OneOverSqrtTwo * (h - t)
	var twoB float64

	if q1 < codysThreshold {
		if q2 < codysThreshold {
			twoB = math.Exp(0.5*x)*erfcxCody(q1) - math.Exp(-0.5*x)*erfcxCody(q2)
		} else {
			twoB = math.Exp(0.5*x)*erfcxCody(q1) - math.Exp(-0.5*(h*h+t*t))*erfcxCody(q2)
		}
	} else {
		if q2 < codysThreshold {
			twoB = math.Exp(-0.5*(h*h+t*t))*erfcxCody(q1) - math.Exp(-0.5*x)*erfcxCody(q2)
		} else {
			twoB = math.Exp(-0.5*(h*h+t*t)) * (erfcxCody(q1) - erfcxCody(q2))
		}
	}
	return math.Abs(math.Max(0.5*twoB, 0.0))
}

func normalisedBlackCall(x, s float64) float64 {
	if x > 0 {
		return normalisedIntrinsicCall(x) + normalisedBlackCall(-x, s)
	}
	var ax = math.Abs(x)

	if s <= ax*DenormalisationCutoff {
		return normalisedIntrinsicCall(x)
	}

	if x < s*AsymptooticExpansionAccuracyThreshold && 0.5*s*s+x < s*(SmallTExpansionOfNormalizedBlackThreshold+AsymptooticExpansionAccuracyThreshold) {
		return asymptoticExpansionOfNormalizedBlackCall(x/s, 0.5*s)
	}

	if 0.5*s < SmallTExpansionOfNormalizedBlackThreshold {
		return smallTExpansionOfNormalizedBlackCall(x/s, 0.5*s)
	}

	// if x+0.5*s*s > s*0.85 {
	//     return normalizedBlackCallUsingNormCdf(x, s)
	// }
	//
	// return normalisedBlackCallUsingErfcx(x/s, 0.5*s)

	// Post-publication updates by author are reflected commented out lines above and the update below

	return normalisedBlackCallWithOptimalUseOfCodysFunctions(x, s)
}

func square(x float64) float64 {
	return x * x
}

func NormalisedVega(x, s float64) float64 {
	var ax = math.Abs(x)
	if ax <= 0 {
		return OneOverSqrtTwoPi * math.Exp(-0.125*s*s)
	} else if s <= 0 || s <= ax*SqrtDblMin {
		return 0
	} else {
		return OneOverSqrtTwoPi * math.Exp(-0.5*(square(x/s)+square(0.5*s)))
	}
}

func normalisedBlack(x, s, q float64) float64 {
	if q < 0 {
		return normalisedBlackCall(-x, s)
	} else {
		return normalisedBlackCall(x, s)
	}
}

// Black returns the Black-Scholes value for a call (q=1) or a put (q=-1)
// given the forward price (F), strike price (K), vol (sigma) and tte (T)
func Black(F, K, sigma, T, q float64) float64 {
	var intrinsic float64
	if q < 0 {
		intrinsic = math.Abs(math.Max(K-F, 0.0))
	} else {
		intrinsic = math.Abs(math.Max(F-K, 0.0))
	}

	if q*(F-K) > 0 {
		return intrinsic + Black(F, K, sigma, T, -q)
	}

	return math.Max(intrinsic, (math.Sqrt(F)*math.Sqrt(K))*normalisedBlack(math.Log(F/K), sigma*math.Sqrt(T), q))
}

func computeFLowerMapAndFirstTwoDerivatives(x, s float64) (f, fp, fpp float64) {
	var ax = math.Abs(x)
	var z = SqrtOneOverThree * ax / s
	var y = z * z
	var s2 = s * s
	var Phi = NormCdf(-z)
	var phi = NormPdf(z)

	fpp = PiOverSix * y / (s2 * s) * Phi * (8*SqrtThree*s*ax + (3*s2*(s2-8)-8*x*x)*Phi/phi) * math.Exp(2*y+0.25*s2)

	if isBelowHorizon(s) {
		fp = 1
		f = 0
	} else {
		var Phi2 = Phi * Phi
		fp = TwoPi * y * Phi2 * math.Exp(y+0.125*s*s)
		if isBelowHorizon(x) {
			f = 0
		} else {
			f = TwoPiOverSqrtTwentySeven * ax * (Phi2 * Phi)
		}
	}

	return
}

func inverseFLowerMap(x, f float64) float64 {
	if isBelowHorizon(f) {
		return 0
	}

	return math.Abs(x / (SqrtThree * inverseNormCdf(math.Pow(f/(TwoPiOverSqrtTwentySeven*math.Abs(x)), 1./3.))))
}

func computeFUpperMapAndFirstTwoDerivatives(x, s float64) (f, fp, fpp float64) {
	f = NormCdf(-0.5 * s)

	if isBelowHorizon(x) {
		fp = -0.5
		fpp = 0
	} else {
		var w = square(x / s)
		fp = -0.5 * math.Exp(0.5*w)
		fpp = SqrtPiOverTwo * math.Exp(w+0.125*s*s) * w / s
	}
	return
}

func inverseFUpperMap(f float64) float64 {
	return -2. * inverseNormCdf(f)
}

func uncheckedNormalisedImpliedVolatilityFromATransformedRationalGuessWithLimitedIterations(beta, x, q float64, N int) float64 {
	if q*x > 0 {
		beta = math.Abs(math.Max(beta-normalisedIntrinsic(x, q), 0.0))
		q = -q
	}
	// Map puts to calls
	if q < 0 {
		x = -x
		q = -q
	}

	if beta <= 0 { // For negative or zero prices we return 0.
		return 0
	}

	if beta < DenormalisationCutoff { // For positive but denormalized (a.k.a. 'subnormal') prices, we return 0 since it would be impossible to converge to full machine accuracy anyway.
		return 0
	}
	var bMax = math.Exp(0.5 * x)

	if beta >= bMax {
		return VolatilityValueToSignalPriceIsAboveMaximum
	}
	var iterations = 0
	var directionReversalCount = 0
	var f = -math.MaxFloat64
	var s = -math.MaxFloat64
	var ds = s
	var dsPrevious float64 = 0
	var sLeft = math.SmallestNonzeroFloat64
	var sRight = math.MaxFloat64
	var sC = math.Sqrt(math.Abs(2 * x))
	var bC = normalisedBlackCall(x, sC)
	var vC = NormalisedVega(x, sC)

	if beta < bC {
		var sL = sC - bC/vC
		var bL = normalisedBlackCall(x, sL)

		if beta < bL {
			var fLowerMapL, dFLowerMapLDBeta, d2FLowerMapLDBeta2 float64
			fLowerMapL, dFLowerMapLDBeta, d2FLowerMapLDBeta2 = computeFLowerMapAndFirstTwoDerivatives(x, sL)
			var rLl = convexRationalCubicControlParameterToFitSecondDerivativeAtRightSide(0., bL, 0., fLowerMapL, 1., dFLowerMapLDBeta, d2FLowerMapLDBeta2, true)

			f = rationalCubicInterpolation(beta, 0., bL, 0., fLowerMapL, 1., dFLowerMapLDBeta, rLl)
			if !(f > 0) {
				var t = beta / bL
				f = (fLowerMapL*t + bL*(1-t)) * t
			}

			s = inverseFLowerMap(x, f)
			sRight = sL

			for ; iterations < N && math.Abs(ds) > DblEpsilon*s; iterations++ {
				if ds*dsPrevious < 0 {
					directionReversalCount++
				}

				if iterations > 0 && (3 == directionReversalCount || !(s > sLeft && s < sRight)) {
					s = 0.5 * (sLeft + sRight)
					if sRight-sLeft <= DblEpsilon*s {
						break
					}

					directionReversalCount = 0
					ds = 0
				}

				dsPrevious = ds
				var b = normalisedBlackCall(x, s)

				var bp = NormalisedVega(x, s)
				if b > beta && s < sRight {
					sRight = s
				} else if b < beta && s > sLeft {
					sLeft = s
				}
				if b <= 0 || bp <= 0 {
					ds = 0.5*(sLeft+sRight) - s
				} else {
					var lnB = math.Log(b)
					var lnBeta = math.Log(beta)
					var bpob = bp / b
					var h = x / s
					var bHalley = h*h/s - s/4
					var newton = (lnBeta - lnB) * lnB / lnBeta / bpob
					var halley = bHalley - bpob*(1+2/lnB)

					var bHh3 = bHalley*bHalley - 3*square(h/s) - 0.25
					var hh3 = bHh3 + 2*square(bpob)*(1+3/lnB*(1+1/lnB)) - 3*bHalley*bpob*(1+2/lnB)
					ds = newton * householderFactor(newton, halley, hh3)
				}

				ds = math.Max(-0.5*s, ds)
				s += ds
			}
			return s
		} else {
			var vL = NormalisedVega(x, sL)
			var rLm = convexRationalCubicControlParameterToFitSecondDerivativeAtRightSide(bL, bC, sL, sC, 1/vL, 1/vC, 0.0, false)
			s = rationalCubicInterpolation(beta, bL, bC, sL, sC, 1/vL, 1/vC, rLm)
			sLeft = sL
			sRight = sC
		}
	} else {
		var sH float64
		if vC > math.SmallestNonzeroFloat64 {
			sH = sC + (bMax-bC)/vC
		} else {
			sH = sC
		}
		var bH = normalisedBlackCall(x, sH)
		if beta <= bH {
			var vH = NormalisedVega(x, sH)
			var rHm = convexRationalCubicControlParameterToFitSecondDerivativeAtLeftSide(bC, bH, sC, sH, 1/vC, 1/vH, 0.0, false)
			s = rationalCubicInterpolation(beta, bC, bH, sC, sH, 1/vC, 1/vH, rHm)
			sLeft = sC
			sRight = sH
		} else {
			var fUpperMapH, dFUpperMapHDBeta, d2FUpperMapHDBeta2 float64
			fUpperMapH, dFUpperMapHDBeta, d2FUpperMapHDBeta2 = computeFUpperMapAndFirstTwoDerivatives(x, sH)
			if d2FUpperMapHDBeta2 > -SqrtDblMax && d2FUpperMapHDBeta2 < SqrtDblMax {
				var rHh = convexRationalCubicControlParameterToFitSecondDerivativeAtLeftSide(bH, bMax, fUpperMapH, 0., dFUpperMapHDBeta, -0.5, d2FUpperMapHDBeta2, true)
				f = rationalCubicInterpolation(beta, bH, bMax, fUpperMapH, 0., dFUpperMapHDBeta, -0.5, rHh)
			}
			if f <= 0 {
				var h = bMax - bH
				var t = (beta - bH) / h
				f = (fUpperMapH*(1-t) + 0.5*h*t) * (1 - t)
			}
			s = inverseFUpperMap(f)
			sLeft = sH
			if beta > 0.5*bMax {
				for ; iterations < N && math.Abs(ds) > DblEpsilon*s; iterations++ {
					if ds*dsPrevious < 0 {
						directionReversalCount++
					}
					if iterations > 0 && (3 == directionReversalCount || !(s > sLeft && s < sRight)) {
						s = 0.5 * (sLeft + sRight)
						if sRight-sLeft <= DblEpsilon*s {
							break
						}
						directionReversalCount = 0
						ds = 0
					}
					dsPrevious = ds
					var b = normalisedBlackCall(x, s)
					var bp = NormalisedVega(x, s)
					if b > beta && s < sRight {
						sRight = s
					} else if b < beta && s > sLeft {
						sLeft = s
					}
					if b >= bMax || bp <= math.SmallestNonzeroFloat64 {
						ds = 0.5*(sLeft+sRight) - s
					} else {
						var bMaxMinusB = bMax - b
						var g = math.Log((bMax - beta) / bMaxMinusB)
						var gp = bp / bMaxMinusB
						var bHalley = square(x/s)/s - s/4
						var bHh3 = bHalley*bHalley - 3*square(x/(s*s)) - 0.25
						var newton = -g / gp
						var halley = bHalley + gp
						var hh3 = bHh3 + gp*(2*gp+3*bHalley)
						ds = newton * householderFactor(newton, halley, hh3)
					}
					ds = math.Max(-0.5*s, ds)
					s += ds
				}
				return s
			}
		}
	}

	for ; iterations < N && math.Abs(ds) > DblEpsilon*s; iterations++ {
		if ds*dsPrevious < 0 {
			directionReversalCount++
		}
		if iterations > 0 && (3 == directionReversalCount || !(s > sLeft && s < sRight)) {
			s = 0.5 * (sLeft + sRight)
			if sRight-sLeft <= DblEpsilon*s {
				break
			}

			directionReversalCount = 0
			ds = 0
		}

		dsPrevious = ds
		var b = normalisedBlackCall(x, s)
		var bp = NormalisedVega(x, s)
		if b > beta && s < sRight {
			sRight = s
		} else if b < beta && s > sLeft {
			sLeft = s
		}
		var newton = (beta - b) / bp
		var halley = square(x/s)/s - s/4
		var hh3 = halley*halley - 3*square(x/(s*s)) - 0.25
		ds = math.Max(-0.5*s, newton*householderFactor(newton, halley, hh3))
		s += ds
	}
	return s
}

func impliedVolatilityFromATransformedRationalGuessWithLimitedIterations(price, F, K, T, q float64, N int) float64 {
	var intrinsic float64
	if q < 0 {
		intrinsic = math.Abs(math.Max(K-F, 0.0))
	} else {
		intrinsic = math.Abs(math.Max(F-K, 0.0))
	}
	if price < intrinsic {
		return VolatilityValueToSignalPriceIsBelowIntrinsic
	}
	var maxPrice float64
	if q < 0 {
		maxPrice = K
	} else {
		maxPrice = F
	}
	if price >= maxPrice {
		return VolatilityValueToSignalPriceIsAboveMaximum
	}
	var x = math.Log(F / K)
	// Map in-the-money to out-of-the-money
	if q*x > 0 {
		price = math.Abs(math.Max(price-intrinsic, 0.0))
		q = -q
	}
	return uncheckedNormalisedImpliedVolatilityFromATransformedRationalGuessWithLimitedIterations(price/(math.Sqrt(F)*math.Sqrt(K)), x, q, N) / math.Sqrt(T)
}

func ImpliedVolatilityFromATransformedRationalGuess(price, F, K, T, q float64) float64 {
	return impliedVolatilityFromATransformedRationalGuessWithLimitedIterations(price, F, K, T, q, ImpliedVolatilityMaximumIterations)
}

func normalisedImpliedVolatilityFromATransformedRationalGuessWithLimitedIterations(beta, x, q float64, N int) float64 {
	// Map in-the-money to out-of-the-money
	if q*x > 0 {
		beta -= normalisedIntrinsic(x, q)
		q = -q
	}
	if beta < 0 {
		return VolatilityValueToSignalPriceIsBelowIntrinsic
	}
	return uncheckedNormalisedImpliedVolatilityFromATransformedRationalGuessWithLimitedIterations(beta, x, q, N)
}

func NormalisedImpliedVolatilityFromATransformedRationalGuess(beta, x, q float64) float64 {
	return normalisedImpliedVolatilityFromATransformedRationalGuessWithLimitedIterations(beta, x, q, ImpliedVolatilityMaximumIterations)
}

// License for original code. This is a GoLang implementation of the original.
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
