package kafka

import (
	"api-onboard-subscriber/host"
	"api-onboard-subscriber/models"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/wvanbergen/kafka/consumergroup"
	"os"
	"time"
)

const (
	zookeeperConn = "localhost:2181"
	consGroup     = "group2"
	topic         = "topic-register"
)

func ConnConsumer() (*consumergroup.ConsumerGroup, error) {
	config := consumergroup.NewConfig()

	config.Offsets.Initial = sarama.OffsetOldest
	config.Offsets.ProcessingTimeout = 10 * time.Second

	consGroup, err := consumergroup.JoinConsumerGroup(consGroup, []string{topic}, []string{zookeeperConn}, config)
	if err != nil {
		os.Exit(1)
	}
	return consGroup, nil
}

func Subscriber(cg *consumergroup.ConsumerGroup) {
	for {
		select {
		case msg := <-cg.Messages():
			if msg.Topic != topic {
				continue
			}

			//fmt.Println("Topic : ", msg.Topic)
			//fmt.Println("Topic : ", string(msg.Value))
			go sendCreateAccount(msg.Value)
			err := cg.CommitUpto(msg)
			if err != nil {
				fmt.Println("Error commit zookeeper :", err.Error())
			}
		}
	}
}

func sendCreateAccount(msg []byte) {
	var register models.UserRegister
	err := json.Unmarshal(msg, &register)
	if err != nil {
		fmt.Println("Message Broken")
	}
	acc := host.Account{
		FullName:    register.FullName,
		Email:       register.Email,
		PhoneNumber: register.PhoneNumber,
	}

	user := host.User{
		UserName: register.UserName,
		Password: register.Password,
		Email:    register.Email,
	}

	go host.SendCreateAccount(acc)
	go host.SendCreateUser(user)
}
