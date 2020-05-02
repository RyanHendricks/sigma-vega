package derivatives

import (
	"testing"
)

func TestStock_SetPrice(t *testing.T) {
	tests := []struct {
		name     string
		price    float64
		newPrice float64
		want     float64
	}{
		{"TestStock_SetPrice1", 300, 100, 100},
		{"TestStock_SetPrice2", 18, 20, 20},
		{"TestStock_SetPrice3", 175, 150, 150},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stock{
				Price: tt.price,
			}
			s.SetPrice(tt.newPrice)
			if got := s.S0(); got != tt.want {
				t.Errorf("Price() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStock_ForwardPrice(t *testing.T) {
	tests := []struct {
		name  string
		stock Stock
		t     float64
		want  float64
	}{
		{"TestStock_ForwardPrice", Stock{
			Price:    40,
			Dividend: 0.02,
			RFrate:   0.05,
		}, .25, 40.50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.stock.F0(tt.t)
			if withinRequiredPrecision(got, tt.want, 0.005) {
				t.Errorf("F0() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStockStruct(t *testing.T) {
	type args struct {
		ExpRet   float64
		Vol      float64
		Price    float64
		Dividend float64
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test Stock",
			args{
				Vol:      0.4,
				Price:    100,
				Dividend: 0.02,
			},
		},
		{"Test Stock no dividend",
			args{
				Vol:   0.4,
				Price: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stock{
				// ExpRet:   tt.args.ExpRet,
				Vol:      tt.args.Vol,
				Price:    tt.args.Price,
				Dividend: tt.args.Dividend,
			}
			if got := s.S0(); got != tt.args.Price {
				t.Errorf("S() = %v, want %v", got, tt.args.Price)
			}
			if got := s.V(); got != tt.args.Vol {
				t.Errorf("V() = %v, want %v", got, tt.args.Vol)
			}
			if got := s.Q(); got != tt.args.Dividend {
				t.Errorf("Q() = %v, want %v", got, tt.args.Dividend)
			}
			if got := s.DeltaS(tt.args.Price + 10); got != 10 {
				t.Errorf("Q() = %v, want %v", got, 10)
			}
		})
	}
}
