package main 

import "fmt"

func main(){
	// operators
	// comparison operators
	// ==, !=, <, >, <=, >=
    // a := 45
	// b := 89
	// fmt.Println(a == b)
	// fmt.Println(a != b)
	// fmt.Println(a < b)
	// fmt.Println(a > b)

	// name := "Omkar"
	// name2 := "vedant"

	// fmt.Println(name == name2)

	// Arithmatic operators
	// +, -, *, /, %
	x , y := 10, 34
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)
	x++
	y--
	fmt.Println(x)
	fmt.Println(y)

	str1 , str2 := "one", "two"
	fmt.Println(str1 + str2)

	// var a float64 = 5.9
	// a++
	// fmt.Print(a)

	var a, b float64 = 24.4, 3.0
        fmt.Println(a / b)
        fmt.Println(int(a) % int(b))



}