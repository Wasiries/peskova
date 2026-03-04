package slau

import (
	"math"
)

type ZeroDivisionError struct {
}

type MatrixNotMatch struct {
}

type ImpossibleMatrixStructure struct {
}

type VectorNotMatch struct {
}

func (zde ZeroDivisionError) Error() string {
	return "Ettempt to divide by zero"
}

func (mnm MatrixNotMatch) Error() string {
	return "Declared size and matrix size does not match"
}

func (vnm VectorNotMatch) Error() string {
	return "Declared size and matrix size does not match"
}

func (ims ImpossibleMatrixStructure) Error() string {
	return "Impossible structure of matrix"
}

func swap[T any](first *[]T, second *[]T) {
	*first, *second = *second, *first
}

func jacobi(matrix *[][]float64, vector []float64, size uint64, iterations uint64) ([]float64, error) {
	var x []float64 = make([]float64, size)
	if uint64(len(vector)) != size {
		return nil, VectorNotMatch{}
	}
	if uint64(len((*matrix))) != size {
		return nil, MatrixNotMatch{}
	}
	var x_next []float64 = make([]float64, size)
	for iteration := uint64(0); iteration < iterations; iteration++ {
		for index := uint64(0); index < size; index++ {
			if uint64(len((*matrix)[index])) != size {
				return nil, ImpossibleMatrixStructure{}
			}
			if math.Abs((*matrix)[index][index]) < 1e-12 {
				return nil, ZeroDivisionError{}
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
	if uint64(len(vector)) != size {
		return nil, VectorNotMatch{}
	}
	if uint64(len((*matrix))) != size {
		return nil, MatrixNotMatch{}
	}
	for iteration := uint64(0); iteration < iterations; iteration++ {
		for index := uint64(0); index < size; index++ {
			if uint64(len((*matrix)[index])) != size {
				return nil, ImpossibleMatrixStructure{}
			}
			if math.Abs((*matrix)[index][index]) < 1e-12 {
				return nil, ZeroDivisionError{}
			}
			var sum float64 = vector[index]
			for i := uint64(0); i < size; i++ {
				if i == index {
					continue
				}
				sum -= (*matrix)[index][i] * x[i]
			}
			x[index] = sum / (*matrix)[index][index]
		}
	}
	return x, nil
}
