package processors

import (
	"github.com/cloudfoundry/sonde-go/events"
	"github.com/pivotal-cf/graphite-nozzle/metrics"
)

type CounterProcessor struct{}

func NewCounterProcessor() *CounterProcessor {
	return &CounterProcessor{}
}

func (p *CounterProcessor) Process(e *events.Envelope) ([]metrics.Metric, error) {
	processedMetrics := make([]metrics.Metric, 1)
	counterEvent := e.GetCounterEvent()

	processedMetrics[0] = metrics.Metric(p.ProcessCounter(counterEvent))

	return processedMetrics, nil
}

func (p *CounterProcessor) ProcessCounter(event *events.CounterEvent) *metrics.CounterMetric {
	stat := "ops." + event.GetName()
	metric := metrics.NewCounterMetric(stat, int64(event.GetDelta()))

	return metric
}
