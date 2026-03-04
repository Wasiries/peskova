package main

import (
	"jacobi_zeidel/slau"
	"math"
	"testing"
)

func TestJacobi(t *testing.T) {
	tests := []struct {
		name       string
		matrix     *[][]float64
		vector     []float64
		size       uint64
		iterations uint64
		want       []float64
		wantErr    bool
	}{
		{
			name: "simple 3x3 diagonal dominant",
			matrix: &[][]float64{
				{5, 1, 1},
				{1, 5, 1},
				{1, 1, 5},
			},
			vector:     []float64{7, 7, 7},
			size:       3,
			iterations: 50,
			want:       []float64{1, 1, 1},
			wantErr:    false,
		},
		{
			name: "zero diagonal element",
			matrix: &[][]float64{
				{0, 2},
				{3, 4},
			},
			vector:     []float64{1, 1},
			size:       2,
			iterations: 10,
			want:       nil,
			wantErr:    true,
		},
		{
			name: "size mismatch (matrix too small)",
			matrix: &[][]float64{
				{1, 2},
				{3, 4},
			},
			vector:     []float64{1, 2, 3},
			size:       3,
			iterations: 5,
			want:       nil,
			wantErr:    true,
		},
		{
			name:       "nil matrix",
			matrix:     nil,
			vector:     []float64{1, 2},
			size:       2,
			iterations: 5,
			want:       nil,
			wantErr:    true,
		},
		{
			name: "nil vector",
			matrix: &[][]float64{
				{1, 2},
				{3, 4},
			},
			vector:     nil,
			size:       2,
			iterations: 5,
			want:       nil,
			wantErr:    true,
		},
		{
			name:       "size zero",
			matrix:     &[][]float64{},
			vector:     []float64{},
			size:       0,
			iterations: 5,
			want:       []float64{},
			wantErr:    false,
		},
		{
			name: "iterations zero (should return initial guess zeros)",
			matrix: &[][]float64{
				{5, 1},
				{1, 5},
			},
			vector:     []float64{6, 6},
			size:       2,
			iterations: 0,
			want:       []float64{0, 0},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := slau.Jacobi(tt.matrix, tt.vector, tt.size, tt.iterations)
			if (err != nil) != tt.wantErr {
				t.Errorf("Jacobi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(got) != len(tt.want) {
					t.Errorf("Jacobi() returned slice of length %d, want %d", len(got), len(tt.want))
					return
				}
				for i := range got {
					if math.Abs(got[i]-tt.want[i]) > 1e-6 {
						t.Errorf("Jacobi() = %v, want %v (diff at index %d)", got, tt.want, i)
						break
					}
				}
			}
		})
	}
}

func TestZeidel(t *testing.T) {
	tests := []struct {
		name       string
		matrix     *[][]float64
		vector     []float64
		size       uint64
		iterations uint64
		want       []float64
		wantErr    bool
	}{
		{
			name: "simple 3x3 diagonal dominant",
			matrix: &[][]float64{
				{5, 1, 1},
				{1, 5, 1},
				{1, 1, 5},
			},
			vector:     []float64{7, 7, 7},
			size:       3,
			iterations: 50,
			want:       []float64{1, 1, 1},
			wantErr:    false,
		},
		{
			name: "zero diagonal element",
			matrix: &[][]float64{
				{0, 2},
				{3, 4},
			},
			vector:     []float64{1, 1},
			size:       2,
			iterations: 10,
			want:       nil,
			wantErr:    true,
		},
		{
			name: "size mismatch (matrix too small)",
			matrix: &[][]float64{
				{1, 2},
				{3, 4},
			},
			vector:     []float64{1, 2, 3},
			size:       3,
			iterations: 5,
			want:       nil,
			wantErr:    true,
		},
		{
			name:       "nil matrix",
			matrix:     nil,
			vector:     []float64{1, 2},
			size:       2,
			iterations: 5,
			want:       nil,
			wantErr:    true,
		},
		{
			name: "nil vector",
			matrix: &[][]float64{
				{1, 2},
				{3, 4},
			},
			vector:     nil,
			size:       2,
			iterations: 5,
			want:       nil,
			wantErr:    true,
		},
		{
			name:       "size zero",
			matrix:     &[][]float64{},
			vector:     []float64{},
			size:       0,
			iterations: 5,
			want:       []float64{},
			wantErr:    false,
		},
		{
			name: "iterations zero (should return initial guess zeros)",
			matrix: &[][]float64{
				{5, 1},
				{1, 5},
			},
			vector:     []float64{6, 6},
			size:       2,
			iterations: 0,
			want:       []float64{0, 0},
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := slau.Zeidel(tt.matrix, tt.vector, tt.size, tt.iterations)
			if (err != nil) != tt.wantErr {
				t.Errorf("Zeidel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(got) != len(tt.want) {
					t.Errorf("Zeidel() returned slice of length %d, want %d", len(got), len(tt.want))
					return
				}
				for i := range got {
					if math.Abs(got[i]-tt.want[i]) > 1e-6 {
						t.Errorf("Zeidel() = %v, want %v (diff at index %d)", got, tt.want, i)
						break
					}
				}
			}
		})
	}
}
