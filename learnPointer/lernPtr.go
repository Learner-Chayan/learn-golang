package learnPointer

import "fmt"


func LearnPTR(){
	fmt.Println("\n.................Pointer..............")

	// pointer basic
	b := 333
	var a *int = &b
	fmt.Printf("Type of a is %T \n ",a)
	fmt.Println("address of b is \n ",a)
	fmt.Println("Value of b is = ",*a)
	*a++
	fmt.Println("new value of b is = ", *a)


	passPointerToFunc(a)
	fmt.Println("after visit pointer into func -  value of b is = ", b)

	pointerUsingNEW()
}

func pointerUsingNEW(){
	size := new(int)
    fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
    *size = 85
    fmt.Println("New size value is", *size)
}

func passPointerToFunc(val *int){
	*val = 55
}


//NOTE
//: Do not pass a pointer to an array as an argument to a function. Use slice instead
//: Go does not support pointer arithmetic

