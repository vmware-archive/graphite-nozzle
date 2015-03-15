package sender

import (
	"github.com/quipo/statsd"
)

type Sender struct {
	StatsdClient *statsd.StatsdClient
}

func NewSender(statsdClient *statsd.StatsdClient) *Sender {
	return &Sender{
		StatsdClient: statsdClient,
	}
}
