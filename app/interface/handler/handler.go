package handler

import (
	"log"
	"net/http"

	"github.com/KumKeeHyun/gin-clean-arch/app/domain/model"
	"github.com/KumKeeHyun/gin-clean-arch/app/interface/presenter"
	"github.com/KumKeeHyun/gin-clean-arch/app/usecase"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	nu usecase.NodeUsecase
	su usecase.SensorUsecase
}

func NewHandler(nu usecase.NodeUsecase, su usecase.SensorUsecase) *Handler {
	return &Handler{
		nu: nu,
		su: su,
	}
}

func (h *Handler) GetAllInfo(c *gin.Context) {
	nodes, err := h.nu.GetAllNodes()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nodes)
}

func (h *Handler) RegisterNode(c *gin.Context) {
	var node presenter.Node

	if err := c.ShouldBindJSON(&node); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	new, err := h.nu.RegisterNode(&node)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	presenter.EnrichmentManager.Submit(presenter.KafkaMessage{
		Type: presenter.NewNode,
		Msg:  *new,
	})
	c.JSON(http.StatusOK, *new)
}

func (h *Handler) GetSensorsInfo(c *gin.Context) {
	sensors, err := h.su.GetAllSensors()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sensors)
}

func (h *Handler) RegisterSensor(c *gin.Context) {
	var sensor model.Sensor

	if err := c.ShouldBindJSON(&sensor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	new, err := h.su.RegisterSensor(&sensor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	presenter.EnrichmentManager.Submit(presenter.KafkaMessage{
		Type: presenter.NewSensor,
		Msg:  *new,
	})
	c.JSON(http.StatusOK, *new)
}

func (h *Handler) KafkaManager(c *gin.Context) {
	listen := make(chan interface{})
	presenter.EnrichmentManager.Register(listen)

	conn, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 0124)
	if err != nil {
		log.Printf("upgrade: %s", err.Error())
	}

	defer func() {
		presenter.EnrichmentManager.Unregister(listen)
		close(listen)
		conn.Close()
	}()

	nodeInfo, err := h.nu.GetRegister()
	sensorInfo, err := h.su.GetRegister()
	conn.WriteJSON(presenter.KafkaMessage{
		Type: presenter.Init,
		Msg: map[string]interface{}{
			"node_info":   nodeInfo,
			"sensor_info": sensorInfo,
		},
	})

	for {
		select {
		case m := <-listen:
			err = conn.WriteJSON(m)
		}
	}
}
