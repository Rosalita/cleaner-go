package main

import (
	"fmt"
	"net/http"
	"strings"
	//	"errors"
)

//////////////////////////////////////////////
// 1. Comments.
// Don't
// some comment
// Do
// Write code comments in gramatically complete sentences and use full stops at the end.


/////////////////////////////////////////////
// 2. Function return arguments.
// Don't name return arguments on functions.
func addOne(number int) (result int) {
	result = 5 // Because it automatically creates a named variable that can be assigned to.
	return number + 1
}
// Do this for cleaner code.
func addAnotherOne(number int) int {
	return number + 1
}

/////////////////////////////////////////////
// 3. Variable names.
// Don't use hungarian names which include the variables type.
var urlStr string
// Do use variable names which don't include types.
var url string


func main() {

	/////////////////////////////////////////////
	// 4. Long conditionals.
	// Dont put long conditional statements all on one line.
	if url != "" && len(url) > 5 && strings.Contains(url, "https") {
		fmt.Println(url)
	}
	// Do this for cleaner code
	if url != "" &&
		len(url) > 5 &&
		strings.Contains(url, "https") {
		fmt.Println(url)
	}

	
    /////////////////////////////////////////////
	// 5. multiple conditional statements.
	// Don't use multiple if statements back to back if only one is required
	name := "Rosie"
	if name == "Rosie" {
		fmt.Println("You Rock!") // if name was changed to Duncan here, it would drop through.
	}
	if name == "Duncan" {
		fmt.Println("You Roll!")
	}
	// Do use an anonymous switch for cleaner code.
	switch {
	case name == "Rosie":
		fmt.Println("You Rock!")
	case name == "Duncan":
		fmt.Println("You Roll")
	}


	/////////////////////////////////////////////
	//6. Function declaration inside a function call
	// Dont declare a function inside a function call.
	err := handle(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("hello world"))
	}))
	// Do declare named functions separately for cleaner code.
	err = handle(http.HandlerFunc(helloWorld))
	fmt.Println(err)


	result := addOne(3)
	fmt.Println(result)
}

func handle(handler http.Handler) error {
	return nil
}

func helloWorld(rw http.ResponseWriter, req *http.Request){
	rw.Write([]byte("hello world"))
}
