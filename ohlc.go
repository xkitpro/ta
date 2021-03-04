package ta

type F func(OHLC) float64

func OC2(c OHLC) float64 {
	return c.OC2()
}

func HL2(c OHLC) float64 {
	return c.HL2()
}

func OHLC4(c OHLC) float64 {
	return c.OHLC4()
}

func Open(c OHLC) float64 {
	return c.Open
}

func Close(c OHLC) float64 {
	return c.Close
}

func High(c OHLC) float64 {
	return c.High
}

func Low(c OHLC) float64 {
	return c.Low
}

func HLC3(c OHLC) float64 {
	return c.HLC3()
}

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

func (t *OHLC) HLC3() float64 {
	return (t.High + t.Low + t.Close) / 3
}
