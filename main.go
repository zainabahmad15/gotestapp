package main // add packages, main is the starting point

import (
	"fmt"
	"runtime"
	//"sync"
	//"time"
)

// "io"
// "log"
//"net/http"

// // defer
// func main() {
// 	// fmt.Println("Hello, it's me")
// 	res, err := http.Get("http://www.google.com/robots.txt") // request the robots.txt file
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer res.Body.Close() // we can defer this so that some other code can work in the function
// 	// defer this now, so we dont forget to close it later

// 	robots, err := io.ReadAll(res.Body) // this is the only working statement here

// 	// res.Body.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("%s", robots)
// }

// web page code
// func main() {
// 	// function listener , with a callback
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Hello Go!"))
// 	})
// 	err := http.ListenAndServe(":8080", nil)
// 	// if the port is blocked, there would be an error and the app would go into panic
// 	// so the developer can add a panic in the code, so the application doesnt break down
// 	// and instead just panics and the error is printed
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

// GO ROUTINES

// // one way to make a go routine
// func main(){

// 	// sayHello() // calling a function
// 	//to make a go routine
// 	go sayHello() // tells Go to spin off on a green thread
// 	// this wouldnt print the message, but we will put sleep here
// 	time.Sleep(100 * time.Millisecond)
// }

// func sayHello(){
// 	fmt.Println("Hello")
// }

// // another way to make a go routine
// func main(){
// 	var msg = "Hello"
// 	go func(){   // anonymous funtion
// 		fmt.Println(msg) //prints Goodbye
// 	}()  //invoked immediately
//  msg = "Goodbye"
// 	time.Sleep(100 * time.Millisecond)
// }

// // another way to make a go routine
// //use arguments to send data to your goroutines so that data doesnt get lost
// func main(){
// 	var msg = "Hello"
// 	go func(msg string){
// 		fmt.Println(msg)   // so prints Hello
// 	}(msg)  //invoked immediately

// 	msg = "Goodbye"
// 	time.Sleep(100 * time.Millisecond)
// }

// //not a good practice to use sleep
// // so we use wait groups
// var wg = sync.WaitGroup{}
// func main(){
// 	var msg = "Hello"
// 	wg.Add(1)
// 	go func(msg string){
// 		fmt.Println(msg)   // so prints Hello
// 		wg.Done()   // decrement the waitgroup by 1
// 	}(msg)  //invoked immediately

// 	msg = "Goodbye"
// 	wg.Wait() // waits for the execution of the goroutine
// }

// // multiple go routines
// var wg = sync.WaitGroup{} // safe to use globally
// var counter = 0

// func main() {
// 	for i := 0; i < 10; i++ {
// 		wg.Add(2)
// 		go sayHello()
// 		go increment()
// 	}
// 	wg.Wait()
// }

// func sayHello() {
// 	fmt.Printf("Hello #%v\n", counter)
// 	wg.Done()
// }

// func increment() {
// 	counter++
// 	wg.Done()
// }

// // there is no synchronization between the go routines, so to do that we have mutex
// // which is a lock that the application is going to honor
// var wg = sync.WaitGroup{} // safe to use globally 
// var counter = 0
// var m = sync.RWMutex{} // read write mutex - lock unclock so one thing can access data at one time
// // RW , infinite reader, only one writer 
// func main() {
// 	runtime.GOMAXPROCS(100) 
// 	for i := 0; i < 10; i++ {
// 		wg.Add(2)
// 		// lock the mutex in a single context, and then asyncronously in the separate routines
// 		m.RLock() // read lock 
// 		go sayHello()
// 		m.Lock() // write lock 
// 		go increment()
// 	}
// 	wg.Wait()
// }

// func sayHello() {
// 	fmt.Printf("Hello #%v\n", counter) // print message 
// 	m.RUnlock()
// 	wg.Done()
// }

// func increment() {
// 	counter++
// 	m.Unlock() // write unlock 
// 	wg.Done() 
// }
// // but this is destroying the parallelism 








