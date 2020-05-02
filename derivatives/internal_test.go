package derivatives

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sigma-vega/lbr"
	"testing"
)

type Opt struct {
	S             float64 `json:"s"`
	K             float64 `json:"k"`
	T             float64 `json:"t"`
	R             float64 `json:"r"`
	V             float64 `json:"v"`
	Q             float64 `json:"q"`
	WantCallPrice float64 `json:"want_CallPrice"`
	WantPutPrice  float64 `json:"want_PutPrice"`
	WantCallDelta float64 `json:"want_CallDelta"`
	WantPutDelta  float64 `json:"want_PutDelta"`
	WantGamma     float64 `json:"want_Gamma"`
	WantVega      float64 `json:"want_Vega"`
	WantCallTheta float64 `json:"want_CallTheta"`
	WantPutTheta  float64 `json:"want_PutTheta"`
	WantD1        float64 `json:"want_d1"`
	WantD2        float64 `json:"want_d2"`
	WantNd1       float64 `json:"want_Nd1"`
	WantNd2       float64 `json:"want_Nd2"`
	WantCmoney    string  `json:"want_Cmoney"`
	WantPmoney    string  `json:"want_Pmoney"`
}

func testStock(o *Opt) Stock {
	return Stock{
		Vol:      o.V,
		Price:    o.S,
		Dividend: o.Q,
		RFrate:   o.R,
	}
}

func newTestOpt(o *Opt) *OptionBuilder {
	return NewOptionBuilder().WithUnderlying(testStock(o)).Strike(o.K).Rate(o.R).TTE(o.T)
}

func (o *Opt) Name() string {
	return fmt.Sprintf("%v %v", o.WantCmoney, o.WantPmoney)
}

func NewTestData() []Opt {
	keysBody := []byte(`[
 {
   "s": 90,
   "k": 100,
   "t": 0.5,
   "r": 0.0025,
   "v": 0.2,
   "q": 0.02,
   "want_CallPrice": 1.581998472,
   "want_PutPrice": 12.35259153,
   "want_CallDelta": 0.230812858,
   "want_PutDelta": -0.769187142,
   "want_Gamma": 0.023903947,
   "want_Vega": 0.193621971,
   "want_CallTheta": -0.010739454,
   "want_PutTheta": -0.010055378,
   "want_d1": -0.736172516,
   "want_d2": -0.877593872,
   "want_Nd1": 0.304247574,
   "want_Nd2": 0.190082078,
   "want_Cmoney": "OTM",
   "want_Pmoney": "ITM"
 },
 {
   "s": 95,
   "k": 100,
   "t": 0.5,
   "r": 0.0025,
   "v": 0.2,
   "q": 0.02,
   "want_CallPrice": 3.040331813,
   "want_PutPrice": 8.860675699,
   "want_CallDelta": 0.361722079,
   "want_PutDelta": -0.638277921,
   "want_Gamma": 0.027892088,
   "want_Vega": 0.251726097,
   "want_CallTheta": -0.014005412,
   "want_PutTheta": -0.013321336,
   "want_d1": -0.353859528,
   "want_d2": -0.495280884,
   "want_Nd1": 0.37473101,
   "want_Nd2": 0.310200931,
   "want_Cmoney": "OTM",
   "want_Pmoney": "ITM"
 },
 {
   "s": 99,
   "k": 100,
   "t": 0.5,
   "r": 0.0025,
   "v": 0.2,
   "q": 0.02,
   "want_CallPrice": 4.696783706,
   "want_PutPrice": 6.556928257,
   "want_CallDelta": 0.475190723,
   "want_PutDelta": -0.524809277,
   "want_Gamma": 0.028439307,
   "want_Vega": 0.278733652,
   "want_CallTheta": -0.015559919,
   "want_PutTheta": -0.014875843,
   "want_d1": -0.062227772,
   "want_d2": -0.203649128,
   "want_Nd1": 0.398170616,
   "want_Nd2": 0.41931385,
   "want_Cmoney": "OTM",
   "want_Pmoney": "ITM"
 },
 {
   "s": 100,
   "k": 100,
   "t": 0.5,
   "r": 0.0025,
   "v": 0.2,
   "q": 0.02,
   "want_CallPrice": 5.181292003,
   "want_PutPrice": 6.05138672,
   "want_CallDelta": 0.503526139,
   "want_PutDelta": -0.496473861,
   "want_Gamma": 0.028208377,
   "want_Vega": 0.282083773,
   "want_CallTheta": -0.015762606,
   "want_PutTheta": -0.01507853,
   "want_d1": 0.008838835,
   "want_d2": -0.132582521,
   "want_Nd1": 0.398926697,
   "want_Nd2": 0.447261778,
   "want_Cmoney": "ATM",
   "want_Pmoney": "ATM"
 },
 {
   "s": 101,
   "k": 100,
   "t": 0.5,
   "r": 0.0025,
   "v": 0.2,
   "q": 0.02,
   "want_CallPrice": 5.693716876,
   "want_PutPrice": 5.57376176,
   "want_CallDelta": 0.531562551,
   "want_PutDelta": -0.468437449,
   "want_Gamma": 0.02784272,
   "want_Vega": 0.284023591,
   "want_CallTheta": -0.015888004,
   "want_PutTheta": -0.015203928,
   "want_d1": 0.079198299,
   "want_d2": -0.062223057,
   "want_Nd1": 0.397693083,
   "want_Nd2": 0.475192601,
   "want_Cmoney": "ITM",
   "want_Pmoney": "OTM"
 },
 {
   "s": 105,
   "k": 100,
   "t": 0.5,
   "r": 0.0025,
   "v": 0.2,
   "q": 0.02,
   "want_CallPrice": 8.013559139,
   "want_PutPrice": 3.933404688,
   "want_CallDelta": 0.638269626,
   "want_PutDelta": -0.361730374,
   "want_Gamma": 0.025235897,
   "want_Vega": 0.27822576,
   "want_CallTheta": -0.015644822,
   "want_PutTheta": -0.014960746,
   "want_d1": 0.353837394,
   "want_d2": 0.212416038,
   "want_Nd1": 0.374733945,
   "want_Nd2": 0.584108763,
   "want_Cmoney": "ITM",
   "want_Pmoney": "OTM"
 },
 {
   "s": 110,
   "k": 100,
   "t": 0.5,
   "r": 0.0025,
   "v": 0.2,
   "q": 0.02,
   "want_CallPrice": 11.46641204,
   "want_PutPrice": 2.436008415,
   "want_CallDelta": 0.752628196,
   "want_PutDelta": -0.247371804,
   "want_Gamma": 0.020312799,
   "want_Vega": 0.245784867,
   "want_CallTheta": -0.013950533,
   "want_PutTheta": -0.013266457,
   "want_d1": 0.682783579,
   "want_d2": 0.541362223,
   "want_Nd1": 0.315992992,
   "want_Nd2": 0.70587103,
   "want_Cmoney": "ITM",
   "want_Pmoney": "OTM"
 },
 {
   "s": 90,
   "k": 100,
   "t": 0.5,
   "r": 0.0025,
   "v": 0.5,
   "q": 0.02,
   "want_CallPrice": 8.549763614,
   "want_PutPrice": 19.32035667,
   "want_CallDelta": 0.441969937,
   "want_PutDelta": -0.558030063,
   "want_Gamma": 0.012404673,
   "want_Vega": 0.25119463,
   "want_CallTheta": -0.0346214,
   "want_PutTheta": -0.033937324,
   "want_d1": -0.145976582,
   "want_d2": -0.499529973,
   "want_Nd1": 0.394714281,
   "want_Nd2": 0.308703038,
   "want_Cmoney": "OTM",
   "want_Pmoney": "ITM"
 },
 {
   "s": 95,
   "k": 100,
   "t": 0.5,
   "r": 0.0025,
   "v": 0.5,
   "q": 0.02,
   "want_CallPrice": 10.88920796,
   "want_PutPrice": 16.70955185,
   "want_CallDelta": 0.502772073,
   "want_PutDelta": -0.497227927,
   "want_Gamma": 0.011877389,
   "want_Vega": 0.267983583,
   "want_CallTheta": -0.036959387,
   "want_PutTheta": -0.036275311,
   "want_d1": 0.006948613,
   "want_d2": -0.346604778,
   "want_Nd1": 0.398932649,
   "want_Nd2": 0.364444128,
   "want_Cmoney": "OTM",
   "want_Pmoney": "ITM"
 },
 {
   "s": 99,
   "k": 100,
   "t": 0.5,
   "r": 0.0025,
   "v": 0.5,
   "q": 0.02,
   "want_CallPrice": 12.97293528,
   "want_PutPrice": 14.83307983,
   "want_CallDelta": 0.549184525,
   "want_PutDelta": -0.450815475,
   "want_Gamma": 0.011311038,
   "want_Vega": 0.277148698,
   "want_CallTheta": -0.038245406,
   "want_PutTheta": -0.03756133,
   "want_d1": 0.123601315,
   "want_d2": -0.229952075,
   "want_Nd1": 0.395906512,
   "want_Nd2": 0.409064505,
   "want_Cmoney": "OTM",
   "want_Pmoney": "ITM"
 },
 {
   "s": 100,
   "k": 100,
   "t": 0.5,
   "r": 0.0025,
   "v": 0.5,
   "q": 0.02,
   "want_CallPrice": 13.5222289,
   "want_PutPrice": 14.39232362,
   "want_CallDelta": 0.560417558,
   "want_PutDelta": -0.439582442,
   "want_Gamma": 0.011154144,
   "want_Vega": 0.278853598,
   "want_CallTheta": -0.038486533,
   "want_PutTheta": -0.037802457,
   "want_d1": 0.152027958,
   "want_d2": -0.201525433,
   "want_Nd1": 0.39435854,
   "want_Nd2": 0.420143872,
   "want_Cmoney": "ATM",
   "want_Pmoney": "ATM"
 },
 {
   "s": 101,
   "k": 100,
   "t": 0.5,
   "r": 0.0025,
   "v": 0.5,
   "q": 0.02,
   "want_CallPrice": 14.08256527,
   "want_PutPrice": 13.96261015,
   "want_CallDelta": 0.57149113,
   "want_PutDelta": -0.42850887,
   "want_Gamma": 0.010992202,
   "want_Vega": 0.28032862,
   "want_CallTheta": -0.038696138,
   "want_PutTheta": -0.038012062,
   "want_d1": 0.180171744,
   "want_d2": -0.173381647,
   "want_Nd1": 0.392519343,
   "want_Nd2": 0.431175725,
   "want_Cmoney": "ITM",
   "want_Pmoney": "OTM"
 },
 {
   "s": 105,
   "k": 100,
   "t": 0.5,
   "r": 0.0025,
   "v": 0.5,
   "q": 0.02,
   "want_CallPrice": 16.43106784,
   "want_PutPrice": 12.35091339,
   "want_CallDelta": 0.614102355,
   "want_PutDelta": -0.385897645,
   "want_Gamma": 0.010303867,
   "want_Vega": 0.284000324,
   "want_CallTheta": -0.039228867,
   "want_PutTheta": -0.038544791,
   "want_d1": 0.290027382,
   "want_d2": -0.063526009,
   "want_Nd1": 0.382511533,
   "want_Nd2": 0.474673824,
   "want_Cmoney": "ITM",
   "want_Pmoney": "OTM"
 },
 {
   "s": 110,
   "k": 100,
   "t": 0.5,
   "r": 0.0025,
   "v": 0.5,
   "q": 0.02,
   "want_CallPrice": 19.59480137,
   "want_PutPrice": 10.56439775,
   "want_CallDelta": 0.663343634,
   "want_PutDelta": -0.336656366,
   "want_Gamma": 0.009385645,
   "want_Vega": 0.283915763,
   "want_CallTheta": -0.039253166,
   "want_PutTheta": -0.03856909,
   "want_d1": 0.421605856,
   "want_d2": 0.068052465,
   "want_Nd1": 0.36501593,
   "want_Nd2": 0.527128065,
   "want_Cmoney": "ITM",
   "want_Pmoney": "OTM"
 },
 {
   "s": 90,
   "k": 100,
   "t": 0.1,
   "r": 0.0025,
   "v": 0.2,
   "q": 0.02,
   "want_CallPrice": 0.111291818,
   "want_PutPrice": 10.26611506,
   "want_CallDelta": 0.048262084,
   "want_PutDelta": -0.951737916,
   "want_Gamma": 0.017614281,
   "want_Vega": 0.028535135,
   "want_CallTheta": -0.007846774,
   "want_PutTheta": -0.007162014,
   "want_d1": -1.661943178,
   "want_d2": -1.725188731,
   "want_Nd1": 0.100262243,
   "want_Nd2": 0.042246733,
   "want_Cmoney": "OTM",
   "want_Pmoney": "ITM"
 },
 {
   "s": 95,
   "k": 100,
   "t": 0.1,
   "r": 0.0025,
   "v": 0.2,
   "q": 0.02,
   "want_CallPrice": 0.690819634,
   "want_PutPrice": 5.855632886,
   "want_CallDelta": 0.209814418,
   "want_PutDelta": -0.790185582,
   "want_Gamma": 0.047942054,
   "want_Vega": 0.086535408,
   "want_CallTheta": -0.02383985,
   "want_PutTheta": -0.023155089,
   "want_d1": -0.807065348,
   "want_d2": -0.870310901,
   "want_Nd1": 0.288051565,
   "want_Nd2": 0.192065262,
   "want_Cmoney": "OTM",
   "want_Pmoney": "ITM"
 },
 {
   "s": 99,
   "k": 100,
   "t": 0.1,
   "r": 0.0025,
   "v": 0.2,
   "q": 0.02,
   "want_CallPrice": 1.964386122,
   "want_PutPrice": 3.137191378,
   "want_CallDelta": 0.438427642,
   "want_PutDelta": -0.561572358,
   "want_Gamma": 0.062955085,
   "want_Vega": 0.123404557,
   "want_CallTheta": -0.034092709,
   "want_PutTheta": -0.033407949,
   "want_d1": -0.154956916,
   "want_d2": -0.218202469,
   "want_Nd1": 0.394181287,
   "want_Nd2": 0.413635681,
   "want_Cmoney": "OTM",
   "want_Pmoney": "ITM"
 },
 {
   "s": 100,
   "k": 100,
   "t": 0.1,
   "r": 0.0025,
   "v": 0.2,
   "q": 0.02,
   "want_CallPrice": 2.4334392,
   "want_PutPrice": 2.608242458,
   "want_CallDelta": 0.501576954,
   "want_PutDelta": -0.498423046,
   "want_Gamma": 0.06307782,
   "want_Vega": 0.126155641,
   "want_CallTheta": -0.034889381,
   "want_PutTheta": -0.034204621,
   "want_d1": 0.003952847,
   "want_d2": -0.059292706,
   "want_Nd1": 0.398939164,
   "want_Nd2": 0.476359485,
   "want_Cmoney": "ATM",
   "want_Pmoney": "ATM"
 },
 {
   "s": 101,
   "k": 100,
   "t": 0.1,
   "r": 0.0025,
   "v": 0.2,
   "q": 0.02,
   "want_CallPrice": 2.965314381,
   "want_PutPrice": 2.14211564,
   "want_CallDelta": 0.564064111,
   "want_PutDelta": -0.435935889,
   "want_Gamma": 0.061646771,
   "want_Vega": 0.125771741,
   "want_CallTheta": -0.03482713,
   "want_PutTheta": -0.03414237,
   "want_d1": 0.161281392,
   "want_d2": 0.098035839,
   "want_Nd1": 0.393787295,
   "want_Nd2": 0.539048082,
   "want_Cmoney": "ITM",
   "want_Pmoney": "OTM"
 },
 {
   "s": 105,
   "k": 100,
   "t": 0.1,
   "r": 0.0025,
   "v": 0.2,
   "q": 0.02,
   "want_CallPrice": 5.673239738,
   "want_PutPrice": 0.858033003,
   "want_CallDelta": 0.780946288,
   "want_PutDelta": -0.219053712,
   "want_Gamma": 0.044476886,
   "want_Vega": 0.098071534,
   "want_CallTheta": -0.027390573,
   "want_PutTheta": -0.026705813,
   "want_d1": 0.775393078,
   "want_d2": 0.712147525,
   "want_Nd1": 0.295361353,
   "want_Nd2": 0.761813286,
   "want_Cmoney": "ITM",
   "want_Pmoney": "OTM"
 },
 {
   "s": 110,
   "k": 100,
   "t": 0.1,
   "r": 0.0025,
   "v": 0.2,
   "q": 0.02,
   "want_CallPrice": 10.00866528,
   "want_PutPrice": 0.203448554,
   "want_CallDelta": 0.934598017,
   "want_PutDelta": -0.065401983,
   "want_Gamma": 0.018312765,
   "want_Vega": 0.044316891,
   "want_CallTheta": -0.012775804,
   "want_PutTheta": -0.012091043,
   "want_d1": 1.510939109,
   "want_d2": 1.447693556,
   "want_Nd1": 0.127402103,
   "want_Nd2": 0.926148614,
   "want_Cmoney": "ITM",
   "want_Pmoney": "OTM"
 },
 {
   "s": 90,
   "k": 100,
   "t": 0.0001,
   "r": 0.0025,
   "v": 0.01,
   "q": 0.02,
   "want_CallPrice": 0,
   "want_PutPrice": 10.000155,
   "want_CallDelta": 0,
   "want_PutDelta": -1,
   "want_Gamma": 0,
   "want_Vega": 0,
   "want_CallTheta": 0,
   "want_PutTheta": 0.000684931,
   "want_d1": -1053.622607,
   "want_d2": -1053.622707,
   "want_Nd1": 0,
   "want_Nd2": 0,
   "want_Cmoney": "OTM",
   "want_Pmoney": "ITM"
 },
 {
   "s": 95,
   "k": 100,
   "t": 0.0001,
   "r": 0.0025,
   "v": 0.01,
   "q": 0.02,
   "want_CallPrice": 0,
   "want_PutPrice": 5.000165,
   "want_CallDelta": 0,
   "want_PutDelta": -1,
   "want_Gamma": 0,
   "want_Vega": 0,
   "want_CallTheta": 0,
   "want_PutTheta": 0.000684931,
   "want_d1": -512.9503939,
   "want_d2": -512.9504939,
   "want_Nd1": 0,
   "want_Nd2": 0,
   "want_Cmoney": "OTM",
   "want_Pmoney": "ITM"
 },
 {
   "s": 99,
   "k": 100,
   "t": 0.0001,
   "r": 0.0025,
   "v": 0.01,
   "q": 0.02,
   "want_CallPrice": 0,
   "want_PutPrice": 1.000173,
   "want_CallDelta": 0,
   "want_PutDelta": -1,
   "want_Gamma": 0,
   "want_Vega": 0,
   "want_CallTheta": 0,
   "want_PutTheta": 0.000684931,
   "want_d1": -100.5208085,
   "want_d2": -100.5209085,
   "want_Nd1": 0,
   "want_Nd2": 0,
   "want_Cmoney": "OTM",
   "want_Pmoney": "ITM"
 },
 {
   "s": 100,
   "k": 100,
   "t": 0.0001,
   "r": 0.0025,
   "v": 0.01,
   "q": 0.02,
   "want_CallPrice": 0.003902529,
   "want_PutPrice": 0.004077529,
   "want_CallDelta": 0.49303881,
   "want_PutDelta": -0.50696119,
   "want_Gamma": 39.88815456,
   "want_Vega": 0.003988815,
   "want_CallTheta": -0.054978978,
   "want_PutTheta": -0.054294047,
   "want_d1": -0.01745,
   "want_d2": -0.01755,
   "want_Nd1": 0.398881546,
   "want_Nd2": 0.492998922,
   "want_Cmoney": "ATM",
   "want_Pmoney": "ATM"
 },
 {
   "s": 101,
   "k": 100,
   "t": 0.0001,
   "r": 0.0025,
   "v": 0.01,
   "q": 0.02,
   "want_CallPrice": 0.999823,
   "want_PutPrice": 0,
   "want_CallDelta": 1,
   "want_PutDelta": 0,
   "want_Gamma": 0,
   "want_Vega": 0,
   "want_CallTheta": -0.000684931,
   "want_PutTheta": 0,
   "want_d1": 99.48585853,
   "want_d2": 99.48575853,
   "want_Nd1": 0,
   "want_Nd2": 1,
   "want_Cmoney": "ITM",
   "want_Pmoney": "OTM"
 },
 {
   "s": 105,
   "k": 100,
   "t": 0.0001,
   "r": 0.0025,
   "v": 0.01,
   "q": 0.02,
   "want_CallPrice": 4.999815,
   "want_PutPrice": 0,
   "want_CallDelta": 1,
   "want_PutDelta": 0,
   "want_Gamma": 0,
   "want_Vega": 0,
   "want_CallTheta": -0.000684931,
   "want_PutTheta": 0,
   "want_d1": 487.8841917,
   "want_d2": 487.8840917,
   "want_Nd1": 0,
   "want_Nd2": 1,
   "want_Cmoney": "ITM",
   "want_Pmoney": "OTM"
 },
 {
   "s": 110,
   "k": 100,
   "t": 0.0001,
   "r": 0.0025,
   "v": 0.01,
   "q": 0.02,
   "want_CallPrice": 9.999805,
   "want_PutPrice": 0,
   "want_CallDelta": 1,
   "want_PutDelta": 0,
   "want_Gamma": 0,
   "want_Vega": 0,
   "want_CallTheta": -0.000684931,
   "want_PutTheta": 0,
   "want_d1": 953.084348,
   "want_d2": 953.084248,
   "want_Nd1": 0,
   "want_Nd2": 1,
   "want_Cmoney": "ITM",
   "want_Pmoney": "OTM"
 },
 {
   "s": 90,
   "k": 100,
   "t": 1,
   "r": 0.0025,
   "v": 0.4,
   "q": 0.02,
   "want_CallPrice": 9.819055509,
   "want_PutPrice": 21.35148715,
   "want_CallDelta": 0.457334479,
   "want_PutDelta": -0.542665521,
   "want_Gamma": 0.011018295,
   "want_Vega": 0.35699277,
   "want_CallTheta": -0.01977033,
   "want_PutTheta": -0.019087109,
   "want_d1": -0.107151289,
   "want_d2": -0.507151289,
   "want_Nd1": 0.396658634,
   "want_Nd2": 0.306024335,
   "want_Cmoney": "OTM",
   "want_Pmoney": "ITM"
 },
 {
   "s": 95,
   "k": 100,
   "t": 1,
   "r": 0.0025,
   "v": 0.4,
   "q": 0.02,
   "want_CallPrice": 12.19345796,
   "want_PutPrice": 18.82489624,
   "want_CallDelta": 0.51117561,
   "want_PutDelta": -0.48882439,
   "want_Gamma": 0.010494362,
   "want_Vega": 0.378846452,
   "want_CallTheta": -0.021001221,
   "want_PutTheta": -0.020318,
   "want_d1": 0.028016764,
   "want_d2": -0.371983236,
   "want_Nd1": 0.398785738,
   "want_Nd2": 0.354952665,
   "want_Cmoney": "OTM",
   "want_Pmoney": "ITM"
 },
 {
   "s": 99,
   "k": 100,
   "t": 1,
   "r": 0.0025,
   "v": 0.4,
   "q": 0.02,
   "want_CallPrice": 14.27868333,
   "want_PutPrice": 16.98932691,
   "want_CallDelta": 0.552161456,
   "want_PutDelta": -0.447838544,
   "want_Gamma": 0.009988065,
   "want_Vega": 0.39157209,
   "want_CallTheta": -0.021725203,
   "want_PutTheta": -0.021041981,
   "want_d1": 0.13112416,
   "want_d2": -0.26887584,
   "want_Nd1": 0.395527364,
   "want_Nd2": 0.394012615,
   "want_Cmoney": "OTM",
   "want_Pmoney": "ITM"
 },
 {
   "s": 100,
   "k": 100,
   "t": 1,
   "r": 0.0025,
   "v": 0.4,
   "q": 0.02,
   "want_CallPrice": 14.82478439,
   "want_PutPrice": 16.5552293,
   "want_CallDelta": 0.562082017,
   "want_PutDelta": -0.437917983,
   "want_Gamma": 0.00985255,
   "want_Vega": 0.394101982,
   "want_CallTheta": -0.021870454,
   "want_PutTheta": -0.021187233,
   "want_d1": 0.15625,
   "want_d2": -0.24375,
   "want_Nd1": 0.394101982,
   "want_Nd2": 0.403712223,
   "want_Cmoney": "ATM",
   "want_Pmoney": "ATM"
 },
 {
   "s": 101,
   "k": 100,
   "t": 1,
   "r": 0.0025,
   "v": 0.4,
   "q": 0.02,
   "want_CallPrice": 15.38054266,
   "want_PutPrice": 16.1307889,
   "want_CallDelta": 0.571865594,
   "want_PutDelta": -0.428134406,
   "want_Gamma": 0.009714151,
   "want_Vega": 0.396376213,
   "want_CallTheta": -0.022001671,
   "want_PutTheta": -0.021318449,
   "want_d1": 0.181125827,
   "want_d2": -0.218874173,
   "want_Nd1": 0.392451696,
   "want_Nd2": 0.413374033,
   "want_Cmoney": "ITM",
   "want_Pmoney": "OTM"
 },
 {
   "s": 105,
   "k": 100,
   "t": 1,
   "r": 0.0025,
   "v": 0.4,
   "q": 0.02,
   "want_CallPrice": 17.69739951,
   "want_PutPrice": 14.52685106,
   "want_CallDelta": 0.609580335,
   "want_PutDelta": -0.390419665,
   "want_Gamma": 0.009138008,
   "want_Vega": 0.402986161,
   "want_CallTheta": -0.022389934,
   "want_PutTheta": -0.021706713,
   "want_d1": 0.27822541,
   "want_d2": -0.12177459,
   "want_Nd1": 0.383796343,
   "want_Nd2": 0.451538769,
   "want_Cmoney": "ITM",
   "want_Pmoney": "OTM"
 },
 {
   "s": 110,
   "k": 100,
   "t": 1,
   "r": 0.0025,
   "v": 0.4,
   "q": 0.02,
   "want_CallPrice": 20.79386961,
   "want_PutPrice": 12.72232779,
   "want_CallDelta": 0.653403429,
   "want_PutDelta": -0.346596571,
   "want_Gamma": 0.008387999,
   "want_Vega": 0.405979132,
   "want_CallTheta": -0.02258555,
   "want_PutTheta": -0.021902329,
   "want_d1": 0.39452545,
   "want_d2": -0.00547455,
   "want_Nd1": 0.369071939,
   "want_Nd2": 0.497815981,
   "want_Cmoney": "ITM",
   "want_Pmoney": "OTM"
 }
]`)
	keys := make([]Opt, 0)

	if err := json.Unmarshal(keysBody, &keys); err != nil {
		panic(err)
	}

	return keys
}

func TestNdOne(t *testing.T) {
	for _, tt := range NewTestData() {
		t.Run(tt.WantCmoney, func(t *testing.T) {
			if got := lbr.NormPdf(d1(tt.S, tt.K, tt.R, tt.T, tt.Q, tt.V)); !withinRequiredPrecision(got, tt.WantNd1, 0.005) {
				t.Errorf("NdOne() = %v, want %v", got, tt.WantNd1)
			}
		})
	}
}

func TestNdTwo(t *testing.T) {
	for _, tt := range NewTestData() {
		t.Run(tt.WantCmoney, func(t *testing.T) {
			if got := lbr.NormCdf(d2(tt.S, tt.K, tt.R, tt.T, tt.Q, tt.V)); !withinRequiredPrecision(got, tt.WantNd2, 0.005) {
				t.Errorf("NdTwo() = %v, want %v", got, tt.WantNd2)
			}
		})
	}
}

func TestWithinRequiredPrecision(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		y    float64
		want bool
	}{
		{"equal", 2, 2, true},
		{"almost equal", 2.001, 2, true},
		{"0.005", 2.005, 2, true},
		{"0.006", 2.006, 2, false},
		{"not equal", 3, 2, false},
	}
	for _, tst := range tests {
		tt := tst
		t.Run(tt.name, func(t *testing.T) {
			if got := withinRequiredPrecision(tt.x, tt.y, 0.005); got != tt.want {
				t.Errorf("withinRequiredPrecision() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GenerateIncrementalArray(t *testing.T) {
	type args struct {
		median    float64
		qty       float64
		increment float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{"GenerateIncrementalArray", args{
			median:    100.00,
			qty:       12.0,
			increment: 5.00,
		},
			[]float64{
				70, 75, 80, 85, 90, 95, 100, 105, 110, 115, 120, 125,
			}},
		{"GenerateIncrementalArray odd qty", args{
			median:    3,
			qty:       4,
			increment: 1,
		},
			[]float64{
				1, 2, 3, 4,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateIncrementalArray(tt.args.median, tt.args.qty, tt.args.increment)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("computeUnderlyingPriceAxis() = %v, want %v", got, tt.want)
			}
			if float64(len(got)) != tt.args.qty {
				t.Errorf("computeUnderlyingPriceAxis() = %v, want %v", len(got), tt.args.qty)
			}
		})
	}
}

func TestSomeFunction(t *testing.T) {
	f := func(in float64) float64 {
		return in * 10
	}

	tests := []struct {
		name   string
		values []float64
		calc   fn
		want   []float64
	}{
		{"some name", []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, f, []float64{10, 20, 30, 40, 50, 60, 70, 80, 90}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SomeFunction(tt.values, tt.calc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SomeFunction() = %v, want %v", got, tt.want)
			}
			fmt.Println(tt.values, tt.want)
		})
	}
}
