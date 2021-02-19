package ta

import (
	"fmt"
	"testing"
)

func TestATR(t *testing.T) {
	o := NewATR(5)
	fmt.Println(o.Write(&OHLC{High: 82.15, Low: 81.29, Close: 81.59}))
	fmt.Println(o.Write(&OHLC{High: 81.89, Low: 80.64, Close: 81.06}))
	fmt.Println(o.Write(&OHLC{High: 83.03, Low: 81.31, Close: 82.87}))
	fmt.Println(o.Write(&OHLC{High: 83.30, Low: 82.65, Close: 83.00}))
	fmt.Println(o.Write(&OHLC{High: 83.85, Low: 83.07, Close: 83.61}))
	fmt.Println(o.Write(&OHLC{High: 83.90, Low: 83.11, Close: 83.15}))
	fmt.Println(o.Write(&OHLC{High: 83.33, Low: 82.49, Close: 82.84}))
	fmt.Println(o.Write(&OHLC{High: 84.30, Low: 82.30, Close: 83.99}))
	fmt.Println(o.Write(&OHLC{High: 84.84, Low: 84.15, Close: 84.55}))
	fmt.Println(o.Write(&OHLC{High: 85.00, Low: 84.11, Close: 84.36}))
	fmt.Println(o.Write(&OHLC{High: 85.90, Low: 84.03, Close: 85.53}))
	fmt.Println(o.Write(&OHLC{High: 86.58, Low: 85.39, Close: 86.54}))
	fmt.Println(o.Write(&OHLC{High: 86.98, Low: 85.76, Close: 86.89}))
	fmt.Println(o.Write(&OHLC{High: 88.00, Low: 87.17, Close: 87.77}))
	fmt.Println(o.Write(&OHLC{High: 87.87, Low: 87.01, Close: 87.29}))
	t.Error("")
}
