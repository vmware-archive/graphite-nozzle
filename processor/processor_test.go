package processor_test

import (
	. "github.com/teddyking/graphite-nozzle/processor"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry/noaa/events"
)

var _ = Describe("Processor", func() {
	var (
		processor *Processor
		event     *events.Envelope
	)

	BeforeEach(func() {
		processor = NewProcessor()

		startTimestamp := int64(1425881484152112140)
		stopTimestamp := int64(1425881484161498528)
		method := events.Method_GET
		uri := "api.10.244.0.34.xip.io/v2/info"
		statusCode := int32(200)

		httpStartStopEvent := events.HttpStartStop{
			StartTimestamp: &startTimestamp,
			StopTimestamp:  &stopTimestamp,
			Method:         &method,
			Uri:            &uri,
			StatusCode:     &statusCode,
		}

		event = &events.Envelope{
			HttpStartStop: &httpStartStopEvent,
		}
	})

	Describe("#ProcessHttpStartStop", func() {
		It("creates a TimingMetric with a Stat and a Value", func() {
			httpStartStopMetric := processor.ProcessHttpStartStop(event)

			Expect(httpStartStopMetric.Stat).To(Equal("http.endpoints.api_10_244_0_34_xip_io"))
			Expect(httpStartStopMetric.Value).To(Equal(int64(9)))
		})
	})
})
