def getlongestsubstringnumb(a,b):
    solutionmax=[ [ 0 for i in range(len(a)) ] for j in range(len(b)) ]
    for i in range(len(a)):
        for j in range(len(b)):
            if (i==0 or j==0): 
                solutionmax[i][j]=0

            if (a[i]==b[j]):
                solutionmax[i][j]=1+solutionmax[i-1][j-1]          
    print solutionmax

getlongestsubstringnumb("adbcdefa","thahdecd")