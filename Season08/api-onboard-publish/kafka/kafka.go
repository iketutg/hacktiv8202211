package kafka

import (
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

func ConnProducer() (sarama.SyncProducer, error) {
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)
	//producer config
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{kafkaConn}, config)
	return producer, err
}

func Publish(data []byte, producer sarama.SyncProducer) error {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.ByteEncoder(data),
		Timestamp: time.Time{},
	}

	partion, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error Publish : ", err.Error())
		return err
	}
	fmt.Println("Partion : ", partion)
	fmt.Println("offset : ", offset)

	return nil
}
