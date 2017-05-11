package processors_test

import (
	. "github.com/pivotal-cf/graphite-nozzle/processors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry/sonde-go/events"
)

var _ = Describe("ValueMetricProcessor", func() {
	var (
		processor        *ValueMetricProcessor
		event            *events.Envelope
		valueMetricEvent *events.ValueMetric
		origin           string
	)

	BeforeEach(func() {
		processor = NewValueMetricProcessor()

		origin = "router__0"
		name := "numCPUS"
		unit := "count"
		value := float64(4)

		valueMetricEvent = &events.ValueMetric{
			Name:  &name,
			Unit:  &unit,
			Value: &value,
		}

		event = &events.Envelope{
			Origin:      &origin,
			ValueMetric: valueMetricEvent,
		}
	})

	Describe("#Process", func() {
		It("returns a Metric for each of the ProcessValueMetric* methods", func() {
			processedMetrics, err := processor.Process(event)

			Expect(err).To(BeNil())
			Expect(processedMetrics).To(HaveLen(1))
		})
	})

	Describe("#ProcessValueMetric", func() {
		It("formats the Stat string to include the ValueMetric's name and Origin", func() {
			metric := processor.ProcessValueMetric(valueMetricEvent, origin)

			Expect(metric.Stat).To(Equal("ops.router__0.numCPUS"))
		})

		It("sets the Metric Value to the value of the ValueMetric", func() {
			metric := processor.ProcessValueMetric(valueMetricEvent, origin)

			Expect(metric.Value).To(Equal(float64(4)))
		})
	})
})
