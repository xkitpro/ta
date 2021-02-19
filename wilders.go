package ta

import "math"

type wilders struct {
	c int
	n int
	s float64
	i Indicator
}

func NewWilders(n int) Indicator {
	return &wilders{
		n: n,
		c: 0,
		i: NewEMA(2*n - 1),
	}
}

func (i *wilders) Write(v float64) float64 {
	if i.c < i.n {
		i.c++
		i.s += v

		if i.c == i.n {
			return i.i.Write(i.s / float64(i.n))
		}

		return math.NaN()
	}

	return i.i.Write(v)
}
