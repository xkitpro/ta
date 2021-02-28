package ta

import (
	"fmt"
	"math"
	"testing"
)

var input = []float64{
	81.59,
	81.06,
	82.87,
	83.00,
	83.61,
	83.15,
	82.84,
	83.99,
	84.55,
	84.36,
	85.53,
	86.54,
	86.89,
	87.77,
	87.29,
}

var result = []float64{
	math.NaN(),
	81.06,
	82.09,
	83.04,
	83.48,
	83.42,
	82.997,
	93.59,
	84.48,
	84.56,
	85.21,
	86.38,
	87.00,
	87.67,
	87.68,
}

func TestZLEMA(t *testing.T) {
	ma := NewZLEMA(5)
	for _, i := range input {
		fmt.Println(ma.Write(i))
	}
	t.Error("")
}
