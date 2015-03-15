package sender_test

import (
	. "github.com/teddyking/graphite-nozzle/sender"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/quipo/statsd"
)

var _ = Describe("Sender", func() {
	const (
		statsdAddress = "10.244.2.2:8125"
		statsdPrefix  = "mycf."
	)

	var (
		statsdClient *statsd.StatsdClient
	)

	BeforeEach(func() {
		statsdClient = statsd.NewStatsdClient(statsdAddress, statsdPrefix)
	})

	Describe("#NewSender", func() {
		It("creates a Sender with a StatsD client", func() {
			sender := NewSender(statsdClient)
			Expect(sender.StatsdClient).ToNot(BeNil())
		})
	})
})
