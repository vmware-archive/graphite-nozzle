package consumer

import (
	"crypto/tls"

	"github.com/cloudfoundry/noaa"
)

type Consumer struct {
	FirehoseConsumer *noaa.Consumer
}

func NewConsumer(dopplerAddress string) *Consumer {
	firehoseConsumer := noaa.NewConsumer(dopplerAddress, &tls.Config{InsecureSkipVerify: true}, nil)

	consumer := &Consumer{
		FirehoseConsumer: firehoseConsumer,
	}

	return consumer
}
