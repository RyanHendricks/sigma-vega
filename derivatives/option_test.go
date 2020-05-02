package derivatives

import (
	"fmt"
	"runtime"
	"sigma-vega/logger"
	"sync"
	"testing"
	"time"

	"go.uber.org/zap"
)

func Test_OptionGreekValues(t *testing.T) {
	for _, tt := range NewTestData() {
		t.Run(fmt.Sprintf("%v %v", tt.WantCmoney, tt.WantPmoney), func(t *testing.T) {
			builder := newTestOpt(&tt)
			o := builder.Call()

			got := o.GreekValues()

			if !withinRequiredPrecision(got.Vega, tt.WantVega, 0.005) {
				t.Errorf("vega() = %v, want %v", got.Vega, tt.WantVega)
			}
			if !withinRequiredPrecision(got.Gamma, tt.WantGamma, 0.005) {
				t.Errorf("gamma() = %v, want %v", got.Gamma, tt.WantGamma)
			}
			if !withinRequiredPrecision(got.Theta, tt.WantCallTheta, 0.005) {
				t.Errorf("callTheta() = %v, want %v", got.Theta, tt.WantCallTheta)
			}
			if !withinRequiredPrecision(got.Delta, tt.WantCallDelta, 0.005) {
				t.Errorf("callDelta() = %v, want %v", got.Delta, tt.WantCallDelta)
			}

			o.SetOptionType(PUT)

			got = o.GreekValues()

			if !withinRequiredPrecision(got.Delta, tt.WantPutDelta, 0.005) {
				t.Errorf("putDelta() = %v, want %v", got.Delta, tt.WantPutDelta)
			}
			if !withinRequiredPrecision(got.Theta, tt.WantPutTheta, 0.005) {
				t.Errorf("putTheta() = %v, want %v", got.Theta, tt.WantPutTheta)
			}
		})
	}
}

func Benchmark_OptionPrice(b *testing.B) {
	opts := NewTestData()
	for i := range opts {
		call := newTestOpt(&opts[i]).Call()
		put := newTestOpt(&opts[i]).Put()
		b.Run(opts[i].Name(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				call.Price()
				put.Price()
			}
		})
	}
}

func Benchmark_OptionGreekValues(b *testing.B) {
	opts := NewTestData()
	for i := range opts {
		call := newTestOpt(&opts[i]).Call()
		put := newTestOpt(&opts[i]).Put()
		b.Run(opts[i].Name(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				call.GreekValues()
				put.GreekValues()
			}
		})
	}
}

func BenchmarkPricePool(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	jobs := make(chan func(), 72)
	done := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		var count int
		for z := 0; z <= 36; z++ {
			<-done
			count++
		}
		wg.Done()
	}()

	for i := 0; i <= 36; i++ {
		go func() {
			for {
				job, more := <-jobs
				if !more {
					done <- struct{}{}
					return
				}
				job()
			}
		}()
	}

	opts := NewTestData()
	for i := range opts {
		start := time.Now()
		o := newTestOpt(&opts[i]).Put()
		for i := 0; i < 1000000; i++ {
			fn := func() { o.Price() }
			jobs <- fn
		}
		logger.Logz.Info("total time", zap.String("", time.Since(start).String()))
	}

	logger.Logz.Info("closing channels")

	close(jobs)

	wg.Wait()
	close(done)
}

func TestOption_GetOptionType(t *testing.T) {
	for _, tt := range NewTestData() {
		t.Run(tt.Name(), func(t *testing.T) {
			o := &Contract{}
			o.SetOptionType(CALL)
			if got := o.GetOptionType(); got != CALL {
				t.Errorf("GetOptionType() = %v, want %v", got, CALL)
			}
			o.SetOptionType(PUT)
			if got := o.GetOptionType(); got != PUT {
				t.Errorf("GetOptionType() = %v, want %v", got, PUT)
			}
		})
	}
}
