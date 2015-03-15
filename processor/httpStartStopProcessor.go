package processor

import (
	"strings"

	"github.com/cloudfoundry/noaa/events"
	"github.com/teddyking/graphite-nozzle/metrics"
)

type HttpStartStopProcessor struct{}

func NewHttpStartStopProcessor() *HttpStartStopProcessor {
	return &HttpStartStopProcessor{}
}

func (p *HttpStartStopProcessor) ProcessHttpStartStop(e *events.Envelope) []metrics.Metric {
	processedMetrics := make([]metrics.Metric, 2)
	httpStartStopEvent := e.GetHttpStartStop()

	processedMetrics[0] = metrics.Metric(p.ProcessHttpStartStopResponseTime(httpStartStopEvent))
	processedMetrics[1] = metrics.Metric(p.ProcessHttpStartStopStatusCodeCount(httpStartStopEvent))

	return processedMetrics
}

func (p *HttpStartStopProcessor) ProcessHttpStartStopResponseTime(event *events.HttpStartStop) *metrics.TimingMetric {
	statPrefix := "http.responsetimes."
	hostname := strings.Replace(strings.Split(event.GetUri(), "/")[0], ".", "_", -1)
	stat := statPrefix + hostname

	startTimestamp := event.GetStartTimestamp()
	stopTimestamp := event.GetStopTimestamp()
	durationNanos := stopTimestamp - startTimestamp
	durationMillis := durationNanos / 1000000 // NB: loss of precision here

	metric := &metrics.TimingMetric{
		Stat:  stat,
		Value: durationMillis,
	}

	return metric
}

func (p *HttpStartStopProcessor) ProcessHttpStartStopStatusCodeCount(event *events.HttpStartStop) *metrics.CounterMetric {
	statPrefix := "http.statuscodes."
	hostname := strings.Replace(strings.Split(event.GetUri(), "/")[0], ".", "_", -1)
	stat := statPrefix + hostname + ".200"

	metric := &metrics.CounterMetric{
		Stat:  stat,
		Value: int64(1),
	}

	return metric
}
