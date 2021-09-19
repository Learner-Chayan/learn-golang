package maps

import "fmt"

func LearnMap(){

	fmt.Println("................map...................")

	//create map
	//var map1 map[string]int // zero value
	employeeSalary := make(map[string]int)
	employeeSalary["emp1"] = 1000;
	employeeSalary["emp2"] = 2000;
	employeeSalary["emp3"] = 2000;

	fmt.Println(employeeSalary)

	// initialize map during decleration
	m := map[string]int{
		"value1":1234,
		"value2":5689,
	}
	m["value3"] = 5895
	fmt.Println(m["value1"])


	//Checking if a key exists
	value,ok := m["value4"]
	if ok == true {
		fmt.Println("value4 value is = ",value)
		return
	}else {
		fmt.Println("Value4 not found in map")
	}

	// delete items from a map
	delete(employeeSalary,"emp2")

	// map iterating all
	for key,value := range employeeSalary {
		fmt.Printf("key = %s and value = %d \n",key,value)
	}


	/// map of struc
	mapANDstruc()

	/// map length
	fmt.Println("length is of the map ", len(employeeSalary))
}

type employee struct {
	salary int
	country string
}
func mapANDstruc(){
	emp1 := employee{
		salary:30000,
		country:"USA",
	}

	emp2 := employee{
		salary:300000,
		country:"Canada",
	}

	employeeInfo := map[string]employee{
		"steve":emp1,
		"Jamie":emp2,
	}

	for name,info := range employeeInfo{
		fmt.Printf("Employee: %s salary:$%d country: %s \n", name,info.salary, info.country)
	}

}