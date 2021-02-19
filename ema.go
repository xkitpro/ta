package ta

type ema struct {
	p float64
	a float64
}

func NewEMA(n int) *ema {
	return &ema{
		a: 2. / (float64(n) + 1),
	}
}

func (i *ema) Write(v float64) float64 {
	if i.p == 0. {
		i.p = v
		return v
	}

	i.p = i.a*(v-i.p) + i.p

	return i.p
}
