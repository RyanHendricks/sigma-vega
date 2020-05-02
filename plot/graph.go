package plot

import (
	"fmt"
	"net/http"
	"sigma-vega/client"
	"sigma-vega/derivatives"
	"sigma-vega/lbr"

	"github.com/gin-gonic/gin"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
	"github.com/wcharczuk/go-chart/seq"
)

func Plot() {
	r := gin.Default()
	r.LoadHTMLFiles(htmltop, htmlform, htmlbasic, htmlbot)

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, htmltop, gin.H{})

		//fmt.Println(ctx.PostForm("strike"))
		//fmt.Println(ctx.PostForm("optiontype"))
	})

	//r.GET("/chain.svg", func(ctx *gin.Context) {
	//	drawOptionChain(greeks)
	//	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{})
	//
	//	//fmt.Println(ctx.PostForm("strike"))
	//	//fmt.Println(ctx.PostForm("optiontype"))
	//})

	r.GET("/call-payoff.svg", drawCallPayoff)
	r.GET("/put-payoff.svg", drawPutPayoff)
	r.GET("/greeks.svg", drawGreeks)
	r.GET("/delta.svg", drawDelta)
	r.GET("/gamma.svg", drawGamma)
	r.GET("/vega.svg", drawVega)
	r.GET("/theta.svg", drawTheta)

	r.Run(":" + "3000")
}

func simValues() SimGreeks {
	prices := derivatives.GenerateIncrementalArray(100, 20, 5)
	return simulateGreeks(prices)
}

func drawGoChart(res http.ResponseWriter, req *http.Request) {
	greeks := simValues()

	prices := derivatives.GenerateIncrementalArray(100, 50, 1)

	deltaValues := make([]float64, len(prices))

	for z, price := range prices {
		g := derivatives.NewOptionBuilder().WithUnderlying(&derivatives.Stock{
			Price:    price,
			Vol:      0.4,
			Dividend: 0,
		}).Strike(100).Rate(0.001).TTE(0.5).Call().GreekValues().Delta

		deltaValues[z] = g
	}

	g := chart.Chart{
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: greeks.Prices,
				YValues: deltaValues,
			},
		},
	}

	// mainSeries := chart.ContinuousSeries{
	//     Name:    "A test series",
	//     XValues: greeks.Prices,
	//     YValues: deltaValues,
	//     // seq.Range(1.0, 100.0),             //generates a []float64 from 1.0 to 100.0 in 1.0 step increments, or 100 elements.
	//     // YValues: seq.RandomValuesWithMax(100, 100), //generates a []float64 randomly from 0 to 100 with 100 elements.
	// }
	//
	// g := chart.Chart{
	//     Width: 1920,
	//     Series: []chart.Series{
	//         mainSeries,
	//         chart.ContinuousSeries{
	//             XValues: []float64{1.0, 2.0, 3.0, 4.0},
	//             YValues: []float64{1.0, 2.0, 3.0, 4.0},
	//         },
	//     },
	// }

	if renderError := g.Render(chart.SVG, res); renderError != nil {
		panic(renderError)
	}
}

func genTicks(vals []float64) chart.Ticks {
	ticks := make([]chart.Tick, len(vals))
	for i, val := range vals {
		t := chart.Tick{
			Value: val,
			Label: fmt.Sprintf("%v", val),
		}
		ticks[i] = t
	}
	return ticks
}

func padding() chart.Box {
	return chart.Box{
		IsSet:  true,
		Top:    10,
		Left:   30,
		Right:  30,
		Bottom: 10,
	}
}

// Return keys of the given map
func Keys(m map[float64][]client.Option) (keys []float64) {
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

var colors = []string{"2689cc", "227bb8", "1e6ea3", "1b608f", "17527a", "134566", "0f3752", "0b293d", "081b29", "040e14", "2689cc", "3c95d1", "51a1d6", "67acdb", "7db8e0", "93c4e6", "a8d0eb", "bedcf0", "d4e7f5", "e9f3fa", "2689cc", "227bb8", "1e6ea3", "1b608f", "17527a", "134566", "0f3752", "0b293d", "081b29", "040e14", "2689cc", "3c95d1", "51a1d6", "67acdb", "7db8e0", "93c4e6", "a8d0eb", "bedcf0", "d4e7f5", "e9f3fa", "2689cc", "227bb8", "1e6ea3", "1b608f", "17527a", "134566", "0f3752", "0b293d", "081b29", "040e14", "2689cc", "3c95d1", "51a1d6", "67acdb", "7db8e0", "93c4e6", "a8d0eb", "bedcf0", "d4e7f5", "e9f3fa"}

func drawOptionChain(ctx *gin.Context, chain *client.OptionChain) {
	var expirations = make(map[float64][]client.Option)

	for _, option := range chain.PutMap {
		expirations[option.DaysToExpiration] = append(expirations[option.DaysToExpiration], option)
	}

	chartSeries := make([]chart.Series, len(expirations))

	keys := Keys(expirations)

	fmt.Println(keys)

	for s, dte := range keys {
		xvals := make([]float64, len(expirations[dte]))
		yvals := make([]float64, len(expirations[dte]))

		for i, option := range expirations[dte] {
			xvals[i] = option.StrikePrice
			impliedvolatility := lbr.ImpliedVolatilityFromATransformedRationalGuess(option.Mark, option.S(), option.K(), option.TTE(), -1)

			fmt.Println(option.Mark, option.S(), option.K(), option.TTE(), impliedvolatility)
			// fmt.Print(impliedvolatility)
			yvals[i] = impliedvolatility
		}

		series := chart.ContinuousSeries{
			Name: fmt.Sprintf("%v", dte),
			Style: chart.Style{
				Show:        true,
				StrokeColor: drawing.ColorFromHex(colors[s]),
			},
			XValues: xvals,
			YValues: yvals,
		}

		chartSeries[s] = series
	}

	c := chart.Chart{
		Title: "Contract Chain",
		XAxis: chart.XAxis{
			Name:      "Strike",
			NameStyle: chart.StyleShow(),
			Style: chart.Style{
				Show: true,
			},
		},
		Height: 1080,
		Width:  1920,
		Canvas: chart.Style{
			Padding: padding(),
		},
		Background: chart.Style{
			Padding: padding(),
		},
		YAxis: chart.YAxis{
			Name:      "Value",
			NameStyle: chart.StyleShow(),
			AxisType:  0,
			Style: chart.Style{
				Show: true,
			},
			Range: &chart.ContinuousRange{
				Min: 0.2,
				Max: 1.2,
			},
		},

		Series: chartSeries,
	}
	// l := chart.Style{
	//     Show:        true,
	//     Padding: chart.Box{
	//         Top:    100,
	//         Left:   100,
	//         Right:  100,
	//         Bottom: 100,
	//     },
	// }
	//
	//
	// c.Elements = []chart.Renderable{
	//     chart.Legend(&c, l),
	// }
	ctx.Header("Content-Type", "image/svg+xml")

	if renderError := c.Render(chart.SVG, ctx.Writer); renderError != nil {
		panic(renderError)
	}
}

func drawPutPayoff(ctx *gin.Context) {
	xvals := seq.Range(1.0, 100.0)
	yvals := make([]float64, len(xvals))
	for i := range xvals {
		yvals[i] = derivatives.PutPayoff(xvals[i], 50)
		fmt.Println(xvals[i], yvals[i])
	}

	c := chart.Chart{
		Width: 1920, // this overrides the default.
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xvals,
				YValues: yvals,
			},
		},
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
		},
	}
	ctx.Header("Content-Type", "image/svg+xml")
	if renderError := c.Render(chart.SVG, ctx.Writer); renderError != nil {
		panic(renderError)
	}
}

func drawCallPayoff(ctx *gin.Context) {
	xvals := seq.Range(1.0, 100.0)
	yvals := make([]float64, len(xvals))
	for i := range xvals {
		if 50 >= xvals[i] {
			yvals[i] = 0
		}
		yvals[i] = xvals[i] - 50
	}

	c := chart.Chart{
		Width: 640, // this overrides the default.
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xvals,
				YValues: yvals,
			},
		},
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
			Range: &chart.ContinuousRange{
				Min: -10,
				Max: 100,
			},
		},
	}
	ctx.Header("Content-Type", "image/svg+xml")

	if renderError := c.Render(chart.SVG, ctx.Writer); renderError != nil {
		panic(renderError)
	}
}

func drawGreeks(ctx *gin.Context) {
	greeks := simValues()

	c := chart.Chart{
		Title: "Greeks",
		XAxis: chart.XAxis{
			Name:      "Strike",
			NameStyle: chart.StyleShow(),
			Style: chart.Style{
				Show: true,
			},
		},
		Height: 500,
		Width:  900,
		Canvas: chart.Style{
			Padding: padding(),
		},
		Background: chart.Style{
			Padding: padding(),
		},
		YAxis: chart.YAxis{
			Name:      "Value",
			NameStyle: chart.StyleShow(),
			AxisType:  0,
			Style: chart.Style{
				Show: true,
			},
			Range: &chart.ContinuousRange{
				Min: -1.0,
				Max: 1.0,
			},
		},

		Series: []chart.Series{

			chart.ContinuousSeries{
				Name: "Delta",
				Style: chart.Style{
					Show:        true,
					StrokeColor: drawing.ColorBlue,
				},
				XValues: greeks.Prices,
				YValues: greeks.Delta,
			},
			chart.ContinuousSeries{
				Name: "Gamma",
				Style: chart.Style{
					Show:        true,
					StrokeColor: drawing.ColorRed,
				},
				XValues: greeks.Prices,
				YValues: greeks.Gamma,
			},
			chart.ContinuousSeries{
				Name: "Vega",
				Style: chart.Style{
					Show:        true,
					StrokeColor: drawing.ColorGreen,
				},
				XValues: greeks.Prices,
				YValues: greeks.Vega,
			},
			chart.ContinuousSeries{
				Name: "Theta",
				Style: chart.Style{
					Show:        true,
					StrokeColor: drawing.ColorBlack,
				},
				XValues: greeks.Prices,
				YValues: greeks.Theta,
			},
		},
	}
	l := chart.Style{
		Show: true,
		Padding: chart.Box{
			Top:    100,
			Left:   100,
			Right:  100,
			Bottom: 100,
		},
	}

	c.Elements = []chart.Renderable{
		chart.Legend(&c, l),
	}
	ctx.Header("Content-Type", "image/svg+xml")

	if renderError := c.Render(chart.SVG, ctx.Writer); renderError != nil {
		panic(renderError)
	}
}

func drawGamma(ctx *gin.Context) {
	greeks := simValues()
	c := chart.Chart{
		Title: "Gamma",
		XAxis: chart.XAxis{
			Name:      "Strike",
			NameStyle: chart.StyleShow(),
			Style: chart.Style{
				Show: true,
			},
		},
		Height: 500,
		Width:  900,
		Canvas: chart.Style{
			Padding: padding(),
		},
		Background: chart.Style{
			Padding: padding(),
		},
		YAxis: chart.YAxis{
			Name:      "Value",
			NameStyle: chart.StyleShow(),
			AxisType:  0,
			Style: chart.Style{
				Show: true,
			},
			Range: &chart.ContinuousRange{
				Min: -1.0,
				Max: 1.0,
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValueFormatter: chart.FloatValueFormatter,

				XValues: greeks.Prices,
				YValues: greeks.Gamma,
			},
		},
	}

	l := chart.Style{
		Show:        true,
		FillColor:   drawing.ColorBlack,
		FontColor:   chart.DefaultTextColor,
		FontSize:    8.0,
		StrokeColor: chart.DefaultAxisColor,
		StrokeWidth: 5,
		Padding:     padding(),
	}

	c.Elements = []chart.Renderable{chart.Legend(&c, l)}
	ctx.Header("Content-Type", "image/svg+xml")

	if renderError := c.Render(chart.SVG, ctx.Writer); renderError != nil {
		panic(renderError)
	}
}

func drawVega(ctx *gin.Context) {
	greeks := simValues()

	c := chart.Chart{
		Title: "Vega",

		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: greeks.Prices,
				YValues: greeks.Vega,
			},
		},
	}
	ctx.Header("Content-Type", "image/svg+xml")

	if renderError := c.Render(chart.SVG, ctx.Writer); renderError != nil {
		panic(renderError)
	}
}

func drawTheta(ctx *gin.Context) {
	greeks := simValues()

	c := chart.Chart{
		Title: "Vega",

		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: greeks.Prices,
				YValues: greeks.Theta,
			},
		},
	}
	ctx.Header("Content-Type", "image/svg+xml")

	if renderError := c.Render(chart.SVG, ctx.Writer); renderError != nil {
		panic(renderError)
	}
}

func drawDelta(ctx *gin.Context) {
	greeks := simValues()

	c := chart.Chart{
		Title: "Vega",

		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: greeks.Prices,
				YValues: greeks.Theta,
			},
		},
	}
	ctx.Header("Content-Type", "image/svg+xml")

	if renderError := c.Render(chart.SVG, ctx.Writer); renderError != nil {
		panic(renderError)
	}
}
