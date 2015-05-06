package processors

import (
	"github.com/cloudcredo/graphite-nozzle/metrics"
	"github.com/cloudfoundry/noaa/events"
)

type HeartbeatProcessor struct{}

func NewHeartbeatProcessor() *HeartbeatProcessor {
	return &HeartbeatProcessor{}
}

func (p *HeartbeatProcessor) Process(e *events.Envelope) []metrics.Metric {
	processedMetrics := make([]metrics.Metric, 4)
	heartbeat := e.GetHeartbeat()
	origin := e.GetOrigin()

	processedMetrics[0] = metrics.Metric(p.ProcessHeartbeatCount(heartbeat, origin))
	processedMetrics[1] = metrics.Metric(p.ProcessHeartbeatEventsSentCount(heartbeat, origin))
	processedMetrics[2] = metrics.Metric(p.ProcessHeartbeatEventsReceivedCount(heartbeat, origin))
	processedMetrics[3] = metrics.Metric(p.ProcessHeartbeatEventsErrorCount(heartbeat, origin))

	return processedMetrics
}

func (p *HeartbeatProcessor) ProcessHeartbeatCount(e *events.Heartbeat, origin string) *metrics.CounterMetric {
	stat := "ops." + origin + ".heartbeats.count"
	metric := metrics.NewCounterMetric(stat, int64(1))

	return metric
}

func (p *HeartbeatProcessor) ProcessHeartbeatEventsSentCount(e *events.Heartbeat, origin string) *metrics.GaugeMetric {
	stat := "ops." + origin + ".heartbeats.eventsSentCount"
	metric := metrics.NewGaugeMetric(stat, int64(e.GetSentCount()))

	return metric
}

func (p *HeartbeatProcessor) ProcessHeartbeatEventsReceivedCount(e *events.Heartbeat, origin string) *metrics.GaugeMetric {
	stat := "ops." + origin + ".heartbeats.eventsReceivedCount"
	metric := metrics.NewGaugeMetric(stat, int64(e.GetReceivedCount()))

	return metric
}

func (p *HeartbeatProcessor) ProcessHeartbeatEventsErrorCount(e *events.Heartbeat, origin string) *metrics.GaugeMetric {
	stat := "ops." + origin + ".heartbeats.eventsErrorCount"
	metric := metrics.NewGaugeMetric(stat, int64(e.GetErrorCount()))

	return metric
}
