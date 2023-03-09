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
	//fmt.Println(addTwo(2, 5))

	// function with return type
	//fmt.Println(greet("Omkar"))

	// function with multiple return values
	fmt.Println(sum(10, 12, 14, 15))
	fmt.Println(sum())
	fmt.Println(sum(10))


	sum, product := sumAndProduct(10, 20)
	fmt.Println(sum, product)

}

func addTwo(a int, b int) int{
	return a + b
}

func greet(name string)string{
	return "Hello " + name
}

// we can also declare function with return type
// we can return multiple return values

func sum (nums ...int)int {
	ans := 0
	for _, num := range nums{
		ans += num
	}
	return ans
}

// function with multiple return values
func sumAndProduct(a int, b int)(int, int){
	return a + b, a * b
}