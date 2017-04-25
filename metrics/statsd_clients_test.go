package metrics_test

import (
	. "github.com/pivotal-cf/graphite-nozzle/metrics"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StatsdClient", func() {
	var (
		clientConf map[string]string
	)

	Describe("#NewStatsdTCPClient", func() {

		BeforeEach(func() {
			clientConf = map[string]string{
				"protocol": "tcp",
				"endpoint": "statsd.test:8125",
				"prefix":   "test.prefix",
			}
		})

		It("returns a new tcp client if a valid conf is specified", func() {
			client, err := NewStatsdTCPClient(clientConf)
			Expect(err).To(BeNil())

			_, ok := client.(*StatsdTCPClient)
			Expect(ok).To(BeTrue())
		})

		It("returns an error if the config is not valid", func() {
			invalidConf := map[string]string{
			}

			_, err := NewStatsdTCPClient(invalidConf)
			Expect(err).NotTo(BeNil())
		})

	})

	Describe("#NewStatsdUDPClient", func() {

		BeforeEach(func() {
			clientConf = map[string]string{
				"protocol": "udp",
				"endpoint": "statsd.test:8125",
				"prefix":   "test.prefix",
			}
		})

		It("returns a new udp client if a valid conf is specified", func() {
			client, err := NewStatsdUDPClient(clientConf)
			Expect(err).To(BeNil())

			_, ok := client.(*StatsdUDPClient)
			Expect(ok).To(BeTrue())
		})

		It("returns an error if the config is not valid", func() {
			invalidConf := map[string]string{
			}

			_, err := NewStatsdUDPClient(invalidConf)
			Expect(err).NotTo(BeNil())
		})

	})

})
