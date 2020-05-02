package derivatives

import (
	"reflect"
	"testing"
)

func TestNewOptionBuilder(t *testing.T) {
	tests := []struct {
		name string
		want *OptionBuilder
	}{
		{"New Contract Builder", &OptionBuilder{&Contract{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOptionBuilder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOptionBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptionBuilder_Call(t *testing.T) {
	tests := []struct {
		name string
		want *Contract
	}{
		{"Contract builder Call()", &Contract{ContractType: CALL}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &OptionBuilder{
				option: &Contract{},
			}
			if got := b.Call(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Call() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptionBuilder_Put(t *testing.T) {
	tests := []struct {
		name string
		want *Contract
	}{
		{"Contract builder Put()", &Contract{ContractType: PUT}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &OptionBuilder{
				option: &Contract{},
			}
			if got := b.Put(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Put() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NewOptionBuilder(t *testing.T) {
	tests := []struct {
		name string
		r    float64
		k    float64
		t    float64
		u    Stock
	}{
		{
			"Contract builder test",
			0.05,
			100,
			0.2,
			Stock{
				Vol:      0.5,
				Price:    95.25,
				Dividend: 0.00,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			option := NewOptionBuilder().WithUnderlying(&tt.u).Strike(tt.k).TTE(tt.t).Rate(tt.r).Call()
			if option.S() != tt.u.S0() {
				t.Errorf("S() = %v, want %v", option.S(), tt.u.S0())
			}
			if option.R() != tt.r {
				t.Errorf("R() = %v, want %v", option.R(), tt.r)
			}
			if option.V() != tt.u.V() {
				t.Errorf("V() = %v, want %v", option.V(), tt.u.V())
			}
			if option.T() != tt.t {
				t.Errorf("T() = %v, want %v", option.T(), tt.t)
			}
			if option.Q() != tt.u.Q() {
				t.Errorf("Q() = %v, want %v", option.Q(), tt.u.Q())
			}
			if option.K() != tt.k {
				t.Errorf("K() = %v, want %v", option.K(), tt.k)
			}
		})
	}
}

func Test_newOption(t *testing.T) {
	for _, tt := range NewTestData() {
		t.Run(tt.Name(), func(t *testing.T) {
			got := NewOptionBuilder().WithUnderlying(&Stock{
				Vol:      tt.V,
				Price:    tt.S,
				Dividend: 0,
			}).Strike(tt.K).Rate(tt.R).TTE(tt.T).Call()
			if got.S() != tt.S {
				t.Errorf("UnderlyingPrice() = %v, want = %v", got.S(), tt.S)
			}
			if got.K() != tt.K {
				t.Errorf("K() = %v, want = %v", got.K(), tt.K)
			}
			if got.T() != tt.T {
				t.Errorf("T() = %v, want = %v", got.T(), tt.T)
			}
			if got.R() != tt.R {
				t.Errorf("IntRate() = %v, want = %v", got.R(), tt.R)
			}
			if got.V() != tt.V {
				t.Errorf("IntRate() = %v, want = %v", got.V(), tt.V)
			}
		})
	}
}

func BenchmarkName_SigmaVega(b *testing.B) {
	var st = Stock{}

	// Which has a price and historical volatility
	st.Price = 100
	st.Vol = 0.15

	// We want the price of a call option with strike 105 expiring in 6 months
	o := NewOptionBuilder().WithUnderlying(st).Strike(105).Rate(0.0025).TTE(0.5).Call()

	b.StartTimer()
	for i := 0; i < 100000000; i++ {
		p := o.Price()
		o.ImpliedVolatility(p)
	}
	b.StopTimer()
}
