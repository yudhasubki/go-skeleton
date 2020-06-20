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
	nsq           *nsq.Config
	consumer      *nsq.Consumer
	Config        *config.Config
	Topic         string
	Channel       string
	HandleMessage func(message *nsq.Message) error
}

func (e *NsqEventConsumer) Start() {
	log.Println("Starting nsq consumer : topic %v - channel %v", e.Topic, e.Channel)
	e.nsq = nsq.NewConfig()
	e.consumer, err := nsq.NewConsumer(e.Topic, e.Channel, e.nsq)
	if err != nil {
		log.Fatalf("error creating consumer %v", err.Error())
		return
	}
}

func (e *NsqEventConsumer) Consume() error {
	e.consumer.AddHandler(nsq.HandlerFunc(e.HandleMessage))
	err = e.consumer.ConnectToNSQD(fmt.Sprintf("%s:%s", e.Config.NsqHost, e.Config.NsqPort))
	if err != nil {
		log.Printf("err connect to nsqd : %v \n", err)
		return err
	}
	waitGroup.Wait()
	return nil
}

func (e *NsqEventConsumer) Stop() error {
	log.Println("stoping consumer...")
	e.consumer.Stop()
}