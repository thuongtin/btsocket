package bittrex

import "encoding/json"

type ExchangeState struct {
	MarketName string   `json:"MarketName"`
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
}

type Negotiate struct {
	Url                     string
	ConnectionToken         string
	ConnectionId            string
	KeepAliveTimeout        float32
	DisconnectTimeout       float32
	ConnectionTimeout       float32
	TryWebSockets           bool
	ProtocolVersion         string
	TransportConnectTimeout float32
	LogPollDelay            float32
}

type Request struct {
	Hub        string        `json:"H"`
	Method     string        `json:"M"`
	Arguments  []string      `json:"A"`
	Identifier int           `json:"I"`
}

type Response struct {
	Hub        string        `json:"H"`
	Method     string        `json:"M"`
	Arguments  []json.RawMessage `json:"A"`
}

type ServerMessage struct {
	Cursor     string            `json:"C"`
	Data       []Response		 `json:"M"`
	Result     json.RawMessage   `json:"R"`
	Identifier string            `json:"I"`
	Error      string            `json:"E"`
}
