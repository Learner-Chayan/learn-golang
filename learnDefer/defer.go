package learnDefer

import "fmt"


func DeferDetails(){
  fmt.Println(".......................Defer.............................")

  deferIdea()
  deferMethod()
  evaluteArguments()
}

// basic , know defer...................................................
func deferIdea(){
	nums := []int{78, 109, 2, 563, 300}
	defer finished()    
    fmt.Println("Started finding largest")
    max := nums[0]
    for _, v := range nums {
        if v > max {
            max = v
        }
    }
    fmt.Println("Largest number in", nums, "is", max)
}

func finished() {  
    fmt.Println("Finished finding largest")
}

//Defer method ........................................................
type person struct {  
    firstName string
    lastName string
}

func (p person) fullName() {  
    fmt.Printf("%s %s \n",p.firstName,p.lastName)
}

func deferMethod() {  
	fmt.Println("............................defer method...................")
    p := person {
        firstName: "John",
        lastName: "Smith",
    }
    defer p.fullName()
    fmt.Printf("Welcome ")  
}

// Defer ... Arguments evaluation .........................................
// note : The arguments of a deferred function are evaluated when the defer 
///statement is executed and not when the actual function call is done.
//        
func printA(a int) {  
    fmt.Println("value of a in deferred function", a) // output : 5 
}
func evaluteArguments() {  
    a := 5
    defer printA(a) 
    a = 10
    fmt.Println("value of a before deferred function call", a) //output : 10
}


// ...............Stack of defers ............................................
