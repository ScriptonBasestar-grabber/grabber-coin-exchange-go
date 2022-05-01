package bithumb

type RestAssetStatus struct {
	Status string `json:"status"`
	Data   map[string]struct {
		WithdrawalStatus int `json:"withdrawal_status"`
		DepositStatus    int `json:"deposit_status"`
	} `json:"data"`
}

//type RestTicker struct {
//	Status string `json:"status"`
//	Data   map[string]struct {
//		OpeningPrice     json.Number `json:"opening_price"`
//		ClosingPrice     json.Number `json:"closing_price"`
//		MinPrice         json.Number `json:"min_price"`
//		MaxPrice         json.Number `json:"max_price"`
//		UnitsTraded      json.Number `json:"units_traded"`
//		AccTradeValue    json.Number `json:"acc_trade_value"`
//		PrevClosingPrice json.Number `json:"prev_closing_price"`
//		UnitsTraded24H   json.Number `json:"units_traded_24H"`
//		AccTradeValue24H json.Number `json:"acc_trade_value_24H"`
//		Fluctate24H      json.Number `json:"fluctate_24H"`
//		FluctateRate24H  json.Number `json:"fluctate_rate_24H"`
//	} `json:"data"`
//}

type RestTicker struct {
	Status string                 `json:"status"`
	Data   map[string]interface{} `json:"data"`
}

type RestTickerContent struct {
	OpeningPrice     string `json:"opening_price" mapstructure:"opening_price"`
	ClosingPrice     string `json:"closing_price" mapstructure:"closing_price"`
	MinPrice         string `json:"min_price" mapstructure:"min_price"`
	MaxPrice         string `json:"max_price" mapstructure:"max_price"`
	UnitsTraded      string `json:"units_traded" mapstructure:"units_traded"`
	AccTradeValue    string `json:"acc_trade_value" mapstructure:"acc_trade_value"`
	PrevClosingPrice string `json:"prev_closing_price" mapstructure:"prev_closing_price"`
	UnitsTraded24H   string `json:"units_traded_24H" mapstructure:"units_traded_24H"`
	AccTradeValue24H string `json:"acc_trade_value_24H" mapstructure:"acc_trade_value_24H"`
	Fluctate24H      string `json:"fluctate_24H" mapstructure:"fluctate_24H"`
	FluctateRate24H  string `json:"fluctate_rate_24H" mapstructure:"fluctate_rate_24H"`
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
