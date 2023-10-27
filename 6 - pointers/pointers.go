package main

import "fmt"

func main() {

	var p *int32 = new(int32)
	var i int32 = 3

	fmt.Printf("The value of p points to is: %v", *p) // >> The value of p points to is: 0
	fmt.Printf("\nThe value of i is: %v", i) // >> The value of i is: 3

	// Assigning a value to the pointer
	*p = 10

	// Assigning a memory address to the pointer via '&'
	p = &i
	fmt.Printf("\nThe value of p points to is: %v", *p) // >> The value of p points to is: 3
	// Changing the value of p would change the variable i itself.
	*p = 11
	fmt.Printf("\nThe value of p points to is: %v", *p) // >> The value of p points to is: 11
	fmt.Printf("\nThe value of i is: %v", i) // >> The value of i is: 11

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing1) // >> The memory location of the thing1 array is: 0xc0000a6090

	var result = square(thing1)
	fmt.Printf("\nThe result is: %v", result) // >> The result is: [1 4 9 16 25]
	fmt.Printf("\nThe value of thing1 is: %v", thing1) // >> The value of thing1 is: [1 2 3 4 5]

	// You see, the original array that we passed in to the function is not changed and also the thing1 and thing2 arrays have different memory locations.
	// The momery allocation is doubled with this. We can use pointers to reduce memory usage.

	var thing = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of the thing array is: %p", &thing) // >> The memory location of the thing1 array is: 0xc0000a6090

	var resultthing = squareWithPointer(&thing)
	fmt.Printf("\nThe result is: %v", resultthing) // >> The result is: [1 4 9 16 25]
	fmt.Printf("\nThe value of thing is: %v", thing) // >> The value of thing is: [1 4 9 16 25]

}

func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing2 array is: %p", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}

func squareWithPointer(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing2 array is: %p", &thing2) // >> The memory location of the thing2 array is: 0xc0000a60c0
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}