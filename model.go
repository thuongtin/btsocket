package bittrex

type ExchangeState struct {
	MarketName string   `json:"MarketName"`
	Nounce     int      `json:"Nounce"`
	Buys       []OrderBook  `json:"Buys"`
	Sells      []OrderBook  `json:"Sells"`
	Fills      []OrderHistory   `json:"Fills"`
}

type OrderBook struct {
	Quantity float64 `json:"Quantity"`
	Rate     float64 `json:"Rate"`
	Type     int     `json:"Type"`
}

type OrderHistory struct {
	OrderType string  `json:"OrderType"`
	Quantity  float64 `json:"Quantity"`
	Rate      float64 `json:"Rate"`
	Price     float64 `json:"Price"`
	TimeStamp string  `json:"TimeStamp"`
}

type SummaryState struct {
	Id	int64 `json:"Nounce"`
	MarketSummaries []MarketSummary `json:"Deltas"`
}

type MarketSummary struct {
	Ask            float64 `json:"Ask"`
	BaseVolume     float64 `json:"BaseVolume"`
	Bid            float64 `json:"Bid"`
	Created        string  `json:"Created"`
	High           float64 `json:"High"`
	Last           float64 `json:"Last"`
	Low            float64 `json:"Low"`
	MarketName     string  `json:"MarketName"`
	OpenBuyOrders  int     `json:"OpenBuyOrders"`
	OpenSellOrders int     `json:"OpenSellOrders"`
	PrevDay        float64 `json:"PrevDay"`
	TimeStamp      string  `json:"TimeStamp"`
	Volume         float64 `json:"Volume"`
	UnixTime		int64
}