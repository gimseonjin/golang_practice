package main

import "fmt"

func main() {
	fmt.Println("This is the first exercise")
	exercise1()
	fmt.Println("This is the seconde exercise")
	exercise2()
	fmt.Println("This is the third exercise")
	exercise3()
	fmt.Println("This is the fourth exercise")
	exercise4()
	fmt.Println("This is the fifth exercise")
	exercise5()
}

func exercise1() {
	// Using the short declaration operator,
	// ASSIGN theses VALUES to VARIABLES with
	// the IDENTIFIERES "x" and "y" and "z"
	x := 42
	y := "James Bond"
	z := true

	// Now print hte values stored in those variables using

	// a. single print satement
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// b. bultiple print statements
	fmt.Printf("%v, %v, %v\n", x, y, z)
}

/*
Use var to DECLARE three VARIABLES. The variables should have package level
scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the
variables and make sure the variables are of the following TYPE (meaning they can
store VALUES of that TYPE).
*/
var x int
var y string
var z bool

func exercise2() {
	// print out the values for each identifier
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// The compiler assigned values to the variables. What are these values called?
	// It's called zero values because they are not initialized, So they got a zero of each types
}

/*
At the package level scope, assign the following values to the three variables
*/
var x1 int = 42
var y1 string = "James Bond"
var z1 bool = true

func exercise3() {
	/*
		use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the
		returned value of TYPE string using the short declaration operator to a
		VARIABLE with the IDENTIFIER “s”
	*/
	s := fmt.Sprintf("%v, %v, %v", x1, y1, z1)

	// print out the value stored by variable “s”
	fmt.Println(s)
}

// Create your own type. Have the underlying type be an int.
type _underlyingType int

// create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
var x2 _underlyingType

func exercise4() {
	// print out the value of the variable “x”
	fmt.Println(x2)

	// print out the type of the variable “x”
	fmt.Printf("%T\n", x2)

	// assign 42 to the VARIABLE “x” using the “=” OPERATOR
	x2 = 42

	// print out the value of the variable “x”
	fmt.Printf("%v\n", x2)
}

// now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
func exercise5() {
	// then use the “=” operator to ASSIGN that value to “y”
	y2 := int(x2)

	// print out the value stored in “y”
	fmt.Println(y2)

	// print out the type of “y”
	fmt.Printf("%T\n", y2)
}
