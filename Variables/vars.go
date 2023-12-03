package main

import "fmt"

func main() {

	// var assigns text to tha a variable
	var a = "initial"
	fmt.Println(a)

	// var can assign multiple variables
	// var can infer the type of the variable
	var b, c int = 1, 2
	fmt.Println(b, c)

	// var can infer the type of the variable
	var d = true
	fmt.Println(d)

	// var can assign a zero value
	var e int
	fmt.Println(e)

	// := is a shorthand for declaring and initializing a variable
	f := "apple"
	fmt.Println(f)
}
