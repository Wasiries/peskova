package main

import (
	"fmt"
	"jacobi_zeidel/slau"
)

func TestJacobi2x2(iterations uint64) {
	var matrix = &[][]float64{
		{2, 1},
		{1, -1},
	}
	var vector = []float64{7, 2}
	var answer = []float64{3, 1}
	var received, err = slau.Jacobi(matrix, vector, 2, iterations)
	fmt.Printf("Jacobi\nerror presence: %v \nanswer: %v\nreceived: %v\n\n", (err == nil), answer, received)
}

func TestJacobi3x3(iterations uint64) {
	var matrix = &[][]float64{
		{5, 1, 1},
		{1, 5, 1},
		{1, 1, 5},
	}
	var vector = []float64{7, 7, 7}
	var answer = []float64{1, 1, 1}
	var received, err = slau.Jacobi(matrix, vector, 3, iterations)
	fmt.Printf("Jacobi\nerror presence: %v \nanswer: %v\nreceived: %v\n\n", (err == nil), answer, received)
}

func TestZeidel2x2(iterations uint64) {
	var matrix = &[][]float64{
		{5, -3},
		{2, 3},
	}
	var vector = []float64{17, 11}
	var answer = []float64{4, 1}
	var received, err = slau.Zeidel(matrix, vector, 2, iterations)
	fmt.Printf("Zeidel\nerror presence: %v \nanswer: %v\nreceived: %v\n\n", (err == nil), answer, received)
}

func TestZeidel3x3(iterations uint64) {
	var matrix = &[][]float64{
		{10, 1, 2},
		{2, 10, 1},
		{3, 1, 10},
	}
	var vector = []float64{15, 14, 24}
	var answer = []float64{1, 1, 2}
	var received, err = slau.Zeidel(matrix, vector, 3, iterations)
	fmt.Printf("Zeidel\nerror presence: %v \nanswer: %v\nreceived: %v\n\n", (err == nil), answer, received)
}

func main() {
	TestJacobi2x2(10)
	TestJacobi3x3(10)
	TestZeidel2x2(10)
	TestZeidel3x3(10)
}
