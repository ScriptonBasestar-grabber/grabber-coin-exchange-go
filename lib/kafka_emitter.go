package lib

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Init() *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		//"client.id":         socket.gethostname(),
		"client.id": "client1",
		"acks":      "all",
	})
	if err != nil {
		panic(err)
	}
	return p
}

func EmitMessage(p *kafka.Producer) {
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()
}
