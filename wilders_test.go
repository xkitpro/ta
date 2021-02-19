package ta

import (
	"fmt"
	"testing"
)

func TestWilders(t *testing.T) {
	o := NewWilders(5)
	fmt.Println(o.Write(81.59))
	fmt.Println(o.Write(81.06))
	fmt.Println(o.Write(82.87))
	fmt.Println(o.Write(83.00))
	fmt.Println(o.Write(83.61))
	fmt.Println(o.Write(83.15))
	fmt.Println(o.Write(82.84))
	fmt.Println(o.Write(83.99))
	fmt.Println(o.Write(84.55))
	t.Error("")
}
