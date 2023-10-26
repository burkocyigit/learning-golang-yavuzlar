package main

import "fmt"

func main() {

	var myMap map[uint8]string = map[uint8]string{1:"Monday", 2:"Tuesday", 3:"Wednesday"}

	// Iterate over maps
	for key, value:= range myMap {
		fmt.Printf("%v: %v\n", key, value)
	}/* 
	>>  1: Monday
		2: Tuesday
		3: Wednesday
	*/

	var myArray [4]float32 = [4]float32{3.14, 2.82, 0.52, 0.099}

	// Iterate ovre array
	for index, value := range myArray {
		fmt.Printf("%v: %v\n", index, value)
	}/* 
	>>  0: 3.14
		1: 2.82
		2: 0.52
		3: 0.099
	*/

	// There is no while loop but there are various ways to implement
	var i int = 0
	for i<10 {
		fmt.Println(i)
		i = i + 1
	}
	// or
	var j int = 1
	for {
		if j >= 10 {
			break
		}
		fmt.Println(j)
		j = j * 3
	}
	// or
	for k:=10; k>=0; k-- {
		fmt.Println(k)
	}

}