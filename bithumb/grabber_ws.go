package bithumb

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ScriptonBasestar-grabber/lib"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gorilla/websocket"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"time"
)

type WS struct {
	wsConn  *websocket.Conn
	kProd   *kafka.Producer
	symbols []string
}

func (w *WS) Init(url string, kafkaServers string) {
	var err error
	w.wsConn, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	//defer w.wsConn.Close()

	// kafka
	w.kProd, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaServers})
	if err != nil {
		panic(err)
	}
	//defer p.Close()

	// grab symbols
	client := &http.Client{}
	var body io.Reader
	req, err := http.NewRequest("GET", "https://api.bithumb.com/public/ticker/all_krw", body)
	lib.Err(err)
	res, err := client.Do(req)
	lib.Err(err)
	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	lib.Err(err)
	var r RestTicker
	err = json.Unmarshal(bytes, &r)

	for k, _ := range r.Data {
		w.symbols = append(w.symbols, k+"_KRW")
	}
}

func (w *WS) Run(topic string) {
	var err error
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Delivery report handler for produced messages
	go func() {
		for e := range w.kProd.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	regexComp, _ := regexp.Compile(`^{"status":"0000".+`)
	go func() {
		for {
			_, msgBArr, _ := w.wsConn.ReadMessage()
			//fmt.Println("msgType ", msgType)

			msg := MsgDesc{}
			//err = w.wsConn.ReadJSON(&msg)
			err = json.Unmarshal(msgBArr, &msg)
			if err != nil {
				log.Println("read:", err)
				return
			}

			//fmt.Println("msg ", string(msgBArr))
			// if....계속 호출되는건 싫은데
			if regexComp.Match(msgBArr) {
				log.Println("connected : ", msg)
			} else {
				w.kProd.ProduceChannel() <- &kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Value: msgBArr}
			}
		}
	}()

	var bArr []byte
	bArr, _ = json.Marshal(WSRequest{
		Type:      WSTypeTicker,
		Symbols:   Symbols,
		TickTypes: Intervals})
	w.wsConn.WriteMessage(websocket.TextMessage, bArr)
	bArr, _ = json.Marshal(WSRequest{
		Type:    WSTypeTransaction,
		Symbols: Symbols,
	})
	w.wsConn.WriteMessage(websocket.TextMessage, bArr)
	bArr, _ = json.Marshal(WSRequest{
		Type:    WSTypeOrderbookdepth,
		Symbols: Symbols,
	})
	w.wsConn.WriteMessage(websocket.TextMessage, bArr)

	defer w.wsConn.Close()
	defer w.kProd.Close()

	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := w.wsConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			//case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
