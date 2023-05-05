package main

import "fmt"

type shape interface{
	area() float64


}

type Circle struct{
	radius float64
	 
}

type rect struct{
	length , breadth float64
	
}

func (c Circle) area() float64{
	return 3.14 * c.radius * c.radius
}

func (r rect) area() float64{
	return r.length * r.breadth
}

func printData (s shape,n string){
	
	fmt.Printf("area of %v=%v\n",n, s.area())
}


func main(){

	r := rect{length: 10, breadth: 30.98 }
	c := Circle{radius: 56}
	printData(r,"rect")
	printData(c, "circle")

}
