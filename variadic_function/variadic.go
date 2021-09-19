package variadic_function

import (  
    "fmt"
)

func find(num int, nums ...int) {  
    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in \n", nums)
    }
    fmt.Printf("\n")
}
func Variadic() {  
	fmt.Println("........................Variadic function......................")
    find(89, 89, 90, 95)
    find(45, 56, 67, 45, 90, 109)
    find(78, 38, 56, 98)
    find(87)


	//variadic function and slice
	welcome := []string{"hello", "world"}
    change(welcome...)
    fmt.Println(welcome)

}

func change(s ...string) {  
    s[0] = "Go"
    s = append(s, "playground")
    fmt.Println(s)
}



// we can can't pass slice to a variadic function as last parameter 
//because - variadic function convert it to slice , so slice + slice going wrong there 
