package main

import (
	"fmt"
	"reflect"
)

func main() {
	test5()
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

func slice_slice() {
	x := []int{1, 2, 3, 4, 5}

	fmt.Println(x[0:3])

	x = append(x, 1, 2, 3, 4)
	fmt.Println(x)

	y := []int{7, 8, 9}
	x = append(x, y...)
	fmt.Println(x)

	// remove
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	// with make
	k := make([]int, 10, 100)
	fmt.Println(k)
}

func multi_slice() {
	jb := []string{"kim"}
	mp := []string{"seonjin"}

	total := [][]string{jb, mp}
	fmt.Println(total)
}

func test_map() {
	m := map[string]int{
		"seonjin": 27,
	}
	fmt.Println(m)

	m["jeni"] = 35

	v, ok := m["jeni"]
	fmt.Println(v, ok)

	delete(m, "jeni")

	if v, ok := m["seonjin"]; ok {
		fmt.Println(v)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func test1() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := make([]int, 5, 10)

	for i, v := range arr {
		fmt.Println(i, v)
	}

	for i, v := range slice {
		fmt.Println(i, v)
	}

	fmt.Println(reflect.TypeOf(arr))
	fmt.Println(len(arr))
	fmt.Println(cap(arr))

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func test2() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	// 1 ~ 5
	fmt.Println(x[:5])

	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 52, 53, 54)
	fmt.Println(x)

	y := []int{8, 912}
	x = append(x, y...)
	fmt.Println(x)

}

func test3() {

	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	x = append(x[:3], x[4:]...)

	fmt.Println(x)
}

func test4() {
	// 강원도ㆍ경기도ㆍ경상도ㆍ전라도ㆍ충청도ㆍ평안도ㆍ함경도ㆍ황해도
	name := make([]string, 0, 8)
	fmt.Println(len(name))
	fmt.Println(cap(name))

	name = append(name, "강원도", "경기도", "경상도", "전라도", "충청도", "평안도", "함경도", "황해도")
	fmt.Println(len(name))
	fmt.Println(cap(name))

	for x := 0; x < len(name); x++ {
		fmt.Println(x, name[x])
	}
}

func test5(){
	x1 := []string{"kim","kim","kim"}
	x2 := []string{"lee","lee","lee"}

	x3 := [][]string{x1,x2}

	for i, v := range(x3){
		for j, v2 := range(v){
			fmt.Println(i, j, v2)
		}
	}
}