package main

import (
	"math"
)

type ZeroDivisionError struct {
}

func (zde ZeroDivisionError) Error() string {
	return "Ettempt to divide by zero"
}

func swap[T any](first *[]T, second *[]T) {
	*first, *second = *second, *first
}

func jacobi(matrix *[][]float64, vector []float64, size uint64, iterations uint64) ([]float64, error) {
	var x []float64 = make([]float64, size)
	var x_next []float64 = make([]float64, size)
	for iteration := uint64(0); iteration < iterations; iteration++ {
		for index := uint64(0); index < size; index++ {
			if math.Abs((*matrix)[index][index]) < 1e-12 {
				return x, ZeroDivisionError{}
			}
			var sum float64 = vector[index]
			for i := uint64(0); i < size; i++ {
				if i == index {
					continue
				}
				sum -= (*matrix)[index][i] * x[i]
			}
			x_next[index] = sum / (*matrix)[index][index]
		}
		swap(&x, &x_next)
	}
	return x, nil
}

func zeidel(matrix *[][]float64, vector []float64, size uint64, iterations uint64) ([]float64, error) {
	var x []float64 = make([]float64, size)
	return x, nil
}
