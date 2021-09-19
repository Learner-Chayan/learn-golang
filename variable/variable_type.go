package variable

import (
	"fmt"
	"math"
)

func VariableType() {
	//type
	var mobile int

	//auto type detect
	var age = 20

	// multiple variable decleration
	var n, n1, n2 = 12, 30, 33
	var n4, n5 int

	// multiple variable decleration with multiple types
	var (
		x int    = 10
		y string = "Chayan"
	)
	//Short hand decleration
	x1 := 10
	//Multiple short hand decleration
	x3, x4 := 10, "learn"
	/////name, age := "naveen" //error
	// at least one of the variables on the left side of := is newly declared
	a, b := 20.20, 30.30 // declare variables a and b
	b, c := 40, 50       // b is already declared but c is new
	d := math.Min(a, b)

	fmt.Println("Hello World")
	fmt.Println("Your age is - ", age)
	fmt.Println("Mobile is not given-", mobile)
	fmt.Println("Multiple values are - ", n, n1, n2)
	fmt.Println("undecared variables - ", n4, n5)
	fmt.Println("multiple - ", x, y)
	fmt.Println("shorhand values - ", x1)
	fmt.Println("shorhand multiple - ", x3, x4)
	fmt.Println(a, b, c, d)
}
