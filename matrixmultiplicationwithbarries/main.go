package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Time before parallel process -> 2.456039577s
	Time before parallel process -> 373.556584ms
*/

const (
	matrixSize = 250
)

var (
	matrixA      = [matrixSize][matrixSize]int{}
	matrixB      = [matrixSize][matrixSize]int{}
	result       = [matrixSize][matrixSize]int{}
	workStart    = NewBarrier(matrixSize + 1) //matrixSize of threads plus the main thread
	workComplete = NewBarrier(matrixSize + 1)
)

func generateRandomMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] += rand.Intn(10) - 5
		}
	}
}

func workoutRow(row int) {
	for {
		workStart.Wait()
		for col := 0; col < matrixSize; col++ {
			for i := 0; i < matrixSize; i++ {
				result[row][col] += matrixA[row][i] * matrixB[i][col]
			}
		}
		workComplete.Wait()
	}
}

func main() {
	fmt.Println("Working...")

	for row := 0; row < matrixSize; row++ {
		go workoutRow(row)
	}
	start := time.Now()
	for i := 0; i < 100; i++ {
		generateRandomMatrix(&matrixA)
		generateRandomMatrix(&matrixB)
		workStart.Wait()
		workComplete.Wait()
	}
	fmt.Printf("Processing took %s\n", time.Since(start))
}
