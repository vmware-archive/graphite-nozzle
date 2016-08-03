package processors

import (
	"errors"
	"github.com/cloudfoundry/noaa/events"
	"github.com/pivotal-cf/graphite-nozzle/metrics"
	"net/url"
	"strconv"
	"strings"
)

type HttpStartStopProcessor struct{}

func NewHttpStartStopProcessor() *HttpStartStopProcessor {
	return &HttpStartStopProcessor{}
}

func (p *HttpStartStopProcessor) Process(e *events.Envelope) (processedMetrics []metrics.Metric, err error) {

	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknown error")
			}
			processedMetrics = nil
		}
	}()

	processedMetrics = make([]metrics.Metric, 4)
	httpStartStopEvent := e.GetHttpStartStop()

	processedMetrics[0] = metrics.Metric(p.ProcessHttpStartStopResponseTime(httpStartStopEvent))
	processedMetrics[1] = metrics.Metric(p.ProcessHttpStartStopStatusCodeCount(httpStartStopEvent))
	processedMetrics[2] = metrics.Metric(p.ProcessHttpStartStopHttpErrorCount(httpStartStopEvent))
	processedMetrics[3] = metrics.Metric(p.ProcessHttpStartStopHttpRequestCount(httpStartStopEvent))

	return
}

func (p *HttpStartStopProcessor) parseEventUri(uri string) string {
	parsed_uri, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}

	hostname := strings.Replace(parsed_uri.Host, ".", "_", -1)
	if !(len(hostname) > 0) {
		panic(errors.New("Hostname cannot be extracted from Event"))
	}

	return hostname
}

func (p *HttpStartStopProcessor) ProcessHttpStartStopResponseTime(event *events.HttpStartStop) *metrics.TimingMetric {
	statPrefix := "http.responsetimes."
	hostname := p.parseEventUri(event.GetUri())
	stat := statPrefix + hostname

	startTimestamp := event.GetStartTimestamp()
	stopTimestamp := event.GetStopTimestamp()
	durationNanos := stopTimestamp - startTimestamp
	durationMillis := durationNanos / 1000000 // NB: loss of precision here
	metric := metrics.NewTimingMetric(stat, durationMillis)

	return metric
}

func (p *HttpStartStopProcessor) ProcessHttpStartStopStatusCodeCount(event *events.HttpStartStop) *metrics.CounterMetric {
	statPrefix := "http.statuscodes."
	hostname := p.parseEventUri(event.GetUri())
	stat := statPrefix + hostname + "." + strconv.Itoa(int(event.GetStatusCode()))

	metric := metrics.NewCounterMetric(stat, isPeer(event))

	return metric
}

func (p *HttpStartStopProcessor) ProcessHttpStartStopHttpErrorCount(event *events.HttpStartStop) *metrics.CounterMetric {
	var incrementValue int64

	statPrefix := "http.errors."
	hostname := p.parseEventUri(event.GetUri())
	stat := statPrefix + hostname

	if 299 < event.GetStatusCode() && 1 == isPeer(event) {
		incrementValue = 1
	} else {
		incrementValue = 0
	}

	metric := metrics.NewCounterMetric(stat, incrementValue)

	return metric
}

func (p *HttpStartStopProcessor) ProcessHttpStartStopHttpRequestCount(event *events.HttpStartStop) *metrics.CounterMetric {
	statPrefix := "http.requests."
	hostname := p.parseEventUri(event.GetUri())
	stat := statPrefix + hostname
	metric := metrics.NewCounterMetric(stat, isPeer(event))

	return metric
}

func isPeer(event *events.HttpStartStop) int64 {
	if event.GetPeerType() == events.PeerType_Client {
		return 1
	} else {
		return 0
	}
}
