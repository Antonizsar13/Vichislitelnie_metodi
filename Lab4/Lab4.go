package main

import (
	"fmt"
	"math"
)

type Element struct {
	number float64
	degree int
}

func enterEquations() []Element {
	// var size int
	// fmt.Print("Введите количество элементов:")
	// fmt.Scanf("%d", &size)

	// elements := make([]Element, size)
	// fmt.Println("Введите элементы:")
	// for i := 0; i < size; i++ {
	// 	fmt.Printf("Число %d: ", i)
	// 	fmt.Scan(&elements[i].number)
	// 	fmt.Printf("Степерь числа %d: ", i)
	// 	fmt.Scan(&elements[i].degree)
	// }

	// elements := []Element{{1, 3},
	// 	{4, 1},
	// 	{-3, 0}}

	elements := []Element{{1, 3},
		{-1, 0},
		{2, 1}}

	// elements := []Element{{1, 2},
	// {-10000, 0}}

	return elements
}

func printEquations(equations []Element) {
	fmt.Println("\nУравнение имеет вид")
	for i, element := range equations {
		if element.number < 0 || i == 0 {
			fmt.Printf("%0.2f*x^%d", element.number, element.degree)
		} else {
			fmt.Printf("+%0.2f*x^%d", element.number, element.degree)
		}
	}
	fmt.Printf("=0")
}

func derivative(equations []Element) []Element {
	var equationsNew []Element
	for _, element := range equations {
		if element.degree > 0 {
			newElement := Element{
				number: element.number * float64(element.degree),
				degree: element.degree - 1,
			}
			equationsNew = append(equationsNew, newElement)
		}
	}
	return equationsNew
}

func calculateFunction(Xn float64, equations []Element) float64 {
	var answerd float64

	for _, element := range equations {
		answerd += (element.number * (math.Pow(Xn, float64(element.degree))))
	}

	return answerd
}

func main() {
	var Xn float64 = 1
	var E float64 = 1e-06

	equations := enterEquations()
	printEquations(equations)

	derivativeEq := derivative(equations)
	printEquations(derivativeEq)
	for i := Xn; i <= 1024; i *= 2 {
		var Xn2 float64 = i
		fmt.Printf("\nНачальное приближение = %f\n", Xn2)
		fmt.Println("X\t| F(Xn)\t| F′(Xn)\t| diff")
		var different float64 = E
		for math.Abs(different) >= E {

			answerd1 := calculateFunction(Xn2, equations)
			answerd2 := calculateFunction(Xn2, derivativeEq)

			different = answerd1 / answerd2
			fmt.Printf("%f\t| %f\t| %f\t| %f\n", Xn2, answerd1, answerd2, different)
			Xn2 = Xn2 - different
		}
		fmt.Printf("\nX = %f\n", Xn2)

	}

}
