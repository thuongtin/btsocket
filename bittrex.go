package bittrex

import (
	"github.com/gorilla/websocket"
	"net/url"
	"time"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"fmt"
	"io/ioutil"
)
const (
	BITTREX_HOST = "socket.bittrex.com"
	BITTREX_HUB = "coreHub"
)

type Bittrex struct {
	mutex sync.Mutex
	socket *websocket.Conn
	sc *Transport
	nextId int
	serverMessage chan []byte
	waitingQueryExchangeStateResponses map[string]string
	OnMessageError func(err error)
	OnUpdateSummaryState func(marketSummaries []MarketSummary)
	OnUpdateExchangeState func(exchangeState ExchangeState)
	OnUpdateAllExchangeState func(exchangeState ExchangeState)
	OnConnected func()
	AutoReconnect bool
	Id string
}

func (this Bittrex) NewClient() *Bittrex {
	scap, _ := NewTransport(http.DefaultTransport)
	serverMessage := make(chan []byte)
	return &Bittrex{sc: scap,
	nextId: 1,
	serverMessage: serverMessage,
	AutoReconnect: true,
	Id : "",
	waitingQueryExchangeStateResponses: make(map[string]string)}
}

func (this *Bittrex) getNegotiate() (Negotiate, error)  {
	var response Negotiate
	url := url.URL{Scheme: "https", Host: BITTREX_HOST, Path: "/signalr/negotiate"}

	c := http.Client{Transport: this.sc}
	res, err := c.Get(url.String())
	if err != nil {
		return response, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, err
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return response, err
	} else {
		return response, nil
	}
}


func (this *Bittrex) Connect()  {
	negotiate, err := this.getNegotiate()
	if err != nil {
		if this.AutoReconnect {
			fmt.Printf("%s - getNegotiate error. Reconnecting...\n", this.Id)
			this.Connect()
			return
		} else {
			panic(err)
		}
	}

	error := this.connectWebsocket(negotiate)
	if error != nil {
		if this.AutoReconnect {
			fmt.Println("%s - connectWebsocket error. Reconnecting...\n", this.Id)
			this.Connect()
			return
		} else {
			panic(err)
		}
	}
	go this.scanServerMessage()

	if this.OnConnected != nil {
		this.OnConnected()
	}

}

func (this *Bittrex) SubscribeToExchangeDeltas(pair string) (error) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	request := Request{
		Hub:        BITTREX_HUB,
		Method:     "SubscribeToExchangeDeltas",
		Arguments:  []string{pair},
		Identifier: this.nextId,
	}
	this.nextId++

	data, err := json.Marshal(request)
	if err != nil {
		return err
	}
	if err := this.socket.WriteMessage(websocket.TextMessage, data); err != nil {
		return err
	}

	return nil
}

func (this *Bittrex) QueryExchangeState(pair string) (error) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	request := Request{
		Hub:        BITTREX_HUB,
		Method:     "QueryExchangeState",
		Arguments:  []string{pair},
		Identifier: this.nextId,
	}
	this.nextId++

	data, err := json.Marshal(request)
	if err != nil {
		return err
	}
	reponseKey := fmt.Sprintf("%d", request.Identifier)
	this.waitingQueryExchangeStateResponses[reponseKey] = pair

	if err := this.socket.WriteMessage(websocket.TextMessage, data); err != nil {
		return err
	}
	return nil
}

func (this *Bittrex) connectWebsocket(negotiation Negotiate) error {
	var connectionParameters = url.Values{}
	connectionParameters.Set("transport", "webSockets")
	connectionParameters.Set("clientProtocol", "1.5")
	connectionParameters.Set("connectionToken", negotiation.ConnectionToken)
	connectionParameters.Set("connectionData", `[{"name":"corehub"}]`)


	header := http.Header{}
	header["Origin"] = []string{"https://bittrex.com"}
	header["User-Agent"] = []string{`Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:55.0) Gecko/20100101 Firefox/55.0`}
	header["Sec-WebSocket-Extensions"] = []string{"permessage-deflate"}
	header["Sec-WebSocket-Version"] = []string{"13"}


	u := url.URL{Scheme: "wss", Host: BITTREX_HOST, Path: "/signalr/connect"}
	u.RawQuery = connectionParameters.Encode()
	c, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		return err
	}

	this.socket = c
	go this.msgListener()
	if this.AutoReconnect {
		go this.ping()
	}
	return nil
}

func (this *Bittrex) msgListener()  {
	done := make(chan struct{})
	go func() {
		defer this.socket.Close()
		defer close(done)
		for {
			_, message, err := this.socket.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				this.socket.Close()
				return
			}
			this.serverMessage <- message
		}
	}()

}

func (this *Bittrex) scanServerMessage() {
	for {
		msg := <- this.serverMessage
		serverMessage := ServerMessage{}
		json.Unmarshal(msg, &serverMessage)
		if len(serverMessage.Identifier) > 0 {
			if pair, ok := this.waitingQueryExchangeStateResponses[serverMessage.Identifier]; ok {
				defer delete(this.waitingQueryExchangeStateResponses, serverMessage.Identifier)
				exchangeState := ExchangeState{}
				jMsg, _ := json.Marshal(serverMessage.Result)
				json.Unmarshal(jMsg, &exchangeState)
				exchangeState.MarketName = pair
				if this.OnUpdateAllExchangeState != nil {
					this.OnUpdateAllExchangeState(exchangeState)
				}
			}
		} else if len(serverMessage.Data) > 0 {
			switch serverMessage.Data[0].Method {
			case "updateSummaryState":
				summaryState := SummaryState{}
				jMsg, _ := json.Marshal(serverMessage.Data[0].Arguments[0])
				json.Unmarshal(jMsg, &summaryState)
				if this.OnUpdateSummaryState != nil {
					this.OnUpdateSummaryState(summaryState.MarketSummaries)
				}
			case "updateExchangeState":
				exchangeState := ExchangeState{}
				jMsg, _ := json.Marshal(serverMessage.Data[0].Arguments[0])
				json.Unmarshal(jMsg, &exchangeState)
				if this.OnUpdateExchangeState != nil {
					this.OnUpdateExchangeState(exchangeState)
				}
			}
		}
	}
}

func (this *Bittrex) Close() {
	this.socket.Close()
}

func (this *Bittrex) ping() {
	ticker := time.NewTicker(time.Second * 5)
	go func() {
		this.mutex.Lock()
		defer this.mutex.Unlock()
		for ; true; <-ticker.C {

			err := this.socket.WriteMessage(websocket.TextMessage, []byte("ping"))
			if err != nil {
				ticker.Stop()
				this.socket.Close()
				go this.Connect()
				break
			}
		}
	}()
}