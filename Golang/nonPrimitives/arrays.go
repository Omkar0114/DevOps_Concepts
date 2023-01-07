package main

import (
	"fmt"
)

func main(){
	// arrays declaration
	var fruits [3]string = [3]string{"apple", "orange", "grapes"}
	fmt.Println(fruits)

	// short-hand declaration
	marks := [4]int{2, 3, 4, 5}
	fmt.Println(marks)

	// declaration without specifying length of the array
	nums := [...]float32{1.2, 4.5, 5.5, 6.7}
	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(nums[3])

	// changing the value from an array
	grades := [5]int{10,100, 200, 300, 500}
	fmt.Println(grades)
	grades[2] = 999
	fmt.Println(grades)

	// arrays in for loop
	students := [3]string{"john", "Brad", "Marino"}
	for i := 0; i < len(students); i++{
		fmt.Println(students[i])
	}

	// using range keyword
	for index, element := range students{
		fmt.Println(index, "=>", element)
	} 


	// multidimentional arrays
	arr := [3][2]int{{2, 1}, {3, 5}, {6, 7}}
	fmt.Println(arr[2][1])
	for index, element := range arr{
		fmt.Println(index, "=>", element)
	}

	arrr := [5]bool{true, true, true, true}
        for i := 0; i < len(arr); i++ {
                if arrr[i] {
                        fmt.Println(i)
                }
        }

	arrbool := [2]bool{}
	fmt.Println(arrbool)
}