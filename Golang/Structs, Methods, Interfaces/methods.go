package main

import "fmt"

func main(){

	c := Circle{radius: 5}

	c.calcarea()
	fmt.Printf("%+v\n", c)
}

type Circle struct{
	radius float64
	area float64
}

func (c *Circle) calcarea(){
	c.area = 3.14 * c.radius * c.radius
}