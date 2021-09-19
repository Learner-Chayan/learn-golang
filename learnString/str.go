package learnString

import "fmt"

func LearnStr(){
	fmt.Println("...................string....................")
	name := "Hello World "
    fmt.Printf("String: %s\n", name)
    printChars(name)
    fmt.Printf("\n")
    printBytes(name)

	//Rune - represent a unicode code point in go
	printAsRune("Señor") // will print same otherwise unicode problem in normal

	// byte to string
	byteToString()

	// string length
	fmt.Println("String length - ",len(name))

	// String compare
	str1 := "Hello Bangladesh"
	str2 := "Hello Bangladesh"
	if str1 == str2 {
		fmt.Printf("String -%s- and -%s- are equal",str1,str2)
	}else{
		fmt.Printf("String -%s- and -%s- are not not not equal",str1,str2)
	}

	// String concatenation
	result := str1 + " " + str2
	fmt.Println("\n Concated string is = ",result)

	// String is immutable
	// We can mute it using rune
	//str1[0] = "R" //error

	newStr := []rune(str1)
	newStr[0] = 'R'
	fmt.Printf("\n Muted string is =  %s",string(newStr))
}


func printBytes(s string) {  
    fmt.Printf("Bytes: ")
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
}

func printChars(s string) {  
    fmt.Printf("Characters: ")
    for i := 0; i < len(s); i++ {
        fmt.Printf("%c ", s[i])
    }
}

func printAsRune(s string){
	fmt.Printf("Characters in Rune \n: ")
	runes :=[]rune(s)
	for i:=0; i<len(runes);i++{
		fmt.Printf("%c",runes[i])
	}
}

func byteToString(){
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
    str := string(byteSlice)
    fmt.Println(str)
}