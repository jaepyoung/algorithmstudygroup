package main

import "fmt"

func fibonacci(x int) int {
	results := make([]int, x)
	results[0] = 0
	results[1] = 1

	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	for i := 2; i < x; i++ {
		results[i] = results[i-1] + results[i-2]
	}
	return results[x-1]
}

func main() {
	fmt.Println(fibonacci(50))
}
