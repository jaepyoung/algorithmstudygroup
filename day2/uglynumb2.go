package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getuglynumber(x int) int {
	if x == 1 {
		return 1
	}
	factors := []int{2, 3, 5}

	h := &IntHeap{1, 2, 3, 5}
	heap.Init(h)
	heap.Pop(h)
	fmt.Print("Test")

	for i := x; i >= 2; i-- {
		uglinum := heap.Pop(h).(int)
		fmt.Print(h)
		for _, factor := range factors {
			newuglinum := uglinum * factor
			i := 0
			for _, element := range *h {
				if newuglinum == element {
					i++
				}
			}
			if i == 0 {
				heap.Push(h, newuglinum)
			}
		}

	}
	return heap.Pop(h).(int)
}

func main() {
	fmt.Println(getuglynumber(13))
}
