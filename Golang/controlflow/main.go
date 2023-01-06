package main

import "fmt"

func main(){
	// if - else statements
	age := 18
	if age < 18 {
		fmt.Println("You are not eligible for vote!")
	} else {
		fmt.Println("You are eligible for vote, congrats!")
	}

	//if-else if -else:
	fmt.Print("Enter the fruit name : ")
	var fruit string
	fmt.Scanf("%s", &fruit)
	if fruit == "apple" || fruit == "apples"{
		fmt.Println("I love apples!")
	} else if fruit == "orange"{
		fmt.Println("Oranges are nice!")
	} else if fruit == ""{
		fmt.Println("Eat any fruit once a day")
	} else {
		fmt.Println("Thnak you :)")
	}

	switch{
	case fruit == "Grapes":
		fmt.Println("fruit is grapes")
	case fruit == "apple":
		fmt.Println("Apple")
	default:
		fmt.Println("None")
	}
}