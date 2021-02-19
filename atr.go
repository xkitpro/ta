package ta

type atr struct {
	tr *tr
	i  Indicator
}

func NewATR(n int) *atr {
	return &atr{
		tr: NewTR(),
		i:  NewWilders(n),
	}
}

func (i *atr) Write(c *OHLC) float64 {
	return i.i.Write(i.tr.Write(c))
}
