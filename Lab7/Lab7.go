package main

import (
	"fmt"
	"math"
)

type Interval struct {
	start float64
	end   float64
}

func F1(x float64) float64 {
	return math.Sin(1 / math.Pow(x, 2))
}
func F2(x float64) float64 {
	return x * math.Sin(1/math.Pow(x, 3))
}
func F3(x float64) float64 {
	return math.Exp(x) * math.Cos(1/math.Pow(x, 3))
}
func F4(x float64) float64 {
	return (math.Exp(x) / x) * math.Cos(1/math.Pow(x, 3))
}
func F5(x float64) float64 {
	return (math.Pow(2, x) / x) * math.Log(x)
}

func main() {

	intervals := []Interval{{1, 2}, {1, 2}, {1, 2}, {2, 4}, {3, 4}}
	var F func(float64) float64

	for numberF, intinterval := range intervals {
		switch numberF {
		case 0:
			F = F1
		case 1:
			F = F2
		case 2:
			F = F3
		case 3:
			F = F4
		case 4:
			F = F5
		}

		e := 0.0001
		eR := e
		n := 5
		var answerd float64
		var answerd2 float64
		for eR >= e {

			h := (intinterval.end - intinterval.start) / float64(n)
			backValue := intinterval.start
			answerd = 0
			for i := intinterval.start; i < intinterval.end; i += h {
				x := (backValue + i) / 2
				answerd += h * F(x)
				backValue = i
			}
			n *= 2
			h = (intinterval.end - intinterval.start) / float64(n)
			backValue = intinterval.start
			answerd2 = 0
			for i := intinterval.start; i < intinterval.end; i += h {
				x := (backValue + i) / 2
				answerd2 += h * F(x)
				backValue = i
			}

			eR = math.Abs(answerd2-answerd) / 3.0

		}

		fmt.Printf("Ответ для %d-го: %f. Точность: %f\n", numberF, answerd2, eR)
	}
}
