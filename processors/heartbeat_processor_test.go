package processors_test

import (
	. "github.com/cloudcredo/graphite-nozzle/processors"
	"github.com/cloudfoundry/noaa/events"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Heartbeat", func() {
	var (
		processor      *HeartbeatProcessor
		event          *events.Envelope
		heartbeatEvent *events.Heartbeat
		origin         string
	)

	BeforeEach(func() {
		processor = NewHeartbeatProcessor()

		origin = "router__0"
		sentCount := uint64(5491)
		receivedCount := uint64(5491)
		errorCount := uint64(1)

		heartbeatEvent = &events.Heartbeat{
			SentCount:     &sentCount,
			ReceivedCount: &receivedCount,
			ErrorCount:    &errorCount,
		}

		event = &events.Envelope{
			Origin:    &origin,
			Heartbeat: heartbeatEvent,
		}
	})

	Describe("#Process", func() {
		It("returns a Metric for each of the ProcessHeartbeat* methods", func() {
			processedMetrics := processor.Process(event)

			Expect(processedMetrics).To(HaveLen(4))
		})
	})

	Describe("#ProcessHeartbeatCount", func() {
		It("formats the Stat string to include the Heartbeat's Origin", func() {
			metric := processor.ProcessHeartbeatCount(heartbeatEvent, origin)

			Expect(metric.Stat).To(Equal("ops.router__0.heartbeats.count"))
		})

		It("sets the increment value for the CounterMetric to 1", func() {
			metric := processor.ProcessHeartbeatCount(heartbeatEvent, origin)

			Expect(metric.Value).To(Equal(int64(1)))
		})
	})

	Describe("#ProcessHeartbeatEventsSentCount", func() {
		It("formats the Stat string to include the Heartbeat's Origin", func() {
			metric := processor.ProcessHeartbeatEventsSentCount(heartbeatEvent, origin)

			Expect(metric.Stat).To(Equal("ops.router__0.heartbeats.eventsSentCount"))
		})

		It("sets the Metric Value to the value of the sentCount", func() {
			metric := processor.ProcessHeartbeatEventsSentCount(heartbeatEvent, origin)

			Expect(metric.Value).To(Equal(int64(5491)))
		})
	})

	Describe("#ProcessHeartbeatEventsReceivedCount", func() {
		It("formats the Stat string to include the Heartbeat's Origin", func() {
			metric := processor.ProcessHeartbeatEventsReceivedCount(heartbeatEvent, origin)

			Expect(metric.Stat).To(Equal("ops.router__0.heartbeats.eventsReceivedCount"))
		})

		It("sets the Metric Value to the value of the receivedCount", func() {
			metric := processor.ProcessHeartbeatEventsReceivedCount(heartbeatEvent, origin)

			Expect(metric.Value).To(Equal(int64(5491)))
		})
	})

	Describe("#ProcessHeartbeatEventsErrorCount", func() {
		It("formats the Stat string to include the Heartbeat's Origin", func() {
			metric := processor.ProcessHeartbeatEventsErrorCount(heartbeatEvent, origin)

			Expect(metric.Stat).To(Equal("ops.router__0.heartbeats.eventsErrorCount"))
		})

		It("sets the Metric Value to the value of the errorCount", func() {
			metric := processor.ProcessHeartbeatEventsErrorCount(heartbeatEvent, origin)

			Expect(metric.Value).To(Equal(int64(1)))
		})
	})
})
