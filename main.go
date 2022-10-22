package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		txt := hello()
		fmt.Println(txt)
	}
}

func hello() string {
	return "안녕!, 세상."
}

// 제어 유형은 3개
// sequence
// roop
// contitional
