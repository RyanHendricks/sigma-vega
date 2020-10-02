package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	urlGetOptionChain = `https://api.tdameritrade.com/v1/marketdata/chains?%v`
	timeOut           = time.Second * 10
)

func NewOptionChainRequest(apiKey, symbol, contractType, numStrikes string, quotes bool, strategy string) *OptionChainRequest {
	return &OptionChainRequest{
		APIKey:       apiKey,
		Symbol:       symbol,
		ContractType: contractType,
		NumStrikes:   numStrikes,
		Quotes:       quotes,
		Strategy:     strategy,
	}
}

func PrettyPrint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
}

func FetchOptionChain() (*OptionChain, error) {
	oc := new(OptionChain)

	params := url.Values{}
	params.Add("apikey", "CATALYSTCAP@AMER.OAUTHAPP")
	params.Add("symbol", "SPY")
	params.Add("contractType", "ALL")
	params.Add("strategy", "SINGLE")
	params.Add("includeQuotes", "TRUE")
	// params.Add("strike", "110")
	// params.Add("strikeCount", "50")
	// params.Add("range", "30")
	// params.Add("expMonth", "APR")
	request, err := http.NewRequest("GET", fmt.Sprintf(urlGetOptionChain, params.Encode()), nil)
	if err != nil {
		return nil, err
	}

	td := http.Client{
		Timeout: timeOut,
	}

	res, doErr := td.Do(request)
	if doErr != nil {
		return nil, doErr
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			panic(errClose)
		}
	}()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	jsonErr := json.Unmarshal(body, &oc)
	if jsonErr != nil {
		return nil, jsonErr
	}

	oc.CallMap = parseExpDateMap(oc.CallExpDateMap, oc.Underlying.Last, oc.InterestRate)
	oc.PutMap = parseExpDateMap(oc.PutExpDateMap, oc.Underlying.Last, oc.InterestRate)

	return oc, nil
}
