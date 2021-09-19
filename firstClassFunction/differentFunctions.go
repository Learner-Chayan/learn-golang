package firstClassFunction

import "fmt"

//user defined function type
type add func(a int, b int) int

func DifferentFunctions(){
	fmt.Println("......FirstClassFunction.....Anonymous...Higher oder....Closure....")

	anonymousFunc()
	anoymousSelfCall()
	//user defined function type
	userDefinedFunc()
	higerOrderFunction()
	//Returning functions from other functions
	returnFuncFromFunc()
	learnClosure()
	closureExample2()
}

func anonymousFunc(){
	a := func(){
		fmt.Println("\nHello i am anoynymous function")
	}

	a()
	fmt.Printf(" and my Type is = %T\n",a)
}

func anoymousSelfCall(){
	func(a int , b int){
		fmt.Printf("Summation from self call function is = %d\n",a+b)
	}(10,5)
}

func userDefinedFunc(){
	var operation add = func(a int, b int) int{
		return a + b
	}

	result := operation(3,4)
	fmt.Println("Result from user defined function = ",result)
}


//higher order function
func higerOrderFunction(){
	f := func(a int , b int) int {
		return a+b
	}
	simple(f)
}

func simple(result func(a,b int) int){
	fmt.Println("Higher order function result = ",result(60,4))
}

// returning function from a function
func returnFuncFromFunc(){
	s := returnFunc()
	fmt.Println("Return function from function , result - ",s(60,7))
}
func returnFunc() func(a,b int) int {
	f := func(a,b int) int {
		return a + b
	}
	return f
}

/// Closure
func learnClosure(){
	a :=5
	func(){
		fmt.Println("from outside func (closure) - a and see down=",a)
	}()


	// more complex or details
	x := appendStr()
	y := appendStr()
	fmt.Println(x("World"))
	fmt.Println(y("Everyone"))

	fmt.Println(x("All time"))
	fmt.Println(y("!"))
}

func appendStr() func(string) string {
	t := "Hello"
	fmt.Printf("........Now t value is = %s\n ",t)
	c := func(b string) string {
		t = t + " " + b
		//return t+" "+ b
		return t
	}
	return c
}


/// closure example 2
func closureExample2(){

	 seriesSum := clExample2()
	 seriesSum2 := clExample2()
	 seriesSum(5);
	 seriesSum2(20)
	 seriesSum(4);
	 seriesSum2(20)
	 seriesSum(3);
	 seriesSum2(20)
	 seriesSum(2);
	 seriesSum2(20)
	 seriesSum(1);
	 seriesSum2(20)
	 fmt.Println("The series sum using closure is = ",seriesSum(0))
	 fmt.Println("The series sum2 using closure is = ",seriesSum2(0))
}
func clExample2() func(n int) int {
	total := 0
	result := func(n int) int {
		total = total+n
		return total
	}
	return result
}