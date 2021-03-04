package ta

import "math"

type sma struct {
	p float64
	f float64
	n int
	c int
	m []float64
}

func NewSMA(n int) Indicator {
	return &sma{
		n: n,
		f: float64(n),
		c: 1,
		m: make([]float64, n, n),
	}
}

func (o *sma) Write(v float64) float64 {
	p := o.m[0]
	o.m = o.m[1:]
	o.m = append(o.m, v)

	o.p = o.p + (v-p)/o.f

	if o.c < o.n {
		o.c++
		return math.NaN()
	}

	return o.p
}
