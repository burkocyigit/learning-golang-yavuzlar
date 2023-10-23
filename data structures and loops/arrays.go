package main

import "fmt"

/*
 1 - Fixed Length
 2 - Same Type
 3 - Indexable
 4 - Contiguous in Memory
*/
func main() {
	var intArr [3]int32

	intArr[1] = 123

	fmt.Println(intArr[0]) // >> 0

	//slice
	fmt.Println(intArr[1:3]) // >> [123 0]

	// int32 is 4 bytes so allocated memory is 4B * 3 = 12 bytes
	fmt.Println(&intArr[0]) // 0xc000096080
	fmt.Println(&intArr[1]) // 0xc000096084
	fmt.Println(&intArr[2]) // 0xc000096088

	// initialize the array
	var intArray [3]int16 = [3]int16{1,2,3}
	// or
	sIntArray := [3]int32{666, 777, 888}
	// or
	sIntArray2 := [...]int32{111, 222, 333, 444}

	fmt.Println(intArray) // >> [1 2 3]
	fmt.Println(sIntArray) // >> [666 777 888]
	fmt.Println(sIntArray2) // >> [111 222 333 444]

	// Slices
	var intSlice []int32 = []int32{4,5,6}
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice)) // >> The length is 3 with capacity 3

	// Append to slice
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice)) // >> The length is 4 with capacity 6

	// Append to slice from another slice
	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice) // >> [4 5 6 7 8 9]

	// initialize with make(type, size, capacity(optional))
	var intSlice3 []int32 = make([]int32, 3, 10)
	fmt.Println(intSlice3) // >> [0 0 0]
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice3), cap(intSlice3)) // >> The length is 3 with capacity 10

}