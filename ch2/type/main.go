package main

import "fmt"

var x bool

func main() {

	a := 7
	b := 42

	fmt.Println(a == b)

	x = 10 < 5

	fmt.Println(x)
}
