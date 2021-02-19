package ta

import (
	"fmt"
	"testing"
)

func TestTR(t *testing.T) {
	o := NewTR()
	fmt.Println(o.Write(&OHLC{High: 82.15, Low: 81.29, Close: 81.59}))
	fmt.Println(o.Write(&OHLC{High: 81.89, Low: 80.64, Close: 81.06}))
	fmt.Println(o.Write(&OHLC{High: 83.03, Low: 81.31, Close: 82.87}))
	fmt.Println(o.Write(&OHLC{High: 83.30, Low: 82.65, Close: 83.00}))
	fmt.Println(o.Write(&OHLC{High: 83.85, Low: 83.07, Close: 83.61}))
	t.Error("")
}
