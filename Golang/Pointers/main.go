package main 

import "fmt"

func main(){
	// pointers are used to share memory address
	// pointers are used to pass references to functions
	// pointers are used to return multiple values from functions
	// pointers are used to return references from functions
	// pointers are used to implement data structures like linked list, trees, graphs etc

	// pointer declaration
	// var var_name *var-type

	 i := 10
	var ptr *int = &i

	fmt.Println(ptr)

	fmt.Println(*ptr)

	fmt.Printf("%T %v \n", &i, &i)

	// maps are used to store key value pairs
	// maps are by default reference type for pointers
	// slices are used to store multiple values of same type
	// slices are by default reference type for pointers


	fmt.Printf("%T %v \n", *(&i), *(&i))



}