#!/bin/python3

import sys


def make_change(coins, n):
    """
        dynamic programming
    """
    dp = [0 for i in range(n+1)]
    dp[0] = 1
    
    cl = len(coins)
    for i in range(cl-1,-1,-1):
        c = coins[i]
        for j in range(c,n+1):
            dp[j] += dp[j-c]
            
    return dp[n]


def make_change_2(coins, n):
    return len(change_ways(coins,n,len(coins)-1))


def change_ways(coins,n,c_i):
    """
        right about concept , but Time Limited Exceeded 
    """   
    ret = []
    for i in range(c_i,-1,-1):
        a = coins[i]
        if a == n:
            ret += [[a]]
        if a < n :
            ret += change_ways(coins,n-a,i)
    ret = [[coins[c_i]]+i for i in ret]
    
    return ret

def make_change_3(coins, n):
    coins.sort()
    c_i = len(coins)-1
    c_s = []
    ret = 0
    while True:
        #print(n , c_s,c_i)
        c =coins[c_i]
        
        if n >= c:
            c_s += [c]
            n -= c
        else:
            c_i -= 1
            
        if c_i < 0 or n == 0:
            # 連續pop到c_i !=0 or len(c_s) ==0 break
            if n == 0:
                ret += 1
                #print("==",n,c_s,c_i)
             
            
            while len(c_s) > 0:
                c = c_s.pop()
                c_i = coins.index(c)
                n += c
                if c_i > 0:
                    break
                    
            
            if len(c_s) == 0 and c_i ==0:
                break
            c_i -= 1
            
    return ret
            
