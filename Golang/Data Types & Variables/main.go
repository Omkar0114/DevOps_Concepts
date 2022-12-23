package main

import (
	"fmt"
	"reflect"
	"strconv"
)

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
	//  var input string
	//  fmt.Println("Enter your name: ")
	//  fmt.Scanf("%s", &input)
	//  fmt.Println("Your name is: ", input)

 
    // use of reflect function

	number1 := 56
	fmt.Println(reflect.TypeOf(number1))

	string := "omkar is the best"
	fmt.Println(reflect.TypeOf(string))

	// typpe conversion

	var number2 int = 59
	number3 := float64(number2)
	fmt.Printf("%.2f \n", number3)

	m := 90.78
	n := int(m)
	fmt.Printf("%d \n", n)

	// strconv package
    Name := "200"
	strname, err := strconv.Atoi(Name)
	fmt.Printf("%q %T \n", strname, strname)
	fmt.Printf("%q \n", strname, err)

	//int to string
	age3 :=21
	age4 := strconv.Itoa(age3)
	fmt.Printf("%q %T \n", age4, age4)


	// Constants
	// constants are the values that cannot be changed
	// constants are declared using const keyword
	// constants are declared outside the function
	// constants are declared using camel case
	//constants are immutable

	const pi float64 = 3.14
	fmt.Printf("%.2f \n", pi)





}