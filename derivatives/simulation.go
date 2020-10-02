package derivatives

import (
	"fmt"
	"sigma-vega/logger"

	"go.uber.org/zap"
)

type Simulation struct {
	Strikes   []float64
	SimGreeks []Greeks
}

func calcGreeks() *Simulation {
	strikes := GenerateIncrementalArray(335, 50, 1)

	var simGreeks []Greeks
	var simPrices []float64
	sim := new(Simulation)

	for z := range strikes {
		k := strikes[z]
		o := NewOptionBuilder().WithUnderlying(Stock{
			Vol:      0.238,
			Price:    335,
			Dividend: 0.016,
		}).Strike(k).Rate(0.001).TTE(0.994).Put()

		d := o.GreekValues().Delta
		g := o.GreekValues().Gamma
		v := o.GreekValues().Vega
		t := o.GreekValues().Theta

		gr := Greeks{
			Delta: d,
			Theta: t,
			Gamma: g,
			Vega:  v,
		}

		simGreeks = append(simGreeks, gr)
		simPrices = append(simPrices, k)

		PrintOption(o)
	}

	sim.SimGreeks = simGreeks
	sim.Strikes = simPrices
	for g := range sim.SimGreeks {
		logger.Logz.Info("underlying", zap.Float64("price", sim.Strikes[g]))
		sim.SimGreeks[g].String()
	}
	// fmt.Println(sim)
	return sim
}

type SimGreeks struct {
	Prices []float64
	Delta  []float64
	Gamma  []float64
	Vega   []float64
	Theta  []float64
}

func SimulateGreeks() SimGreeks {
	prices := GenerateIncrementalArray(100, 20, 5)
	var simGreeks SimGreeks

	deltaValues := make([]float64, len(prices))
	gammaValues := make([]float64, len(prices))
	vegaValues := make([]float64, len(prices))
	thetaValues := make([]float64, len(prices))

	for z, price := range prices {
		o := NewOptionBuilder().WithUnderlying(Stock{
			Vol:      0.5,
			Price:    100,
			Dividend: 0,
		}).Strike(price).Rate(0.001).TTE(0.5).Put()

		d := o.GreekValues().Delta
		g := o.GreekValues().Gamma
		v := o.GreekValues().Vega
		t := o.GreekValues().Theta

		thetaValues[z] = t
		deltaValues[z] = d
		gammaValues[z] = g
		vegaValues[z] = v
	}
	simGreeks.Delta = deltaValues
	simGreeks.Gamma = gammaValues
	simGreeks.Vega = vegaValues
	simGreeks.Theta = thetaValues
	simGreeks.Prices = prices

	fmt.Println(simGreeks)

	return simGreeks
}
