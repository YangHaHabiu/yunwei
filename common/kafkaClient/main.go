package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strings"
)

//var wg sync.WaitGroup

func main() {

	kafkaTest()
}

func kafkaTest() {
	config := sarama.NewConfig()
	config.Version = sarama.V0_10_2_0
	config.Net.DialTimeout = 20
	list := strings.Split("192.168.1.244:9999", ",")
	consumer, err := sarama.NewConsumer(list, config)
	if err != nil {
		fmt.Println("consumer connect error:", err)
		return
	}

	fmt.Println("connnect success...")
	fmt.Println(consumer.Topics())

}
