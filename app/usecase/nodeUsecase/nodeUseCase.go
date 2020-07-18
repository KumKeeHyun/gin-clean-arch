package nodeUsecase

import (
	"github.com/KumKeeHyun/gin-clean-arch/app/domain/model"
	"github.com/KumKeeHyun/gin-clean-arch/app/domain/repository"
	"github.com/KumKeeHyun/gin-clean-arch/app/usecase"
)

type nodeUsecase struct {
	nr repository.NodeRepository
	sr repository.SensorRepository
}

func NewNodeUsecase(nr repository.NodeRepository, sr repository.SensorRepository) *nodeUsecase {
	return &nodeUsecase{
		nr: nr,
		sr: sr,
	}
}

func (nu *nodeUsecase) GetAllNodes() ([]usecase.Node, error) {
	ns, err := nu.nr.GetAll()
	if err != nil {
		return nil, err
	}
	nodes := usecase.ToNodes(ns)
	for i := range nodes {
		nodes[i].Sensors, err = nu.sr.GetByNodeUUID(nodes[i].UUID)
		if err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (nu *nodeUsecase) GetRegister() ([]model.Node, error) {
	nodes, err := nu.nr.GetAll()
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func (nu *nodeUsecase) RegisterNode(n *usecase.Node) error {
	newNode := model.NewNode(n.Name, n.Location)
	if err := nu.nr.Create(&newNode); err != nil {
		return err
	}
	for _, s := range n.Sensors {
		ns := &model.NodeSensor{
			NodeUUID:   newNode.UUID,
			SensorUUID: s.UUID,
		}
		if err := nu.nr.CreateNS(ns); err != nil {
			return err
		}
	}
	return nil
}
