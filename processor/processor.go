package processor

import (
	"strings"

	"github.com/cloudfoundry/noaa/events"
	"github.com/teddyking/graphite-nozzle/metric"
)

// A Processor is responsible for converting a 'raw' Event from the Firehose
// into a Metric that can be sent to StatsD.
type Processor struct{}

// Responsible for creating a new Processor.
func NewProcessor() *Processor {
	return &Processor{}
}

// Takes an HttpStartStop Event from the Firehose, processes it, and converts
// it into a TimingMetric that can then be sent to StatsD.
// The hostname is extracted from the provided URI and '.' characters
// are replaced with '_' characters. This is so that the stats don't appear
// ridiculously nested in the Graphite web UI.
// Note that there is a loss of precision here as the StatsD server only operates in
// millisecond timings and the StatsD client only accepts int64s for Timing metrics.
func (p *Processor) ProcessHttpStartStop(e *events.Envelope) *metric.TimingMetric {
	httpStartStopEvent := e.GetHttpStartStop()

	statPrefix := "http.hostnames."
	hostname := strings.Replace(strings.Split(httpStartStopEvent.GetUri(), "/")[0], ".", "_", -1)
	stat := statPrefix + hostname

	startTimestamp := httpStartStopEvent.GetStartTimestamp()
	stopTimestamp := httpStartStopEvent.GetStopTimestamp()
	durationNanos := stopTimestamp - startTimestamp
	durationMillis := durationNanos / 1000000 // NB: loss of precision here

	timingMetric := &metric.TimingMetric{
		Stat:  stat,
		Value: durationMillis,
	}

	return timingMetric
}
