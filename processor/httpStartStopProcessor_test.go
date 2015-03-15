package processor_test

import (
	. "github.com/teddyking/graphite-nozzle/processor"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry/noaa/events"
)

var _ = Describe("HttpStartStopProcessor", func() {
	var (
		processor          *HttpStartStopProcessor
		event              *events.Envelope
		httpStartStopEvent *events.HttpStartStop
	)

	BeforeEach(func() {
		processor = NewHttpStartStopProcessor()

		startTimestamp := int64(1425881484152112140)
		stopTimestamp := int64(1425881484161498528)
		method := events.Method_GET
		uri := "api.10.244.0.34.xip.io/v2/info"
		statusCode := int32(200)

		httpStartStopEvent = &events.HttpStartStop{
			StartTimestamp: &startTimestamp,
			StopTimestamp:  &stopTimestamp,
			Method:         &method,
			Uri:            &uri,
			StatusCode:     &statusCode,
		}

		event = &events.Envelope{
			HttpStartStop: httpStartStopEvent,
		}
	})

	Describe("#Process", func() {
		It("returns a Metric for each of the ProcessHttpStartStop* methods", func() {
			processedMetrics := processor.Process(event)

			Expect(processedMetrics).To(HaveLen(2))
		})
	})

	Describe("#ProcessHttpStartStopResponseTime", func() {
		It("formats the Stat string to include the hostname", func() {
			metric := processor.ProcessHttpStartStopResponseTime(httpStartStopEvent)

			Expect(metric.Stat).To(Equal("http.responsetimes.api_10_244_0_34_xip_io"))
		})

		It("calculates the HTTP response time in milliseconds", func() {
			metric := processor.ProcessHttpStartStopResponseTime(httpStartStopEvent)

			Expect(metric.Value).To(Equal(int64(9)))
		})
	})

	Describe("#ProcessHttpStartStopStatusCodeCount", func() {
		It("formats the Stat string to include the hostname and the status code", func() {
			metric := processor.ProcessHttpStartStopStatusCodeCount(httpStartStopEvent)

			Expect(metric.Stat).To(Equal("http.statuscodes.api_10_244_0_34_xip_io.200"))
		})

		It("sets the increment value for the CounterMetric to 1", func() {
			metric := processor.ProcessHttpStartStopStatusCodeCount(httpStartStopEvent)

			Expect(metric.Value).To(Equal(int64(1)))
		})
	})
})
