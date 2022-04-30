package main

import (
	"fmt"
)

func main() {
	// config from .env
	//kafakServers := os.Getenv("KAFKA_SERVERS")
	//zkServers := os.Getenv("ZK_SERVERS")

	fmt.Println("read config success")
	// 생략 config from zookeeper
	// connect ws + publish kafka
	fmt.Println("connect to websocket")
	// go func
	fmt.Println("daemon running")
}
