package main

func getnumberofisland(grid [][]int) int {
	result := 0
	if len(grid) == 0 {
		return 0
	}

	return result
}

func dfs(grid [][]int, i int, j int) {
	numrows=len(grid)
	numcolumns=len(grid[0])

    if i < 0  or j < 0 or i > numrows-1 or j > numcolums-1 or grid[i][j] != 1:
        return
    if grid[i][j]==1:
        grid[i][j]=2
        dfs(grid,i,j-1)
        dfs(grid,i,j+1)
        dfs(grid,i-1,j)
        dfs(grid,i+1,j) 

}
