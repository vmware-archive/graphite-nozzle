package metrics_test

import (
	"errors"

	. "github.com/cloudcredo/graphite-nozzle/metrics"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"
)

type FakeStatsdClient struct {
	timingCalled          bool
	precisionTimingCalled bool
	incrCalled            bool
	gaugeCalled           bool
	fGaugeCalled          bool
	stat                  string
	value                 int64
	fValue                float64
	precisionTimingValue  time.Duration
}

func (f *FakeStatsdClient) Timing(stat string, delta int64) error {
	f.timingCalled = true
	f.stat = stat
	f.value = delta
	return nil
}

func (f *FakeStatsdClient) PrecisionTiming(stat string, delta time.Duration) error {
	f.precisionTimingCalled = true
	f.stat = stat
	f.precisionTimingValue = delta
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

func (f *FakeStatsdClient) FGauge(stat string, value float64) error {
	f.fGaugeCalled = true
	f.stat = stat
	f.fValue = value

	return errors.New("StatsdClientSendError")
}

var _ = Describe("Metric", func() {
	var (
		fakeStatsdClient *FakeStatsdClient
	)

	Describe("#NewCounterMetric", func() {
		It("creates a new CounterMetric", func() {
			metric := NewCounterMetric("my.counter.metric", 1)

			Expect(metric.Stat).To(Equal("my.counter.metric"))
			Expect(metric.Value).To(Equal(int64(1)))
		})
	})

	Describe("#NewGaugeMetric", func() {
		It("creates a new GaugeMetric", func() {
			metric := NewGaugeMetric("my.gauge.metric", 20)

			Expect(metric.Stat).To(Equal("my.gauge.metric"))
			Expect(metric.Value).To(Equal(int64(20)))
		})
	})

	Describe("#NewFGaugeMetric", func() {
		It("creates a new FGaugeMetric", func() {
			metric := NewFGaugeMetric("my.fgauge.metric", 20.25)

			Expect(metric.Stat).To(Equal("my.fgauge.metric"))
			Expect(metric.Value).To(Equal(float64(20.25)))
		})
	})

	Describe("#NewTimingMetric", func() {
		It("creates a new TimingMetric", func() {
			metric := NewTimingMetric("my.timing.metric", 100)

			Expect(metric.Stat).To(Equal("my.timing.metric"))
			Expect(metric.Value).To(Equal(int64(100)))
		})
	})

	Describe("#NewPrecisionTimingMetric", func() {
		It("creates a new PrecisionTimingMetric", func() {
			metric := NewPrecisionTimingMetric("my.precision.timing.metric", 100*time.Millisecond)

			Expect(metric.Stat).To(Equal("my.precision.timing.metric"))
			Expect(metric.Value).To(Equal(100 * time.Millisecond))
		})
	})

	Describe("#Send", func() {
		BeforeEach(func() {
			fakeStatsdClient = new(FakeStatsdClient)
		})

		Context("with a PrecisionTimingMetric", func() {
			It("sends the Metric to StatsD with time.Duration precision", func() {
				metric := NewPrecisionTimingMetric("http.responsetimes.api_10_244_0_34_xip_io", 50*time.Millisecond)
				metric.Send(fakeStatsdClient)

				Expect(fakeStatsdClient.precisionTimingCalled).To(BeTrue())
				Expect(fakeStatsdClient.stat).To(Equal("http.responsetimes.api_10_244_0_34_xip_io"))
				Expect(fakeStatsdClient.precisionTimingValue).To(Equal(50 * time.Millisecond))
			})
		})

		Context("with a CounterMetric", func() {
			It("sends the Metric to StatsD with int64 precision", func() {
				metric := NewCounterMetric("http.statuscodes.api_10_244_0_34_xip_io.200", 1)
				metric.Send(fakeStatsdClient)

				Expect(fakeStatsdClient.incrCalled).To(BeTrue())
				Expect(fakeStatsdClient.stat).To(Equal("http.statuscodes.api_10_244_0_34_xip_io.200"))
				Expect(fakeStatsdClient.value).To(Equal(int64(1)))
			})
		})

		Context("with a GaugeMetric", func() {
			It("sends the Metric to StatsD with int64 precision", func() {
				metric := NewGaugeMetric("router__0.numCPUS", 4)
				metric.Send(fakeStatsdClient)

				Expect(fakeStatsdClient.gaugeCalled).To(BeTrue())
				Expect(fakeStatsdClient.stat).To(Equal("router__0.numCPUS"))
				Expect(fakeStatsdClient.value).To(Equal(int64(4)))
			})
		})

		Context("with an FGaugeMetric", func() {
			It("sends the Metric to StatsD with float64 precision", func() {
				metric := NewFGaugeMetric("router__0.numCPUS", 4)
				metric.Send(fakeStatsdClient)

				Expect(fakeStatsdClient.fGaugeCalled).To(BeTrue())
				Expect(fakeStatsdClient.stat).To(Equal("router__0.numCPUS"))
				Expect(fakeStatsdClient.fValue).To(Equal(float64(4)))
			})
		})

		Context("when the StatsdClient doesn't return an error", func() {
			It("doesn't return an error", func() {
				metric := NewGaugeMetric("router__0.numCPUS", 4)
				err := metric.Send(fakeStatsdClient)

				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("when the StatsdClient returns an error", func() {
			It("returns the error", func() {
				metric := NewFGaugeMetric("router__0.numCPUS", 4)
				err := metric.Send(fakeStatsdClient)

				Expect(err).To(MatchError(errors.New("StatsdClientSendError")))
			})
		})
	})
})
