package main

import (
	"fmt"
	"runtime"
)

var x bool

func main() {

	a := 7
	b := 42

	fmt.Println(a == b)

	x = 10 < 5

	fmt.Println(x)

	y := 12.3
	fmt.Printf("%T\n", y)

	var test uint
	fmt.Printf("%T\n", test)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
