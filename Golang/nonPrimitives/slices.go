package main

import "fmt"

func main(){
	arr := [10]int {1,2,3,4,5, 6, 7, 8, 9, 10}
	// Slices are like pointers to arrays
	// Slices are like references to arrays
	// Slices are like views to arrays
	slice := arr[0:8]

	fmt.Println(slice)
	sub_slice := slice[0:5]
	fmt.Println(sub_slice)


	// using make function
	// make(type, length, capacity)
	slice2 := make([]int, 5)

	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	// append function
	slice2 = append(slice2, 1,2,3)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))


	arr1 := [5]int{5, 10, 15, 20, 25}
	slice3 := arr1[0:2]

	fmt.Println(slice3)

	arr2 := [4]int{40,45,50,55}
	slice4 := arr2[0:2]

	fmt.Println(slice4)

	new_slice := append(slice3, slice4...)
	fmt.Println(new_slice)
}