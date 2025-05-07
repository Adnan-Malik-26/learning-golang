package main

import (
	"fmt"
)

func main(){
	var num1 float64
	var num2 float64
	var choice int

	fmt.Println("***************************************")
	fmt.Println("**************CALCULATOR***************")
	fmt.Println("Give 2 numbers separated by a space")
	fmt.Scanln(&num1, &num2)
	fmt.Println("1. For Addition\n2. For Subtraction\n3. For Multiplication\n4. For Division")
	
	start:
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		sum:=num1+num2
		fmt.Println("num1+num2=",sum)
	case 2:
		diff:=num1-num2
		fmt.Println("num1-num2=",diff)
	case 3:
		product:=num1*num2
		fmt.Println("num1*num2=",product)
	case 4:
		quotient:=num1/num2
		fmt.Println("num1/num2=",quotient)
	default:
		fmt.Println("Wrong choice try again")
		goto start
	}
}
