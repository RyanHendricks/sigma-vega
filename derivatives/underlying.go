package derivatives

// Underlying is an interface representing the underlying asset of a derivative
type Underlying interface {
	S0() float64        // The current price
	Q() float64         // Dividend yield or 0 if none
	V() float64         // Volatility of returns
	F0(float64) float64 // The forward price of the underlying at time t
}

// Stock is a financial instrument with a price, historical volatility, and expected return
// type Stock is a struct that implements the Underlying interface
type Stock struct {
	Vol      float64 // historical volatility of returns for Stock
	Price    float64 // current price of the Stock
	Dividend float64 // Dividend yield or 0 if none
	RFrate   float64 // Risk-free interest rate
}

// S0 returns the current price of the asset
func (s Stock) S0() float64 {
	return s.Price
}

// F0 is the Forward/Futures price of the Stock and is calculated
// using equation F0 = S0e^(r-q)t if the the Stock pays a dividend
// or using equation F0 = S0e^rt. t is the time to maturity.
func (s Stock) F0(t float64) float64 {
	return forwardPrice(s.Price, t, s.RFrate, s.Dividend)
}

// V returns the volatility of the stock
func (s Stock) V() float64 {
	return s.Vol
}

// Q returns the dividend yield of the stock
func (s Stock) Q() float64 {
	return s.Dividend
}

// SetPrice sets the current price of the stock
func (s *Stock) SetPrice(price float64) {
	s.Price = price
}

// DeltaS returns the difference between a given future price p
// and the current price of the stock.
func (s Stock) DeltaS(p float64) float64 {
	return p - s.Price
}
