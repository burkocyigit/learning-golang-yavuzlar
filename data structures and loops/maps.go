package main

import "fmt"

func main() {
	// initialize map with make()
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap) // >> map[]

	// initialize map immediately
	var myMap2 = map[string]uint8{"Fenasi":22,"Kerim":20}
	fmt.Println(myMap2) // >> map[Fenasi:22 Kerim:20]

	// indexing
	var guy = myMap2["Fenasi"]
	fmt.Println(guy) // >> 22

	// what if index doesn't exist
	var noGuy = myMap2["Fatih"]
	fmt.Println(noGuy) // >> 0
	// It returns the default value of the data type of the map

	// Error handling with maps, ok returns true or false based on the index exists or not.
	var age, ok = myMap2["Portakal"]
	
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid name!")
	}
	
	// delete a key by reference
	delete(myMap2, "Fenasi")
	fmt.Println(myMap2) // >> map[Kerim:20]

}