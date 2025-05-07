package main

import (
	"fmt"
)

func main(){
	const pi = 3.14

	fmt.Println("Area Calculator")
	fmt.Println("1. For Square\n2.For Rectangle\n3.For Circle\n4.For Triangle")

	var choice int
	start:
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		var side int
		println("Give side of Square")
		fmt.Scanln(&side)
		area:=side*side
		println("Area=",area)
	case 2:
		var length, breadth int
		println("Give me length and breadth of Rectangle")
		fmt.Scanln(&length,&breadth)
		area:=length*breadth
		println("Area=",area)
	case 3:
		var radius float64
		fmt.Println("Give radius of circle")
		fmt.Scanln(&radius)
		area:=pi*radius*radius
		println("Area=",area)
	case 4:
		var base,height float64
		fmt.Println("Give base and height of Triangle")
		fmt.Scanln(&base,&height)
		area:=0.5*base*height
		println("Area=",area)
	default:
		fmt.Println("Wrong choice try again")
		goto start
	}

}
