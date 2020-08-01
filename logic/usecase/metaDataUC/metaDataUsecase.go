package metaDataUC

import (
	"github.com/KumKeeHyun/gin-clean-arch/logic/domain/model"
	"github.com/KumKeeHyun/gin-clean-arch/logic/domain/repository"
	"github.com/KumKeeHyun/gin-clean-arch/logic/domain/service"
)

type metaDataUsecase struct {
	mr repository.MetaRepo
	ls service.LogicCore
}

func (mu *metaDataUsecase) NewNode(key string, n *model.Node) (*model.Node, error) {
	err := mu.mr.NewNode(key, n)
	if err != nil {
		return nil, err
	}
	return n, nil
}

func (mu *metaDataUsecase) NewSensor(key string, s *model.Sensor) (*model.Sensor, error) {
	err := mu.mr.NewSensor(key, s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (mu *metaDataUsecase) DeleteNode(key string) error {
	return mu.mr.DelNode(key)
}

func (mu *metaDataUsecase) DeleteSensor(key string) error {
	err := mu.mr.DelSensor(key)
	if err != nil {
		return err
	}
	return mu.ls.RemoveLogicsBySID(key)
}
