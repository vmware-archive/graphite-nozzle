package metrics

import (
	"time"
)

type StatsdTCPClient struct {
	StatsdSender
}

func NewStatsdTCPClient(conf map[string]string) (sender StatsdClient, err error) {
	statsd_sender, err := newStatsdClient(conf)

	if err != nil {
		return nil, err
	}

	return &StatsdTCPClient{StatsdSender: statsd_sender}, err
}

func (sender *StatsdTCPClient) Connect() (err error) {
	err = sender.CreateTCPSocket()
	return
}

func (sender *StatsdTCPClient) Reconnect() (err error) {
	sender.Close()
	return Retry(10, 5 * time.Second, sender.Connect)
}
