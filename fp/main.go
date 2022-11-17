package main

import (
	"fmt"
)

type fn func(int) int

func main() {

	x := func(a int) int{
		return a + 10
	}

	fmt.Println(test(x))

}

func test(x fn) int {
	return x(10)
}