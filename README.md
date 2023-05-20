## For installation on mac

__Step 1:__
Terminal: `brew install go` \
**Step 2:**
VS code: create file main.go, vs code will suggest to install go extensions, click install and then install all

## To run file from terminal
`go run main.go`

A [good tutorial](https://www.youtube.com/watch?v=YS4e4q9oBaU) to follow 

## Variables
- 3 ways to declare variables

  - `var i int` variable i of type integer \
  when you want to declare a variable but dont want to initialize \
  `i = 42` initialize later

  - `var j int = 27` \
  if Go doesnt have their info to assign a type that we want \
  `fmt.Printf("%v, %T", j, j)` output = 27, int (value, type) \
  `var j float32 = 27` \
  `fmt.Printf("%v, %T", j, j)` output = 27, float32

  - `k:= 99` compiler can figure out the data type, mostly used \
  `fmt.Printf("%v, %T", k, k)` output = 99 , int \
  `k:= 99.` \
  `fmt.Printf("%v, %T", k, k)` output = 99 , float64 //wont print float32, unless defined

- If declared at package level, only use the first method \
  `var i int = 42` then it can be used in function

- Variable blocks can be declared. \
  This method allows us to declare multiple variables as a group. This does not look cluttered. These variables may or may not related. It is done outside the function.
  ```
  var
    (
        actorName string = "Elisabeth Sladen"
        companion string = "Sarah Jane Smith"
        doctorNumber int = 3
        season int = 11
    )
    ```
- Redeclaration of variables \
  Variable can be re-assigned, but **not** redeclared within one scope \
  `var i int = 20` is correct but cannot do `i:=10` later, it will give error. You can only do `i = 10` which will give output 10. \
  If a variable with the same identifier is declared in package scope and main scope aswell then precedence is given to the variable in inner most scope. This is called shadowing. \
  ```
  var i int = 11

  func main () {
        var i int = 9
        `fmt.Println(i)` // prints 9
        // shadowing by inner most scope
  }
    ```
    11 exists, but it is being overshadowed by 9 now.
    ```
    var i int = 11

    func main () {
            `fmt.Println(i)` // prints 11
            var i int = 9
            `fmt.Println(i)` // prints 9
            // shadowing by inner most scope
    }
        ```

- Declared variables in Go should always be used, otherwise you get compile time error.
- How to name variables:
  - visibility \
  a lower case variable declared in package scope can only be used inside the package but an uppercase variable declared inside package scope is usable outside the package. \
  package level - lowercase : package scope \
  package level - uppercase : global scope \
  function level - block scope \
  no private scope 
  - naming convention \
  length of variable reflects the life of the variable e.g `i` or `seasonNumber` \
  package level variables should be meaningful \
  acronyms should be all caps e.g `theHTTPRequest`
- Type conversions - using conversion functions \
  - `var j float32 = float32(i)` // float32 is a conversion function here \
  explicitly change types, because Go doesnt want to lose information by itself 
  - string converstion package
    ```
    import "strconv"

    func main () {
            var i int =42
            fmt.Printf("%v, %T", i , i)  // prints 42, int 
            
            var k string = string(i)
            fmt.Printf("%v, %T", k, k)  // prints *, string 
            // In Go, string is basically an alias for a string of bytes. While 
            converting Go looks for what unicode character 42 represents and 
            assigns that to string variable
            
            var j string = strconv.Itoa(i)
            fmt.Printf("%v, %T", j, j)  // prints 42, string
    }
    ```

## Primitives data types 
- Boolean
  - Two states : true/false
  - Flags \
  `var n bool = false` \
  `fmt.Printf("%v,%T\n",n,n)` // prints false, bool
  - Logical tests \
  `n := 1==1` \
  This method is a very common logical test. it compares values on right side to see if they're equal and returns boolean, on run time. \
  `m := 1==2` \
  `fmt.Printf("%v,%T\n",n,n)` // true, bool \
  `fmt.Printf("%v,%T\n",m,m)` // false, bool
  - In Go, variables are intialized with 0, if not initialized explicitly \
  `var n bool` // false
- Numeric
  - Integer - signed and unsigned
  - int is basically signed int. for signed int we int8,int16,int32 and int64. minimun 32 bit
  - `var i uint16=42` unsigned int16. for unsigned integers we have  uint8,uint16 and uint32, byte (unsigned 8 bit integer)
  - Basic arithematic operations for int
	```
    a:=10
	b:=3
	fmt.Println(a+b) // 13
	fmt.Println(a-b) // 7
	fmt.Println(a*b) // 30
	fmt.Println(a/b) // 3 
    //division will always result in an int output so quotient will be returned in answer. We cannot display a float by parsing the result of the division. remainder is dropped.  
	fmt.Println(a%b)// 1 - mod function 
	
	//we are not allowed to perform arithematic operations using different types of int. For example we cannot add an int8 with an int. we'll have to parse one of them.
	
    var a int = 10
	var b int8 = 3
	fmt.Println(a+b) //this will not work
	fmt.Println(a+int(b)) //this will work
    ```
  - Binary operations
    ```
    a:=10//1010
    b:=3//0011
    fmt.Println(a&b)  //bitwise AND. 0010
    fmt.Println(a|b)  //bitwise OR.  1011
    fmt.Println(a^b)  //bitwise XOR. 1001
    fmt.Println(^b)   //bitwise NOT. It is same as (1111^0011). result wil be a signed number
    
    fmt.Println(a&^b) //bitwise AND NOT. 1000. 
    // Note that it is not the opposite of AND. First bitwise NOT of b is performed. The result of this is operation is used in calculating bitwise AND with a
    
    //bit shifting
    c:=8
    fmt.Println(c<<3) //shifts each bit 3 places to the left => 2^3 * 2^3 = 2^6=64
    fmt.Println(c>>3) //shifts each bit 3 places to the right => 2^3 / 2^3 = 2^0=1
    ```
  - Floating Point numbers \
    IEEE 754 standard \
    32 bit, 64 bit \
    ```
    n:=3.14     //data type will be defaulted to float64
	n = 2.1E14  // 2.1 * 10^14
	
    //both e and E represent (10 raised to power) 
	fmt.Printf("%v,%T\n",n,n)
	
	//arithematic operations with floating point numbers
	a:=10.2
	b:=3.7
	fmt.Println(a+b) // 13.89999999999
	fmt.Println(a-b) // 6.499999999999
	fmt.Println(a*b) // 37.74 
	fmt.Println(a/b) // 2.7567567567567566
    //Note that data type mismatch error does not occur with floats since this time in division we are expected to get a float output.
    ```
	we do not have mod function available for float. Also we cannot apply bitwise logical operations and bit shifting operations on float
  - Complex numbers \
    complex64, complex128 - This is because we are working with either two 32bit floating points or two 64bit floating points \
    zero value is 0+0i 
    ```
	var n complex64 = 1+2i
    fmt.Printf("%v,%T\n",real(n),real(n)) // 1, float 32
	fmt.Printf("%v,%T\n",imag(n),imag(n)) // 2. float 32
	
    //arithematic operations with imaginary numbers
	a:=1+2i
	b:=2+5.2i
	fmt.Println(a+b) // 3+ 7.2i
	fmt.Println(a-b) // -1 -3.2i
 	fmt.Println(a*b) // -8.4 + 9.2i
	fmt.Println(a/b) // 0.39390259295 -0.037138812984i

    n:=complex(5,12) //another method to declare a complex number is by complex function. This defaults to complex128.
	fmt.Printf("%v,%T\n",n,n) // (5+12i), complex128
    ```
- Text
  - String \
    strings are utf8 8bit characters \
    we can treat strings like arrays
    ```
    s:= "this is a string"
    fmt.Printf("%v,%T\n",s[2],s[2])  // 105, uint8
    //this will return an int as strings are aliases for bytes. we'll need to parse it
    fmt.Printf("%v,%T\n",string(s[2]),s[2])  // i, unit8

    s[2] = "u" 
    //this will not work as first strings are immutable plus since "u" is a string and s[2] is a byte therefore we cannot assign a string to a byte so there is a type mismatch
    
    s1 :="this is a string"
    s2 :="this is also a string"
    fmt.Printf("%v,%T\n",s1 + s2, s1 + s2) //string concatenation

    s := "this is a string"
    b := []byte(s) //byte slices
    this function converts each character into an array of bytes. this helps send messages accross servers. we can also use this function to send text documents
    fmt.Printf("%v,%T\n",b,b) // [ascii valuses of each character] , []unit8
    ```
  - Rune \
    runes are utf8 32bit characters \
    Runes are alias for int32 \
	`r:='a'` or `var r rune = 'a'` \
    Any utf8 8bit character can also be a valid utf8 32bit character \
	For declaring runes we use single quotes but for strings we use double quotes \
	`fmt.Printf("%v,%T\n",r,r)` // 97, int32 
   
## Constants
 - Naming constants \
  `const myConst int = 42 ` // use camel casing, change first character to uppercase to make it global \
  `myConst = 45` // compile time error \
  `fmt.Printf("%v,%T\n",myConst,myConst)` //42, int \
- Constants should be assignable at compile time. you cannot do \
  `const myConst float64 = math.Sin(1.57)` // this doesnt compile since it requires function to be executed which is done on runtime where as constants are to be given values before runtime 
- Constants can be of any type i.e. int, string, float, bool, etc.
- Arrays are always going to variable types. you cannot declare an array constant.
- Constants are immutable, but can be shadowed
    ```
    const i int16 = 11

    func main () {
            const i int = 9
            fmt.Printf("%v,%T\n",a,a) // prints 9, int 
            // shadowing by inner most scope
    }
    ```
- Operations 
    ```
    const a = 42 //another method to declare constants
    var b int = 10

	// if const a int = 42, var b int16 = 10 // type mismatch error 

	fmt.Printf("%v,%T\n",a+b,a+b) //prints 52, int
    / /compiler reads it as fmt.Printf("%v,%T\n", 42+b, 42+b)
    //in case of constants, type mismatch error does not occur, if we have not specified data type of constant. Go implicitly parses the constant
    ```
- Enumerated constants
  ```
  package main

  import (
    "fmt"
  )

  // const a = iota 
  // iota is a counter that can be used while creating enumerated constants
  // An enum is a special "class" that represents a group of constants

  // const block can be used to created enumerated constants
  const (
    a = iota
    b = iota
    c = iota

    // if we dont assign the vaule of constant after 1st one, compiler will establish the pattern 
    a = iota
    b // compiler will assign b = iota itself
    c
    
  )

  // iota is scoped to its own constant block
  const (
    a2 = iota
  )

  const (
    _ = iota // Go's only write only variable "underscore"
    // memory is not used
    // if we dont need 0 value of iota

    y 
  )

  func main () {
    fmt.Printf(%v, %T\n, a, a) // prints 0, int 

    // iota changes value as the constants are evaluated
    //fmt.Printf("%v,%T\n",a,a)   // prints 0 
	//fmt.Printf("%v,%T\n",b,b)   // prints 1
	//fmt.Printf("%v,%T\n",c,c)   // prints 2


	//fmt.Printf("%v,%T\n",a2,a2) // prints 0

    var x int
	fmt.Printf("%v\n",x == a2) // prints true since an uninitialised variable = 0 

    fmt.Printf("%v\n",y ) // prints 1 - we can also shift by using _ = iota + 5 so that it prints 6 
  }
  ```
- Bit shifting example
  ```
  package main

  import (
    "fmt"
  )

  const // bit shifting 
    (
    _=iota
    KB=1<<(10*iota)
    MB//all of these constants will be assigned value through bit shifting
    GB
    TB
    PB
    EB
    ZB
    YB
    )
    
  func main (){

    //GO allows us to assign results of arithematic and logical calculations and bit shifting to constants at time of decleration execpt raised to power as it is a function in math package

	fileSize:= 4000000000.
	fmt.Printf("%.2fGB",fileSize/GB) // prints 3.73GB
    //.2f means float will be displayed upto 2 decimal places
    }

  ```

  
  



