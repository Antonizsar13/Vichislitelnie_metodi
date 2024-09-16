package main

import (
	"fmt"
	"math"
)

func craeteMatrix() [][]float32 {
	var size int
	fmt.Printf("Введите размер матрицы: ")
	fmt.Scanf("%d\n", &size)

	fmt.Println("Заполнение матрицы")
	var matrix = make([][]float32, size)
	for i := range matrix {
		matrix[i] = make([]float32, size+1)
		for j := range matrix[i] {
			fmt.Printf("[%d][%d]=", i, j)
			fmt.Scanf("%f\n", &matrix[i][j])
		}
	}

	// matrix := [][]float32{
	// 	{2, 3, -1, 7},
	// 	{1, -1, 6, 14},
	// 	{6, -2, 1, 11},
	// }

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

func choosingMainElement(matrix *[][]float32, column int) {
	maxValue := (*matrix)[column][column]
	maxLine := column
	for i := column; i < len((*matrix)); i++ {
		if float32(math.Abs(float64((*matrix)[i][column]))) > maxValue {
			maxValue = (*matrix)[i][column]
			maxLine = i
		}
	}

	if maxLine != column {
		for i := column; i < len((*matrix)[0]); i++ {
			time := (*matrix)[maxLine][i]
			(*matrix)[maxLine][i] = (*matrix)[column][i]
			(*matrix)[column][i] = time
		}
	}
}

func dataNormalization(matrix *[][]float32, column int) {
	divider := (*matrix)[column][column]
	for i := column; i < len((*matrix)[0]); i++ {
		(*matrix)[column][i] = (*matrix)[column][i] / divider
	}
}

func excludingElements(matrix *[][]float32, column int) {

	for i := column + 1; i < len((*matrix)); i++ {
		factor := (*matrix)[i][column]
		(*matrix)[i][column] = 0
		for j := column + 1; j < len((*matrix)[i]); j++ {
			(*matrix)[i][j] = (*matrix)[i][j] - factor*(*matrix)[column][j]
		}
	}
}

func reverseStroke(matrix *[][]float32, step int) {
	var sum float32 = 0
	len := len(*matrix)
	for i := 0; i < step; i++ {
		sum += (*matrix)[len-1-step][len-1-i] * (*matrix)[len-1-i][len]
		(*matrix)[len-1-step][len-1-i] = 0;
	}
	(*matrix)[len-1-step][len] = ((*matrix)[len-1-step][len]) - sum
}

func main() {
	var matrix = craeteMatrix()
	outputMatrix(matrix)
	for step := 0; step < len(matrix); step++ {
		choosingMainElement(&matrix, step)
		outputMatrix(matrix)
		dataNormalization(&matrix, step)
		outputMatrix(matrix)
		excludingElements(&matrix, step)
		outputMatrix(matrix)
	}
	for step := 1; step < len(matrix); step++ {
		reverseStroke(&matrix, step)
		outputMatrix(matrix)
	}

}
