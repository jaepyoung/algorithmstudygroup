package main

import "testing"

func TestSum(t *testing.T) {
	result := getnumberofisland([][]int{{}})
	if result != 0 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", result, 0)
	}
}

func TestSum1(t *testing.T) {
	result := getnumberofisland([][]int{{1, 0}})
	if result != 1 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", result, 1)
	}
}
