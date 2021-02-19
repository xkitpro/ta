package ta

import (
	"math"
)

type hma struct {
	ma1    Indicator
	ma2    Indicator
	masqrt Indicator
}

func NewHMA(n int) Indicator {
	return &hma{
		ma1:    NewWMA(n),
		ma2:    NewWMA(n / 2),
		masqrt: NewWMA(int(math.Floor(math.Sqrt(float64(n))))),
	}
}

func (i *hma) Write(v float64) float64 {
	wma := i.ma1.Write(v)
	wma2 := i.ma2.Write(v)

	if math.IsNaN(wma) {
		return math.NaN()
	}

	return i.masqrt.Write(2*wma2 - wma)
}
