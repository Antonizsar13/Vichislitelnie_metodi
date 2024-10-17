package main

import (
	"fmt"
	"math"
)

// Функция для метода Эйлера
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

func f1(x float64, u float64) float64 {
	return x*x + 3*u
}

func f2(x float64, u float64) float64 {
	return 2*u + 4*x
}

func f3(x float64, u float64) float64 {
	return 2*u + math.Exp(x)
}

func f4(x float64, u float64) float64 {
	return 3*x*x*u+x*x*math.Exp(math.Pow(x,3))
}


func main() {
	hValues := []float64{0.1, 0.01}

	fmt.Println("Уравнение 1:")
	for _, h := range hValues {
		uValues := eulerMethod(f1, 2.0, h, 0, 2)
		fmt.Printf("h = %.2f: %v\n", h, uValues)
	}

	fmt.Println("Уравнение 2:")
	for _, h := range hValues {
		uValues := eulerMethod(f2, 1.0, h, 0, 1)
		fmt.Printf("h = %.2f: %v\n", h, uValues)
	}

	fmt.Println("Уравнение 3:")
	for _, h := range hValues {
		uValues := eulerMethod(f3, 0.0, h, 0, 1)
		fmt.Printf("h = %.2f: %v\n", h, uValues)
	}

	fmt.Println("Уравнение 4:")
	for _, h := range hValues {
		uValues := eulerMethod(f4, 0.0, h, 0, 1)
		fmt.Printf("h = %.2f: %v\n", h, uValues)
	}
}
