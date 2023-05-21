package main // add packages, main is the starting point

import (
	// "io"
	// "log"
	"net/http"
)

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
func main() {
	// function listener , with a callback
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	// if the port is blocked, there would be an error and the app would go into panic
	// so the developer can add a panic in the code, so the application doesnt break down
	// and instead just panics and the error is printed
	if err != nil {
		panic(err.Error())
	}
}
