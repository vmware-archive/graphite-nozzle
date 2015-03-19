package metrics_test

import (
	. "github.com/CloudCredo/graphite-nozzle/metrics"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type FakeStatsdClient struct {
	timingCalled bool
	incrCalled   bool
	gaugeCalled  bool
	stat         string
	value        int64
}

func (f *FakeStatsdClient) Timing(stat string, delta int64) error {
	f.timingCalled = true
	f.stat = stat
	f.value = delta
	return nil
}

func (f *FakeStatsdClient) Incr(stat string, count int64) error {
	f.incrCalled = true
	f.stat = stat
	f.value = count
	return nil
}

func (f *FakeStatsdClient) Gauge(stat string, value int64) error {
	f.gaugeCalled = true
	f.stat = stat
	f.value = value
	return nil
}

var _ = Describe("Metric", func() {
	var (
		fakeStatsdClient *FakeStatsdClient
	)

	Describe("#Send", func() {
		Context("with a TimingMetric", func() {
			It("sends the Metric to StatsD", func() {
				fakeStatsdClient = new(FakeStatsdClient)
				metric := TimingMetric{
					Stat:  "http.responsetimes.api_10_244_0_34_xip_io",
					Value: 10,
				}

				metric.Send(fakeStatsdClient)
				Expect(fakeStatsdClient.timingCalled).To(BeTrue())
				Expect(fakeStatsdClient.stat).To(Equal("http.responsetimes.api_10_244_0_34_xip_io"))
				Expect(fakeStatsdClient.value).To(Equal(int64(10)))
			})
		})

		Context("with a CounterMetric", func() {
			It("sends the Metric to StatsD", func() {
				fakeStatsdClient = new(FakeStatsdClient)
				metric := CounterMetric{
					Stat:  "http.statuscodes.api_10_244_0_34_xip_io.200",
					Value: 1,
				}

				metric.Send(fakeStatsdClient)
				Expect(fakeStatsdClient.incrCalled).To(BeTrue())
				Expect(fakeStatsdClient.stat).To(Equal("http.statuscodes.api_10_244_0_34_xip_io.200"))
				Expect(fakeStatsdClient.value).To(Equal(int64(1)))
			})
		})

		Context("with a GaugeMetric", func() {
			It("sends the Metric to StatsD", func() {
				fakeStatsdClient = new(FakeStatsdClient)
				metric := GaugeMetric{
					Stat:  "router__0.numCPUS",
					Value: 4,
				}

				metric.Send(fakeStatsdClient)
				Expect(fakeStatsdClient.gaugeCalled).To(BeTrue())
				Expect(fakeStatsdClient.stat).To(Equal("router__0.numCPUS"))
				Expect(fakeStatsdClient.value).To(Equal(int64(4)))
			})
		})
	})
})
