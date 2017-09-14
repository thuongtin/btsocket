package main

import (
	"../../bittrex"
	"fmt"
	"time"
)

func main() {
	bt := bittrex.Bittrex{}

	bt.OnUpdateExchangeState = func(exchangeState bittrex.ExchangeState) {
		fmt.Printf("%#v", exchangeState)
	}
	bt.OnUpdateSummaryState = func(summaryState bittrex.SummaryState) {
		fmt.Printf("%#v", summaryState.MarketSummaries)
	}

	bt.Connect()
	bt.SubscribeToExchangeDeltas("BTC-ETC")

	for {
		time.Sleep(time.Minute)
	}
}
