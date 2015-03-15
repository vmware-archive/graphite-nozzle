package processor

import (
	"github.com/cloudfoundry/noaa/events"
	"github.com/teddyking/graphite-nozzle/metrics"
)

type Processor interface {
	Process(e *events.Envelope) []metrics.Metric
}
