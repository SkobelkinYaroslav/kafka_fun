package main

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"kafka_fun/domain"
	"log"
	"math/rand"
	"time"
)

func main() {

	producerArr := make([]sarama.AsyncProducer, 10)
	for i := 0; i < 10; i++ {
		producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, nil)
		if err != nil {
			panic(err)
		}
		producerArr[i] = producer
	}

	for i := 0; i < 100; i++ {
		sensorId := rand.Intn(100)
		timestamp := time.Now()
		temperature := rand.Float64() * 50

		data := domain.Temperature{
			SensorId:    sensorId,
			Timestamp:   timestamp,
			Temperature: temperature,
		}

		dataBytes, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
			continue
		}
		msg := &sarama.ProducerMessage{
			Topic: "temperature_topic",
			Value: sarama.StringEncoder(dataBytes),
		}
		producerArr[i%10].Input() <- msg
		time.Sleep(2 * time.Second)
	}

}
