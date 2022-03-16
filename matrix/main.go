package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Time before parallel process -> 2.456039577s
*/

const (
	matrixSize = 250
)

var (
	matrixA = [matrixSize][matrixSize]int{}
	matrixB = [matrixSize][matrixSize]int{}
	result  = [matrixSize][matrixSize]int{}
)

func generateRandomMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] += rand.Intn(10) - 5
		}
	}
}

func workoutRow(row int) {
	for col := 0; col < matrixSize; col++ {
		for i := 0; i < matrixSize; i++ {
			result[row][col] += matrixA[row][i] * matrixB[i][col]
		}
	}
}

func main() {
	fmt.Println("Working...")
	start := time.Now()
	for i := 0; i < 100; i++ {
		generateRandomMatrix(&matrixA)
		generateRandomMatrix(&matrixB)
		for row := 0; row < matrixSize; row++ {
			workoutRow(row)
		}
	}
	fmt.Printf("Processing took %s\n", time.Since(start))
}
