package main

import (
	"github.com/KumKeeHyun/gin-clean-arch/app/interface/db/orm"
	"github.com/KumKeeHyun/gin-clean-arch/app/interface/handler"
	"github.com/KumKeeHyun/gin-clean-arch/app/interface/presenter"
	"github.com/KumKeeHyun/gin-clean-arch/app/setting"
	"github.com/KumKeeHyun/gin-clean-arch/app/usecase/nodeUsecase"
	"github.com/KumKeeHyun/gin-clean-arch/app/usecase/sensorUsecase"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	presenter.KafkaSetup()
	setting.Setup()
	orm.Setup()

	nr := orm.NewNodeRepository()
	sr := orm.NewSensorRepository()

	nu := nodeUsecase.NewNodeUsecase(nr, sr)
	su := sensorUsecase.NewSensorUsecase(sr)

	h := handler.NewHandler(nu, su)

	r := gin.Default()

	ng := r.Group("/node")
	{
		ng.GET("", h.GetAllInfo)
		ng.POST("", h.RegisterNode)
	}
	sg := r.Group("/sensor")
	{
		sg.GET("", h.GetSensorsInfo)
		sg.POST("", h.RegisterSensor)
	}
	r.GET("/kafkaManager", h.KafkaManager)

	r.Run()
}
