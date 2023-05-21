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

## Arrays
- Arrays are collections of items with the same type 
- Fixed size  
- Initialization \
  `grades := [3]int {97, 85, 93}` // size of array, data type, values \
  `fmt.Printf("Grades: %v",grades)` // prints Grades: [97 85 93] \
  or `grades := [...]int {97, 85, 93}` // [...] means the size of the array should be just enough to hold the given elements in it \
  `var students [3]string` // empty array \
  `students [0] = "Lisa"` // a student has been added to the empty array
- An array can be of any type but each element has to be of the same type
- Array elements are stored in contiguous memory blocks GO sets a pointer to the first element of array and uses the index we pass to determine how many elements the pointer has to move forward
- Length of array `len(students)`
- Array of arrays
  `var identitymatrix [3][3]int = [3][3]int{[3]int{1,0,0},[3]int{0,1,0},[3]int{0,0,1}}` \
  `fmt.Printf("Identity Matrix: %v\n",identitymatrix)` \
  or
  ```
  var identitymatrix [3][3]int
  identitymatrix[0] = [3]int{1,0,0}
  identitymatrix[1] = [3]int{0,1,0}
  identitymatrix[2] = [3]int{0,0,1} 
  ```
- Arrays in Go work as values. we can declare an array as variable. 
- When we copy an array to another variable, it poitns to another set of data 
  ```
    a := [...] int {1, 2, 3}
    b := a
    b[1] = 5 
    fmt.Println(a) // prints [1 2 3]
    fmt.Println(b) // prints [1 5 3]
  ```
  We can use pointers, by using the "address of" literal `&`   
  ```
    a := [...] int {1, 2, 3}
    b := &a // now b is pointing to the address of a 
    b[1] = 5 
    fmt.Println(a) // prints [1 5 3]
    fmt.Println(b) // prints [1 5 3]
  ```
## Slices
- Just a projection of the underlying array - reference type 
- Initialization
  ```
    a := [] int {1, 2, 3} // size is not defined
    fmt.Println(a) // prints [1 2 3]
    fmt.Printf("Length: %v\n" , len(a)) // prints Length: 3
    fmt.Printf("Capacity: %v\n" , cap(a)) // prints Capacity: 3 
    // Capacity might be much larger than the number of elements inside of an array
  ```
- Dont need to use pointers
  ```
    a := [] int {1, 2, 3}
    b := a 
    b[1] = 5 
    fmt.Println(a) // prints [1 5 3]
    fmt.Println(b) // prints [1 5 3]
  ```
- Slicing operations - can be used on arrays as well 
  ```
  // indexing starts from 0
  // first number is inclusive index and second number is exclusive index
  // ':' creates slices. slicing operations work with arrays too
  a := []int{1,2,3,4,5,6,7,8,9,10}
  b := a[:]   //slice of all elements
  c := a[3:]  //slice from 4th element to end
  d := a[:6]  //slice of first 6 elements
  e := a[3:6] //slice of 4th,5th and 6th element
  
  fmt.Println(a) // prints [1,2,3,4,5,6,7,8,9,10] 
  fmt.Println(b) // prints [1,2,3,4,5,6,7,8,9,10] 
  fmt.Println(c) // prints [4,5,6,7,8,9,10] 
  fmt.Println(d) // prints [1,2,3,4,5,6] 
  fmt.Println(e) // prints [4,5,6] 

  a[5] = 42   //since all these slices point to the same data changing a value in one slice will reflect in all the other slices

  fmt.Println(a) // prints [1,2,3,4,5,42,7,8,9,10] 
  fmt.Println(b) // prints [1,2,3,4,5,42,7,8,9,10] 
  fmt.Println(c) // prints [4,5,42,7,8,9,10] 
  fmt.Println(d) // prints [1,2,3,4,5,42] 
  fmt.Println(e) // prints [4,5,42]
  ```
- make() function 
  `a := make([]int,3,100)` // type, length, capacity \
  `fmt.Printf("Length: %v\n" , len(a))`   // prints Length: 3 \
  `fmt.Printf("Capacity: %v\n" , cap(a))` // prints Capacity: 100 
- Unlike arrays, slices dont need to have a fixed size in their entire lifetime
  ```
  a := []int {} // type, length, capacity 
  fmt.Println(a) // prints []
  fmt.Printf("Length: %v\n" , len(a))`   // prints Length: 0
  fmt.Printf("Capacity: %v\n" , cap(a))` // prints Capacity: 0 

  a = append (a, 1) 
  //creates a new array with larger size and then copies into it the elements of the previous array and then re-assigns it to the same pointer

  fmt.Println(a) // prints [1]
  fmt.Printf("Length: %v\n" , len(a))`   // prints Length: 1
  fmt.Printf("Capacity: %v\n" , cap(a))` // prints Capacity: 2 
  ```
- Variatic function \
  `a = append(a,2,3,4,5)` \
  append function can take more than 2 parameters. This type of function is known as a variatric function. Everything after the the first parameter is interpreted as a value to interpret.
	```
    fmt.Println(a) // prints [1 2 3 4 5]
	fmt.Printf("Length: %v\n",len(a)) // prints Length: 5
	fmt.Printf("Capacity: %v\n",cap(a)) // prints Capacity: 8 
    // It is not clear as to how GO resizes the array but generally if the underlying array is filled GO creates a new array with double the size 
    ```
- Like an array each element in a slice has to be of the same type therefore we cannot append a slice into a slice of integer regardless of the slice being appended containg integers so `a = append(a, []int{2,3,4,5})` will not work \
  Spread operator "..." decomposes a slice into its elements so `a = append(a,[]int{2,3,4,5}...)` will work
- Stack operations - slices can be treated as stacks
	```
    a := []int{1,2,3,4,5}
	//append function allows us to push elements to the top of the stack
	b := a[1:]    //this can also operate as shift opertor. It removes the very first element from the array
	c := a[:len(a)-1]     // we can use this to remove last element of the array
	d := append(a[:2],a[3:]...)   // we can use this to remove the middle element from the array

	fmt.Println(b) // prints [2 3 4 5]
	fmt.Println(c) // prints [1 2 3 4]
	fmt.Println(d) // prints [1 2 4 5]
    ```
    Caution must be taken in creating multiple references to the same slice while using a command such as `d := append(a[:2],a[3:]...)` because otherwise we get strange behaviors and since we do not have any way to keep the data same we must create a copy the original slice using loops before creating a reference

## Maps
- Declare using the map literal \
  `mapName := map[string]int `// map[type of key] type of value 
    ``` 
    package main

    import (
        "fmt"
    )

    func main() {
        statePopulation := map[string]int{
        // this syntax suggests that every key is of string type and every value is of int type

        // type should be consistant for every key-value pair in a map
        
        "California":39250017,
        "Texas":27862596,
        "Florida":20612039,
        "New York":19745289,
        "Pennsylvania":12802503,
        "Illinois":12801539,
        "Ohio":11614373,
        }

        fmt.Println(statePopulation)
    }
   ```
- Slices cannot be keys to a map \
  `m := map[[]int]string` // invalid \
  but arrays can be keys to a map \
  `m := map[[...]int]string`
- Using the make() function  
  ```
  mapName := make(map[string]int) 
  statePopulation = map[string]int{
       "California":39250017,
       "Texas":27862596,
       }
  ```
- Manipulation of values in a map \
  Get one value from a map using [] \
  `fmt.Println(statePopulation["Ohio"])` // prints 11614373 \
  Add one value to a map \
  `statePopulation["Georgia"] = 10310371` \
  `fmt.Println(statePopulation["Georgia"])` // prints 10310371 \
- Return order of a map is not guaranteed
- Delete entries from a map using builtin delete function : `delete(mapName, "key")` \
  `delete(statePopulation, "Georgia")`
- When returnung a value, if we mispell a key or enter a deleted key then the returned value will always be 0
- Use the ,ok syntax to see if a key exists in the map \
  ok returns false if key does not exist otherwise it returns true \
  ok is not a keyword but it is a convention to not use ok for variable name
  ```
    population, ok:= statePopulation["Georgia"]
    fmt.Println(population,ok) // prints 0 false

    //to just check for presence we can use the following syntax
    //_, ok = statePopulation["Ohio"]
    fmt.Println(ok) // prints true
  ```
- Find out the number of keys inside of map using len function \
`fmt.Println(len(statePopulation))`
- Maps are passed by reference \
  if we initialize a variable with another map variable then both variables will point to the same map. It also applies to when passing maps to a function
  ```
    sp := statePopulation
    delete (sp, "Ohio")
    fmt.Println(sp)
    fmt.Println(statePopulation)

    // ohio is deleted from both
  ```
## Struct
- Structure is also a collection type - it gathers similar info together. \
  It doesnt need to have the same type of data 
- Creation
  ```
    package main 

    import (
        "fmt"
    )

    type Doctor struct {
        number int
    	actorname string
    	episodes []string
        companions []string
    }

    func main(){
        aDoctor := Doctor {
            number : 3,
            actorName: "Jon",
            companions: []string{
                "liz" ,
                "jo", 
                "sarah"
            }
        }
        fmt.Println(aDoctor) 
        // prints { 3 Jon [liz jo sarah]}
    }
  ```
- Dot syntax - to get one value from a struct 
  `fmt.Println(aDoctor.name)` // prints Jon
- Positional syntax to instantiate a struct
  ```
  aDoctor := Doctor {
            3,
            "Jon",
            []string{
                "liz" ,
                "jo", 
                "sarah"
            }
        }
  ```
  but not recommended, because it can become a maintenance problem. In case of any changes to struct body Go will not map the values to fields properly so we will be required to add a placeholder every time
- Naming conventions
  same rules apply to struct when declaring as other variables \
  - If first letter is Capital then the struct will be exported outside of the package \
  - If the field names start with a capital letter only then will the other packages be able to see the fields of the struct that is being exported \
  - Go does not allow _ to be in the name of struct and the field names
- Anonymous struct
  ```
    func main(){
        aDoctor := struct {name string} {name : "Jon"} // structure of struct ,initializer which provides data into the struct
    }
  ```
  rarely used 
- Structs are value types \
  we can assign a struct to another struct, but maniplation will only be done in the new one, since it is value type 
  ```
    anotherDoctor := aDoctor
    anotherDoctor.name = "Tom"
    fmt.Println(aDoctor)        // prints Jon
    fmt.Println(anotherDoctor)  // prints Tom
  ```
  but we can create pointer to struct using  `&`
  ```
    anotherDoctor := aDoctor
    anotherDoctor.name = "Tom"
    fmt.Println(aDoctor)        // prints Tom
    fmt.Println(anotherDoctor)  // prints Tom
  ```
- Embedding
  - Go does not support the traditional rules of OOP for example it does not have inheritance \
  - Instead Go has a model similar to inheritance called composition through embedding 
    ```
        type Animal struct{
            Name string 
            Origin string
        }

        type Bird struct{
            Animal // embedding animal struct in bird struct
            SpeedKPH float32
            CanFly bool
        }

        func main(){
            //
            b := Bird{}

            b.Name = "Emu"
            b.Origin = "Australia"
            b.SpeedKPH = 48
            b.CanFly = false
            fmt.Println(b) // prints {{Emu Australia} 48 false}
            fmt.Println(b.Name) // prints Emu
        }
    ```
  - We can access the fields of an embedded struct the same way we access the fields of a normal struct that is by using . notation
  - Composition does not work like inheritance i.e. Bird is not type Animal so they cannot be used interchangeably. Bird is an independant struct. Bird has Animal like characteristics only
  - To allow interchangeable usage we use interfaces
  - To use literal syntax
    ```
        b := Bird{
            Animal : Animal{Name:"Emu",Origin:"Austrailia"},

            //incase of literal syntax, we need to have some knowledge of the internal structure of the struct that has embedding. We have to explicitly initialize the embedded struct

            SpeedKPH : 48,
            CanFly : false,
        }
        fmt.Println(b.Name)  // prints 
    ```
  - Embedding is a not a good choice for modelling behaviors for that we use interfaces but to provide accessibility to custom structs we use embedding
- Tags \
  Tagging is a concept used when we have to define some specific rules for a field inside of a struct \
  Space delimited subtags are used

  ```
    package main

    import (
        "fmt"
        "reflect"
         // we have to use reflect package to get tags which have been attached to a field
    )

    type Animal struct {
        Name string `required max:"100" ` 
        // so the name field is required and is not more than 100 char
        Origin string
    }

    func main(){
        t := reflect.TypeOf(Animal{})  // type of an object 
        field, _ := t.FeildByName("Name") // FieldByName function returns a field 
        fmt.Println(field.Tag) // Tag function simply returns a string of Tags. It is for the validation function to see what to do with the Tags 
    }
  ```

## Control Flow - if statements
- if statement
  ```
    package main

    import (
        "fmt"
    )

    func main() {
    	if true {
            //{} are necessary even in case of one statement 
    		fmt.Println("The test is True!")
    	}
    }
  ```
- Initializer syntax
  ```
    statePopulation := map[string]int{
        "california":39250017,
        "Texas":27862596,
        "Florida":20612039,
        "New York":19745289,
        "Pennsylvania":12802503,
        "Illinois":12801539,
        "Ohio":11614373,
	}
	
	if pop,ok := statePopulation["Florida"]; ok {
    		fmt.Println(pop)
    }

    fmt.Println(pop) // prints undefined
  ```
  The statement before the `;` is the initializer statement which uses a function call to generate a boolean result and store it in a variable `ok` \
  The variable `ok` is typed after the `;` and if statement runs on basis of the value stored in `ok` \
  `pop` and `ok` variables are locked to the scope of if statement i.e. they are inaccessible outside the scope of if statement
- Comparison operators
    ```
        number := 50
        guess := 70

        if guess < number{
            fmt.Println("too low!")
        }
        if guess > number{
            fmt.Println("too high!")
        }
        if guess == number{
            fmt.Println("You got it!")
        }

        fmt.Println(number<=guess,number>=guess,number!=guess) // true false true 
        // these only workfor int types, not string types

        // string type works with == or !=  
    ```

    ```
        if guess < 1{
    		fmt.Println("guess should be greater than or equal to 1")
    	}else if guess > 100{
    		fmt.Println("guess should be less than or equal to 100")
    	}else{
            if guess < number{
                fmt.Println("too low!")
            }
            if guess > number{
                fmt.Println("too high!")
            }
            if guess == number{
                fmt.Println("You got it!")
            }
    	}
        
        fmt.Println(!true) // prints false
    ```
    the `else` and `else if` statement should be written in the same line as the ending brace of the previous `if` or `else if` statement
- Short circuiting
    ```
        if guess<1 || guess>100{
            fmt.Println("Guess should be between 1 and 100")
        }
    ```
    In case of multiple OR, go checks the boolean value returned by each statement in order they are written. If the first statement returns true then the remaining conditions are ignored since end result will always be true. This is called shortcicuiting in Go. 

- equality operators
    ```
        import (
            "fmt"
            "math"
        )

        func main() {

            myNum := 0.1234567
            if myNum == math.Pow(math.Sqrt(myNum),2) {
                fmt.Println("These are same!")
            }else{
                fmt.Println("These are different!")
            }

            // Go uses floating points to represent decimal values which are an approximation of the original values therefore when Go calculates the square root and then calculates exponent of the value the new value turn out to be different. Better approach is to introduce an error and check if the error is less than a certain threshold, so 

            myNum := 0.1234567
            if math.Abs(myNum / math.Pow(math.Sqrt(myNum),2)-1) < 0.001 {
                fmt.Println("These are same!")
            }else{
                fmt.Println("These are different!")
            }
        }
    ```
## Control Flow - switch statements
- simple cases
- multiple tests
- falling through
- type switches



