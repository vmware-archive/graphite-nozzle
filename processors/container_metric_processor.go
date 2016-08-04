package processors

import (
	"strconv"
	"github.com/cloudfoundry/sonde-go/events"
	"github.com/pivotal-cf/graphite-nozzle/metrics"
)

type ContainerMetricProcessor struct{}

func NewContainerMetricProcessor() *ContainerMetricProcessor {
	return &ContainerMetricProcessor{}
}

func (p *ContainerMetricProcessor) Process(e *events.Envelope) ([]metrics.Metric, error) {
	processedMetrics := make([]metrics.Metric, 3)
	containerMetricEvent := e.GetContainerMetric()

	processedMetrics[0] = metrics.Metric(p.ProcessContainerMetricCPU(containerMetricEvent))
	processedMetrics[1] = metrics.Metric(p.ProcessContainerMetricMemory(containerMetricEvent))
	processedMetrics[2] = metrics.Metric(p.ProcessContainerMetricDisk(containerMetricEvent))

	return processedMetrics, nil
}

func (p *ContainerMetricProcessor) ProcessContainerMetricCPU(e *events.ContainerMetric) metrics.GaugeMetric {
	appID := e.GetApplicationId()
	instanceIndex := strconv.Itoa(int(e.GetInstanceIndex()))

	stat := "apps." + appID + ".cpu." + instanceIndex
	metric := metrics.NewGaugeMetric(stat, int64(e.GetCpuPercentage()))

	return *metric
}

func (p *ContainerMetricProcessor) ProcessContainerMetricMemory(e *events.ContainerMetric) metrics.GaugeMetric {
	appID := e.GetApplicationId()
	instanceIndex := strconv.Itoa(int(e.GetInstanceIndex()))

	stat := "apps." + appID + ".memoryBytes." + instanceIndex
	metric := metrics.NewGaugeMetric(stat, int64(e.GetMemoryBytes()))

	return *metric
}

func (p *ContainerMetricProcessor) ProcessContainerMetricDisk(e *events.ContainerMetric) metrics.GaugeMetric {
	appID := e.GetApplicationId()
	instanceIndex := strconv.Itoa(int(e.GetInstanceIndex()))

	stat := "apps." + appID + ".diskBytes." + instanceIndex
	metric := metrics.NewGaugeMetric(stat, int64(e.GetDiskBytes()))

	return *metric
}
