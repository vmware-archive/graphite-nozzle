package processors

import (
	"github.com/cloudcredo/graphite-nozzle/metrics"
	"github.com/cloudfoundry/noaa/events"
)

type CounterProcessor struct{}

func NewCounterProcessor() *CounterProcessor {
	return &CounterProcessor{}
}

func (p *CounterProcessor) Process(e *events.Envelope) []metrics.Metric {
	processedMetrics := make([]metrics.Metric, 1)
	counterEvent := e.GetCounterEvent()

	processedMetrics[0] = metrics.Metric(p.ProcessCounter(counterEvent))

	return processedMetrics
}

func (p *CounterProcessor) ProcessCounter(event *events.CounterEvent) *metrics.CounterMetric {
	stat := "ops." + event.GetName()
	metric := metrics.NewCounterMetric(stat, int64(event.GetDelta()))

	return metric
}
