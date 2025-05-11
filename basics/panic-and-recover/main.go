package main
import "fmt"

func main(){
	defer func(){
		if r:=recover(); r!=nil{
			fmt.Println("Recovered from Panic",r)
		}
	}()

	fmt.Println("Before Panic")
	panic("Something went wrong!")
	fmt.Println("After Panic")
}
