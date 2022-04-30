package manual

import (
	"fmt"
	"github.com/ScriptonBasestar-grabber/bithumb"
	"github.com/ScriptonBasestar-grabber/lib"
	"testing"
)

func TestRest(t *testing.T) {
	t.Skipped()
	options := lib.RequestOptions{
		Url:    "https://api.bithumb.com/public/orderbook/BTC_KRW",
		Method: "GET",
	}
	res := bithumb.RestOrderbook{}
	lib.Request(&options, &res)
	fmt.Println(res.Data)
}

func TestWS(t *testing.T) {
	t.Skipped()
	ws := lib.WS{}
	ws.Init("wss://pubwss.bithumb.com/pub/ws", "localhost:9092")
	ws.Run("topic1")
}
