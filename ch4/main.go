package main

import (
	"fmt"
)

type person struct{
	first string
	second string
}

type secretAgent struct{
	person
	ltk bool
}

func main(){

	p1 := secretAgent{
		person: person{
			first : "kim",
			second : "seonjin",
		},
		ltk : true, 
	} 

	p2 := person{
		first : "miss",
		second : "jin",
	}

	p3 := struct{
		first string
		second string
	}{
		first : "kim",
		second : "jin",
	}


	test2()

	fmt.Println(p1, p2.first, p3)
}


func test1(){
	type person struct {
		first string
		last string
		favoriteIceCream []string
	}

	p := person{
		first : "kim",
		last : "jijn",
		favoriteIceCream : []string {"merona"},
	}

	fmt.Println(p)

}

func test2(){

	type person struct {
		first string
		last string
		favoriteIceCream []string
	}

	p := person{
		first : "kim",
		last : "jijn",
		favoriteIceCream : []string {"merona"},
	}

	pmap := map[string]person{
		p.last : p,
	}

	for k, v := range(pmap){
		fmt.Println(k, v)
	}
	
}