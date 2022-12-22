package main

import "fmt"

func main (){
	var s string = "Hello world"

	var i int = 10

	var f float64 = 70.89

	var b bool = true

	fmt.Println(s, i, f, b)

	// var greet string = "Let's learn Golang"
	// var name string = "omkar"
	// fmt.Print("Hey", name , "," ,  greet, "\n")


	// format specifier 
	var name string = "john"
	var age int = 21

	fmt.Printf("My name is %v and my age is %d", name, age)
	fmt.Println()
	
	var h float32 = 899.09
	fmt.Printf("type of h is %T", h)
	fmt.Println()


	// shorthand declaration
	var (
		name1 string = "Amit"
		age2 int = 78
	)

	fmt.Println(name1, age2)

	language := "Golang"
	Trade := "DevOps"

	fmt.Println(language, Trade)
	fmt.Print("Hey there \n")

	 // input from user
	 var input string
	 fmt.Println("Enter your name: ")
	 fmt.Scanf("%s", &input)
	 fmt.Println("Your name is: ", input)



}