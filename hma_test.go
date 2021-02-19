package ta

import (
	"fmt"
	"testing"
)

func TestHMA(t *testing.T) {
	o := NewHMA(5)
	fmt.Println(o.Write(81.59))
	fmt.Println(o.Write(81.06))
	fmt.Println(o.Write(82.87))
	fmt.Println(o.Write(83.00))
	fmt.Println(o.Write(83.61))
	fmt.Println(o.Write(83.15))
	fmt.Println(o.Write(82.84))
	fmt.Println(o.Write(83.99))
	fmt.Println(o.Write(84.55))
	fmt.Println(o.Write(84.36))
	fmt.Println(o.Write(85.53))
	fmt.Println(o.Write(86.54))
	fmt.Println(o.Write(86.89))
	fmt.Println(o.Write(87.77))
	fmt.Println(o.Write(87.29))
	t.Error("")
}
