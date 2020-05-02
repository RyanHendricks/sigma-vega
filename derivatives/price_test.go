package derivatives

import (
	"testing"
)

func Test_forwardPrice(t *testing.T) {
	tests := []struct {
		name      string
		spotPrice float64
		tte       float64
		intRate   float64
		want      float64
		q         float64
	}{
		{"forward price 1", 300, 1, 0.05, 315.38132891280725, 0.000},
		{"forward price 2", 100, 1, 0.05, 105.12710963760242, 0.000},
		{"forward price 3", 25, 0.5, 0.1, 26.281777409400604, 0.000},
		{"forward with q (Hull 5.3)", 25, 0.5, 0.1, 25.766516136769308, 0.0396},
		{"forward with q (Hull 5.3)", 25, 0.5, 0.1, 25.766516136769308, 0.0396},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := forwardPrice(tt.spotPrice, tt.tte, tt.intRate, tt.q); got != tt.want {
				t.Errorf("forwardPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_callPayoff(t *testing.T) {
	tests := []struct {
		name string
		s    float64
		k    float64
		want float64
	}{
		{"call payoff", 100, 75, 25},
		{"call payoff", 100, 125, 0},
		{"call payoff", 100, 100, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CallPayoff(tt.s, tt.k); got != tt.want {
				t.Errorf("CallPayoff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_putPayoff(t *testing.T) {
	tests := []struct {
		name string
		s    float64
		k    float64
		want float64
	}{
		{"put payoff", 100, 75, 0},
		{"put payoff", 100, 125, 25},
		{"put payoff", 100, 100, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PutPayoff(tt.s, tt.k); got != tt.want {
				t.Errorf("PutPayoff() = %v, want %v", got, tt.want)
			}
		})
	}
}

//
// func TestFetchandPriceOptionChain(t *testing.T) {
//     tests := []struct {
//         name    string
//         want    *client.OptionChain
//         wantErr bool
//     }{
//         {"TestFetchOptionChain", &client.OptionChain{}, false},
//     }
//     for _, tt := range tests {
//         t.Run(tt.name, func(t *testing.T) {
//             got, err := client.FetchOptionChain()
//             if (err != nil) != tt.wantErr {
//                 t.Errorf("FetchOptionChain() error = %v, wantErr %v", err, tt.wantErr)
//                 return
//             }
//
//             for _, call := range got.CallMap {
//                 c := NewOptionBuilder().WithUnderlying(&Stock{
//                     Vol:      call.Volatility / 100,
//                     Price:    call.S(),
//                     Dividend: 0.02,
//                 }).Strike(call.K()).Rate(0.0025).TTE(call.TTE()).Call()
//
//                 iv := c.ImpliedVolatility(call.Mark)
//
//                 impliedvolatility := lbr.ImpliedVolatilityFromATransformedRationalGuess(call.Mark, c.Underlying.F0(c.T()), c.K(), c.T(), 1)
//                 if impliedvolatility != iv {
//                     t.Errorf("iv() = %v, want %v", iv, impliedvolatility)
//                 }
//                 if c.Price() != c.Price() {
//                     t.Errorf("c.BlackScholes(iv)() = %v, mark %v", c.Price(), call.Mark)
//                 }
//
//                 // g := (call.S(), call.K(), 0.0025, call.TTE(), 0.02, iv)
//                 g := c.GreekValues()
//                 client.PrettyPrint(g)
//
//                 fmt.Println(call.S(), call.K(), 0.0025, call.TTE(), 0.02, impliedvolatility)
//                 today := callTheta(call.S(), call.K(), 0.0025, call.TTE(), 0.02, impliedvolatility)
//                 yesterday := callTheta(call.S(), call.K(), 0.0025, call.TTE(), 0.02, impliedvolatility)
//                 fmt.Println("THETA:", (today-yesterday)*call.TTE())
//                 // p := NewCall(call.S(), call.K(), 0.0025, call.TTE(), 0.02, call.Volatility/100)
//                 //
//                 price := c.Price()
//                 //
//                 fmt.Println(call.Delta, call.Theta, call.Gamma, call.Vega)
//
//                 // client.PrettyPrint(g)
//             }
//             // for _, put := range got.PutMap {
//             //     impliedvolatility := lbr.ImpliedVolatilityFromATransformedRationalGuess(put.Mark, put.S(), put.K(), put.TTE(), -1)
//             //     g := PutGreeks(put.S(), put.K(), 0.0025, put.TTE(), 0.02, impliedvolatility)
//             //     p := NewPut(put.S(), put.K(), 0.0025, put.TTE(), 0.02, put.Volatility/100)
//             //
//             //     price := p.Price(put.Volatility / 100)
//             //
//             //     fmt.Println(put.Delta, put.Theta, put.Gamma, put.Vega)
//             //
//             //     client.PrettyPrint(g)
//             // }
//
//         })
//     }
// }
//
// func TestPriceFromIVandIVfromPriceOptionChain(t *testing.T) {
//     tests := []struct {
//         name    string
//         wantErr bool
//     }{
//         {"TestPriceFromIVandIVfromPriceOptionChain", false},
//     }
//     for _, tt := range tests {
//         t.Run(tt.name, func(t *testing.T) {
//
//
//             got, err := client.FetchOptionChain()
//             if (err != nil) != tt.wantErr {
//                 t.Errorf("FetchOptionChain() error = %v, wantErr %v", err, tt.wantErr)
//                 return
//             }
//
//             for _, call := range got.CallMap {
//                 t.Run(tt.name, func(t *testing.T) {
//
//                     c := NewOptionBuilder().WithUnderlying(&Stock{
//                         Price:    call.S(),
//                         Vol:      call.V(),
//                         Dividend: 0.02,
//                     }).Strike(call.K()).Rate(call.R()).TTE(call.TTE()).Call()
//
//                     pricebs := c.Price()
//                     if !util.withinRequiredPrecision(pricebs, call.Mark) {
//                         t.Errorf("pricebs() = %v, want %v", pricebs, call.Mark)
//
//                     }
//
//                 })
//             }
//             for _, put := range got.PutMap {
//                 t.Run(tt.name, func(t *testing.T) {
//
//                     p := NewOptionBuilder().WithUnderlying(&Stock{
//                         Price:    put.S(),
//                         Vol:      put.V(),
//                         Dividend: 0.02,
//                     }).Strike(put.K()).Rate(put.R()).TTE(put.TTE()).Put()
//                     impliedvolatility := p.ImpliedVolatility(put.ClosePrice)
//                     PrintOption(p)
//                     pricebs := p.Price()
//                     if !util.withinRequiredPrecision(pricebs, put.Mark) {
//                         t.Errorf("pricebs() = %v, want %v", pricebs, put.V())
//
//                         client.PrettyPrint(put)
//                     }
//                     fmt.Println(impliedvolatility, pricebs, put.Mark)
//                 })
//             }
//             //
//         })
//     }
//     //
// }
