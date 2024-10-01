package main

import (
	"fmt"
	"math"
	"sort"
)

type Element struct {
	number float64
	degree int
}
type ByDegree []Element

func (a ByDegree) Len() int           { return len(a) }
func (a ByDegree) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDegree) Less(i, j int) bool { return a[i].degree > a[j].degree }

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

	// elements := []Element{{1, 3},
	// 	{-1, 0},
	// 	{2, 1}}

	// elements := []Element{{1, 2},
	// {-10000, 0}}

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

type Interval struct {
	start float64
	end   float64
}

func findIntervals(equations []Element) Interval {
	sort.Sort(ByDegree(equations))
	maxA := math.Abs(equations[1].number)
	maxB := math.Abs(equations[0].number)
	for i, value := range equations {
		if i > 1 {
			if maxA < math.Abs(value.number) {
				maxA = math.Abs(value.number)
			}
		}
		if i < (len(equations) - 1) {
			if maxB < math.Abs(value.number) {
				maxB = math.Abs(value.number)
			}
		}
	}

	var interval Interval
	start := 1 / (1 + (maxB / math.Abs(equations[len(equations)-1].number)))
	end := 1 + (maxA / math.Abs(equations[0].number))
	interval.start, interval.end = start, end

	return interval
}

func main() {

	equations := enterEquations()
	printEquations(equations)

	interval := findIntervals(equations)

	fmt.Print(interval)

	// var a, b float64 = 1, 5
	// var c float64
	// var E float64 = 0.01

	// Fa := calculateFunction(a, equations)
	// Fb := calculateFunction(b, equations)

	// if Fa*Fb >= 0 {
	// 	fmt.Println("\nФункция не меняет знак на данном интервале. Выберите другие точки.")
	// 	return
	// }

	// var index int = 0
	// // Fa := calculateFunction(a, equations)
	// for math.Abs(b-a) > E {
	// 	index++
	// 	c = (a + b) / 2
	// 	Fc := calculateFunction(c, equations)

	// 	if Fa*Fc < 0 {
	// 		b = c
	// 		Fb = Fc
	// 	} else {
	// 		a = c
	// 		Fa = Fc
	// 	}
	// }

	// fmt.Printf("\nОтвет: %f. Точность: %f. Количество итераций: %d", c, E, index)
}
