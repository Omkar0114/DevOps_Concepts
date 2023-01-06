package main

import "fmt"

func main(){
	// for i := 1; i <= 5; i++{
	// 	fmt.Println(i*i)
	// }

	// i := 3
    //     for i > 10 {
    //             fmt.Println(i * 2)
    //             i += 1
    //     }

		for i := 0; i <= 5; i++ {
			fmt.Println(i * i)
			if i == 3 {
					break
			}
	}
}