package service

import "github.com/KumKeeHyun/gin-clean-arch/logic/domain/model"

type ElasticClient interface {
	GetInput() chan<- model.Document
}
