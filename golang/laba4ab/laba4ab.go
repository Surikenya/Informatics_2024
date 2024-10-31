package laba4ab

import (
	"fmt"
	"math"
)

func Calculate(x, a float64) float64 {
	y := math.Tan(math.Pow(math.Log10(a+x), 3)) / math.Pow((a+x), 2.0/7.0)
	return y
}

func CalculateB(xValues []float64, a float64) []float64 {
	var spb = []float64{}
	for _, x := range xValues {
		spb = append(spb, Calculate(x, a))
	}
	return spb
}

func CalculateA(xn, xk, deltaX, a float64) []float64 {
	var spa = []float64{}
	for x := xn; x <= xk; x += deltaX {
		spa = append(spa, Calculate(x, a))
	}
	return spa
}

func Readylaba4ab() {
	fmt.Print("A\n")
	fmt.Print(CalculateA(1.08, 1.88, 0.16, 2), "\n")
	fmt.Print("B\n")
	var xValues = []float64{1.16, 1.35, 1.48, 1.52, 1.96}
	fmt.Print(CalculateB(xValues, 2.0))