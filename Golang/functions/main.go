package main

import "fmt"

func main() {
	// function is a block of code that can be called by name
	// function can take input and return output
	// function can be passed as parameter to another function
	// function can be returned from another function
	// function can be assigned to a variable
	// function can be declared inside another function
	// function can be declared anonymously

	// function declaration
	// func function_name( [parameter list] ) [return_types] {
	// 	// function body
	// }

	// function call
	fmt.Println(addTwo(2, 5))

	// function with return type
	fmt.Println(greet("Omkar"))

}

func addTwo(a int, b int) int{
	return a + b
}

func greet(name string)string{
	return "Hello " + name
}