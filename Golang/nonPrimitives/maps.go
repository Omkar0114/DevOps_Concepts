package main

import "fmt"

func main() {
	// Maps are like dictionaries in Python
	// Maps are like hash tables in C++
	// Maps are like associative arrays in PHP
	// Maps are like objects in JS
	// Maps are like HashMaps in Java

	mymap := map[int]string{1: "omkar", 2:"sagar", 3:"sachine"}
	fmt.Println(mymap)

	//using make function
	map1 := make(map[string]string)
	fmt.Println(map1)

	// adding elements to map
	map1["name"] = "brad"
	map1["age"] = "37"
	map1["address"] = "123 main st"
	fmt.Println(map1)

	fmt.Println(len(map1))

	delete(map1, "age")
	fmt.Println(map1)

	fmt.Println(mymap[1])
}