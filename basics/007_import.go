// 1)
// package main

// import (
// 	"fmt"
// 	// "net/http"
// 	foo "net/http"
// )

// func main() {
// 	fmt.Println("Hello, Go standard library")
// 	// res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
// 	res, err := foo.Get("https://jsonplaceholder.typicode.com/posts/1")

// 	if err != nil {
// 		fmt.Println("Error", err)
// 		return
// 	}
// 	defer res.Body.Close()

//		fmt.Println("HTTP Response status: ", res.Status)
//	}
//
// -----------------------------------------
// 2)
package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Go standard library")
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	defer res.Body.Close()

	fmt.Println("HTTP response status: ", res.Status)
}
