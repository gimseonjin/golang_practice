package main

import (
	"fmt"
)

func main(){
	a1 := agent{
		person : person{
			first : "kim",
			last : "seonjin",
		},
		ltk : false,
	}

	a1.speak()
}

// maybe defer using stack 
// [tset] -> [tset, tset2] -> [tset] -> []

// defer를 쓰는 케이스 -> 파일을 열고, 로직이 끝날 떄 까지 기다리면 컴퓨터 자원 낭비!! -> 따라서 파일 읽고 defer로 로직을 미룬 다음, 파일을 닫아버릴 수 있다!!

// default function form
// func (r receiver) identifier(parameters) (return(s)) {...}

func foo(){
	fmt.Println("with out param")
}

func bar(s string){
	fmt.Println(s)
}

func woo(s string) string{
	return fmt.Sprint("hello", s)
}

func returnTwoValue(first string, last string) (name string, result bool){
	new_name := fmt.Sprint(first, " ", last)
	new_result := true

	return new_name, new_result
}

func variableArgument(x ...int) int{

	sum := 0

	for _, v := range x{
		sum += v
	}

	return sum
}

// #################################################

type person struct{
	first string
	last string
}

type agent struct{
	person
	ltk bool
}

func (a agent) speak(){
	fmt.Println(a.first, a.last, a.ltk)
}
