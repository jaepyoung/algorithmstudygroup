
def getUglyNum(n):
    if n <=1 : 
        return n
    factors = [2,3,5]
    
    heap = factors + [1] 

    import heapq
    heapq.heapify(heap)

    while n > 1: 
        uglNum = heapq.heappop(heap)
        for factor in factors: 
            newUglyNum = factor * uglNum
            if newUglyNum not in heap:
                heapq.heappush(heap,newUglyNum)
        n -= 1
    return heapq.heappop(heap)


print getUglyNum(10)