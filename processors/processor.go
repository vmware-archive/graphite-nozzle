package processors

import (
	"github.com/cloudfoundry/sonde-go/events"
	"github.com/pivotal-cf/graphite-nozzle/metrics"
)

type Processor interface {
	Process(e *events.Envelope) ([]metrics.Metric, error)
}
