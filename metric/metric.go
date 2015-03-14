package metric

type Metric interface {
	Send()
}

type GaugeMetric struct {
	Stat  string
	Value float64
}

type TimingMetric struct {
	Stat  string
	Value int64
}
