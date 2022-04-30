package test

import (
	"regexp"
	"testing"
)

func BenchmarkRegex(b *testing.B) {
	var msgArr = []string{
		"{\"status\":\"0000\",\"resmsg\":\"Filter Registered Successfully\"}",
		"{\"type\":\"orderbookdepth\",\"content\":{\"list\":[{\"symbol\":\"ETH_KRW\",\"orderType\":\"ask\",\"price\":\"3625000\",\"quantity\":\"85.9733\",\"total\":\"4\"}],\"datetime\":\"1651337500115563\"}}",
		"{\"type\":\"orderbookdepth\",\"content\":{\"list\":[{\"symbol\":\"BTC_KRW\",\"orderType\":\"bid\",\"price\":\"49602000\",\"quantity\":\"0\",\"total\":\"0\"},{\"symbol\":\"BTC_KRW\",\"orderType\":\"bid\",\"price\":\"42612000\",\"quantity\":\"0.0743\",\"total\":\"1\"}],\"datetime\":\"1651337502110445\"}}",
	}

	//regexComp, _ := regexp.Compile(`^(\s+)?{(\s+)?"status"(\s+)?:(\s+)?"0000".+`)
	regexComp, _ := regexp.Compile(`^{"status":"0000".+`)

	for _, msg := range msgArr {
		matched := regexComp.MatchString(msg)
		//matched := regexp.MatchString(msg)
		println(matched)
	}

	//msg := bithumb.MsgDesc{}
	//err = w.wsConn.ReadJSON(&msg)
	//err = json.Unmarshal(msgBArr, &msg)
}
