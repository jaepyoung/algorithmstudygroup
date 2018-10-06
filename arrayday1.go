package main

type query struct {
	left  int
	right int
}

func main() {
	calquery([]int{1, 1, 3, 4, 4}, []query{query{1, 3}, query{2, 3}})
	calquery([]int{1, 1, 1, 4, 4}, []query{query{1, 3}, query{2, 3}})

}

var Queries = make([]query, 10)

func calquery(targetarray []int, queries []query) {
	diffarray := make([]int, len(targetarray))
	diffarray[0] = 0
	for i := 0; i < len(targetarray)-1; i++ {
		if targetarray[i] == targetarray[i+1] {
			diffarray[i+1] = diffarray[i] + 1
		} else {
			diffarray[i+1] = diffarray[i]
		}
	}
	//fmt.Printf("%v", diffarray)
	for i := 0; i < len(queries); i++ {
		//print(diffarray[queries[i].right])
		//print(diffarray[queries[i].left])
		print(diffarray[queries[i].right] - diffarray[queries[i].left])
	}
}
