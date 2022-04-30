package bithumb

type WSType string

const (
	WSTypeTicker         WSType = "ticker"
	WSTypeOrderbookdepth WSType = "orderbookdepth"
	WSTypeTransaction    WSType = "transaction"
)

type TickerWrapper struct {
	Type    string `json:"type"`
	Content `json:"content"`
}

type OrderbookdepthWrapper struct {
	Type    string `json:"type"` // orderbookdepth
	Content struct {
		List []Orderbookdepth `json:"list"`
	} `json:"content"`
	Datetime uint64 `json:"datetime"` // 일시
}

type Orderbookdepth struct {
	Symbol    string `json:"symbol"`
	OrderType string `json:"orderType"` // 주문타입 bid/ask
	Price     string `json:"price"`     // 호가
	Quantity  string `json:"quantity"`  //잔량
	Total     int    `json:"total"`     // 건수
}

type TransactionWrapper struct {
	Type    string `json:"type"`
	Content struct {
		List []TransactionContent `json:"list"`
	} `json:"content"`
}

type TransactionContent struct {
	Symbol    string `json:"symbol"`    // 통화코드
	BuySellGb string `json:"buySellGb"` // 체결종류 (1: 매도체결, 2: 매수체결)
	ContPrice string `json:"contPrice"` // 체결가격
	ContQty   string `json:"contQty"`   // 체결수량
	ContAmt   string `json:"contAmt"`   // 체결금액
	ContDtm   string `json:"contDtm"`   // 체결시각
	Updn      string `json:"updn"`      // 직전시세와 비교: up-상승, dn-하락
}
type AssetsStatus struct {
	Status string `json:"status"`
	Data   []struct {
		DepositStatus    int `json:"deposit_status"`
		WithdrawalStatus int `json:"withdrawal_status"`
	} `json:"data"`
}

type BithumbCryptoIndex struct {
	Status string `json:"status"`
	Data   struct {
		Btai MarketIndex `json:"btai"`
		Btmi MarketIndex `json:"btmi"`
		Date string      `json:"date"` // 타임스템프
	} `json:"data"`
}

type MarketIndex struct {
	MarketIndex string `json:"market_index"` // 시장지수기준
	Width       string `json:"width"`        // 등락폭
	Rate        string `json:"rate"`         // 등락율
}
