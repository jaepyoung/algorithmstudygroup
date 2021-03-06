package main

import "fmt"

func maxfunc(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getlongesequence(a string, b string) int {
	aArray := []rune(a) //Create array from string a
	bArray := []rune(b) //Create array from string b
	max := 0
	solutionmax := make([][]int, len(aArray)+1)
	fmt.Println(solutionmax)
	for i := range solutionmax {
		solutionmax[i] = make([]int, len(bArray)+1)
	}
	fmt.Println(solutionmax)

	for i := 0; i <= len(aArray); i++ {
		solutionmax[i][0] = 0
	}

	for i := 0; i <= len(bArray); i++ {
		solutionmax[0][i] = 0
	}

	for i := 1; i <= len(aArray); i++ {
		for j := 1; j <= len(bArray); j++ {
			if aArray[i-1] == bArray[j-1] {
				solutionmax[i][j] = 1 + solutionmax[i-1][j-1]
				if solutionmax[i][j] > max {
					max = solutionmax[i][j]
				}
			} else {
				solutionmax[i][j] = maxfunc(solutionmax[i-1][j], solutionmax[i][j-1])
			}

		}
	}
	fmt.Print(solutionmax)
	return max
}

func main() {
	fmt.Println(getlongesequence("testhahatest1111", "testtesthahathhh"))
}
