package logicCore

import "github.com/KumKeeHyun/gin-clean-arch/logic/domain/model"

func getRinger(logic string) chainRing {
	switch logic {
	case "range":
		return &rangeRing{}
	case "loc":
		return &locFilterRing{}
	case "elastic":
		return &elasticRing{}
	default:
		return nil
	}
}

func chainFactory(rs []model.LogicRing) chainRing {
	return &chainRingBase{}
}
