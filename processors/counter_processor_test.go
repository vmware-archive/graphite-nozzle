package processors_test

import (
	. "github.com/pivotal-cf/graphite-nozzle/processors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry/sonde-go/events"
)

var _ = Describe("CounterEventProcessor", func() {
	var (
		processor    *CounterProcessor
		event        *events.Envelope
		counterEvent *events.CounterEvent
	)

	BeforeEach(func() {
		processor = NewCounterProcessor()

		name := "pollCount"
		delta := uint64(1)

		counterEvent = &events.CounterEvent{
			Name:  &name,
			Delta: &delta,
		}

		event = &events.Envelope{
			CounterEvent: counterEvent,
		}
	})

	Describe("#Process", func() {
		It("returns a Metric for each of the ProcessCounter* methods", func() {
			processedMetrics, err := processor.Process(event)

			Expect(err).To(BeNil())
			Expect(processedMetrics).To(HaveLen(1))
		})
	})

	Describe("#ProcessCounter", func() {
		It("formats the Stat string to include the Counter's name", func() {
			metric := processor.ProcessCounter(counterEvent)

			Expect(metric.Stat).To(Equal("ops." + "pollCount"))
		})

		It("sets the Metric Value to the value of the Counter's delta", func() {
			metric := processor.ProcessCounter(counterEvent)

			Expect(metric.Value).To(Equal(int64(1)))
		})
	})
})
