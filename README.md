## For installation on mac

__Step 1:__
Terminal: `brew install go` \
**Step 2:**
VS code: create file main.go, vs code will suggest to install go extensions, click install and then install all \

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
   
  
  
  



