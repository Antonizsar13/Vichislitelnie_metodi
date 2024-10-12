package main

import (
	"fmt"
	"math"
)

func f1(x float64) float64 {
	return math.Sin(1 / math.Pow(x, 2))
}
func f2(x float64) float64 {
	return x * math.Sin(1/math.Pow(x, 3))
}

func gaussQuadrature(a, b float64, n int, F func(float64) float64) float64 {
	t := map[int][]float64{
		1: {0},
		2: {-0.57735026918962576450, 0.57735026918962576450},
		3: {-0.77459666924148337703, 0.77459666924148337703, 0},
		4: {
			-0.33998104358485626480, 0.33998104358485626480,
			-0.86113631159405257522, 0.86113631159405257522,
		},
		5: {
			-0.53846931010568309103, 0.53846931010568309103,
			-0.90617984593866399279, 0.90617984593866399279, 0,
		},
		6: {
			-0.23861918608319690863, 0.23861918608319690863,
			-0.66120938646626451366, 0.66120938646626451366,
			-0.93246951420315202781, 0.93246951420315202781,
		},
		7: {
			-0.40584515137739716690, 0.40584515137739716690,
			-0.74153118559939443986, 0.74153118559939443986,
			-0.94910791234275852452, 0.94910791234275852452, 0,
		},
		8: {
			-0.18343464249564980493, 0.18343464249564980493,
			-0.52553240991632898581, 0.52553240991632898581,
			-0.79666647741362673959, 0.79666647741362673959,
			-0.96028985649753623168, 0.96028985649753623168,
		},
	}
	w := map[int][]float64{
		1: {2},
		2: {1},
		3: {5.0 / 9.0, 8.0 / 9.0},
		4: {0.65214515486254614262, 0.34785484513745385737},
		5: {0.47862867049936646804, 0.23692688505618908751, 0.56888888888888888889},
		6: {0.46791393457269104738, 0.36076157304813860756, 0.17132449237917034504},
		7: {0.38183005050511894495, 0.27970539148927666790, 0.12948496616886969327, 0.41795918367346938775},
		8: {0.36268378337836198296, 0.31370664587788728733, 0.22238103445337447054, 0.10122853629037625915},
	}

	integral := 0.0
	timeI := 0
	for i := 0; i < n; i++ {
		fmt.Print(timeI)
		integral += w[n][timeI] * F(0.5*(a+b+t[n][i]*(b-a)))
		if i%2 != 0 {
			timeI++
		}

	}
	c := (b - a) / 2
	integral *= c

	return integral
}

func main() {
	n := 4
	for i := 0; i < 2; i++ {
		var F func(float64) float64
		var a float64
		var b float64
		switch i {
		case 0:
			F = f1
			a = 1
			b = 2
		case 1:
			F = f2
			a = 1
			b = 2
		}
		fmt.Printf("Уравнение %d\n", i+1)
		result := gaussQuadrature(a, b, n, F)
		fmt.Printf("Приближенное значение интеграла от %.2f до %.2f при n=%d: %f\n", a, b, n, result)
		result1 := gaussQuadrature(a, b, n+1, F)
		fmt.Printf("Приближенное значение интеграла от %.2f до %.2f при n=%d: %f Разница = %f\n", a, b, n+1, result1, math.Abs(result1-result))
		result2 := gaussQuadrature(a, b, n+2, F)
		fmt.Printf("Приближенное значение интеграла от %.2f до %.2f при n=%d: %f Разница = %f\n", a, b, n+2, result2, math.Abs(result2-result1))

	}

}
