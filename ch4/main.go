package main

import (
	"fmt"
)

type person struct{
	first string
	second string
}

type secretAgent struct{
	person
	ltk bool
}

func main(){

	p1 := secretAgent{
		person: person{
			first : "kim",
			second : "seonjin",
		},
		ltk : true, 
	} 

	p2 := person{
		first : "miss",
		second : "jin",
	}

	fmt.Println(p1, p2.first)
}