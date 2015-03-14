package consumer_test

import (
	. "github.com/teddyking/graphite-nozzle/consumer"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Consumer", func() {
	It("creates a Firehose consumer", func() {
		dopplerAddress := "wss://doppler.10.244.0.34.xip.io:443"

		consumer := NewConsumer(dopplerAddress)
		Expect(consumer.FirehoseConsumer).ToNot(BeNil())
	})
})
