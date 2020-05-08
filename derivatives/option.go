package derivatives

import (
	"math"
	"sigma-vega/lbr"
	"sigma-vega/logger"

	"go.uber.org/zap"
)

const (
	DAYS_PER_YEAR float64 = 365.25
	// DAYS_PER_YEAR float64 = 252 //

	ONE_DAY_IN_YEARS = 1 / DAYS_PER_YEAR

	CALL = "call"
	PUT  = "put"
)

// There are six factors affecting the price of a stock option:
// 1. The current stock price, s
// 2. The strike price, k
// 3. The time to expiration, t
// 4. The volatility of the stock price, v
// 5. The risk-free interest rate, r
// 6. The dividends that are expected to be paid, q.

// s: Current underlying price
// r: Continuously compounded risk-free rate of interest for an investment maturing in time t

type Contract struct {
	k float64 // strike price
	t float64 // time to expiration in years
	r float64 // risk free rate (continuously compounded)

	ContractType string // call or put

	Underlying Underlying
}

func PrintOption(o Option) {
	logger.Logz.Info("",
		zap.String("type", o.GetOptionType()),
		zap.Float64("s", o.S()),
		zap.Float64("k", o.K()),
		zap.Float64("t", o.T()),
		zap.Float64("r", o.R()),
		zap.Float64("sigma", o.V()),
		zap.Float64("q", o.Q()))
}

type Option interface {
	GetOptionType() string
	S() float64
	K() float64
	T() float64
	Q() float64
	V() float64
	R() float64
}

// S returns the current price of the underlying as a decimal
func (o Contract) S() float64 {
	return o.Underlying.S0()
}

// K returns the strike price of the option as a decimal
func (o Contract) K() float64 {
	return o.k
}

// T returns the time to expiration in years and returns the
// smallest non-zero float64 if T is zero.
func (o Contract) T() float64 {
	return math.Max(math.SmallestNonzeroFloat64, o.t)
}

// R returns the risk free interest rate as a decimal
func (o Contract) R() float64 {
	return o.r
}

// V returns the price volatility of the underlying
func (o Contract) V() float64 {
	return o.Underlying.V()
}

// Q returns the dividend for the underlying asset of the option
func (o Contract) Q() float64 {
	return o.Underlying.Q()
}

// F0 returns the forward price of the Underlying asset using the
// time to expiration of the option contract
func (o Contract) F0() float64 {
	return o.Underlying.F0(o.T())
}

// GetOptionType returns the option type either "call" or "put"
func (o Contract) GetOptionType() string {
	return o.ContractType
}

// SetOptionType assigns the CALL or PUT designation to an option
func (o *Contract) SetOptionType(optionType string) {
	if optionType != CALL && optionType != PUT {
		panic("invalid option type")
	}
	o.ContractType = optionType
}

// ImpliedVolatility returns the IV for an option given a price as
// computed using the methodology from Peter Jaeckel's Let's Be Rational
func (o Contract) ImpliedVolatility(price float64) float64 {
	var q float64
	if o.GetOptionType() == CALL {
		q = 1
	} else {
		q = -1
	}

	return lbr.ImpliedVolatilityFromATransformedRationalGuess(price, o.F0(), o.K(), o.T(), q)
}

func (o Contract) GreekValues() *Greeks {
	g := &Greeks{}

	g.Gamma = gamma(o.S(), o.K(), o.R(), o.T(), o.Q(), o.V())
	g.Vega = vega(o.S(), o.K(), o.R(), o.T(), o.Q(), o.V())

	if o.GetOptionType() == CALL {
		g.Delta = callDelta(o.S(), o.K(), o.R(), o.T(), o.Q(), o.V())
		g.Theta = callTheta(o.S(), o.K(), o.R(), o.T(), o.Q(), o.V())
	} else {
		g.Delta = putDelta(o.S(), o.K(), o.R(), o.T(), o.Q(), o.V())
		g.Theta = putTheta(o.S(), o.K(), o.R(), o.T(), o.Q(), o.V())
	}

	return g
}

func (o Contract) Intrinsic() float64 {
	if o.GetOptionType() == CALL {
		return CallPayoff(o.S(), o.K())
	}

	return PutPayoff(o.S(), o.K())
}

// Price returns the price of the option computed using Let's Be Rational
func (o Contract) Price() float64 {
	if o.GetOptionType() == CALL {
		return lbr.Black(o.F0(), o.K(), o.V(), o.T(), 1)
	}

	if o.GetOptionType() == PUT {
		return lbr.Black(o.F0(), o.K(), o.V(), o.T(), -1)
	}

	return 0
}
