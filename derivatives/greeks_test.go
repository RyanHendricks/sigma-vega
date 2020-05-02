package derivatives

import (
	"fmt"
	"sigma-vega/logger"
	"testing"

	"go.uber.org/zap"
)

func TestComputeGreeksPriceChange(t *testing.T) {
	options := NewTestData()
	for _, tt := range options {
		t.Run(tt.Name(), func(t *testing.T) {
			prices := GenerateIncrementalArray(tt.K, 20, 0.25)

			for _, s := range prices {
				o := NewOptionBuilder().WithUnderlying(testStock(&tt)).Strike(tt.K).Rate(tt.R).TTE(tt.T).Call()
				cp := o.Price()
				logger.Logz.Info("CALL", zap.Float64("", s), zap.Float64("price", cp))
				// logger.Pretty(o)
				o.SetOptionType(PUT)
				pp := o.Price()
				logger.Logz.Info("PUT", zap.Float64("", s), zap.Float64("price", pp))
			}
		})
	}
}

func Test_Price(t *testing.T) {
	for _, tt := range NewTestData() {
		t.Run(tt.Name(), func(t *testing.T) {
			builder := newTestOpt(&tt)

			o := builder.Call()

			if !withinRequiredPercent(o.Price(), tt.WantCallPrice, 0.01) {
				t.Errorf("callPrice() = %v, want %v", o.Price(), tt.WantCallPrice)
			}

			o.SetOptionType(PUT)

			if !withinRequiredPercent(o.Price(), tt.WantPutPrice, 0.01) {
				t.Errorf("putPrice() = %v, want %v", o.Price(), tt.WantPutPrice)
			}
		})
	}
}

func TestBSCallDelta(t *testing.T) {
	opts := NewTestData()
	fmt.Println(opts)
	for i := range opts {
		tt := opts[i]
		t.Run(tt.WantCmoney, func(t *testing.T) {
			if got := callDelta(tt.S, tt.K, tt.R, tt.T, tt.Q, tt.V); !withinRequiredPrecision(got, tt.WantCallDelta, 0.005) {
				t.Errorf("CallDelta() = %v, want %v", got, tt.WantCallDelta)
			}
		})
	}
}

func TestCallTheta(t *testing.T) {
	for _, tt := range NewTestData() {
		t.Run(tt.WantCmoney, func(t *testing.T) {
			if got := callTheta(tt.S, tt.K, tt.R, tt.T, tt.Q, tt.V); !withinRequiredPrecision(got, tt.WantCallTheta, 0.005) {
				t.Errorf("CallTheta() = %v, want %v", got, tt.WantCallTheta)
			}
		})
	}
}

func TestOGamma(t *testing.T) {
	for _, tt := range NewTestData() {
		t.Run(tt.WantCmoney, func(t *testing.T) {
			if got := gamma(tt.S, tt.K, tt.R, tt.T, tt.Q, tt.V); !withinRequiredPrecision(got, tt.WantGamma, 0.005) {
				t.Errorf("OGamma() = %v, want %v", got, tt.WantGamma)
			}
		})
	}
}

func TestPutDelta(t *testing.T) {
	for _, tt := range NewTestData() {
		t.Run(tt.WantCmoney, func(t *testing.T) {
			if got := putDelta(tt.S, tt.K, tt.R, tt.T, tt.Q, tt.V); !withinRequiredPrecision(got, tt.WantPutDelta, 0.005) {
				t.Errorf("PutDelta() = %v, want %v", got, tt.WantPutDelta)
			}
		})
	}
}

func TestPutTheta(t *testing.T) {
	for _, tt := range NewTestData() {
		t.Run(tt.WantCmoney, func(t *testing.T) {
			if got := putTheta(tt.S, tt.K, tt.R, tt.T, tt.Q, tt.V); !withinRequiredPrecision(got, tt.WantPutTheta, 0.005) {
				t.Errorf("PutTheta() = %v, want %v", got, tt.WantPutTheta)
			}
		})
	}
}

func TestVega(t *testing.T) {
	for _, tt := range NewTestData() {
		t.Run(tt.WantCmoney, func(t *testing.T) {
			if got := vega(tt.S, tt.K, tt.R, tt.T, tt.Q, tt.V); !withinRequiredPrecision(got, tt.WantVega, 0.005) {
				t.Errorf("Vega() = %v, want %v", got, tt.WantVega)
			}
		})
	}
}
