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
	testIf()
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

		test()
	*/
	prob9()
}

// Go는 while 문이 없다. Only for!!!

func test() {
	/*
			for i := 0; i <= 10; i++ {
				for j := 0; j <= 10; j++{

				}
				fmt.Println("test", i)
			}


		a := 3
		b := 10

		for a < b {
			fmt.Println(a)
			a++
		}

		for {
			if a > 10 {
				break
			}
			a++
			fmt.Println("test")
		}
	*/

	x := 83 / 40
	y := 83 % 40
	fmt.Println(x, y)

	for i := 0; i <= 100; i++ {
		if i%2 != 0 {
			continue
		}

		fmt.Println(i)
	}
}

func testIf() {
	if true {
		fmt.Println("true")
	}

	if !false {
		fmt.Println("false")
	}

	if x := 2; x == 2 {
		fmt.Println("True")
	} else if x == 1 {
		fmt.Println("Else if")
	} else {
		fmt.Println("Else")
	}

	switch "Bond" {
	case "Jin", "TEst":
		fmt.Println("false")
		fallthrough
	case "false":
		fmt.Println("test2")
	case "Bond":
		fmt.Println("test")
	default:
		fmt.Println("test default ")
	}
}

// ============ 실습 ========

func prob1() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
}

func prob2() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}

func prob3() {
	bd := 1996

	for bd <= 2022 {
		fmt.Println(bd)
		bd++
	}
}

func prob4() {
	bd := 1996

	for {
		if bd > 2022 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}

func prob5() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("%v\n", i%4)
	}
}

func prob6() {
	x := "James Bond"

	if x == "James Bond" {
		fmt.Println("True")
	}
}

func prob7() {
	x := "James Bond2"

	if x == "James Bond" {
		fmt.Println("True")
	} else if x == "Jmes Bond2" {
		fmt.Println("Wrong")
	} else {
		fmt.Println("No")
	}
}

func prob8() {
	switch {
	case true:
		fmt.Println("True")
	case false:
		fmt.Println("False")
	}
}

func prob9() {
	favSport := "suffing"

	switch favSport {
	case "swimming":
		fmt.Println("go to river")
	case "soccer":
		fmt.Println("go to ground")
	case "suffing":
		fmt.Println("go to hawaii")
	}
}

func prob10() {
	favSport := "suffing"

	switch favSport {
	case "swimming":
		fmt.Println("go to river")
	case "soccer":
		fmt.Println("go to ground")
	case "suffing":
		fmt.Println("go to hawaii")
	}
}
