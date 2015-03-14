package processor

import (
	"strings"

	"github.com/cloudfoundry/noaa/events"
	"github.com/teddyking/graphite-nozzle/metric"
)

type Processor struct{}

func NewProcessor() *Processor {
	return &Processor{}
}

func (p *Processor) ProcessHttpStartStop(e *events.Envelope) *metric.TimingMetric {
	httpStartStopEvent := e.GetHttpStartStop()

	statPrefix := "http.endpoints."
	uri := strings.Replace(strings.Split(httpStartStopEvent.GetUri(), "/")[0], ".", "_", -1)
	stat := statPrefix + uri

	startTimestamp := httpStartStopEvent.GetStartTimestamp()
	stopTimestamp := httpStartStopEvent.GetStopTimestamp()
	durationNanos := stopTimestamp - startTimestamp
	durationMillis := durationNanos / 1000000 // NB: loss of precision here

	timingMetric := &metric.TimingMetric{
		Stat:  stat,
		Value: durationMillis,
	}

	return timingMetric
}
