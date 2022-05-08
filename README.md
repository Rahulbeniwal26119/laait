#LAAIT Interpreter
Written in Go from scratch

## <i> Functionalities of Laait </i>
#### Dynamic Type Language 
<e> There is no need of defining Data type for variable explictly </e>
<i>input : </i><br> 
```javascript
let a = 10;
puts(a+1)
```
<i> Output -> </i>
<code> 11 null </code>

<i>input: </i>
```javascript
let b = "String"
puts(b + " is a string")
```
<i> Output: </i>
<code>String is a string null </code>
<i>input: </i>
```javascript
let c = [1,2,3]
puts(c[1])
```
<i> Output: </i>
<code>2 null</code>

#### Variety of DataTypes 
<li> String </li>
<li> Integer </li>
<li> Boolean </li>
<li> Array </li>
<li> Hash Table </li>


#### Support Conditional Statements and Comparision Operators
<i>input: </i>
```javascript
puts(1 < 2)
puts(1 > 2)
puts(1 < 1)
puts(1 == 2)
puts(1 != 2)
puts(true == true)
puts(false == true)
puts(true != true)
puts((1 < 2) == false)
```
<i>Output: </i>
<code>true false false false true true false false false null</code>

<i>input: </i>
```javascript
let a = 10;
if ( a > 9) {
   if (a < 11){
        puts(" a is 10")
   }
}
else{
   puts("a is less than 10");
}
```
<i>Output: </i>
<code>a is 10 null</code>

<i>input: </i>
<u> false and true are seprate datatype </u>
```javascript
1 == true
```
<i>Output: </i>
<code> false null </code>

<i>input: </i>
```javascript
puts(!true)
puts(!false)
puts(!!true)
puts(!!false)</code>
```
<i>Output: </i>
<code>false true true false null</code>


<b><u> Every statement is an expression
</u></b>
<i>input: </i>
```javascript
let a = if (true) { 10 }
let b = if (false) { 10 }
let c = if (1) { 10 }
let d = if ( 1 < 2 ) { 10 }
let e = if ( 1 > 2) { 10 } else { 20}
let f = if ( 1 < 2) { 10 } else { 20}

puts(a,b,c,d,e,f)
```

<i>Output: </i>
<code>10 null 10 10 20 10 null </code>

#### Support Functions 
<i>input: </i>
```javascript
let add = function(x,y) { 
          x + y;}
 add(5+5, add(5,5));
```
<i>Output: </i>
<code>20 null </code>


#### Support global and local variables
<i>input: </i>
```javascript
let global = 10;
let testFunction = function(){
    puts(global);
    let c = 100;
    let global = 11;}
testFunction();
puts(global)
puts(c)
```
<i>Output: </i>
<code>10 10 ERROR: identifier not found: c </code>

#### Supports Clousure 
<i>input: </i>
```
let newAdder = function(x){
                function(y){
                        x + y;
                };
       };
let addTwo = newAdder(2);       
addTwo(2);
```
<i>Output: </i>
<code>4 </code>

#### Inbuilt functions
<i> Array in LAAIT are like list in python </i>
-  last(arr) : Give the last element of list
-  push(arr,10) : Return a new the array with appending the element at last
-  first : Give the first element of list
-  rest :  Give the rest element of list except first element

<i>input: </i>
```javascript
let a = [1,2,3,4,5];
puts(a[0], last(a));
let s = "String"
puts(len(s))
let c = push(a, 100)
puts(c);
rest(a)
```
<i>Output: </i>
<code>1 5 6 [1,2,3,4,5,100] [2,3,4,5] </code>

#### Dictionary or Hash Table in LAAIT

<i>input: </i>
```javascript
let c = {"Name" : "Rahul", true: false, 1:20}
puts(c["Name"], c[true], c[1])
```
<i>Output: </i>
Rahul false 20 null 


#### Examples 

<b> Implementing Fibonacci series </b>
<i>input: </i>
```javascript
let fib = function(a){
    if(a == 0){
        0;
    }
    else{
        if( a == 1 ) {
            1;
        }
        else{
            fib(a - 1) + fib(a -2);
        }
    }
    puts(a);
}

puts(fib(10));
```
<i>Output: </i>
<code>55 null</code> 

<b> Implementing Factorial Program</b>
<i>input: </i>
```javascript
let fact = function(a){
    if ( a == 0) { 1}
    else { a * fact(a-1) }
} 

fact(5)
```
<i>Output: </i>
<code>120 null </code>
