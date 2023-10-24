package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	var indexed = myString[5]
	fmt.Println(indexed) // >> 195

	fmt.Printf("%v, %T", indexed, indexed) // >> 195, uint8

	for i,v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("\nThe length of '%v' is %v", myString, len(myString)) // >> The length of 'résumé' is 8

	var myStringWithRunes = []rune("résumé")
	fmt.Printf("\nThe length of '%v' is %v", myStringWithRunes, len(myString)) // >> The length of '[114 233 115 117 109 233]' is 8

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune) // >> myRune = 97

	// String concatenate
	var strSlice = []string{"b", "e", "y", "t", "i"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("\n%v", catStr) // >> beyti
	// Import "strings" instead of doing this.

	// String Builder from strings
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var builtStr = strBuilder.String()
	fmt.Printf("\n%v", builtStr) // >> beyti
}