def isPrime(n):
    """ O(n^(0.5)) """
    ret = True
    def printIsPrime():
        if ret:
            print("Prime")
        else:
            print("Not prime")
        
        if n == 2:
            return printIsPrime
        if n <= 1 or n % 2 == 0 :
            ret = False
            return printIsPrime
        
        i=3
        smax = math.sqrt(n)
        while i <= smax:
            if n % i == 0:
                ret = False 
                return printIsPrime
            i+=1
        return printIsPrime


