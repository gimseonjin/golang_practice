package main

import (
	"fmt"
)

func main(){

	s := "test"

	foo()
	bar(s)

	name, r := returnTwoValue("kim", "seonjin")

	fmt.Println(name, r)
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