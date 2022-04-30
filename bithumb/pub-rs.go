package bithumb

type RestTicker struct {
	Status string `json:"status"`
	Data   struct {
		OpeningPrice     string `json:"opening_price"`
		ClosingPrice     string `json:"closing_price"`
		MinPrice         string `json:"min_price"`
		MaxPrice         string `json:"max_price"`
		UnitsTraded      string `json:"units_traded"`
		AccTradeValue    string `json:"acc_trade_value"`
		PrevClosingPrice string `json:"prev_closing_price"`
		UnitsTraded24H   string `json:"units_traded_24H"`
		AccTradeValue24H string `json:"acc_trade_value_24H"`
		Fluctate24H      string `json:"fluctate_24H"`
		FluctateRate24H  string `json:"fluctate_rate_24H"`
		Date             string `json:"date"`
	} `json:"data"`
}
type RestOrderbook struct {
	Status string `json:"status"`
	Data   struct {
		Timestamp       string   `json:"timestamp"`
		OrderCurrency   string   `json:"order_currency"`
		PaymentCurrency string   `json:"payment_currency"`
		Bids            []BidAsk `json:"bids"`
		Asks            []BidAsk `json:"asks"`
	} `json:"data"`
}

type BidAsk struct {
	Quantity string `json:"quantity"`
	Price    string `json:"price"`
}

type RestTransaction struct {
	Status string `json:"status"`
	Data   []struct {
		TransactionDate string `json:"transaction_date"`
		Type            string `json:"type"`
		UnitsTraded     string `json:"units_traded"`
		Price           string `json:"price"`
		Total           string `json:"total"`
	} `json:"data"`
}
