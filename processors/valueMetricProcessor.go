package processors

import (
	"github.com/CloudCredo/graphite-nozzle/metrics"
	"github.com/cloudfoundry/noaa/events"
)

type ValueMetricProcessor struct{}

func NewValueMetricProcessor() *ValueMetricProcessor {
	return &ValueMetricProcessor{}
}

func (p *ValueMetricProcessor) Process(e *events.Envelope) []metrics.Metric {
	processedMetrics := make([]metrics.Metric, 1)
	valueMetricEvent := e.GetValueMetric()

	processedMetrics[0] = p.ProcessValueMetric(valueMetricEvent, e.GetOrigin())

	return processedMetrics
}

func (p *ValueMetricProcessor) ProcessValueMetric(event *events.ValueMetric, origin string) *metrics.GaugeMetric {
	statPrefix := "ops." + origin + "."
	valueMetricName := event.GetName()
	stat := statPrefix + valueMetricName

	metric := metrics.GaugeMetric{
		Stat:  stat,
		Value: int64(event.GetValue()),
	}

	return &metric
}
