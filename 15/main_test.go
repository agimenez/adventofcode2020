package main

import "testing"

func TestGame(t *testing.T) {
	tests := []struct {
		input    []int
		limit    int
		expected int
	}{
		{[]int{0, 3, 6}, 2020, 436},
		{[]int{1, 3, 2}, 2020, 1},
		{[]int{2, 1, 3}, 2020, 10},
		{[]int{1, 2, 3}, 2020, 27},
		{[]int{2, 3, 1}, 2020, 78},
		{[]int{3, 2, 1}, 2020, 438},
		{[]int{3, 1, 2}, 2020, 1836},
	}

	for _, tt := range tests {
		v := MemoryGame(tt.input, tt.limit)
		if v != tt.expected {
			t.Errorf("Got %d, expected %d", v, tt.expected)
		}
	}
}
