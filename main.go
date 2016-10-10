package main

// Inspired by the noaa firehose sample script
// https://github.com/cloudfoundry/noaa/blob/master/firehose_sample/main.go

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/cloudfoundry/noaa/consumer"
	"github.com/cloudfoundry/sonde-go/events"
	"github.com/pivotal-cf/graphite-nozzle/metrics"
	"github.com/pivotal-cf/graphite-nozzle/processors"
	"github.com/pivotal-cf/graphite-nozzle/token"
	"github.com/quipo/statsd"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	dopplerEndpoint   = kingpin.Flag("doppler-endpoint", "Doppler endpoint").Default("wss://doppler.10.244.0.34.xip.io:443").OverrideDefaultFromEnvar("DOPPLER_ENDPOINT").String()
	uaaEndpoint       = kingpin.Flag("uaa-endpoint", "UAA endpoint").Default("https://uaa.10.244.0.34.xip.io").OverrideDefaultFromEnvar("UAA_ENDPOINT").String()
	subscriptionId    = kingpin.Flag("subscription-id", "Id for the subscription.").Default("firehose").OverrideDefaultFromEnvar("SUBSCRIPTION_ID").String()
	statsdEndpoint    = kingpin.Flag("statsd-endpoint", "Statsd endpoint").Default("10.244.11.2:8125").OverrideDefaultFromEnvar("STATSD_ENDPOINT").String()
	statsdPrefix      = kingpin.Flag("statsd-prefix", "Statsd prefix").Default("mycf.").OverrideDefaultFromEnvar("STATSD_PREFIX").String()
	prefixJob         = kingpin.Flag("prefix-job", "Prefix metric names with job.index").Default("false").OverrideDefaultFromEnvar("PREFIX_JOB").Bool()
	username          = kingpin.Flag("username", "Firehose username.").Default("admin").OverrideDefaultFromEnvar("FIREHOSE_USERNAME").String()
	password          = kingpin.Flag("password", "Firehose password.").Default("admin").OverrideDefaultFromEnvar("FIREHOSE_PASSWORD").String()
	skipSSLValidation = kingpin.Flag("skip-ssl-validation", "Please don't").Default("false").OverrideDefaultFromEnvar("SKIP_SSL_VALIDATION").Bool()
	debug             = kingpin.Flag("debug", "Enable debug mode. This disables forwarding to statsd and prints to stdout").Default("false").OverrideDefaultFromEnvar("DEBUG").Bool()
)

func main() {
	kingpin.Parse()

	tokenFetcher := &token.UAATokenFetcher{
		UaaUrl:                *uaaEndpoint,
		Username:              *username,
		Password:              *password,
		InsecureSSLSkipVerify: *skipSSLValidation,
	}

	authToken, err := tokenFetcher.FetchAuthToken()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	consumer := consumer.New(*dopplerEndpoint, &tls.Config{InsecureSkipVerify: *skipSSLValidation}, nil)
	consumer.RefreshTokenFrom(tokenFetcher)

	httpStartStopProcessor := processors.NewHttpStartStopProcessor()
	valueMetricProcessor := processors.NewValueMetricProcessor()
	containerMetricProcessor := processors.NewContainerMetricProcessor()
	counterProcessor := processors.NewCounterProcessor()

	sender := statsd.NewStatsdClient(*statsdEndpoint, *statsdPrefix)
	sender.CreateSocket()

	var processedMetrics []metrics.Metric
	var proc_err error

	msgChan, errorChan := consumer.Firehose(*subscriptionId, authToken)

	go func() {
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
			processedMetrics, proc_err = containerMetricProcessor.Process(msg)
		case events.Envelope_CounterEvent:
			processedMetrics, proc_err = counterProcessor.Process(msg)
		case events.Envelope_HttpStartStop:
			processedMetrics, proc_err = httpStartStopProcessor.Process(msg)
		case events.Envelope_ValueMetric:
			processedMetrics, proc_err = valueMetricProcessor.Process(msg)
		default:
			// do nothing
		}

		if proc_err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", proc_err.Error())
			continue
		}

		if !*debug {
			if len(processedMetrics) > 0 {
				for _, metric := range processedMetrics {
					var prefix string
					if *prefixJob {
						prefix = msg.GetJob() + "." + msg.GetIndex()
					}
					metric.Send(sender, prefix)
				}
			}
		} else {
			for _, msg := range processedMetrics {
				fmt.Println(msg)
			}
		}

		processedMetrics = nil
	}
}
