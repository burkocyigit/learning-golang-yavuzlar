package main

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

type rectangle struct {
	height int
	width int
}

type square struct {
	height int
}

type shape interface {
	calculateArea()
	calculatePerimeter()
}

func (r rectangle) calculateArea() int {
	return r.height * r.width
}

func (s square) calculateArea() int {
	return int(math.Pow(float64(s.height), 2))
}

func split(str string) string {
	var strBuilder strings.Builder

	for _, v := range str {
		strBuilder.WriteRune(v)
		strBuilder.WriteRune(' ')
	}

	return strBuilder.String()
}

func reverseString(str string) string {
	var strBuilder strings.Builder

	for i := range str {
		strBuilder.WriteString(string(str[utf8.RuneCountInString(str)-i-1]))
	}

	return strBuilder.String()
}


func main() {
	var rect rectangle = rectangle{5, 11}
	fmt.Printf("The area of the rectangle is: %v", rect.calculateArea())

	var square square = square{5}
	fmt.Printf("\nThe area of the square is: %v", square.calculateArea())

	fmt.Printf("\nsplit(\"Split me!\"): %v\n", split("Split me!"))

	fmt.Printf("\nreverseString(\"Burak\"): %v\n", reverseString("Burak"))
}

