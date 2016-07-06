package processors

import (
	"github.com/cloudfoundry/noaa/events"
	"github.com/pivotal-cf/graphite-nozzle/metrics"
)

type Processor interface {
	Process(e *events.Envelope) []metrics.Metric
}
