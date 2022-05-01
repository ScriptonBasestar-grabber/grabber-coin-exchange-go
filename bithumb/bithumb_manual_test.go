package bithumb

import (
	"fmt"
	"testing"
)

func TestRest(t *testing.T) {
	t.Skipped()
	options := RequestOptions{
		Url:    "https://api.bithumb.com/public/orderbook/BTC_KRW",
		Method: "GET",
	}
	res := RestOrderbook{}
	Request(&options, &res)
	fmt.Println(res.Data)
}

func TestWS(t *testing.T) {
	t.Skipped()
	ws := WS{}
	ws.Init("wss://pubwss.bithumb.com/pub/ws", "localhost:9092")
	ws.Run("topic1")
}
