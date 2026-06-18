package main

import (
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("Hello, Go Standard libray")	
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("HTTP Response Status:", res.Status)
}