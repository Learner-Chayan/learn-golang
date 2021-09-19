package learnMethods


import (
	"fmt"
	"math"
)


type Employee struct{
	name string
	salary int 
	currency string
}
type Rectangle struct {
	length int
	width int 
}
type Circle struct{
	radius float64
}

// type alias
type myInt int


func Methods(){
	fmt.Println("...........................Methods.......................")


	//Sample Method .................................................
	emp1 := Employee {
	name: "Adam smith",
	salary: 343434,
	currency:"$",
	}
	emp1.displaySalary()

	// pass value and pointer on method .............................
	//value -
	fmt.Printf("\nEmployee name before change: %s", emp1.name)
    emp1.changeName("Michael Andrew")
    fmt.Printf("\nEmployee name after change: %s", emp1.name)
	//pointer -
    fmt.Printf("\n\nEmployee salary before change: %d", emp1.salary)
    (&emp1).changeSalary(500000)
    fmt.Printf("\nEmployee salary after change: %d", emp1.salary)

	//function - same example .......................................
	displaySalaryFunc(emp1)

	//same name methods
	r := Rectangle{
		length:10,
		width:5,
	}
	fmt.Printf("Area of rectangle %d\n",r.Area())
	c := Circle{
		radius:12,
	}
	fmt.Printf("Area of circle %f\n",c.Area())

	/* *** Differnt & methods , function , pointer
	note : When function has a value argument , it will accept only a value argument
	        but
	      When a method has a value receiver , it will accept both pointer and value receivers
	code - example : */

	calculateArea(r) // function call without pointer
	r.calculateArea() // method call without pointer

	p :=&r //make pointer
	//calculateArea(p) // error , because value argument doesn't accpet pointer in func
	p.calculateArea() // method call with pointer

	/*
		Similar to value arguments, functions with pointer arguments 
		will accept only pointers whereas 
		methods with pointer receivers will accept both pointer and value receiver.
	*/
	p2 := &r
	calculateAreaWithPointer(p2) // function call with pointer
	p2.calculateAreaWithPointer() // method call with pointer

	// calculateAreaWithPointer(r) // error, the function accept only pointer argument
	r.calculateAreaWithPointer() //method works again


	/*Methods with non-struct receivers*/
	//To define a method on a type, the definition of the receiver type and the definition 
	//of the method should be present in the same package.
	num1 := myInt(5)
    num2 := myInt(10)
    sum := num1.add(num2)
    fmt.Println("\nSum is", sum)

}

//methods
func (e Employee) displaySalary(){
	fmt.Printf("\nFrom method : Salary of %s is %s%d",e.name,e.currency,e.salary)
}

//functions
func displaySalaryFunc(e Employee){
	fmt.Printf("\n From function : Salary of %s is %s%d",e.name,e.currency,e.salary)
}

// same name methods
func (r Rectangle) Area() int{
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}


// method with value receiver and pointer receiver
// value receiver
func (e Employee) changeName(newName string){
	e.name = newName
}
//pointer receiver
func ( e *Employee) changeSalary(newSalary int){
	e.salary = newSalary
}

/* *** Differnt & methods , function , pointer
note : When function has a value argument , it will accept only a value argument
        but
      When a method has a value receiver , it will accept both pointer and value receivers
 code - example : 
*/
func calculateArea( r Rectangle) { // function
	fmt.Printf("\n Area Function result: %d\n", (r.length * r.width))
}

func (rec Rectangle) calculateArea(){ // method
	fmt.Printf("Area Method result: %d\n", (rec.length * rec.width))
}

/*
	Similar to value arguments, functions with pointer arguments 
	will accept only pointers whereas methods with pointer 
	receivers will accept both pointer and value receiver.
*/

func calculateAreaWithPointer( r *Rectangle) { // function
	fmt.Printf("\n Area Function result: %d\n", (r.length * r.width))
}

func (rec *Rectangle) calculateAreaWithPointer(){ // method
	fmt.Printf("Area Method result: %d\n", (rec.length * rec.width))
}

// method  without struct 

func (a myInt) add(b myInt) myInt {  
    return a + b
}