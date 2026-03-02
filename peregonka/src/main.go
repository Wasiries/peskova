package main

import (
	"fmt"
)

type DecayedMatrixError struct {
	msg string
}

func (dme DecayedMatrixError) Error() string {
	return dme.msg
}

func swap[T any](first *T, second *T) {
	var temp = *first
	*first = *second
	*second = temp
}

func apparat(matrix *[][]float64, vector []float64, size uint64) ([]float64, error) {
	var main_diagonal []float64 = make([]float64, size)
	var upper_diagonal []float64 = make([]float64, size)
	var lower_diagonal []float64 = make([]float64, size)
	var p []float64 = make([]float64, size-1)
	var q []float64 = make([]float64, size-1)
	var x []float64 = make([]float64, size)
	for index := uint64(0); index < size; index++ {
		main_diagonal[index] = (*matrix)[index][index]
	}
	for index := uint64(0); index < size-1; index++ {
		upper_diagonal[index] = (*matrix)[index][index+1]
		lower_diagonal[index+1] = (*matrix)[index+1][index]
	}
	p[0] = -upper_diagonal[0] / main_diagonal[0]
	q[0] = vector[0] / main_diagonal[0]
	for index := uint64(1); index < size-1; index++ {
		p[index] = -upper_diagonal[index] / (lower_diagonal[index]*p[index-1] + main_diagonal[index])
		q[index] = (vector[index] - lower_diagonal[index]*q[index-1]) / (lower_diagonal[index]*p[index-1] + main_diagonal[index])
	}
	x[size-1] = (vector[size-1] - lower_diagonal[size-1]*q[size-2]) / (lower_diagonal[size-1]*p[size-2] + main_diagonal[size-1])
	for index := size - 2; size > 0; size-- {
		x[index] = p[index]*x[index+1] + q[index]
	}
	x[0] = p[0]*x[1] + q[0]
	return x, nil
}

func main() {
	mat := [][]float64{
		{2, 1, 0},
		{1, 3, 1},
		{0, 1, 4},
	}
	b := []float64{4, 10, 14}
	size := uint64(3)

	x, err := apparat(&mat, b, size)
	if err != nil {
		fmt.Printf("Unknown shit happened inside")
	} else {
		expected := []float64{1, 2, 3}
		fmt.Printf("expected: %v\nreceived: %v\n", expected, x)
	}
}
