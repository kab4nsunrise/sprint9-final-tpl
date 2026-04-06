package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"zero", 0, 0},
		{"positive", 100, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.n)
			if len(got) != tt.want {
				t.Errorf("length = %d, want %d", len(got), tt.want)
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  int
	}{
		{"empty", []int{}, 0},
		{"single", []int{42}, 42},
		{"multiple", []int{1, 5, 3, 9, 2}, 9},
		{"all equal", []int{7, 7, 7}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.slice); got != tt.want {
				t.Errorf("maximum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  int
	}{
		{"empty", []int{}, 0},
		{"single", []int{100}, 100},
		{"less than chunks", []int{3, 1, 4, 2}, 4},
		{"normal", []int{10, 20, 5, 30, 15, 25, 35, 40, 0}, 40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.slice); got != tt.want {
				t.Errorf("maxChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}
