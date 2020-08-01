package kafkaConsumer

import (
	"fmt"
	"time"

	"github.com/KumKeeHyun/gin-clean-arch/logic/domain/model"
)

type consumer struct {
	// c *kafka.Consumer
	ctx chan struct{}
}

func NewConsumer() *consumer {
	return &consumer{
		ctx: make(chan struct{}, 1),
	}
}

func (c *consumer) run(out chan<- model.KafkaData) {
	for {
		select {
		case <-c.ctx:
			fmt.Printf("consumer stop\n")
			return
		case <-time.After(3 * time.Second):
			// doSomthing
		}
	}
}

func (c *consumer) stop() {
	c.ctx <- struct{}{}
}
