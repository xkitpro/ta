package ta

type OHLC struct {
	Open  float64
	High  float64
	Low   float64
	Close float64
}

func (t *OHLC) OC2() float64 {
	return (t.Open + t.Close) / 2
}

func (t *OHLC) HL2() float64 {
	return (t.High + t.Low) / 2
}

func (t *OHLC) OHLC4() float64 {
	return (t.Open + t.High + t.Low + t.Close) / 4
}
