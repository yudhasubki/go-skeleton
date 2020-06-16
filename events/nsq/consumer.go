package nsq

import (
	"fmt"
	"log"
	"sync"

	"github.com/nsqio/go-nsq"
	"github.com/yudhasubki/go-skeleton/config"
)

var (
	waitGroup sync.WaitGroup
)

type NsqEventConsumer struct {
	Config        *config.Config
	Topic         string
	Channel       string
	HandleMessage func(message *nsq.Message) error
}

func (e *NsqEventConsumer) Consume() error {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(e.Topic, e.Channel, config)
	if err != nil {
		log.Printf("err create consumer : %v \n", err)
		return err
	}

	consumer.AddHandler(nsq.HandlerFunc(e.HandleMessage))
	err = consumer.ConnectToNSQD(fmt.Sprintf("%s:%s", e.Config.NsqHost, e.Config.NsqPort))
	if err != nil {
		log.Printf("err connect to nsqd : %v \n", err)
		return err
	}
	waitGroup.Wait()
	return nil
}
