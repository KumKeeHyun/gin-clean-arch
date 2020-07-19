package usecase

import (
	"github.com/KumKeeHyun/gin-clean-arch/app/domain/model"
	"github.com/KumKeeHyun/gin-clean-arch/app/interface/presenter"
)

type NodeUsecase interface {
	GetAllNodes() ([]presenter.Node, error)
	GetRegister() ([]model.Node, error)
	RegisterNode(*presenter.Node) (*model.Node, error)
}

type SensorUsecase interface {
	GetAllSensors() ([]model.Sensor, error)
	GetRegister() ([]model.Sensor, error)
	RegisterSensor(*model.Sensor) (*model.Sensor, error)
}

// type SinkUsecase interface {
// }
