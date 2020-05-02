package derivatives

import "fmt"

type Simulation struct {
	Prices    []float64
	SimGreeks []Greeks
}

func calcGreeks() *Simulation {
	prices := GenerateIncrementalArray(1000, 100000, 0.0001)

	var simGreeks []Greeks
	sim := new(Simulation)

	for z := range prices {
		o := NewOptionBuilder().WithUnderlying(Stock{
			Vol:      0.5,
			Price:    prices[z],
			Dividend: 0,
		}).Strike(100).Rate(0.001).TTE(0.5).Put()

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
	}
	sim.SimGreeks = simGreeks
	//simGreeks.Prices = prices
	fmt.Println(sim)
	return sim
}
