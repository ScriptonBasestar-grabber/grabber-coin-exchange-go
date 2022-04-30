Grabber Coin Exchange GO
========================

## 기능

거래소 데이터 모아서 kafka에 전달하는 기능까지만 후 처리는 별도

## Install

```bash
go get -u
go mod tidy
go get github.com/gorilla/websocket
go get -u ./...
```

## Exchanges

추가중
* Bithumb
  https://apidocs.bithumb.com/docs
* Upbit
  https://docs.upbit.com/docs/upbit-quotation-websocket https://docs.upbit.com/v1.0.3/reference

예정
* Coinone
* Korbit
* Bittrex
* OKEX
* Bitmex
* Kraken
* CoinbasePro
...
