package metrics

import (
	"errors"
	"github.com/quipo/statsd"
	"time"
)

type StatsdSender interface {
	Gauge(stat string, value int64) error
	FGauge(stat string, value float64) error
	Incr(stat string, count int64) error
	Timing(string, int64) error
	PrecisionTiming(stat string, delta time.Duration) error
	Close() error
	CreateTCPSocket() error
	CreateSocket() error
}

type StatsdClient interface {
	StatsdSender
	Connect() error
	Reconnect() error
}

func newStatsdClient(conf map[string]string) (sender StatsdSender, err error) {
	var ok bool
	var endpoint, prefix string

	if endpoint, ok = conf["endpoint"]; !ok {
		err = errors.New("Missing endpoint for statsd client")
		return nil, err
	}
	if prefix, ok = conf["prefix"]; !ok {
		err = errors.New("Missing prefix for statsd client")
		return nil, err
	}
	return statsd.NewStatsdClient(endpoint, prefix), nil
}
