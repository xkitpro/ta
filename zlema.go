package ta

import "math"

type zlema struct {
	period int
	lag    int
	per    float64
	c      int
	buf    []float64
	prev   float64
}

func NewZLEMA(n int) Indicator {
	lag := (n - 1) / 2
	return &zlema{
		period: n,
		lag:    lag,
		per:    2. / (float64(n) + 1),
		buf:    make([]float64, lag, lag),
	}
}

func (i *zlema) Write(v float64) float64 {
	p := i.buf[0]
	i.buf = i.buf[1:]
	i.buf = append(i.buf, v)
	if i.c < i.lag {
		i.prev = v
		i.c++
		if i.c == i.lag {
			return i.prev
		}
		return math.NaN()
	}

	i.prev = ((v+(v-p))-i.prev)*i.per + i.prev

	return i.prev
}
