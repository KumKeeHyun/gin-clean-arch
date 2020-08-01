package service

import "github.com/KumKeeHyun/gin-clean-arch/logic/domain/model"

type KafkaConsumerGroup interface {
	GetOutput() <-chan model.KafkaData

	// IncreaseConsumer() error
	// DecreaseConsumer() error
}
