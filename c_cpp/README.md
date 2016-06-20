# C_CPP

sizeof , size of bytes  check each type of bytes at sizeof.cpp
- int 4 bytes
- char 1 bytes
- float 4 bytes
- double 8 bytes

use sizeof to know memory size(byte) , but there is no way to know a pointer array size.



[pointer mapping](http://stackoverflow.com/questions/3920729/in-c-c-is-char-arrayname-a-pointer-to-a-pointer-to-a-pointer-or-a-pointe?answertab=active#3925968)


swap continuous memory => segmentation fault
```
    int a[10];
    swap_address(a[1],a[5]);
```



### build
	
##### build shared object
`gcc -shared -o libhello.so -fPIC hello.c`

##### build
`gcc`

`g++`







# CPP


### garbege collection
If it use **malloc** or **new** , it need to be **free()** or **delete**. 

##### new
If new is assigned to pointer or reference. After end of the scope , pointer and reference will be recycled, but not the new object.
In other situation , if you reassign pointer to a usual variable , it'll be recycled after end of the scope.

##### malloc
?



### Class
check example [here](class.cpp)

**this** is a pointer.

##### Different between new and no-new
- `Point p1 = Point(0,0)` , collect when scope is gone
- `Point p1 = new Point(0,0)` , need to delete p1

##### Differenet between . and ->
`*a.b` = `a->b`



### function parameter 

##### pass by value
```
void <func_name>(int p)
```

##### pass by reference
```
void <func_name>(int &p)
```
treat this like alias.


##### pass by address
```
void <func_name>(int *p)
```





##### virtual function
	ex: virtual void function_name(){};

	操作將延後到執行階段決定
		Wolf wolf = new Wolf();
		wolf.eat();		// wolf eat
		Animal* animal = wolf;
		animal.eat();	// animal eat

##### pure virtual void function
	ex: pure virtual void function_name()=0;

	child class 必定需要 overide

##### abstract class
	含有 pure virtual function 即稱為 , 如果用此class new instance會發生編譯錯誤


Casting example:
	ref. http://stackoverflow.com/questions/5313322/c-cast-to-derived-class
	class Animal { /* Some virtual members */ }
	class Dog: public Animal {};
	class Cat: public Animal {};

	Dog     dog;
	Cat     cat;
	// Notice no cast required. (Dogs and cats are animals).
	Animal& AnimalRef1 = dog;
	Animal& AnimalRef2 = cat;
	Animal* AnimalPtr1 = &dog;
	Animal* AnimlaPtr2 = &cat;

	// Throws an exception  AnimalRef1 is a dog
	Cat&    catRef1 = dynamic_cast<Cat&>(AnimalRef1);
	// Returns NULL         AnimalPtr1 is a dog
	Cat*    catPtr1 = dynamic_cast<Cat*>(AnimalPtr1);
	Cat&    catRef2 = dynamic_cast<Cat&>(AnimalRed2);  // Works
	Cat*    catPtr2 = dynamic_cast<Cat*>(AnimalPtr2);  // Works

	// This on the other hand makes no sense
	// An animal object is not a cat. Therefore it can not be treated like a Cat.
	Animal  a;
	// Throws an exception  Its not a CAT
	Cat&    catRef1 = dynamic_cast<Cat&>(a);
	// Returns NULL         Its not a CAT.	
	Cat*    catPtr1 = dynamic_cast<Cat*>(&a);


##### public private protected 差異

|作用域　 |當前類|同一package|子孫類|其他package|
|---------|------|-----------|------|-----------|
|public 　|　√　|　　√　　 |　√　|　　√     |
|protected|  √　|　　√　　 |　√　|　　×     |
|friendly |　√　|　　√　　 |　×　|　  ×     |
|private　|　√　|　　×　　 |　×　|　  ×     |
 



