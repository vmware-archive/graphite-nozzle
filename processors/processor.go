package processors

import (
	"github.com/CloudCredo/graphite-nozzle/metrics"
	"github.com/cloudfoundry/noaa/events"
)

type Processor interface {
	Process(e *events.Envelope) []metrics.Metric
}
