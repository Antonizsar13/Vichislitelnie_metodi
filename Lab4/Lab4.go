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
	var size int
	fmt.Print("Введите количество элементов:")
	fmt.Scanf("%d", &size)

	elements := make([]Element, size)
	fmt.Println("Введите элементы:")
	for i := 0; i < size; i++ {
		fmt.Printf("Число %d: ", i)
		fmt.Scan(&elements[i].number)
		fmt.Printf("Степерь числа %d: ", i)
		fmt.Scan(&elements[i].degree)
	}

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
	for i, element := range equations {
		if element.degree == 0 {
			copy(equations[i:], equations[i+1:])
			equations = equations[:len(equations)-1]
		} else {
			equations[i].number *= float64(element.degree)
			equations[i].degree--
		}
	}
	return equations
}

func calculateFunction(Xn float64, equations []Element) float64 {
	var answerd float64;

	for i, element := range equations {
		answerd += (element.number * (math.Pow(Xn, float64(element.degree))))
	}
	
	return answerd
}

func main() {
	var Xn float64 = 2
	var E float64 = 0.001

	equations := enterEquations()
	printEquations(equations)

	derivativeEq := derivative(equations)
	printEquations(derivativeEq)

	for i := Xn; i <= 1024; i *= 2 {
		var Xn2 float32 = i;
			fmt.Println("\nНачальное приближение = %f", Xn2)
		for Xn2 

	}

}
