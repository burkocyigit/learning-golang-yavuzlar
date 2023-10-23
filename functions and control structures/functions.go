package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe() // >> Hello World!

	printMyName("Burak") // >> Hello Burak!

	printValue := "Beytullah"
	printMyName(printValue) // >> Hello Beytullah!

	result := intDivision(99, 11)
	fmt.Println(result) // >> 9

	result, remainder := intDivisionWithRemainder(95, 11)
	fmt.Printf("The result of the integer division is %v with remainder %v\n", result, remainder) // >> The result of the integer division is 8 with remainder 7

	// resultZero := intDivision(99, 0)
	// fmt.Println(resultZero) // >> panic: runtime error: integer divide by zero

	resultWithError, remainderError, err := intDivisionWithErrorHandling(54, 7)
	// Error handling and if-else if-else
	if err!=nil {
		fmt.Printf(err.Error())
	} else if remainderError == 0{
		fmt.Printf("The result of the integer division is %v", resultWithError)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v\n", resultWithError, remainderError)
	}

	// Swith case with no parameter
	switch {
		case err!=nil:
			fmt.Printf(err.Error())
		case remainderError == 0:
			fmt.Printf("The result of the integer division is %v", resultWithError)
		default:
			fmt.Printf("The result of the integer division is %v with remainder %v\n", resultWithError, remainderError)
	}
	
	// Switch case with parameter
	switch remainderError {
	case 0:
		fmt.Printf("The division was exact!")
	case 1,2:
		fmt.Printf("The division was close!")
	default:
		fmt.Printf("The division was not close.")
	}
}


func printMe() { // <-- curly braces needs to be this way
	fmt.Println("Hello World!")
}

/*
func printMe(str string) { // GO doesn't have function overloading!

} */

func printMyName(str string){
	fmt.Println("Hello " + str + "!")
}

func intDivision(numerator int, denominator int) int {
	var result int = numerator / denominator
	return result
}

func intDivisionWithRemainder(numerator int, denominator int) (int, int) {
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder
}

func intDivisionWithErrorHandling(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, nil
}

