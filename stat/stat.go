package stat

type Stat interface {
	Send()
}

type GaugeStat struct {
	Stat  string
	Value float64
}

type TimingStat struct {
	Stat  string
	Value int64
}
