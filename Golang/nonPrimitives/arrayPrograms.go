package main

import "fmt"

func main(){
	// students := [3]string{"Omkar", "Brad", "Kumar"}

	// fmt.Println(students)

	// marks := [...]int{21, 23, 44444444, 65}
	// fmt.Println(marks)

	// for i:= 0; i<len(marks); i++{
	// 	fmt.Println(marks[i])
	// }

	// for index, element := range marks{
	// 	fmt.Println(index, "=", element)
	// }


	//sum()
	max()
	min()
	sum()
}

func twosum(){
	nums := [...]int{2, 3, 4, 5, 90, 8}
	target := 95

	arr := [2]int{}

	for i := 0; i<len(nums); i++{
		for j := i+1; j <len(nums); j++{
			if nums[i] + nums[j] == target {
				arr[0] = i
				arr[1] = j
			}
		}
	}
	fmt.Println(arr)
}

func max(){
	arr := [...]int{2,4,6,8, 100}

	max := arr[0]
	for i := 0; i<len(arr); i++{
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Println(max)
}

func min(){
	nums := [...]int{-1, 3, -2, 5}
	min := nums[0]
	for i := 0; i<len(nums); i++{
		if nums[i] < min {
			min = nums[i]
		}
	}
	fmt.Println(min)
}

func sum() {
	arrsum := 0

	arr := [...]int{2, 4, 50, 8, 0}
	for i := 0; i <len(arr); i++{
		arrsum += arr[i]
	}

	fmt.Println(arrsum)
}