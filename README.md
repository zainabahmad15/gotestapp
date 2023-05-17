## For installation on mac

### Step 1: 
Terminal: brew install go 
### Step 2: 
VS code: create file main.go, vs code will suggest to install go extensions, click install and then install all
### Step 3:
Create directories

`go run main.go`

## Good tutorial
https://www.youtube.com/watch?v=YS4e4q9oBaU

## Variables
3 ways to declare variables

`var i int` variable i of type integer 

when you want to declare a variable but dont want to initialize

`i = 42` // initialize later

`var j int = 27`

// if go doesnt have ther info to assign a type that we want 

`fmt.Printf("%v, %T", j, j)`

// 27, int (value, type)

`var j float32 = 27`

`fmt.Printf("%v, %T", j, j)`

// 27, float32

`k:= 99` // compiler can figure out the data type

`fmt.Printf("%v, %T", k, k)`

// 99 , int

`k:= 99.`

`fmt.Printf("%v, %T", k, k)`

// 99 , float64 //wont print float32, unless defined

if delared at package level, only use the first method

`var i int = 42 `

then it can be used in function