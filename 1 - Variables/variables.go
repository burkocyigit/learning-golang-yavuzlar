package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 32767
    intNum = intNum + 1
    fmt.Println(intNum) // >> 32768

    var floatNum float64 = 12345678.9
    fmt.Println(floatNum) // >> 1.23456789e+07

	/* Can't do that.
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + intNum32
	*/ 

	// You can do this.
	var floatNum32 float32 =  10.1
	var intNum32 int32 = 2
	var result float32 = float32(intNum32) + floatNum32
	fmt.Println(result) // >> 12.1

	// int division result is int and rounded down.
	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2) // >> 1

	// Strings
	var myString string = "Hello"
	var myString1 string = `World!
	Goodbye `
	var myString3 string = myString + myString1 + " concat!"
	fmt.Println(myString, myString1, myString3)
	/*
	>> Hello World!
			Goodbye  HelloWorld!
			Goodbye  concat!
	*/

	// length of a string
	fmt.Println(len(myString)) // >> 5
	fmt.Println(len("öç")) // >> 4
	//import "unicode/utf8"
	fmt.Println(utf8.RuneCountInString("öç")) // >> 2

	// Runes
	var myRune rune = 'a'
	fmt.Println(myRune) // >> 97

	// Booleans
	var myBoolean bool = false
	fmt.Println(myBoolean, !myBoolean) // >> false true

	// Default values
	var intNum3 int
	var myString2 string // ''
	var myBool bool
	fmt.Println(intNum3, myString2, myBool) // >> 0  false

	// omit the type and var keyword
	var myVar = "string"
	myVariable := "string!"
	fmt.Println(myVar, myVariable) // >> string string!

	// multiple variable declerations :O
	var1, var2 := 1, 2
	fmt.Println(var1, var2) // >> 1 2

	// Constants
	const myConst string = "Constant Value"
	// You can't change its value anymore
	// myConst = "something else"
}