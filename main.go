package main

import "fmt"

func main() {
	txt := hello()
	x, _ := fmt.Println(txt)
	fmt.Println(x)
	test3()
}

func hello() string {
	return "안녕!, 세상."
}

// 제어 유형은 3개
// sequence
// roop
// contitional

// 단축 선언자, var, const
func test1() {
	// 단축 선언자
	x := 42
	fmt.Println(x)

	var y = 43
	fmt.Println(y)

	// := vs var ==>
	/*
		:=(단축 선언자)의 경우, 함수 내부{}에서만 사용할 수 있다.
		반면 var의 경우, 함수 내부에서만이 아니라 패키지, go 파일 어느 곳에서도 다 선언할 수 있다.
		가능한 단축 선언자를 쓰자!!! 글로벌 변수를 쓰면 예상치 못한 에러, 또는 일을 만들 수 있다.
	*/
}

func test2() {
	/*
		Go에서 타입은 정말 중요하다!!!

		Go는 정적 타입 언어다.

		따라서 타입을 잘 정의해야 한다.

		타입 지정 방식은 아래와 같다.

		var a type(string, int ...)

		여기서 값을 선언할 때, 초기화 시킬 수 있다.

		이때 초시화 과정을 생략하면 모든 변수에는 해당 타입의 제로값이 선언된다.

		예를 들어 var a int 하면 a 에는 0이 배정된다.

		다른 예로 var a string 하면 a 에는 ""이 배정된다.
	*/

}

func test3() {
	// fmt 실습

	y := 42
	/*
		go printf 에서 사용되는 형식 프린팅 내용
		https://pkg.go.dev/fmt 참고

		백 슬러시 다음 문자열 내용
		\a   U+0007 alert or bell
		\b   U+0008 backspace
		\f   U+000C form feed
		\n   U+000A line feed or newline
		\r   U+000D carriage return
		\t   U+0009 horizontal tab
		\v   U+000B vertical tab
		\\   U+005C backslash
		\'   U+0027 single quote  (valid escape only within rune literals)
		\"   U+0022 double quote  (valid escape only within string literals)
	*/
	fmt.Println(y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)

	s := fmt.Sprintf("%x\n", y)
	fmt.Println(s)
}
