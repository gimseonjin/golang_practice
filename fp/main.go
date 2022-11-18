package main

import (
	"fmt"
)

func main() {

	x := func(a int) int {
		return a + 10
	}

	test := []int{1,2,3,4,5}

	result := mapper(x, test)

	fmt.Println(result)

}

func mapper(f func(a int) int, iter []int) []int {
	result := make([]int, len(iter), len(iter))
	for i, v := range iter{
		result[i] = f(v)
	}
	return result
}