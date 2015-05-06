package processors_test

import (
	. "github.com/cloudcredo/graphite-nozzle/processors"
	"github.com/cloudfoundry/noaa/events"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ContainerMetricProcessor", func() {
	var (
		processor            *ContainerMetricProcessor
		event                *events.Envelope
		containerMetricEvent *events.ContainerMetric
	)

	BeforeEach(func() {
		processor = NewContainerMetricProcessor()

		applicationId := "60a13b0f-fce7-4c02-b92a-d43d583877ed"
		instanceIndex := int32(0)
		cpuPercentage := float64(70.75)
		memoryBytes := uint64(1024)
		diskBytes := uint64(2048)

		containerMetricEvent = &events.ContainerMetric{
			ApplicationId: &applicationId,
			InstanceIndex: &instanceIndex,
			CpuPercentage: &cpuPercentage,
			MemoryBytes:   &memoryBytes,
			DiskBytes:     &diskBytes,
		}

		event = &events.Envelope{
			ContainerMetric: containerMetricEvent,
		}
	})

	Describe("#Process", func() {
		It("returns a Metric for each of the ProcessContainerMetric* methods", func() {
			processedMetrics := processor.Process(event)

			Expect(processedMetrics).To(HaveLen(3))
		})
	})

	Describe("#ProcessContainerMetricCPU", func() {
		It("formats the Stat string to include the ContainerMetric's app ID and instance index", func() {
			metric := processor.ProcessContainerMetricCPU(containerMetricEvent)

			Expect(metric.Stat).To(Equal("apps.60a13b0f-fce7-4c02-b92a-d43d583877ed.cpu.0"))
		})

		It("sets the Metric Value to the value of the ContainerMetric cpuPercentage", func() {
			metric := processor.ProcessContainerMetricCPU(containerMetricEvent)

			Expect(metric.Value).To(Equal(int64(70)))
		})
	})

	Describe("#ProcessContainerMetricMemory", func() {
		It("formats the Stat string to include the ContainerMetric's app ID and instance index", func() {
			metric := processor.ProcessContainerMetricMemory(containerMetricEvent)

			Expect(metric.Stat).To(Equal("apps.60a13b0f-fce7-4c02-b92a-d43d583877ed.memoryBytes.0"))
		})

		It("sets the Metric Value to the value of the ContainerMetric memoryBytes", func() {
			metric := processor.ProcessContainerMetricMemory(containerMetricEvent)

			Expect(metric.Value).To(Equal(int64(1024)))
		})
	})

	Describe("#ProcessContainerMetricDisk", func() {
		It("formats the Stat string to include the ContainerMetric's app ID and instance index", func() {
			metric := processor.ProcessContainerMetricDisk(containerMetricEvent)

			Expect(metric.Stat).To(Equal("apps.60a13b0f-fce7-4c02-b92a-d43d583877ed.diskBytes.0"))
		})

		It("sets the Metric Value to the value of the ContainerMetric diskBytes", func() {
			metric := processor.ProcessContainerMetricDisk(containerMetricEvent)

			Expect(metric.Value).To(Equal(int64(2048)))
		})
	})

})
