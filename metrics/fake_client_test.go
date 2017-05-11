package metrics_test

import (
	"errors"
	"time"

	. "github.com/pivotal-cf/graphite-nozzle/metrics"
)

type FakeStatsdClient struct {
	Name                  string
	timingCalled          bool
	precisionTimingCalled bool
	incrCalled            bool
	gaugeCalled           bool
	fGaugeCalled          bool
	stat                  string
	value                 int64
	fValue                float64
	precisionTimingValue  time.Duration
}

func NewFakeClient(conf map[string]string) (fake StatsdClient, err error) {
	name, ok := conf["name"]
	if !ok {
		name = "fake"
	}

	return &FakeStatsdClient{Name: name}, nil
}

func (f *FakeStatsdClient) Timing(stat string, delta int64) error {
	f.timingCalled = true
	f.stat = stat
	f.value = delta
	return nil
}

func (f *FakeStatsdClient) PrecisionTiming(stat string, delta time.Duration) error {
	f.precisionTimingCalled = true
	f.stat = stat
	f.precisionTimingValue = delta
	return nil
}

func (f *FakeStatsdClient) Incr(stat string, count int64) error {
	f.incrCalled = true
	f.stat = stat
	f.value = count
	return nil
}

func (f *FakeStatsdClient) Gauge(stat string, value int64) error {
	f.gaugeCalled = true
	f.stat = stat
	f.value = value
	return nil
}

func (f *FakeStatsdClient) FGauge(stat string, value float64) error {
	f.fGaugeCalled = true
	f.stat = stat
	f.fValue = value

	return errors.New("StatsdClientSendError")
}

func (f *FakeStatsdClient) CreateSocket() error {
	return nil
}

func (f *FakeStatsdClient) CreateTCPSocket() error {
	return nil
}


func (f *FakeStatsdClient) Close() error {
	return nil
}

func (f *FakeStatsdClient) Connect() error {
	return nil
}

func (f *FakeStatsdClient) Reconnect() error {
	return nil
}
