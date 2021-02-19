package ta

import (
	"math"
)

type wma struct {
	p  float64 // sum
	w  float64 // weights
	n  float64 // period
	c  float64
	m  []float64
	ws float64 // weight sum
}

func NewWMA(n int) Indicator {
	return &wma{
		n: float64(n),
		w: float64(n*(n+1)) / 2,
		c: 1.,
		m: make([]float64, n, n),
	}
}

func (i *wma) Write(v float64) float64 {
	p := i.m[0]
	i.m = i.m[1:]
	i.m = append(i.m, v)

	if i.c < i.n {
		i.p += v
		i.ws += v * i.c

		i.c++
		return math.NaN()
	}

	i.p -= p

	i.ws += v * i.n
	i.p += v

	r := i.ws / i.w

	i.ws -= i.p

	return r
}
