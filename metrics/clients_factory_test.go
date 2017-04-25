package metrics_test

import (
	. "github.com/pivotal-cf/graphite-nozzle/metrics"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ClientsFactory", func() {
	var (
		clientConf map[string]string
	)

	BeforeEach(func() {
		clientConf = map[string]string{
			"protocol": "udp",
			"endpoint": "statsd.test:8125",
			"prefix":   "test.prefix",
		}
	})

	Describe("#CreateClient", func() {

		It("returns a new udp client if udp protocol is specified", func() {
			client, err := CreateClient(clientConf)
			Expect(err).To(BeNil())

			_, ok := client.(*StatsdUDPClient)
			Expect(ok).To(BeTrue())
		})

		It("returns a new tcp client if tcp protocol is specified", func() {
			clientConf["protocol"] = "tcp"

			client, err := CreateClient(clientConf)
			Expect(err).To(BeNil())

			_, ok := client.(*StatsdTCPClient)
			Expect(ok).To(BeTrue())
		})

		It("returns an error if the client is not registered", func() {
			clientConf := map[string]string{
				"protocol": "not_existent",
			}

			_, err := CreateClient(clientConf)
			Expect(err).NotTo(BeNil())
		})

	})

})
