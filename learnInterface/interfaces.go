package learnInterface

import "fmt"

//interface definition
type VowelsFinder interface {  
    FindVowels() []rune
}

type MyString string

// part -2 .......................
type SalaryCalculator interface {
	CalculateSalary() int
}
type Permanent struct {
	empId int 
	basicpay int 
	pf int
}
type Contract struct{
	empId int 
	basicpay int
}

type Freelancer struct {
	empId int 
	ratePerHour int 
	totlaPerHours int 
}

func Interfaces(){
	fmt.Println("............................Interface........................")
	name := MyString("Sam Anderson")
    var v VowelsFinder // interface VoewlsFinder type
    v = name // possible since MyString implements VowelsFinder
    fmt.Printf("\nVowels are %c", v.FindVowels()) // function call



	// part - 2 .................................
	pemp1 := Permanent{
        empId:    1,
        basicpay: 5000,
        pf:       20,
    }
    pemp2 := Permanent{
        empId:    2,
        basicpay: 6000,
        pf:       30,
    }
    cemp1 := Contract{
        empId:    3,
        basicpay: 3000,
    }
	freelancer1 := Freelancer {
		empId: 4,
		ratePerHour: 100,
		totlaPerHours: 10,
	}
	freelancer2 := Freelancer {
		empId: 5,
		ratePerHour: 200,
		totlaPerHours: 6,
	}

	employees := []SalaryCalculator{pemp1, pemp2, cemp1,freelancer1,freelancer2}
    totalExpense(employees)


}


//MyString implements VowelsFinder
func (ms MyString) FindVowels() []rune {  
    var vowels []rune
    for _, rune := range ms {
        if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
            vowels = append(vowels, rune)
        }
    }
    return vowels
}

// part -2 ...................................
func (p Permanent) CalculateSalary() int {  
    return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int {  
    return c.basicpay
}

func (f Freelancer) CalculateSalary() int {
	return f.totlaPerHours * f.ratePerHour
}
func totalExpense(s []SalaryCalculator) {  
    expense := 0
    for _, v := range s {
        expense = expense + v.CalculateSalary()
    }
    fmt.Printf("\n Total Expense Per Month $%d", expense)
}