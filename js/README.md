# js

### function arguments passed by 
Arguments are passed by Value
Objects are passed by Reference

### variable assignment
copy by reference

### Memory leak
- [js info](http://javascript.info/tutorial/memory-leaks)
- [MSDN](https://msdn.microsoft.com/en-us/library/ms976398.aspx)
- [js mem leak](http://blogger.gtwang.org/2014/01/javascript-memory-leak-patterns.html)

issue:

will cause leak or not?
```
    function A(){}
    A.prototype.x = function(){
        var y = copy of A
        return y
    }
    var a = new A();
    a = a.x()

```


### condition
- `==`  , weak type comparsion please check [boolean](#boolean)
- `===` , 


### boolean

true:
- 1
- true
- "anyword"

false:
- 0
- false
- ""


### string
```
str = "joe:x:100:100:see";
str.split(":")  => ["joe","x",100,100,"see"]
```

### number
All numbers in JavaScript are 64-bit (double-precision) floating point numbers.


### bits op
negative number right shift bits with 1 and postive with 0


# nodejs

### install
check [here](https://github.com/BizShuk/env_setup/blob/master/setup/nodejs.sh)

### npm

##### install module
find packages at [npm](https://www.npmjs.com/package/package)

`npm install <module>` 
- -g , for install globally




