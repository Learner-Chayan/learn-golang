package exported_struct 

import (
	"fmt"
	"learngo/learnStruct"
	)


func ExportStruct(){
	spec := learnStruct.Spec {
		Maker: "Apple",
		Price: 50000,
	}

	fmt.Println("Maker",spec.Maker)
	fmt.Println("Price",spec.Price)
}