package bithumb

import (
	"encoding/json"
	"fmt"
	"github.com/ScriptonBasestar-grabber/lib"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestUnmarshalAssetStatus(t *testing.T) {
	var err error
	dat, err := ioutil.ReadFile("bithumb_assetstatus_all.json")
	lib.Err(err)
	//fmt.Println(string(dat))

	r := RestAssetStatus{}
	err = json.Unmarshal(dat, &r)
	lib.Err(err)
	fmt.Println(r)
}

func TestUnmarshalTicker_all_btc(t *testing.T) {
	var err error
	dat, err := ioutil.ReadFile("bithumb_ticker_all_btc.json")
	lib.Err(err)
	//fmt.Println(string(dat))

	r := RestTicker{}
	err = json.Unmarshal(dat, &r)
	lib.Err(err)
	fmt.Println(r.Data["date"])
	delete(r.Data, "date")
	assert.True(t, r.Data["date"] == nil)
	fmt.Println(r.Data["date"])

	m := r.Data["1INCH"].(map[string]interface{})
	fmt.Println(reflect.TypeOf(m["acc_trade_value"]))
	var r1 RestTickerContent
	err = mapstructure.Decode(m, &r1)
	lib.Err(err)
	fmt.Println("r1", r1)
	assert.True(t, m["opening_price"] == r1.OpeningPrice)
}

func TestUnmarshalTicker_all_krw(t *testing.T) {
	var err error
	dat, err := ioutil.ReadFile("bithumb_ticker_all_krw.json")
	lib.Err(err)
	//fmt.Println(string(dat))

	r := RestTicker{}
	err = json.Unmarshal(dat, &r)
	lib.Err(err)
	fmt.Println(r.Data["date"])
	delete(r.Data, "date")
	assert.True(t, r.Data["date"] == nil)
	fmt.Println(r.Data["date"])

	m := r.Data["BTC"].(map[string]interface{})
	fmt.Println(reflect.TypeOf(m["acc_trade_value"]))
	var r1 RestTickerContent
	err = mapstructure.Decode(m, &r1)
	lib.Err(err)
	fmt.Println("r1", r1)
	assert.True(t, m["opening_price"] == r1.OpeningPrice)
}
