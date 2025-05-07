package main

import (
	"fmt"
)

func main(){
	const pi = 3.14
	const(
		monday = iota 
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)
	var radius float64	
	fmt.Scanln(&radius)
	area := pi*radius*radius

	fmt.Printf("The area of circle with radius %v is %v", radius, area)
	fmt.Println(monday,"\n", tuesday,"\n", wednesday,"\n",thursday,"\n",friday,"\n",saturday,"\n",sunday )
}
