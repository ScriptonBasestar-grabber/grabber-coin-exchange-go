package lib

import (
	"flag"
	"fmt"
	"github.com/ScriptonBasestar-grabber/bithumb"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
	"time"
)

//var addr = flag.String("addr", "localhost:8080", "http service address")

func Connect(url string) {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	//u := url.URL{Scheme: "wss", Host: host, Path: "/echo"}
	//log.Printf("connecting to %s", u.String())

	//c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			msg := bithumb.MsgDesc{}
			_, msgStr, _ := c.ReadMessage()
			fmt.Println("msg1111", string(msgStr))
			err = c.ReadJSON(&msg)
			if err != nil {
				log.Println("read:", err)
				return
			}
			if msg.Status != "0000" {
				fmt.Printf("failed to connect %s", msg)
			}
			log.Printf("connected : %s", msg)
		}
	}()

	c.WriteMessage(1, []byte("{\"type\":\"ticker\", \"symbols\": [\"BTC_KRW\", \"ETH_KRW\"], \"tickTypes\": [\"30M\", \"1H\", \"12H\", \"24H\", \"MID\" ]}"))
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
