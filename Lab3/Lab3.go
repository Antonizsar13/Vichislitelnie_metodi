package main

import (
	"fmt"
	"math"
)

func craeteMatrix() [][]float32 {
	// var size int
	// fmt.Printf("Введите размер частот матрицы: ")
	// fmt.Scanf("%d\n", &size)
	//
	// fmt.Println("Заполнение матрицы")
	// var matrix = make([][]float32, size)
	// for i := range matrix {
	// 	matrix[i] = make([]float32, size+1)
	// 	for j := range matrix[i] {
	// 		fmt.Printf("[%d][%d]=", i, j)
	// 		fmt.Scanf("%f\n", &matrix[i][j])
	// 	}
	// }

	matrix := [][]float32{
		{10, 1, 1, 12},
		{2, 1, 1, 13},
		{2, 2, 1, 14},
	}

	return matrix
}

func outputMatrix(matrix [][]float32) {
	fmt.Println("Вывод матрицы")
	for _, part := range matrix {
		for _, elemet := range part {
			// fmt.Printf("[%d][%d]=%f", i, j, elemet)
			fmt.Print(elemet, "\t")
		}
		fmt.Println()
	}
}

func countMatrix(matrix [][]float32, maxIterations int)  []float32 {
	for i := range matrix {
		divider := matrix[i][i]
		for j := range matrix[i] {
			matrix[i][j] /= divider
		}
	}

	approximation := make([]float32, len(matrix))

	for i := 0; i < maxIterations; i++ {
		newApproximation := make([]float32, len(matrix))
		for i := range matrix {
			sum := float32(0)
			for j := range matrix {
				if i != j {
					sum += matrix[i][j] * approximation[j]
				}
			}


			newApproximation[i] = matrix[i][len(matrix)] - sum

		}

		coincidence := true
		for i := range approximation {
			if math.Abs(float64(approximation[i]-newApproximation[i])) > 1e-6 {
				coincidence = false
				break
			}
		}
		approximation = newApproximation
		if coincidence {
			fmt.Println("Прошло итераций: ", i)
			break
		}
	}
	fmt.Println("Прошла все заданные итерации")
	return approximation

}

func main() {
	var matrix = craeteMatrix()
	fmt.Println("Вывод матрицы")
	outputMatrix(matrix)

	answerd := countMatrix(matrix, 1000)

	for i, value := range answerd {
		fmt.Printf("x[%d] = %f\n", i, value)
	}

}
