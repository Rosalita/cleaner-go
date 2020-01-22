package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

//////////////////////////////////////////////
// 1. Comments.
// 1.1 Don't just write some comment
// Do write code comments in gramatically complete sentences and use full stops at the end.

/////////////////////////////////////////////
// 2. Function return arguments.
// 2.1. Don't name return arguments on functions.
func addOne(number int) (result int) {
	result = 5 // Because this automatically creates a named variable that can be assigned to.
	return number + 1
}

// Do this for cleaner code.
func addAnotherOne(number int) int {
	return number + 1
}

//2.2. Don't return an error as an argument that isn't the final argument.
func doThings() (error, string) {
	return nil, "something"
}

// Do this for cleaner code.
func doMoreThings() (string, error) {
	return "something", nil
}

// 2.3. If returning an error and another value, dont return a non-nil value if there is an error.
func doStuff(stuff string) (string, error) {
	if stuff != "fun" {
		return "not fun", errors.New("was not fun")
	}
	return "that was fun", nil
}

// Do this for cleaner code.
func doMoreStuff(stuff string) (string, error) {
	if stuff != "fun" {
		return "", errors.New("was not fun")
	}
	return "that was fun", nil
}

/////////////////////////////////////////////
// 3. Variable names.
// 3.1 Don't use hungarian notation for names which include the variables type.
var urlStr string

// Do this for cleaner code.
var url string

func main() {

	/////////////////////////////////////////////
	// 4. Conditionals.
	// 4.1 Dont put long conditional statements all on one line.
	if url != "" && len(url) > 5 && strings.Contains(url, "https") {
		fmt.Println(url)
	}
	// Do this for cleaner code.
	if url != "" &&
		len(url) > 5 &&
		strings.Contains(url, "https") {
		fmt.Println(url)
	}

	// 4.2 Don't use multiple if statements back to back if only one is required.
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
	//5. Function declaration.
	// Dont declare a function inside a function call.
	err := handle(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("hello world"))
	}))
	// Do declare named functions separately for cleaner code.
	err = handle(http.HandlerFunc(helloWorld))
	fmt.Println(err)

	/////////////////////////////////////////////
	// 6. Structs.
	// 6.1 Don't change values on structs immediately after constructing them.
	person := New()
	person.Name = "foo"
	// Do pass values to constructor functions for cleaner code.
	cleanerPerson := CleanerNew("foo", 0)
	fmt.Println(cleanerPerson.Name)

	/////////////////////////////////////////////
	// 7. Logic.
	// 7.1 Don't use fuzzy logic.
	for i := 0; i < 10; i++ {
		if i != 5 && i != 6 && i != 7 {
			fmt.Println("i is not 5, 6 or 7")
		}
	}
	// Do say conditionals outloud using positive statements if possible for cleaner code.
	for i := 0; i < 10; i++ {
		if i == 5 || i == 6 || i == 7 {
			fmt.Println("i is 5, 6 or 7")
		}
	}

}

func handle(handler http.Handler) error {
	return nil
}

func helloWorld(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("hello world"))
}

func New() *Person {
	person := Person{
		Name: "Someone",
		Age:  99,
	}
	return &person
}

func CleanerNew(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}
