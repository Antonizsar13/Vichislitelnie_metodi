package main

import (
	"fmt"
	"math"
)

func eulerMethod(f func(float64, float64) float64, u0 float64, h float64, xStart float64, xEnd float64) []float64 {
	x := xStart
	u := u0
	uValues := []float64{u}

	for x < xEnd {
		u += h * f(x, u)    
		x += h     
		uValues = append(uValues, u) 
	}

	return uValues
}

func simmetrMethod(f func(float64, float64) float64, u0 float64, h float64, xStart float64, xEnd float64) []float64 {
	x := xStart
	u := u0
	uValues := []float64{u}

	for x < xEnd {
		uPrev := uValues[len(uValues)-1]
		uNow := uPrev + h*f(x-h, uPrev)
		u := uPrev + h/2*(f(x-h, uPrev)+f(x+h, uNow))
		x += h
		uValues = append(uValues, u)
	}

	return uValues
}

func accurate(exactFunc func(float64) float64, h float64, xStart float64, xEnd float64) []float64 {
	x := xStart
	accurateValues := []float64{}

	for x <= xEnd {
		accurateValues = append(accurateValues, exactFunc(x))
		x += h
	}

	return accurateValues
}

func f1(x float64, u float64) float64 {
	return x*x + 3*u
}
func f2(x float64, u float64) float64 {
	return 2*u + 4*x
}
func f3(x float64, u float64) float64 {
	return 2*u + math.Exp(x)
}

func exactSolution1(x float64) float64 {
	return (56.0/27.0)*math.Exp(3*x) - (1.0/3.0)*x*x - (2.0/9.0)*x - (2.0 / 27.0)
}

func exactSolution2(x float64) float64 {
	return -2*x + 2*math.Exp(2*x) - 1
}

func exactSolution3(x float64) float64 {
	return (math.Exp(x) - (math.Exp(x) - 1)) 
}

func main() {
	hValues := []float64{0.1, 0.01}

	fmt.Println("Уравнение 1:")
	for _, h := range hValues {
		uValues := eulerMethod(f1, 2.0, h, 0, 2)
		uValues2 := simmetrMethod(f1, 2.0, h, 0, 2)
		uValues3 := accurate(exactSolution1, h, 0, 2)

		fmt.Printf("Эйлер h = %.2f: %v\n\n", h, uValues)
		fmt.Printf("Симметричный h = %.2f: %v\n\n", h, uValues2)
		fmt.Printf("Точный h = %.2f: %v\n\n", h, uValues3)

	}

	// fmt.Println("Уравнение 2:")
	// for _, h := range hValues {
	// 	uValues := eulerMethod(f2, 1.0, h, 0, 1)
	// 	uValues2 := simmetrMethod(f2, 1.0, h, 0, 1)
	// 	fmt.Printf("Эйлер h = %.2f: %v\n", h, uValues)
	// 	fmt.Printf("Симметричный h = %.2f: %v\n", h, uValues2)

	// 	// Вывод точного решения
	// 	exactValues := []float64{}
	// 	for i := 0.0; float64(i) <= 1; i += h {
	// 		exactValues = append(exactValues, exactSolution2(float64(i)*h))
	// 	}
	// 	fmt.Printf("Точные решения: h = %.2f: %v\n", h, exactValues)
	// }

	// fmt.Println("Уравнение 3:")
	// for _, h := range hValues {
	// 	uValues := eulerMethod(f3, 0.0, h, 0, 1)
	// 	uValues2 := simmetrMethod(f3, 0.0, h, 0, 1)
	// 	fmt.Printf("Эйлер h = %.2f: %v\n", h, uValues)
	// 	fmt.Printf("Симметричный h = %.2f: %v\n", h, uValues2)

	// 	// Вывод точного решения
	// 	exactValues := []float64{}
	// 	for i := 0.0; float64(i) <= 1; i += h {
	// 		exactValues = append(exactValues, exactSolution3(float64(i)*h))
	// 	}
	// 	fmt.Printf("Точные решения: h = %.2f: %v\n", h, exactValues)
	// }

}
