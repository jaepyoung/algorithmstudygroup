package main
import (
	"fmt"
)

func numSquares(n int) int {
	if n<= 0 {
		return 0
	}

	f := make([]int, n+1)

	f[0]=0

    for i:=1; i<=n;i++{
		f[i] = f[i-1] + 1 
		for j:=1; j*j<= i; j++{
			f[i] = min(f[i],f[i-j*j] + 1) 
		}
	}
	
	return f[n]
}

func min(a int, b int) int {
	if a<b {
		return a
	}
	return b
}

func main() {
	fmt.Println(numSquares(133))
}