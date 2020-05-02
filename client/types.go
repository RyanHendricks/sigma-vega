package client

import (
	"math"
	"time"
)

// "apikey=CATALYSTCAP" +
//     "&symbol=SPY" +
//     "&contractType=ALL" +
//     "&strikeCount=2" +
//     "&includeQuotes=TRUE" +
//     "&strategy=SINGLE" +
//     "&strike=250" +
//     "&volatility=0.50" +
//     "&underlyingPrice=250" +
//     "&interestRate=0.005" +
//     "&daysToExpiration=30" +
//     "&expMonth=ALL" +
//     "&optionType=ATM"

type OptionChainRequest struct {
	ApiKey       string `url:"apikey,omitempty"`
	Symbol       string `url:"symbol,omitempty"`
	ContractType string `url:"contractType,omitempty"`
	NumStrikes   string `url:"strikeCount,omitempty"`
	Quotes       bool   `url:"includeQuotes,omitempty"`
	Strategy     string `url:"strategy,omitempty"`
}

type PositionsResponse []struct {
	SecuritiesAccount struct {
		Type                    string `json:"type"`
		AccountID               string `json:"accountId"`
		RoundTrips              int    `json:"roundTrips"`
		IsDayTrader             bool   `json:"isDayTrader"`
		IsClosingOnlyRestricted bool   `json:"isClosingOnlyRestricted"`
		Positions               []struct {
			ShortQuantity                  int     `json:"shortQuantity"`
			AveragePrice                   float64 `json:"averagePrice"`
			CurrentDayProfitLoss           int     `json:"currentDayProfitLoss"`
			CurrentDayProfitLossPercentage float64 `json:"currentDayProfitLossPercentage"`
			LongQuantity                   int     `json:"longQuantity"`
			SettledLongQuantity            int     `json:"settledLongQuantity"`
			SettledShortQuantity           int     `json:"settledShortQuantity"`
			Instrument                     struct {
				AssetType        string `json:"assetType"`
				Cusip            string `json:"cusip"`
				Symbol           string `json:"symbol"`
				Description      string `json:"description"`
				PutCall          string `json:"putCall"`
				UnderlyingSymbol string `json:"underlyingSymbol"`
			} `json:"instrument"`
			MarketValue            int `json:"marketValue"`
			MaintenanceRequirement int `json:"maintenanceRequirement"`
		} `json:"positions"`
		InitialBalances struct {
			AccruedInterest                  int     `json:"accruedInterest"`
			AvailableFundsNonMarginableTrade float64 `json:"availableFundsNonMarginableTrade"`
			BondValue                        int     `json:"bondValue"`
			BuyingPower                      float64 `json:"buyingPower"`
			CashBalance                      float64 `json:"cashBalance"`
			CashAvailableForTrading          int     `json:"cashAvailableForTrading"`
			CashReceipts                     int     `json:"cashReceipts"`
			DayTradingBuyingPower            int     `json:"dayTradingBuyingPower"`
			DayTradingBuyingPowerCall        int     `json:"dayTradingBuyingPowerCall"`
			DayTradingEquityCall             int     `json:"dayTradingEquityCall"`
			Equity                           float64 `json:"equity"`
			EquityPercentage                 int     `json:"equityPercentage"`
			LiquidationValue                 float64 `json:"liquidationValue"`
			LongMarginValue                  int     `json:"longMarginValue"`
			LongOptionMarketValue            int     `json:"longOptionMarketValue"`
			LongStockValue                   int     `json:"longStockValue"`
			MaintenanceCall                  int     `json:"maintenanceCall"`
			MaintenanceRequirement           int     `json:"maintenanceRequirement"`
			Margin                           float64 `json:"margin"`
			MarginEquity                     float64 `json:"marginEquity"`
			MoneyMarketFund                  int     `json:"moneyMarketFund"`
			MutualFundValue                  int     `json:"mutualFundValue"`
			RegTCall                         int     `json:"regTCall"`
			ShortMarginValue                 int     `json:"shortMarginValue"`
			ShortOptionMarketValue           int     `json:"shortOptionMarketValue"`
			ShortStockValue                  int     `json:"shortStockValue"`
			TotalCash                        float64 `json:"totalCash"`
			IsInCall                         bool    `json:"isInCall"`
			PendingDeposits                  int     `json:"pendingDeposits"`
			MarginBalance                    int     `json:"marginBalance"`
			ShortBalance                     int     `json:"shortBalance"`
			AccountValue                     float64 `json:"accountValue"`
		} `json:"initialBalances"`
		CurrentBalances struct {
			AccruedInterest                  int     `json:"accruedInterest"`
			CashBalance                      float64 `json:"cashBalance"`
			CashReceipts                     int     `json:"cashReceipts"`
			LongOptionMarketValue            int     `json:"longOptionMarketValue"`
			LiquidationValue                 float64 `json:"liquidationValue"`
			LongMarketValue                  int     `json:"longMarketValue"`
			MoneyMarketFund                  int     `json:"moneyMarketFund"`
			Savings                          int     `json:"savings"`
			ShortMarketValue                 int     `json:"shortMarketValue"`
			PendingDeposits                  int     `json:"pendingDeposits"`
			AvailableFunds                   float64 `json:"availableFunds"`
			AvailableFundsNonMarginableTrade float64 `json:"availableFundsNonMarginableTrade"`
			BuyingPower                      float64 `json:"buyingPower"`
			BuyingPowerNonMarginableTrade    float64 `json:"buyingPowerNonMarginableTrade"`
			DayTradingBuyingPower            int     `json:"dayTradingBuyingPower"`
			Equity                           float64 `json:"equity"`
			EquityPercentage                 int     `json:"equityPercentage"`
			LongMarginValue                  int     `json:"longMarginValue"`
			MaintenanceCall                  int     `json:"maintenanceCall"`
			MaintenanceRequirement           int     `json:"maintenanceRequirement"`
			MarginBalance                    int     `json:"marginBalance"`
			RegTCall                         int     `json:"regTCall"`
			ShortBalance                     int     `json:"shortBalance"`
			ShortMarginValue                 int     `json:"shortMarginValue"`
			ShortOptionMarketValue           int     `json:"shortOptionMarketValue"`
			Sma                              float64 `json:"sma"`
			MutualFundValue                  int     `json:"mutualFundValue"`
			BondValue                        int     `json:"bondValue"`
		} `json:"currentBalances"`
		ProjectedBalances struct {
			AvailableFunds                   float64 `json:"availableFunds"`
			AvailableFundsNonMarginableTrade float64 `json:"availableFundsNonMarginableTrade"`
			BuyingPower                      float64 `json:"buyingPower"`
			DayTradingBuyingPower            int     `json:"dayTradingBuyingPower"`
			DayTradingBuyingPowerCall        int     `json:"dayTradingBuyingPowerCall"`
			MaintenanceCall                  int     `json:"maintenanceCall"`
			RegTCall                         int     `json:"regTCall"`
			IsInCall                         bool    `json:"isInCall"`
			StockBuyingPower                 float64 `json:"stockBuyingPower"`
		} `json:"projectedBalances"`
	} `json:"securitiesAccount,omitempty"`
}

type TxResponse []TransactionsResponse

type TransactionsResponse []struct {
	Type                  string  `json:"type"`
	SubAccount            string  `json:"subAccount"`
	SettlementDate        string  `json:"settlementDate"`
	OrderID               string  `json:"orderId,omitempty"`
	NetAmount             float64 `json:"netAmount"`
	TransactionDate       string  `json:"transactionDate"`
	OrderDate             string  `json:"orderDate,omitempty"`
	TransactionSubType    string  `json:"transactionSubType"`
	TransactionID         int64   `json:"transactionId"`
	CashBalanceEffectFlag bool    `json:"cashBalanceEffectFlag"`
	Description           string  `json:"description"`
	Fees                  struct {
		RFee          float64 `json:"rFee"`
		AdditionalFee float64 `json:"additionalFee"`
		CdscFee       float64 `json:"cdscFee"`
		RegFee        float64 `json:"regFee"`
		OtherCharges  float64 `json:"otherCharges"`
		Commission    float64 `json:"commission"`
		OptRegFee     float64 `json:"optRegFee"`
		SecFee        float64 `json:"secFee"`
	} `json:"fees"`
	TransactionItem struct {
		AccountID      int     `json:"accountId"`
		Amount         float64 `json:"amount"`
		Price          float64 `json:"price"`
		Cost           float64 `json:"cost"`
		Instruction    string  `json:"instruction"`
		PositionEffect string  `json:"positionEffect"`
		Instrument     struct {
			Symbol               string `json:"symbol"`
			UnderlyingSymbol     string `json:"underlyingSymbol"`
			OptionExpirationDate string `json:"optionExpirationDate"`
			PutCall              string `json:"putCall"`
			Cusip                string `json:"cusip"`
			Description          string `json:"description"`
			AssetType            string `json:"assetType"`
		} `json:"instrument"`
	} `json:"transactionItem,omitempty"`
}

type OptionChain struct {
	Symbol     string     `json:"symbol"`
	Status     string     `json:"status"`
	Underlying Underlying `json:"underlying"`
	// enum[SINGLE, ANALYTICAL, COVERED, VERTICAL, CALENDAR, STRANGLE,
	//     STRADDLE, BUTTERFLY, CONDOR, DIAGONAL, COLLAR, ROLL]
	Strategy          string      `json:"strategy"`
	Interval          float64     `json:"interval"`
	IsDelayed         bool        `json:"isDelayed"`
	IsIndex           bool        `json:"isIndex"`
	DaysToExpiration  float64     `json:"daysToExpiration"`
	InterestRate      float64     `json:"interestRate"`
	UnderlyingPrice   float64     `json:"underlyingPrice"`
	Volatility        float64     `json:"volatility"`
	NumberOfContracts int64       `json:"numberOfContracts"`
	CallExpDateMap    interface{} `json:"callExpDateMap"`
	PutExpDateMap     interface{} `json:"putExpDateMap"`
	CallMap           []Option
	PutMap            []Option
}

type StrikePriceMap struct {
}

type Option struct {
	InterestRate           float64              `json:"interestRate"`
	UnderlyingPrice        float64              `json:"underlyingPrice"`
	PutCall                string               `json:"putCall"` // enum[PUT, CALL]
	Symbol                 string               `json:"symbol"`
	Description            string               `json:"description"`
	ExchangeName           string               `json:"exchangeName"`
	Bid                    float64              `json:"bid"`
	Ask                    float64              `json:"ask"`
	Last                   float64              `json:"last"`
	Mark                   float64              `json:"mark"`
	BidSize                int64                `json:"bidSize"`
	AskSize                int64                `json:"askSize"`
	BidAskSize             string               `json:"bidAskSize"`
	LastSize               int64                `json:"lastSize"`
	HighPrice              float64              `json:"highPrice"`
	LowPrice               float64              `json:"lowPrice"`
	OpenPrice              float64              `json:"openPrice"`
	ClosePrice             float64              `json:"closePrice"`
	TotalVolume            int64                `json:"totalVolume"`
	QuoteTimeInLong        int64                `json:"quoteTimeInLong"`
	TradeTimeInLong        int64                `json:"tradeTimeInLong"`
	NetChange              float64              `json:"netChange"`
	Volatility             float64              `json:"volatility"`
	Delta                  float64              `json:"delta"`
	Gamma                  float64              `json:"gamma"`
	Theta                  float64              `json:"theta"`
	Vega                   float64              `json:"vega"`
	Rho                    float64              `json:"rho"`
	TimeValue              float64              `json:"timeValue"`
	OpenInterest           float64              `json:"openInterest"`
	TheoreticalOptionValue float64              `json:"theoreticalOptionValue"`
	TheoreticalVolatility  float64              `json:"theoreticalVolatility"`
	OptionDeliverablesList []OptionDeliverables `json:"optionDeliverablesList"`
	StrikePrice            float64              `json:"strikePrice"`
	ExpirationDate         int64                `json:"expirationDate"`
	DaysToExpiration       float64              `json:"daysToExpiration"`
	ExpirationType         string               `json:"expirationType"`
	LastTradingDay         int64                `json:"lastTradingDay"`
	Multiplier             float64              `json:"multiplier"`
	SettlementType         string               `json:"settlementType"`
	DeliverableNote        string               `json:"deliverableNote"`
	IsIndexOption          bool                 `json:"isIndexOption"`
	PercentChange          float64              `json:"percentChange"`
	MarkChange             float64              `json:"markChange"`
	MarkPercentChange      float64              `json:"markPercentChange"`
	InTheMoney             bool                 `json:"inTheMoney"`
	Mini                   bool                 `json:"mini"`
	NonStandard            bool                 `json:"nonStandard"`
}

// S returns the current price of the underlying as a decimal
func (o Option) S() float64 {
	return o.UnderlyingPrice
}

// K returns the strike price of the option as a decimal
func (o Option) K() float64 {
	return o.StrikePrice
}

// T returns the year fraction time to expiration as a decimal
func (o Option) T() float64 {
	return o.DaysToExpiration / 365
}

// TTE returns the time to expiration as a year fraction with millisecond accuracy
func (o Option) TTE() float64 {
	return ((float64(o.ExpirationDate) / 1000.0) - float64(time.Now().Unix())) / 60 / 60 / 24 / 365
}

// V returns the implied volatility of the option
func (o Option) V() float64 {
	return o.Volatility / 100
}

// R returns the risk free interest rate as a decimal
func (o Option) R() float64 {
	return o.InterestRate
}

// Q returns the dividend for the underlying asset of the option
func (o Option) Q() float64 {
	return 0
}

func (o Option) GetOptionType() string {
	if o.PutCall == "CALL" {
		return "call"
	}

	return "put"
}

type OptionDeliverables struct {
	Symbol           string `json:"symbol"`
	AssetType        string `json:"assetType"`
	DeliverableUnits string `json:"deliverableUnits"`
	CurrencyType     string `json:"currencyType"`
}

type Underlying struct {
	Ask               float64 `json:"ask"`
	AskSize           int64   `json:"askSize"`
	Bid               float64 `json:"bid"`
	BidSize           int64   `json:"bidSize"`
	Change            float64 `json:"change"`
	Close             float64 `json:"close"`
	Delayed           bool    `json:"delayed"`
	Description       string  `json:"description"`
	ExchangeName      string  `json:"exchangeName"` // enum[IND, ASE, NYS, NAS, NAP, PAC, OPR, BATS]
	FiftyTwoWeekHigh  float64 `json:"fiftyTwoWeekHigh"`
	FiftyTwoWeekLow   float64 `json:"fiftyTwoWeekLow"`
	HighPrice         float64 `json:"highPrice"`
	Last              float64 `json:"last"`
	LowPrice          float64 `json:"lowPrice"`
	Mark              float64 `json:"mark"`
	MarkChange        float64 `json:"markChange"`
	MarkPercentChange float64 `json:"markPercentChange"`
	OpenPrice         float64 `json:"openPrice"`
	PercentChange     float64 `json:"percentChange"`
	QuoteTime         int64   `json:"quoteTime"`
	Symbol            string  `json:"symbol"`
	TotalVolume       int64   `json:"totalVolume"`
	TradeTime         int64   `json:"tradeTime"`
}

func (u Underlying) Price() float64 {
	if (u.Ask - u.Bid) >= math.Abs(u.Last-u.Mark) {
		return u.Last
	}
	return u.Mark
}

func (u Underlying) Yield() float64 {
	return 0
}

type ExpirationDate struct {
	Date string `json:"date"`
}

type Deliverable struct{}

type StrikeQuote struct {
	Ask                    float32
	AskSize                int32
	Bid                    float32
	BidSize                int32
	ClosePrice             float32
	DaysToExpiration       int
	Delta                  float32
	ExchangeName           string
	ExpirationDate         int64
	ExpirationType         string
	Gamma                  float32
	HighPrice              float32
	InTheMoney             bool
	IsIndexOption          bool
	Last                   float32
	LastSize               int32
	LastTradingDay         int64
	LowPrice               float32
	Mark                   float32
	MarkChange             float32
	MarkPercentChange      float32
	Mini                   bool
	Multiplier             float32
	NetChange              float32
	NonStandard            bool
	OpenInterest           float32
	OpenPrice              float32
	PercentChange          float32
	PutCall                string
	QuoteTimeInLong        int64
	Rho                    float32
	StrikePrice            float32
	Symbol                 string
	TheoreticalOptionValue float32
	TheoreticalVolatility  float32
	Theta                  float32
	TimeValue              float32
	TotalVolume            int32
	TradeTimeInLong        int64
	Vega                   float32
	Volatility             float32
}

type OptionChainResponse struct {
	CallExpDateMap   map[string]map[string][]Option
	DaysToExpiration float32
	InterestRate     float32
	Interval         float32
	IsDelayed        bool
	IsIndex          bool
	PutExpDateMap    map[string]map[string][]Option
	Status           string
	Strategy         string
	Symbol           string
	UnderlyingPrice  float32
	Volatility       float32
}
