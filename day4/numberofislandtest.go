package main

import "testing"

func TestSum(t *testing.T) {

	values := [][]int{}
	rows1 := []int{1, 0}
	values = append(values, rows1)
	result := getnumberofisland(values)
	if result != 0 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 0)
	}
}
