package plot

import (
	"net/http"
	"sigma-vega/client"
	"sigma-vega/derivatives"

	"github.com/gin-gonic/gin"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

type SimGreeks struct {
	Prices []float64
	Delta  []float64
	Gamma  []float64
	Vega   []float64
	Theta  []float64
}

func Plot() {
	r := gin.Default()

	r.LoadHTMLFiles("./plot/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/greeks.svg", drawGreeks)
	r.GET("/delta.svg", drawDelta)
	r.GET("/gamma.svg", drawGamma)
	r.GET("/vega.svg", drawVega)
	r.GET("/theta.svg", drawTheta)

	r.Run(":" + "3000")
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

func drawGreeks(ctx *gin.Context) {
	greeks := derivatives.SimulateGreeks()

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
				Min: -1.5,
				Max: 1.5,
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
	greeks := derivatives.SimulateGreeks()

	c := chart.Chart{
		Title: "Gamma",
		XAxis: chart.XAxis{
			Name:      "Strike",
			NameStyle: chart.StyleShow(),
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
				Name: "Gamma",
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: greeks.Prices,
				YValues: greeks.Gamma,
			},
		},
	}
	ctx.Header("Content-Type", "image/svg+xml")

	if renderError := c.Render(chart.SVG, ctx.Writer); renderError != nil {
		panic(renderError)
	}
}

func drawVega(ctx *gin.Context) {
	greeks := derivatives.SimulateGreeks()

	c := chart.Chart{
		Title: "Theta",
		XAxis: chart.XAxis{
			Name:      "Strike",
			NameStyle: chart.StyleShow(),
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
				Name: "Theta",
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

func drawTheta(ctx *gin.Context) {
	greeks := derivatives.SimulateGreeks()

	c := chart.Chart{
		Title: "Vega",
		XAxis: chart.XAxis{
			Name:      "Strike",
			NameStyle: chart.StyleShow(),
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
				Name: "Vega",
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

func drawDelta(ctx *gin.Context) {
	greeks := derivatives.SimulateGreeks()

	c := chart.Chart{
		Title: "Delta",
		XAxis: chart.XAxis{
			Name:      "Strike",
			NameStyle: chart.StyleShow(),
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
				Name: "Delta",
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: greeks.Prices,
				YValues: greeks.Delta,
			},
		},
	}
	ctx.Header("Content-Type", "image/svg+xml")

	if renderError := c.Render(chart.SVG, ctx.Writer); renderError != nil {
		panic(renderError)
	}
}
