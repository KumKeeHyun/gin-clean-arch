package model

import "time"

type SensorData struct {
	NID       string
	Values    []float64
	Timestamp time.Time
}

type KafkaData struct {
	Key   string
	Value SensorData
}
