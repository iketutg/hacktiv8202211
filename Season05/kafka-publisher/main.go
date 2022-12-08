package main

import (
	"bufio"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"time"
)

const (
	kafkaConn = "localhost:9092"
	topic     = "msg-golang"
)

func connProducer() (sarama.SyncProducer, error) {
	//producer config
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{kafkaConn}, config)
	return producer, err
}

func main() {
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)
	producer, err := connProducer()
	if err != nil {
		fmt.Println("Can't Connect to kafka Error ", err.Error())
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Enter Message to Publsih topic %s : ", topic)
		readString, _ := reader.ReadString('\n')
		publish(readString, producer)
	}

}

func publish(readString string, producer sarama.SyncProducer) {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.StringEncoder(readString),
		Timestamp: time.Time{},
	}

	partion, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error Publish : ", err.Error())
	}
	fmt.Println("Partion : ", partion)
	fmt.Println("offset : ", offset)

}
