package metric

type StatsdClient interface {
	Timing(string, int64) error
	Incr(stat string, count int64) error
	Gauge(stat string, value int64) error
}

type Metric interface {
	Send(StatsdClient) error
}

type GaugeMetric struct {
	Stat  string
	Value int64
}

type TimingMetric struct {
	Stat  string
	Value int64
}

type CounterMetric struct {
	Stat  string
	Value int64
}

func (m TimingMetric) Send(statsdClient StatsdClient) error {
	statsdClient.Timing(m.Stat, m.Value)
	return nil
}

func (m CounterMetric) Send(statsdClient StatsdClient) error {
	statsdClient.Incr(m.Stat, m.Value)
	return nil
}

func (m GaugeMetric) Send(statsdClient StatsdClient) error {
	statsdClient.Gauge(m.Stat, m.Value)
	return nil
}
