package lbr

import "math"

const (
	// TWO_PI represents 2 * pi
	TwoPi                    float64 = 6.283185307179586476925286766559005768394338798750               // 2π
	SqrtPiOverTwo            float64 = 1.253314137315500251207882642405522626503493370305               // √(π/2) to avoid misinterpretation.
	SqrtThree                float64 = 1.732050807568877293527446341505872366942805253810               // √(3)
	SqrtOneOverThree         float64 = 0.577350269189625764509148780501957455647601751270               // √(1/3)
	TwoPiOverSqrtTwentySeven float64 = 1.209199576156145233729385505094770488189377498728               // 2π/√(27)
	PiOverSix                float64 = 0.523598775598298873077107230546583814032861566563               // π/6
	OneOverSqrtTwo           float64 = 0.7071067811865475244008443621048490392848359376887              // 1/√(2)
	OneOverSqrtTwoPi         float64 = 0.3989422804014326779399460599343818684758586311649              // 1/√(2π)
	SqrtTwoPi                float64 = 2.506628274631000502415765284811045253006986740610               // √(2π)
	Exp                      float64 = 2.71828182845904523536028747135266249775724709369995957496696763 // https://oeis.org/A001113

	DenormalisationCutoff                        float64 = 0
	VolatilityValueToSignalPriceIsBelowIntrinsic float64 = -math.MaxFloat64
	VolatilityValueToSignalPriceIsAboveMaximum   float64 = math.MaxFloat64
	ImpliedVolatilityMaximumIterations                   = 2
	AsymptooticExpansionAccuracyThreshold        float64 = -10
	NormCdfAsymptoticExpansionFirstThreshold     float64 = -10.0
)

var (
	DblEpsilon              = math.Nextafter(1, 2) - 1
	SqrtDblEpsilon          = math.Sqrt(DblEpsilon)
	FourthRootDbtEpsilon    = math.Sqrt(SqrtDblEpsilon)
	EighthRootDbtEpsilon    = math.Sqrt(FourthRootDbtEpsilon)
	SixteenthRootDbtEpsilon = math.Sqrt(EighthRootDbtEpsilon)
	SqrtDblMin              = math.Sqrt(math.SmallestNonzeroFloat64)
	SqrtDblMax              = math.Sqrt(math.MaxFloat64)

	NormCdfAsymptoticExpansionSecondThreshold = -1 / math.Sqrt(DblEpsilon)
	SmallTExpansionOfNormalizedBlackThreshold = 2 * SixteenthRootDbtEpsilon
)
