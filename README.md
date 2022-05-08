#LAAIT Interpreter
Written in Go from scratch

## <i> Functionalities of Laait </i>
#### Dynamic Type Language 
<e> There is no need of defining Data type for variable explictly </e>
<code><i>input : </i>
let a = 10;
puts(a+1)
<i> Output -> </i> 11 null
</code>

<code><i>input: </i>
let b = "String"
puts(b + " is a string")
</code>
<i> Output: </i>
String is a string null

<code> 
let c = [1,2,3]
puts(c[1])

2 null
</code>

#### Variety of DataTypes 
<li> String </li>
<li> Integer </li>
<li> Boolean </li>
<li> Array </li>
<li> Hash Table </li>


#### Support Conditional Statements and Comparision Operators
<i>input: </i>
<code>
puts(1 < 2)
puts(1 > 2)
puts(1 < 1)
puts(1 == 2)
puts(1 != 2)
puts(true == true)
puts(false == true)
puts(true != true)
puts((1 < 2) == false)
</code>
<i>Output: </i>
true false false false true true false false false null

<i>input: </i>
<code>
let a = 10;
if ( a > 9) {
   if (a < 11){
        puts(" a is 10")
   }
}
else{
   puts("a is less than 10");
}
</code>
<i>Output: </i>
a is 10 null

<i>input: </i>
<code>1 == true</code>
<i>Output: </i>
false null

<i>input: </i>
<code>
puts(!true)
puts(!false)
puts(!!true)
puts(!!false)</code>
<i>Output: </i>
false true true false null


<b> Every statement is an expression mode </b>
<i>input: </i>
<code>
let a = if (true) { 10 }
let b = if (false) { 10 }
let c = if (1) { 10 }
let d = if ( 1 < 2 ) { 10 }
let e = if ( 1 > 2) { 10 } else { 20}
let f = if ( 1 < 2) { 10 } else { 20}

puts(a,b,c,d,e,f)
</code>
<i>Output: </i>
10 null 10 10 20 10 null

#### Support Functions 
<i>input: </i>
<code>
let add = function(x,y) { 
          x + y;}
 add(5+5, add(5,5));
</code>
<i>Output: </i>
20 null


#### Support global and local variables
<i>input: </i>
<code>
let global = 10;
let testFunction = function(){
    puts(global);
    let c = 100;
    let global = 11;}
testFunction();
puts(global)
puts(c)
</code>
<i>Output: </i>
10 10 ERROR: identifier not found: c

#### Supports Clousure 
<i>input: </i>
<code>
let newAdder = function(x){
                function(y){
                        x + y;
                };
       };
let addTwo = newAdder(2);       
addTwo(2);
</code>
<i>Output: </i>
4

#### Inbuilt functions
<i> Array in LAAIT are like list in python </i>
<u> last(arr) : Give the last element of list</u>
<u> push(arr,10) : Return a new the array with appending the element at last</u>
<u> first : Give the first element of list</u>
<u> rest :  Give the rest element of list except first element</u>
<code>
let a = [1,2,3,4,5];
puts(a[0], last(a));
let s = "String"
puts(len(s))
let c = push(a, 100)
puts(c);
rest(a)
</code>
<i>Output: </i>
1 5 6 [1,2,3,4,5,100] [2,3,4,5]

#### Dictionary or Hash Table in LAAIT

<i>input: </i>
<code>
let c = {"Name" : "Rahul", true: false, 1:20}
puts(c["Name"], c[true], c[1])
</code>
<i>Output: </i>
Rahul false 20 null 


#### Examples 

<b> Implementing Fibonacci series </b>
<i>input: </i>
<code>
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
</code>
<i>Output: </i>
55 null 

<b> Implementing Factorial Program</b>
<i>input: </i>
<code>

let fact = function(a){
    if ( a == 0) { 1}
    else { a * fact(a-1) }
} 

fact(5)

</code>
<i>Output: </i>
120 null
