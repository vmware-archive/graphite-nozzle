package processors

import (
	"github.com/cloudfoundry/sonde-go/events"
	"github.com/pivotal-cf/graphite-nozzle/metrics"
)

type ValueMetricProcessor struct{}

func NewValueMetricProcessor() *ValueMetricProcessor {
	return &ValueMetricProcessor{}
}

func (p *ValueMetricProcessor) Process(e *events.Envelope) ([]metrics.Metric, error) {
	processedMetrics := make([]metrics.Metric, 1)
	valueMetricEvent := e.GetValueMetric()

	processedMetrics[0] = p.ProcessValueMetric(valueMetricEvent, e.GetOrigin())

	return processedMetrics, nil
}

func (p *ValueMetricProcessor) ProcessValueMetric(event *events.ValueMetric, origin string) *metrics.FGaugeMetric {
	statPrefix := "ops." + origin + "."
	valueMetricName := event.GetName()
	stat := statPrefix + valueMetricName
	metric := metrics.NewFGaugeMetric(stat, event.GetValue())

	return metric
}
