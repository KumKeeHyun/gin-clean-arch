package presenter

import "github.com/dustin/go-broadcast"

const (
	Init = iota
	NewNode
	UpdateNode
	DeleteNode
	NewSensor
	DeleteSensor
)

type KafkaMessage struct {
	Type int         `json:"type"`
	Msg  interface{} `json:"message"`
}

type kafkaManager struct {
	broadcast.Broadcaster
}

var EnrichmentManager *kafkaManager

func KafkaSetup() {
	EnrichmentManager = &kafkaManager{
		broadcast.NewBroadcaster(10),
	}
}
