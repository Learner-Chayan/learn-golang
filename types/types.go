package types

import (
	"fmt"
	"unsafe"
)

func LearnTypes() {

	/*
		bool
		Numeric Types
			int8, int16, int32, int64, int
			uint8, uint16, uint32, uint64, uint
			float32, float64
			complex64, complex128
			byte
			rune
		string
	*/

	//bool
	a1 := true
	b1 := false
	fmt.Println("a:", a1, "b:", b1)
	c := a1 && b1
	fmt.Println("c:", c)
	d := a1 || b1
	fmt.Println("d:", d)

	//int
	var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))   //type and size of a
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) //type and size of b
	fmt.Println("'Learn types'")

	//floating point number
	a3, b3 := 5.67, 8.97
	fmt.Printf("type of a3 %T b3 %T\n", a3, b3)
	sum := a3 + b3
	diff := a3 - b3
	fmt.Println("sum", sum, "diff", diff)

	//complex type
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)

	//Type Conversion
	i := 55   //int
	j := 67.8 //float64
	//sum2 := i + j //int + float64 not allowed
	sum2 := i + int(j) //j is converted to int
	fmt.Println(sum2)
}
