# Python


### note
- execute with sequential , call function after define it
- weak closure , which can modify global variable with `global [variable]` , but variable in outer function can't be modify by inner function
- use [itertools](http://wklken.me/posts/2013/08/20/python-extra-itertools.html#itertoolscountstart0-step1),operations , functools
-

### function
- closure:
- inner function: 
- coroutine (yield):
- decorator





### import sys
sys.getrecursionlimit() , default:1000
sys.setrecursionlimit(limit)


### build-in
listA = [1,2,3,4]
listB = [5,6,7,8]



zip
i = iter(list) , i(result:1) i(result:2) i(result:3) i(result:4)
dict(izip(i,i))     , from itertools import izip




### condition
- 0 or "" or [] or {} == False
- 1 == True 
- "1" == False
- bool("123") == True
- bool([1]) == True
- bool([]) == False
- bool({1:1}) == True
- bool({}) == False



### operator
- `//` == `%`
- `n**i` == n to power of i
 

### string
"".join(list)


### list
concate two list ,  [] + []
length , len([])
new , []
list.sort(function) , sort with integer has problem ,  9 10 => 10 9
[a::b] , list start with a , and jump with b difference
[a:b] , list start with a and end before b
[i for i in range(x)]

list(str)
list(tuple)



### dictionary
new , {}
{i:j for i in range(x)}
dict

### special function

- `__init__` , object initializatioon
- `__doc__` , class document
- `__str__` , str()
- `__repr__` , repr()
- `__init__`pythone 2 vs python 3
- `__main__` 

在ubuntu 14.xx之前 很多系統的嘟星都是用python2建的
所以有些系統吃死python2



x//y

parent class  
    self.v1

child class
    self.v2


dir_name file_name class_name for init class instance



try catch finally的使用


