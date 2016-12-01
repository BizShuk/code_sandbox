# js


### jsonp
Actually , it's not a form request , but a script tag which is to get js content and put it to callback function. This could bypass the cross-domain issue which is called "CORS" stand by "Cross-Origin Resourece Sharing".



### tools
- **npm**: Node JS package manager, helps you to manage all the libraries your software relays on. You would define your needs in a file called package.json and run npm install in the command line... BANG your packages are downloaded and are ready to use. Could be used both for front-end and back-end.
- Bower: Front-end package manager. Not that much good and is not useful these days since most of the devs are relaying on NPM.
- **Webpack**: The coolest kid in the town. It will bundle your app into other bundling patters so you can use all the libraries available in NPM right into your front-end code, load different modules and do many other things. It is amazingly flexible and can create strong development environments. Very trendy these days.
- **Gulp**: Automation just like Grunt but instead of configurations you can write JavaScript with streams like it's a node application. Much better than Grunt.
- Grunt: You can create automation for your development environment to pre process codes or create build scripts with a not very simple config file. It was great back there in 2013 but not much these days.
- Browserify: Similar to Webpack but less powerful.
- **Express**: Node JS web application framework. Could be used for routing and anything else needed is cover through middle wares. Very popular and beautifully designed so if you want to create a web application project with node, you are probably using it.
- Slush and Yeoman: Project scaffolding systems. You can create starter projects with them. Not good that much, use a Github boilerplate instead.
- **Babel** : transcompiler for ES6 to ES5


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

### === Class ===
In JavaScript, class methods are not bound by default.


### string
```
str = "joe:x:100:100:see";
str.split(":")  => ["joe","x",100,100,"see"]
```

### number
All numbers in JavaScript are 64-bit (double-precision) floating point numbers.


### bits op
negative number right shift bits with 1 and postive with 0


# Browser
### eventhandler(e) 
e.preventDefault => prevent default event like a link or a form submit button


### event trigger
trigger start from dom root (html -> body ->element) to child and bubble back

##### event Capturing

##### event Bubbling


# nodejs

### install
check [here](https://github.com/BizShuk/env_setup/blob/master/setup/nodejs.sh)

### npm

##### install module
find packages at [npm](https://www.npmjs.com/package/package)

`npm install <module>` 
- -g , for install globally




