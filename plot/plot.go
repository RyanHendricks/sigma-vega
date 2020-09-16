package plot

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sigma-vega/derivatives"
	"strconv"
)

type SimGreeks struct {
	Prices []float64
	Delta  []float64
	Gamma  []float64
	Vega   []float64
	Theta  []float64
}

func simulateGreeks(prices []float64) SimGreeks {
	var simGreeks SimGreeks

	deltaValues := make([]float64, len(prices))
	gammaValues := make([]float64, len(prices))
	vegaValues := make([]float64, len(prices))
	thetaValues := make([]float64, len(prices))

	for z, price := range prices {
		o := derivatives.NewOptionBuilder().WithUnderlying(derivatives.Stock{
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
	return simGreeks
}

func Plots() {
	listen := flag.String("listen", ":3000", "Address for HTTP listener")
	flag.Parse()

	// http.HandleFunc("/price.svg", func(w http.ResponseWriter, r *http.Request) {
	//     w.Header().Set("Content-Type", "image/svg+xml")
	//     prices := util.GenerateIncrementalArray(100, 50, 1)
	//     greeks := simulateGreeks(prices)
	//     drawPrice(w, r, prices, greeks)
	// })
	//options, err := client.FetchOptionChain()
	//if err != nil {
	//	panic(err)
	//}
	//http.HandleFunc("/chain.svg", func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Content-Type", "image/svg+xml")
	//	drawOptionChain(w, r, options)
	//})
	//
	//http.HandleFunc("/call-payoff.svg", func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Content-Type", "image/svg+xml")
	//	drawCallPayoff(w, r)
	//})
	//http.HandleFunc("/put-payoff.svg", func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Content-Type", "image/svg+xml")
	//	drawPutPayoff(w, r)
	//})
	//
	//http.HandleFunc("/greeks.svg", func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Content-Type", "image/svg+xml")
	//	prices := derivatives.GenerateIncrementalArray(100, 180, 1)
	//	greeks := simulateGreeks(prices)
	//	drawGreeks(w, r, prices, greeks)
	//})
	//
	//http.HandleFunc("/delta.svg", func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Content-Type", "image/svg+xml")
	//	prices := derivatives.GenerateIncrementalArray(100, 20, 5)
	//	greeks := simulateGreeks(prices)
	//	drawTheta(w, r, prices, greeks)
	//})
	//
	//http.HandleFunc("/gamma.svg", func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Content-Type", "image/svg+xml")
	//	prices := derivatives.GenerateIncrementalArray(100, 20, 5)
	//	greeks := simulateGreeks(prices)
	//	drawGamma(w, r, prices, greeks)
	//})
	//
	//http.HandleFunc("/vega.svg", func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Content-Type", "image/svg+xml")
	//	prices := derivatives.GenerateIncrementalArray(100, 20, 5)
	//	greeks := simulateGreeks(prices)
	//	drawVega(w, r, prices, greeks)
	//})
	//
	//http.HandleFunc("/theta.svg", func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Content-Type", "image/svg+xml")
	//	prices := derivatives.GenerateIncrementalArray(100, 20, 5)
	//	greeks := simulateGreeks(prices)
	//	drawTheta(w, r, prices, greeks)
	//})

	//tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.FormValue("strike"))
		fmt.Println(r.FormValue("optiontype"))
		fmt.Fprint(w, htmltop)

		fmt.Fprint(w, htmlform)
		fmt.Fprint(w, htmlbasic)

		fmt.Fprint(w, htmlbot)
		//if r.Method != http.MethodPost {
		//	tmpl.Execute(w, nil)
		//	return
		//}
		//
		//details := ContactDetails{
		//	Email:   r.FormValue("email"),
		//	Subject: r.FormValue("subject"),
		//	Message: r.FormValue("message"),
		//}
		//
		//// do something with details
		//_ = details
		//
		//tmpl.Execute(w, struct{ Success bool }{true})
	})

	//http.HandleFunc("/", html)

	log.Printf("Listening on %s", *listen)
	log.Fatal(http.ListenAndServe(*listen, nil))
}

//http.HandleFunc("/chartthemed.png", func(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "image/png")
//	drawChartThemed(w, r, png.New(), r.FormValue("theme"), r.FormValue("scheme"))
//})
//
//http.HandleFunc("/", html)
//
//log.Printf("Listening on %s", *listen)
//log.Fatal(http.ListenAndServe(*listen, nil))
//}
//
func fFormValue(req *http.Request, name string) float64 {
	x := req.FormValue(name)
	r, _ := strconv.ParseFloat(x, 64)
	return r
}

//
func iFormValue(req *http.Request, name string) int64 {
	x := req.FormValue(name)
	r, _ := strconv.ParseInt(x, 10, 64)
	return r
}

//
//func html(w http.ResponseWriter, r *http.Request) {
//	p := r.FormValue("p")
//	fmt.Fprint(w, htmltop)
//	switch p {
//	case "basic":
//		fmt.Fprint(w, htmlbasic)
//	case "colors":
//		fmt.Fprint(w, htmlcolors)
//	case "random":
//		fmt.Fprint(w, htmlrandom)
//	}
//	fmt.Fprint(w, htmlbot)
//}

const htmlform = `
<h1>Parameters</h1>
<form action="/" method="POST" novalidate>
  <div>
    <p><label>Strike</label></p>
    <p><input type="number" name="strike"></p>
  </div>
  <div>
    <p><label>Call or Put</label></p>
    <p><input type="radio" name="optiontype"></textarea></p>
  </div>
  <div>
    <input type="submit" value="Update">
  </div>
</form>
`

func html(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, htmltop)

	fmt.Fprint(w, htmlform)
	fmt.Fprint(w, htmlbasic)

	fmt.Fprint(w, htmlbot)
}

const htmltop = `
<html>
<body>
<h1>Chart Examples</h1>
`

const htmlbasic = `
<div class="row">
    <div class="col">
        <h2>Call Payoff</h2>
        <img src="/call-payoff.svg"></img>
    </div>
    <div class="col">
        <h2>Put Payoff</h2>
        <img src="/put-payoff.svg"></img>
    </div>
</div>
<div>
    <h2>Greeks</h2>
    <img src="/greeks.svg"></img>
</div>
<div>
    <h2>Delta</h2>
    <img src="/delta.svg"></img>
</div>
<div>
    <h2>Gamma</h2>
    <img src="/gamma.svg"></img>
</div>
<div>
    <h2>Vega</h2>
    <img src="/vega.svg"></img>
</div>
<div>
    <h2>Theta</h2>
    <img src="/theta.svg"></img>
</div>
<div>
    <h2>Chain</h2>
    <img src="/chain.svg"></img>
</div>
`

const htmlbot = `
</body>
</html>
`
