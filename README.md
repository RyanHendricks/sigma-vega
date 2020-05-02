# SigmaVega


```go
    // Let's assume there is some Stock 'st'
    var st = Stock{}

    // Which has a price and historical volatility
    st.Price = 100
    st.Vol = 0.15

    // We want the price of a call option with strike 105 expiring in 6 months
    o := NewOptionBuilder().WithUnderlying(st).Strike(105).Rate(0.0025).TTE(0.5).Call()
    callPrice := o.Price()
    fmt.Println(callPrice)

    // How about the price of the put with same strike and expiration?
    o.SetOptionType(PUT) // Equivalent to -> NewOptionBuilder().WithUnderlying(st).Strike(105).Rate(0.0025).TTE(0.5).Put()
    putPrice := o.Price() 
    fmt.Println(putPrice)

    // Suppose we know the price but want the implied volatility?
    fmt.Println(o.ImpliedVolatility(7.285379466082457))

    // We also want to know the Greek values
    greeks := o.GreekValues()
    greeks.String()
    // {"delta":-0.6536449585063238,"gamma":0.03478744464149236,"theta":-0.010209263637407541,"vega":0.2609058348111927}


```