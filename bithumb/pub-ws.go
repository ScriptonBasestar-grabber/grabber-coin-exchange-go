package bithumb

type WSType string

const (
	WSTypeTicker         WSType = "ticker"
	WSTypeOrderbookdepth WSType = "orderbookdepth"
	WSTypeTransaction    WSType = "transaction"
)

var ConversationConn = []MsgDesc{
	{Status: "0000", Msg: "Connected Successfully"},
	// 필터 설정응답
	// 성공
	{Status: "0000", Msg: "Filter Registered Successfully"},
	// 실패
	{Status: "5100", Msg: "Invalid Filter Syntax"},
}

//현재가(ticker)
//{"type":"ticker", "symbols": ["BTC_KRW", "ETH_KRW"], "tickTypes": ["30M", "1H", "12H", "24H", "MID" ]}
//체결(transaction)
//{"type":"transaction", "symbols":["BTC_KRW" , "ETH_KRW"]}
//변경호가(orderbookdepth)
//{"type":"orderbookdepth", "symbols":["BTC_KRW" , "ETH_KRW"]}
type WSRequest struct {
	Type      string   `json:"type"` // ticker, transaction, orderbookdepth
	Symbols   []string `json:"symbols"`
	TickTypes []string `json:"tickTypes"`
}
type WSTicker struct {
	Type    string `json:"type"` // ticker
	Content struct {
		Symbol         string `json:"symbol"`
		TickType       string `json:"tickType"`
		Date           string `json:"date"`
		Time           string `json:"time"`
		OpenPrice      string `json:"openPrice"`
		ClosePrice     string `json:"closePrice"`
		LowPrice       string `json:"lowPrice"`
		HighPrice      string `json:"highPrice"`
		Value          string `json:"value"`
		Volume         string `json:"volume"`
		SellVolume     string `json:"sellVolume"`
		BuyVolume      string `json:"buyVolume"`
		PrevClosePrice string `json:"prevClosePrice"`
		ChgRate        string `json:"chgRate"`
		ChgAmt         string `json:"chgAmt"`
		VolumePower    string `json:"volumePower"`
	} `json:"content"`
}

type WSOrderbookdepth struct {
	Type    string `json:"type"` // orderbookdepth
	Content struct {
		List     []WSOrderbookdepthContent `json:"list"`
		Datetime int64                     `json:"datetime"` // 일시
	} `json:"content"`
}

type WSOrderbookdepthContent struct {
	Symbol    string `json:"symbol"`
	OrderType string `json:"orderType"` // 주문타입 bid/ask
	Price     string `json:"price"`     // 호가
	Quantity  string `json:"quantity"`  //잔량
	Total     int    `json:"total"`     // 건수
}

type WSTransaction struct {
	Type    string `json:"type"` // transaction
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
