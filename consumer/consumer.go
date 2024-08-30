package main

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"kafka_fun/domain"
	"log"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("temperature_topic", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	defer partitionConsumer.Close()

	for msg := range partitionConsumer.Messages() {
		var recvVal domain.Temperature
		err := json.Unmarshal(msg.Value, &recvVal)
		if err != nil {
			log.Printf("unmarshal error: %v\n", err)
			continue
		}

		fmt.Printf("Получено: SensorID: %d, Timestamp: %s, Temperature: %.2f°C\n",
			recvVal.SensorId, recvVal.Timestamp, recvVal.Temperature)

		if recvVal.Temperature < -5 || recvVal.Temperature > 35 {
			fmt.Printf("ВНИМАНИЕ! Температура вне допустимого диапазона: %.2f°C\n", recvVal.Temperature)
		}
	}
}
