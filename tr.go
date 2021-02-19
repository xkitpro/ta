package ta

import "math"

type tr struct {
	p *OHLC
}

func NewTR() *tr {
	return &tr{}
}

func (i *tr) Write(c *OHLC) float64 {
	if i.p == nil {
		i.p = c
		return c.High - c.Low
	}

	v := math.Max(
		c.High-c.Low,
		math.Max(
			math.Abs(c.High-i.p.Close),
			math.Abs(c.Low-i.p.Close),
		),
	)

	i.p = c

	return v
}
