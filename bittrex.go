package bittrex

import (
	"encoding/json"
)

type Bittrex struct {
	Client *Client
	OnMessageError func(err error)
	OnUpdateSummaryState func(summaryState SummaryState)
	OnUpdateExchangeState func(exchangeState ExchangeState)
}



func (self *Bittrex) Connect() error {
	self.Client = NewClient()
	self.Client.OnMessageError = self.OnMessageError
	self.Client.OnClientMethod = func(hub, method string, arguments []json.RawMessage) {
		switch method {
		case "updateSummaryState":
			if self.OnUpdateSummaryState != nil {
				sumarystate := SummaryState{}
				jMsg, _ := json.Marshal(arguments[0])
				json.Unmarshal(jMsg, &sumarystate)
				self.OnUpdateSummaryState(sumarystate)
			}
		case "updateExchangeState":
			if self.OnUpdateExchangeState != nil {
				exchangeState := ExchangeState{}
				jMsg, _ := json.Marshal(arguments[0])
				json.Unmarshal(jMsg, &exchangeState)
				self.OnUpdateExchangeState(exchangeState)
			}
		}
	}

	return self.Client.Connect("https", "socket.bittrex.com", []string{"CoreHub"})
}

func (self *Bittrex) SubscribeToExchangeDeltas(pair string) error {
	_, err := self.Client.CallHub("CoreHub", "SubscribeToExchangeDeltas", pair)
	return err
}

func (self *Bittrex) QueryExchangeState(pair string) (json.RawMessage, error)  {
	msg, err := self.Client.CallHub("CoreHub", "QueryExchangeState", pair)
	return msg, err
}