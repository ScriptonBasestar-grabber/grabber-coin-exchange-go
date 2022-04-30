package manual

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"testing"
)

func TestConsumer(t *testing.T) {
	t.Skipped()

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "testconsumer112",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	//c.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)
	c.SubscribeTopics([]string{"topic1"}, nil)

	//r1 := int32(3)
	for {
		msg, err := c.ReadMessage(-1)
		//utf8.EncodeRune(msg.Value, r1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}
