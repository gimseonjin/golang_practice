package main

import (
	"fmt"
)

const (
	a      = 10
	b int8 = 15
)

const (
	target = 2022
	first  = target + iota
	second
	third
	forth
)

func main() {

	/*

		a := 7
		b := 42

		fmt.Println(a == b)

		x = 10 < 5

		fmt.Println(x)

		y := 12.3
		fmt.Printf("%T\n", y)

		var test uint
		fmt.Printf("%T\n", test)
		fmt.Println(runtime.GOOS)   // darwin
		fmt.Println(runtime.GOARCH) // amd64 -> int64 이다!!

		s := "hello world!"
		fmt.Println(s)

		bs := []byte(s)
		fmt.Println(bs)
		fmt.Printf("%T\n", bs)

		for i := 0; i < len(s); i++ {
			fmt.Printf("%#U ", s[i])
		}

		for i, v := range s {
			fmt.Printf("%d, %T\n", i, v)
		}

		fmt.Println(string(bs))
	*/

	x := 20150190

	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%#x\n", x)

	fmt.Println(a)
	fmt.Println(b)

	var k int8

	k = a
	fmt.Println(k)

	k = b
	fmt.Println(k)

	fmt.Printf("%d\t\t\t", x<<1)
	fmt.Printf("%b\t\t\t", x<<1)
	fmt.Printf("%#x\t\t\t", x<<1)

	a := `here is something
	as 
	a 
	raw string
	literal
	"you see"
	another thing`
	fmt.Println(a)

	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(forth)
}
