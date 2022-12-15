package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/wvanbergen/kafka/consumergroup"
	"log"
	"os"
	"time"
)

const (
	zookeeperConn = "localhost:2181"
	consGroup     = "group1"
	topic         = "msg-golang"
)

func connConsumer() (*consumergroup.ConsumerGroup, error) {
	config := consumergroup.NewConfig()

	config.Offsets.Initial = sarama.OffsetOldest
	config.Offsets.ProcessingTimeout = 10 * time.Second

	consGroup, err := consumergroup.JoinConsumerGroup(consGroup, []string{topic}, []string{zookeeperConn}, config)
	if err != nil {
		os.Exit(1)
	}
	return consGroup, nil
}

func main() {
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	consumer, _ := connConsumer()

	defer consumer.Close()

	subscriber(consumer)
}

func subscriber(cg *consumergroup.ConsumerGroup) {
	for {
		select {
		case msg := <-cg.Messages():
			if msg.Topic != topic {
				continue
			}

			fmt.Println("Topic : ", msg.Topic)
			fmt.Println("Topic : ", string(msg.Value))

			err := cg.CommitUpto(msg)
			if err != nil {
				fmt.Println("Error commit zookeeper :", err.Error())
			}
		}
	}
}
