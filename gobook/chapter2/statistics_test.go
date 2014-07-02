package chapter2_test

import (
	"fmt"
	"gobook/chapter2/statistics"
	"testing"
)

func TestAll(t *testing.T) {
	var b int32 = 6
	var d byte = 5
	c := b / int32(d)
	fmt.Printf("%d\n", c)
	a := float64(32) / float64(b)
	fmt.Printf("%5.02f\n", a)
	a = float64(32) / 8
	fmt.Printf("%5.02f\n", a)
	numbers := []float64{3.0, 4, 1, 11, 6}
	result := main.GetStatsFor(numbers)
	fmt.Printf("sum = %.02f, mdian = %.02f\n", result.GetMean(), result.GetMdian)
}
