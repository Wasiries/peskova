package main

import (
	"fmt"
)

func test3x3_v1() {
	var matrix = &[][]float64{
		{5, 1, 1},
		{1, 5, 1},
		{1, 1, 5},
	}
	var vector = []float64{7, 7, 7}
	var answer = []float64{1, 1, 1}
	var result = Holetsky(matrix, vector)
	fmt.Printf("\nanswer: %v\nreceived from algorithm: %v\n", answer, result)

}

func test3x3_v2() {
	var matrix = &[][]float64{
		{10, 1, 3},
		{1, 10, 1},
		{3, 1, 10},
	}
	var vector = []float64{17, 13, 24}
	var answer = []float64{1, 1, 2}
	var result = Holetsky(matrix, vector)
	fmt.Printf("\nanswer: %v\nreceived from algorithm: %v\n", answer, result)
}

func test2x2_v1() {
	var matrix = &[][]float64{
		{5, 2},
		{2, 3},
	}
	var vector = []float64{22, 11}
	var answer = []float64{4, 1}
	var result = Holetsky(matrix, vector)
	fmt.Printf("\nanswer: %v\nreceived from algorithm: %v\n", answer, result)
}

func main() {
	test3x3_v1()
	test3x3_v2()
	test2x2_v1()
}
