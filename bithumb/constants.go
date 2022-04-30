package bithumb

type Currency string
type TimeInterval string
type SearchType string

const (
	KRW Currency = "krw"
	ALL Currency = "all" // -> 어떻게 짤 것인지 생각해봐야 함

	Min1   TimeInterval = "1m"
	Min3   TimeInterval = "3m"
	Min5   TimeInterval = "5m"
	Min10  TimeInterval = "10m"
	Min30  TimeInterval = "30m"
	Hour1  TimeInterval = "1h"
	Hour6  TimeInterval = "6h"
	Hour12 TimeInterval = "12h"
	Hour24 TimeInterval = "24h"

	All          SearchType = "0"
	BuyComplete  SearchType = "1"
	SellComplete SearchType = "2"
	InWidrawal   SearchType = "3"
	Deposit      SearchType = "4"
	Withdraw     SearchType = "5"
	InKRWDeposit SearchType = "9"
)

var urlWs = "wss://pubwss.bithumb.com/pub/ws"
var urlRest = "https://api.bithumb.com/public"
