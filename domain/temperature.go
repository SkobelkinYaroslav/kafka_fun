package domain

import "time"

type Temperature struct {
	SensorId    int       `json:"sensorId"`
	Timestamp   time.Time `json:"timestamp"`
	Temperature float64   `json:"temperature"`
}
