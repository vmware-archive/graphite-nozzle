package main

// Inspired by the noaa firehose sample script
// https://github.com/cloudfoundry/noaa/blob/master/firehose_sample/main.go

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/cloudfoundry/noaa"
	"github.com/cloudfoundry/noaa/events"
	"github.com/quipo/statsd"
	"github.com/teddyking/graphite-nozzle/processor"
)

const DopplerAddress = "wss://doppler.10.244.0.34.xip.io:443"
const firehoseSubscriptionId = "firehose-a"
const statsdAddress = "10.244.2.2:8125"
const statsdPrefix = "mycf."

var authToken = os.Getenv("CF_ACCESS_TOKEN")

func main() {
	consumer := noaa.NewConsumer(DopplerAddress, &tls.Config{InsecureSkipVerify: true}, nil)
	processor := processor.NewProcessor()
	sender := statsd.NewStatsdClient(statsdAddress, statsdPrefix)
	sender.CreateSocket()

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

		// graphite-nozzle can only handle HttpStartStop events at the moment
		if eventType == events.Envelope_HttpStartStop {
			metric := processor.ProcessHttpStartStop(msg)
			fmt.Printf("Processed HttpStartStopEvent\n")
			fmt.Printf("\t%s => %d\n", metric.Stat, metric.Value)

			sender.Timing(metric.Stat, metric.Value)
		}
	}
}
