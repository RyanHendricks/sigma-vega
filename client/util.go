package client

import (
	"fmt"
)

func AssertInt64(obj interface{}) int64 {
	res, ok := obj.(float64)
	if ok {
		return int64(res)
	}
	return int64(0)
}

func AssertFloat64(obj interface{}) float64 {
	res, ok := obj.(float64)
	if ok {
		return res
	}
	return float64(0)
}

func AssertBool(obj interface{}) bool {
	res, ok := obj.(bool)
	if ok {
		return res
	}
	return false
}

func AssertString(obj interface{}) string {
	res, ok := obj.(string)
	if ok {
		return res
	}
	return ""
}

func parseExpDateMap(obj interface{}, underlyingPrice float64, interestRate float64) []Option {
	var opts []Option
	if obj == nil {
		return opts
	}
	expDateMap := obj.(map[string]interface{})
	for _, strikePriceMap := range expDateMap {
		switch strikePriceMap.(type) {
		case interface{}:
			optionMap := strikePriceMap.(map[string]interface{})
			for _, optionEntry := range optionMap {
				option := optionEntry.([]interface{})[0].(map[string]interface{})
				var opt Option
				opt.InterestRate = AssertFloat64(interestRate)
				opt.UnderlyingPrice = AssertFloat64(underlyingPrice)
				opt.PutCall = AssertString(option["putCall"])
				opt.Symbol = AssertString(option["symbol"])
				opt.Description = AssertString(option["description"])
				opt.ExchangeName = AssertString(option["exchangeName"])
				opt.Bid = AssertFloat64(option["bid"])
				opt.Ask = AssertFloat64(option["ask"])
				opt.Last = AssertFloat64(option["last"])
				opt.Mark = AssertFloat64(option["mark"])
				opt.BidSize = AssertInt64(option["bidSize"])
				opt.AskSize = AssertInt64(option["askSize"])
				opt.BidAskSize = AssertString(option["bidAskSize"])
				opt.LastSize = AssertInt64(option["lastSize"])
				opt.HighPrice = AssertFloat64(option["highPrice"])
				opt.LowPrice = AssertFloat64(option["lowPrice"])
				opt.OpenPrice = AssertFloat64(option["openPrice"])
				opt.ClosePrice = AssertFloat64(option["closePrice"])
				opt.TotalVolume = AssertInt64(option["totalVolume"])
				opt.QuoteTimeInLong = AssertInt64(option["quoteTimeInLong"])
				opt.TradeTimeInLong = AssertInt64(option["tradeTimeInLong"])
				opt.NetChange = AssertFloat64(option["netChange"])
				opt.Volatility = AssertFloat64(option["volatility"])
				opt.Delta = AssertFloat64(option["delta"])
				opt.Gamma = AssertFloat64(option["gamma"])
				opt.Theta = AssertFloat64(option["theta"])
				opt.Vega = AssertFloat64(option["vega"])
				opt.Rho = AssertFloat64(option["rho"])
				opt.TimeValue = AssertFloat64(option["timeValue"])
				opt.OpenInterest = AssertFloat64(option["openInterest"])
				opt.TheoreticalOptionValue = AssertFloat64(option["theoreticalOptionValue"])
				opt.TheoreticalVolatility = AssertFloat64(option["theoreticalVolatility"])
				opt.StrikePrice = AssertFloat64(option["strikePrice"])
				opt.ExpirationDate = AssertInt64(option["expirationDate"])
				opt.DaysToExpiration = AssertFloat64(option["daysToExpiration"])
				opt.ExpirationType = AssertString(option["expirationType"])
				opt.LastTradingDay = AssertInt64(option["lastTradingDay"])
				opt.Multiplier = AssertFloat64(option["multiplier"])
				opt.SettlementType = AssertString(option["settlementType"])
				opt.DeliverableNote = AssertString(option["deliverableNote"])
				opt.IsIndexOption = AssertBool(option["isIndexOption"])
				opt.PercentChange = AssertFloat64(option["percentChange"])
				opt.MarkChange = AssertFloat64(option["markChange"])
				opt.MarkPercentChange = AssertFloat64(option["markPercentChange"])
				opt.InTheMoney = AssertBool(option["inTheMoney"])
				opt.Mini = AssertBool(option["mini"])
				opt.NonStandard = AssertBool(option["nonStandard"])
				opts = append(opts, opt)
			}
		default:
			fmt.Println("Expecting a JSON object, got something wrong")
		}
	}
	return opts
}
