package main

import (
	"fmt"
	"reflect"
)

func main() {
	slice()
}

func arr1() {
	// 배열을 쓸 때 주의해야 할 점!!!!
	// 배열의 길이 또한 타입의 일부이기 때문에 길이가 다르면 다른 타입니다.
	// x & y는 서로 다른 타입
	var x [5]int
	var y [6]int
	x[0] = 1

	type_x := reflect.TypeOf(x)
	type_y := reflect.TypeOf(y)

	fmt.Println(type_x)
	fmt.Println(type_y)
	// 따라서 둘의 타입을 비교하면 False!
	fmt.Println(type_x == type_y)
}

func slice() {
	x := []int{4, 5, 6, 7}
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}

	m := map[string]string{
		"1": "test",
	}

	for i, v := range m {
		fmt.Println(i, v)
	}
}
