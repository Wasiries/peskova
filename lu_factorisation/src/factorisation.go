package main

import (
	"math"
)

func solvingLU(lower *[][]float64, vector []float64) []float64 {
	var size = len(*lower)
	var y []float64 = make([]float64, size)
	for i := 0; i < size; i++ {
		for j := 0; j < i; j++ {
			vector[i] -= (*lower)[i][j] * y[j]
		}
		y[i] = vector[i] / (*lower)[i][i]
	}
	var x []float64 = make([]float64, size)
	for i := size - 1; i >= 0; i-- {
		for j := i + 1; j < size; j++ {
			y[i] -= (*lower)[j][i] * x[j]
		}
		x[i] = y[i] / (*lower)[i][i]
		if i == 0 {
			break
		}
	}
	return x
}

func Holetsky(matrix *[][]float64, vector []float64) []float64 {
	var size = len(*matrix)
	var lower [][]float64 = make([][]float64, size)
	for i := 0; i < size; i++ {
		lower[i] = make([]float64, size)
	}
	for i := 0; i < size; i++ {
		var mainSum float64 = 0.0
		for p := 0; p < i; p++ {
			mainSum += lower[i][p] * lower[i][p]
		}
		lower[i][i] = math.Pow((*matrix)[i][i]-mainSum, 0.5)
		for j := i + 1; j < size; j++ {
			var sideSum float64 = 0.0
			for p := 0; p < i; p++ {
				sideSum += lower[i][p] * lower[j][p]
			}
			lower[j][i] = ((*matrix)[j][i] - sideSum) / lower[i][i]
		}
	}
	var answer []float64 = solvingLU(&lower, vector)
	return answer
}
