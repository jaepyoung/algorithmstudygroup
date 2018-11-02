import unittest
from sets import Set

def getnumberofilsand(grid): 
    numrows=len(grid)
            
    result = 0
    if numrows==0: 
        return 0
    else: 
        if numrows > 0:
            if isinstance(grid[0],list) : 
                numcolums=len(grid[0])
            else: 
                numrows = 1
                numcolums = len(grid)
                for i in range(0,numcolums):
                    if i == 0 and grid[i]==1:
                        result = 1
                    if i > 0: 
                        if grid[i]==1 and grid[i-1]==0:
                            result = result + 1
                return result


        print("rows"+str(numrows)+"colums"+str(numcolums))    

    for i in range(0,numrows):
        for j in range(0,numcolums):
            if grid[i][j]==1:
                result = result +1 
                dfs(grid,i,j)

    return result


def dfs(grid, i, j): 
    numrows=len(grid)
    if numrows > 0:
        if isinstance(grid[0],list) : 
            numcolums=len(grid[0])
        else:
            numcolums = 0
    print(grid)

    if i < 0  or j < 0 or i > numrows-1 or j > numcolums-1 or grid[i][j] != 1:
        return
    if grid[i][j]==1:
        grid[i][j]=2
        dfs(grid,i,j-1)
        dfs(grid,i,j+1)
        dfs(grid,i-1,j)
        dfs(grid,i+1,j) 
    

class TestStringMethods(unittest.TestCase):

    def test_zero(self):
        self.assertEqual(getnumberofilsand([]), 0)
    def test_first(self):
        self.assertEqual(getnumberofilsand([1,0]), 1)
    
    def test_second(self):
        self.assertEqual(getnumberofilsand([1,1]), 1)
    def test_second_one(self):
        self.assertEqual(getnumberofilsand([1,0,1]), 2)    
    
    def test_third(self):
        self.assertEqual(getnumberofilsand([[1,0],[0,1]]), 2)

    def test_third_two(self):
        self.assertEqual(getnumberofilsand([[1,0],[1,1]]), 1)

    def test_third_four(self):
        self.assertEqual(getnumberofilsand([[1,0,0],[0,1,0],[0,0,1]]), 3)
    
    

if __name__ == '__main__':
    unittest.main()