package main

import (
	"fmt"
	"learngo/array"
	"learngo/types"
	"learngo/variable"
	"learngo/variadic_function"
	"learngo/maps"
	"learngo/learnString"
	"learngo/learnPointer"
	"learngo/learnStruct"
	"learngo/exported_struct"
	"learngo/learnMethods"
	"learngo/learnInterface"
	"learngo/learnInterfaces2"
	"learngo/firstClassFunction"
	"learngo/learnDefer"
)

func main() {

	fmt.Println("Hello world")
	variable.VariableType()
	types.LearnTypes()
	array.LearnArray()
	variadic_function.Variadic()
	maps.LearnMap()
	learnString.LearnStr()
	learnPointer.LearnPTR()
	learnStruct.LearnStructDetails()
	exported_struct.ExportStruct()
	learnMethods.Methods()
	learnInterface.Interfaces()
	learnInterfaces2.Interfaces2()
	firstClassFunction.DifferentFunctions()
	learnDefer.DeferDetails()

}
