package main

// Inspired by the noaa firehose sample script
// https://github.com/cloudfoundry/noaa/blob/master/firehose_sample/main.go

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/cloudcredo/graphite-nozzle/metrics"
	"github.com/cloudcredo/graphite-nozzle/processors"
	"github.com/cloudfoundry/noaa"
	"github.com/cloudfoundry/noaa/events"
	"github.com/quipo/statsd"
)

const DopplerAddress = "wss://doppler.10.244.0.34.xip.io:443"
const firehoseSubscriptionId = "firehose-a"
const statsdAddress = "10.244.2.2:8125"
const statsdPrefix = "mycf."

var authToken = os.Getenv("CF_ACCESS_TOKEN")

func main() {
	consumer := noaa.NewConsumer(DopplerAddress, &tls.Config{InsecureSkipVerify: true}, nil)

	httpStartStopProcessor := processors.NewHttpStartStopProcessor()
	valueMetricProcessor := processors.NewValueMetricProcessor()
	containerMetricProcessor := processors.NewContainerMetricProcessor()
	heartbeatProcessor := processors.NewHeartbeatProcessor()
	counterProcessor := processors.NewCounterProcessor()

	sender := statsd.NewStatsdClient(statsdAddress, statsdPrefix)
	sender.CreateSocket()

	var processedMetrics []metrics.Metric

	msgChan := make(chan *events.Envelope)
	go func() {
		defer close(msgChan)
		errorChan := make(chan error)
		go consumer.Firehose(firehoseSubscriptionId, authToken, msgChan, errorChan, nil)

		for err := range errorChan {
			fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		}
	}()

	for msg := range msgChan {
		eventType := msg.GetEventType()

		// graphite-nozzle can handle CounterEvent, ContainerMetric, Heartbeat,
		// HttpStartStop and ValueMetric events
		switch eventType {
		case events.Envelope_ContainerMetric:
			processedMetrics = containerMetricProcessor.Process(msg)
		case events.Envelope_CounterEvent:
			processedMetrics = counterProcessor.Process(msg)
		case events.Envelope_Heartbeat:
			processedMetrics = heartbeatProcessor.Process(msg)
		case events.Envelope_HttpStartStop:
			processedMetrics = httpStartStopProcessor.Process(msg)
		case events.Envelope_ValueMetric:
			processedMetrics = valueMetricProcessor.Process(msg)
		default:
			// do nothing
		}

		if len(processedMetrics) > 0 {
			for _, metric := range processedMetrics {
				metric.Send(sender)
			}
		}
		processedMetrics = nil
	}
}
