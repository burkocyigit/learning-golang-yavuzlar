package main

import "fmt"

// define a struct
type gasEngine struct {
	mpg uint8
	gallons uint8
	ownerInfo owner // Another struct can be used as a field
}

// 
type electricEngine struct {
	mpkwh uint8
	kwh uint8
}

// define another struct
type owner struct{
	name string
}

// Structs can have methods
func (e gasEngine) milesLeft() uint8{
	return e.gallons * e.mpg
}

//
func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface{
	milesLeft() uint8
}

// We define a interface (engine)
func canMakeIt (e engine, miles uint8) {
	if miles<=e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {

	// With zero values
	var myEngine gasEngine
	fmt.Println(myEngine.mpg, myEngine.gallons) // >> 0 0

	// With parameters
	var yourEngine gasEngine = gasEngine{mpg:25, gallons: 15}
	fmt.Println(yourEngine.mpg, yourEngine.gallons) // >> 25 15

	var theirEngine gasEngine = gasEngine{mpg: 25, gallons:15, ownerInfo: owner{"Alex"}}
	fmt.Println(theirEngine) // {25 15 {Alex}}

	// Struct function
	fmt.Printf("Total miles left in tank: %v\n", theirEngine.milesLeft()) // >> Total miles left in tank: 119
	
	// With implemented interface
	canMakeIt(yourEngine, 50) // >> You can make it there!
	}	