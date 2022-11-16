package main

import (
	"fmt"
)

func main(){

	sum := variableArgument(1,2,3,4,5)
	fmt.Println(sum)

}

// default function form
// func (r receiver) identifier(parameters) (return(s)) {...}

func foo(){
	fmt.Println("with out param")
}

func bar(s string){
	fmt.Println(s)
}

func woo(s string) string{
	return fmt.Sprint("hello", s)
}

func returnTwoValue(first string, last string) (name string, result bool){
	new_name := fmt.Sprint(first, " ", last)
	new_result := true

	return new_name, new_result
}

func variableArgument(x ...int) int{

	sum := 0

	for _, v := range x{
		sum += v
	}

	return sum
}