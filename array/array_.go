package array

import "fmt"

func LearnArray() {
	a := [3]int{12, 78, 50} // short hand declaration to create array
	fmt.Println(a)          //[12,78,50]

	b := [3]int{12}
	fmt.Println(b) ////[12,0,0]

	c := [...]int{12, 78, 50} // ... makes the compiler determine the length
	fmt.Println(c)

	//var d [2]int
	//note
	//d = a //not possible because size is different

	e := [...]string{"USA", "China", "India", "Germany", "France"}
	f := e // a copy of e is assigned to f
	f[0] = "Singapore"
	fmt.Println("e - the main array is ", e)
	fmt.Println("b - new array created using old array and modified it is ", f)

	//Similarly when arrays are passed to functions as parameters, they are passed by value and the original array is unchanged.
	changeLocal(a)
	fmt.Println("after passing to function ", a)

	//Length of an array
	fmt.Println("Lenth of array a is =  ", len(a))

	//iterate array values
	showArrValue(e)

	//Multidimensional arrays
	learnMultidimensionalArray()

	//slice  // same as array but little different
	sliceFunc()

}

func changeLocal(a [3]int) {
	a[0] = 55
	fmt.Println("inside function where changed arr", a)
}

func showArrValue(arr [5]string) {

	//using loop
	fmt.Println("Array values using loop")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	//using range
	fmt.Println("Array values using range")
	for index, value := range arr {
		fmt.Printf("index = %d , value = %s \n", index, value)
	}

	//using range and ignoring index
	fmt.Println("Array values using range and ignoring index")
	for _, value := range arr {
		fmt.Printf(" value = %s \n", value)
	}
}


func learnMultidimensionalArray(){
	fmt.Println()
	a := [3][2]string{
        {"lion", "tiger"},
        {"cat", "dog"},
        {"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
    }

	for _, v1 := range a {
        for _, v2 := range v1 {
            fmt.Printf("%s ", v2)
        }
        fmt.Printf("\n")
    }

	var b [3][2]string
    b[0][0] = "apple"
    b[0][1] = "samsung"
    b[1][0] = "microsoft"
    b[1][1] = "google"
    b[2][0] = "AT&T"
    b[2][1] = "T-Mobile"
    fmt.Printf("\n")
	fmt.Println(b);
}


func sliceFunc(){
	a := [5]int{76, 77, 78, 79, 80} //array
    var b []int = a[1:4] //creates a slice from a[1] to a[3]
    fmt.Println(b) 

	// after creating slice ... just reference is passed / that means slice change will effect on main array
    fmt.Println("array before",a)
    for i := range b {
        b[i]++
    }
    fmt.Println("array after slice modified",a) 


	//slice effect in main array
	numa := [3]int{78, 79 ,80}
    nums1 := numa[:] //creates a slice which contains all elements of the array
    nums2 := numa[:]
    fmt.Println("array before change 1",numa)
    nums1[0] = 100
    fmt.Println("array after modification to slice nums1", numa)
    nums2[1] = 101
    fmt.Println("array after modification to slice nums2", numa)

	//appending in a slice array
	cars := []string{"Ferrari", "Honda", "Ford"}
    fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
    cars = append(cars, "Toyota")
    fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6

	//Passing a slice to a function
	nos := []int{8, 7, 6}
    fmt.Println("slice before function call", nos)
    subtactOne(nos)                               //function modifies the slice
    fmt.Println("slice after function call", nos)


}

func subtactOne(numbers []int) {  
	for i := range numbers {
		numbers[i] -= 2
	}

}

