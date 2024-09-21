package main

import (
	"fmt"
)

func craeteMatrix() [][]float32 {
	// var size int
	// fmt.Printf("Введите размер матрицы: ")
	// fmt.Scanf("%d\n", &size)

	// fmt.Println("Заполнение матрицы")
	// var matrix = make([][]float32, size)
	// for i := range matrix {
	// 	matrix[i] = make([]float32, size+1)
	// 	var start int
	// 	var end int
	// 	if i == 0 {
	// 		start = 0
	// 		end = 1
	// 	} else if i == size-1 {
	// 		start = size - 2
	// 		end = size - 1
	// 	} else {
	// 		start = i - 1
	// 		end = i + 1
	// 	}
	// 	for j := start; j <= end; j++ {
	// 		fmt.Printf("[%d][%d]=", i, j)
	// 		fmt.Scanf("%f\n", &matrix[i][j])
	// 	}

	// }
	// fmt.Println("Корни матрицы")
	// for i := 0; i < size; i++ {
	// 	fmt.Printf("[%d]=", i)
	// 	fmt.Scanf("%f\n", &matrix[i][size])
	// }

	matrix := [][]float32{
		{2, -1, 0, 0, 0, -25},
		{-3, 8, -1, 0, 0, 72},
		{0, -5, 12, 2, 0, -69},
		{0, 0, -6, 18, -4, -156},
		{0, 0, 0, -5, 10, 20},
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

func countMatrix(matrix [][]float32) []float32 {
	size := len(matrix)
	var a = make([]float32, size)
	var b = make([]float32, size)
	for i := range matrix {
		var y float32
		if i == 0 {
			y = matrix[i][0]
			a[i] = -(matrix[i][1]) / y
			b[i] = (matrix[i][size]) / y
		} else if i == (size - 1) {
			y = matrix[i][i] + matrix[i][i-1]*a[i-1]
			// a[i] = -(matrix[i][1]) / y
			b[i] = (matrix[i][size] - matrix[i][i-1]*b[i-1]) / y
		} else {
			y = matrix[i][i] + matrix[i][i-1]*a[i-1]
			a[i] = -(matrix[i][i+1]) / y
			b[i] = (matrix[i][size] - matrix[i][i-1]*b[i-1]) / y
		}
	}

	var answerds = make([]float32, size)
	for i := (size - 1); i >= 0; i-- {
		if i == (size - 1) {
			answerds[i] = b[i]
		} else {
			answerds[i] = a[i]*answerds[i+1] + b[i]
		}
	}

	return answerds
}

func main() {
	var matrix = craeteMatrix()
	outputMatrix(matrix)

	answerd := countMatrix(matrix)

	for i, value := range answerd {
		fmt.Printf("x[%d] = %f\n", i, value)
	}

}
