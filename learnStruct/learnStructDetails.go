package learnStruct

import "fmt"


//create struct
type Employee struct {
	firstName string
	lastName  string 
	age int
	salary int
}

//Exported structs - lets access it from exported_struct pack
type Spec struct {
	Maker string
	Price int 
	model string 
}

func LearnStructDetails(){
	fmt.Println("..............Struct...............")

	//struct with specifying field name
	emp1 := Employee{
		firstName: "Chayan",
		lastName : "Roy",
		age 	 : 25,
		salary	 : 5000,
	}

	//creating struct without specifying names
	emp2 := Employee{"Roy","Chayan",29,800}

	fmt.Println("Employee 1",emp1)
	fmt.Println("Employee 2",emp2)


	// anonymous structs
	emp3 := struct {
		firstName string
		lastName string
		age int 
		salary int 
	}{
		firstName: "Andreah",
		lastName: "Nikola",
		age : 31,
		salary:5000,
	}
	fmt.Println("Empoyee 3 from anonymous struct", emp3)
	fmt.Println("Empoyee 3 first name", emp3.firstName)

	///Pointers to a struct
	pointerANDstruct()

	//snonymous struct fields
	anonymousField()

	// Nested structs
	nestedStruct()
}



func pointerANDstruct(){
	emp4 := &Employee{
		firstName : "Sam",
		lastName  : "Anderson",
		age       : 55,
		salary 	  : 6000,
	}

	fmt.Println("First name from struct pointer",(*emp4).firstName)
	fmt.Println("We can see first name from struct pointer without * - ",emp4.firstName)
}

//.....................................................
//struct
	type Person struct {
		string
		int
	}
//func
func anonymousField(){
	p1 := Person{
		string : "CR",
		int : 50,
	}

	fmt.Println("From anonymouse field - ",p1.string,p1.int)
}
//.........................................................

//struct
type Address struct {
	city string
	state string 
}

type People struct {
	name string 
	age int
	address Address
}

func nestedStruct(){
	p := People{
		name :"Chayan Roy",
		age : 23,
		address:Address{
			city:"Dhaka",
			state:"Khilkhet",
		},
	}

	fmt.Println("Nested struct - ",p)
	//fmt.Println("Promoted fields",p.city)
}




//NOTE : 
// If a struct type starts with a capital letter,
// then it is an exported type and it can be accessed from 
// other packages. Similarly, if the fields of a struct start
//  with caps, they can be accessed from other packages.