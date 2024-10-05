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

	elements := []Element{
		{1, 2},
		{-16, 0}}

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

func calculateFunction(Xn float64, equations []Element) float64 {
	var answerd float64

	for _, element := range equations {
		answerd += (element.number * (math.Pow(Xn, float64(element.degree))))
	}

	return answerd
}

func F(x float64) float64 {
	answerd := math.Exp(x)
	return answerd
}

type Table struct {
	Xi  float64
	FXi float64
}

func aitkenInterpolation(table []Table, X float64, E float64) float64 {
	n := len(table)
	Q := make([][]float64, n)

	var prevValue float64
	for i := 0; i < n; i++ {
		Q[i] = make([]float64, n)
		Q[i][0] = table[i].FXi // Initial values are fx
	}

	for j := 1; j < n; j++ {
		for i := 0; i < n-j; i++ {
			Q[i][j] = ((X-table[i+j].Xi)*Q[i][j-1] - (X-table[i].Xi)*Q[i+1][j-1]) / (table[i+j].Xi - table[i].Xi)
		}

		if j > 1 {
			relativeError := math.Abs((Q[0][j]-prevValue))
			if relativeError < E {
				return Q[0][j] // Return the interpolated value if error is below tolerance
			}
		}
		prevValue = Q[0][j] // Update previous value for next iteration
	}
	

	return Q[0][n-1] // Return the interpolated value
}

func main() {
	// equations := enterEquations()
	// printEquations(equations)

	xStart, xEnd, step := 1.0, 2.0, 0.1
	XforCount := 1.43
	E := 0.001

	table := []Table{}
	for i := xStart; i <= xEnd+step; i += step {
		fx := F(i)
		table = append(table, Table{Xi: i, FXi: fx})
	}

	for _, row := range table {
		fmt.Printf("Xi: %.1f, F(Xi): %.4f\n", row.Xi, row.FXi)
	}

	var answerdL float64
	for i, value := range table {
		var timeItog float64 = 1
		for j, value2 := range table {
			if i != j {
				timeItog *= ((XforCount - value2.Xi) / (value.Xi - value2.Xi))
			}
		}
		timeItog *= value.FXi

		answerdL += timeItog
	}

	fmt.Printf("Ответ по Лагранжу: %.20f\n", answerdL)
	// fmt.Print(answerdL)

	answerdE := aitkenInterpolation(table, XforCount, E)
	fmt.Printf("Ответ по Эйткена: %.20f\n", answerdE)
	// fmt.Print(answerdE)

}
