package main

import (
	"../../btsocket"
	"fmt"
	"time"
)

func main() {
	bt := bittrex.Bittrex{}.NewClient()
	bt.OnUpdateSummaryState = func(summaries []bittrex.MarketSummary) {
		for _, summary := range summaries {
			fmt.Printf("%s\n", summary.MarketName)
		}
	}
	bt.OnUpdateExchangeState = func(exchangeState bittrex.ExchangeState) {
		fmt.Printf("%s - Update Exchange State\n", exchangeState.MarketName)
	}

	bt.OnUpdateAllExchangeState = func(exchangeState bittrex.ExchangeState) {
		fmt.Printf("======================Pair %s, Buy: %#v\n", exchangeState.MarketName, exchangeState.Buys[0])
	}

	bt.Connect()
	pairs := []string{"BTC-1ST","BTC-2GIVE","BTC-ABY","BTC-ADT","BTC-ADX","BTC-AEON","BTC-AGRS","BTC-AMP","BTC-ANT","BTC-APX","BTC-ARDR","BTC-ARK","BTC-AUR","BTC-BAT","BTC-BAY","BTC-BCC","BTC-BCY","BTC-BITB","BTC-BLITZ","BTC-BLK","BTC-BLOCK","BTC-BNT","BTC-BRK","BTC-BRX","BTC-BSD","BTC-BTA","BTC-BTCD","BTC-BTS","BTC-BURST","BTC-BYC","BTC-CANN","BTC-CFI","BTC-CLAM","BTC-CLOAK","BTC-CLUB","BTC-COVAL","BTC-CPC","BTC-CRB","BTC-CRW","BTC-CURE","BTC-CVC","BTC-DAR","BTC-DASH","BTC-DCR","BTC-DCT","BTC-DGB","BTC-DGD","BTC-DMD","BTC-DOGE","BTC-DOPE","BTC-DRACO","BTC-DTB","BTC-DYN","BTC-EBST","BTC-EDG","BTC-EFL","BTC-EGC","BTC-EMC","BTC-EMC2","BTC-ENRG","BTC-ERC","BTC-ETC","BTC-ETH","BTC-EXCL","BTC-EXP","BTC-FAIR","BTC-FCT","BTC-FLDC","BTC-FLO","BTC-FTC","BTC-FUN","BTC-GAM","BTC-GAME","BTC-GBG","BTC-GBYTE","BTC-GCR","BTC-GEO","BTC-GLD","BTC-GNO","BTC-GNT","BTC-GOLOS","BTC-GRC","BTC-GRS","BTC-GUP","BTC-HMQ","BTC-INCNT","BTC-INFX","BTC-IOC","BTC-ION","BTC-IOP","BTC-KMD","BTC-KORE","BTC-LBC","BTC-LGD","BTC-LMC","BTC-LSK","BTC-LTC","BTC-LUN","BTC-MAID","BTC-MCO","BTC-MEME","BTC-MLN","BTC-MONA","BTC-MTL","BTC-MUE","BTC-MUSIC","BTC-MYST","BTC-NAV","BTC-NBT","BTC-NEO","BTC-NEOS","BTC-NLG","BTC-NMR","BTC-NXC","BTC-NXS","BTC-NXT","BTC-OK","BTC-OMG","BTC-OMNI","BTC-PART","BTC-PAY","BTC-PDC","BTC-PINK","BTC-PIVX","BTC-PKB","BTC-POT","BTC-PPC","BTC-PTC","BTC-PTOY","BTC-QRL","BTC-QTUM","BTC-QWARK","BTC-RADS","BTC-RBY","BTC-RDD","BTC-REP","BTC-RISE","BTC-RLC","BTC-SAFEX","BTC-SBD","BTC-SC","BTC-SEQ","BTC-SHIFT","BTC-SIB","BTC-SLR","BTC-SLS","BTC-SNGLS","BTC-SNRG","BTC-SNT","BTC-SPHR","BTC-SPR","BTC-START","BTC-STEEM","BTC-STORJ","BTC-STRAT","BTC-SWIFT","BTC-SWT","BTC-SYNX","BTC-SYS","BTC-THC","BTC-TIME","BTC-TKN","BTC-TKS","BTC-TRIG","BTC-TRST","BTC-TRUST","BTC-TX","BTC-UBQ","BTC-UNB","BTC-UNO","BTC-VIA","BTC-VOX","BTC-VRC","BTC-VRM","BTC-VTC","BTC-VTR","BTC-WAVES","BTC-WINGS","BTC-XAUR","BTC-XCP","BTC-XDN","BTC-XEL","BTC-XEM","BTC-XLM","BTC-XMG","BTC-XMR","BTC-XMY","BTC-XRP","BTC-XST","BTC-XVC","BTC-XVG","BTC-XWC","BTC-XZC","BTC-ZCL","BTC-ZEC","BTC-ZEN","ETH-1ST","ETH-ADT","ETH-ADX","ETH-ANT","ETH-BAT","ETH-BCC","ETH-BNT","ETH-BTS","ETH-CFI","ETH-CRB","ETH-CVC","ETH-DASH","ETH-DGB","ETH-DGD","ETH-ETC","ETH-FCT","ETH-FUN","ETH-GNO","ETH-GNT","ETH-GUP","ETH-HMQ","ETH-LGD","ETH-LTC","ETH-LUN","ETH-MCO","ETH-MTL","ETH-MYST","ETH-NEO","ETH-NMR","ETH-OMG","ETH-PAY","ETH-PTOY","ETH-QRL","ETH-QTUM","ETH-REP","ETH-RLC","ETH-SC","ETH-SNGLS","ETH-SNT","ETH-STORJ","ETH-STRAT","ETH-TIME","ETH-TKN","ETH-TRST","ETH-WAVES","ETH-WINGS","ETH-XEM","ETH-XLM","ETH-XMR","ETH-XRP","ETH-ZEC","USDT-BCC","USDT-BTC","USDT-DASH","USDT-ETC","USDT-ETH","USDT-LTC","USDT-NEO","USDT-XMR","USDT-XRP","USDT-ZEC"}
	for _, pair := range pairs {
		go bt.SubscribeToExchangeDeltas(pair)
	}

	bt.QueryExchangeState("BTC-NEO") // Việc lấy hết orderbook và orderhistory về tốn thời gian > 1s nên tốt nhất cho nó vào 1 luồng khác

	for {
		time.Sleep(time.Second)
	}
}
