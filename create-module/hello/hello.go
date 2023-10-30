package main

import (
	"fmt"

	"github.com/burkocyigit/learning-golang-yavuzlar/create-module/greetings"
)

func main() {
	hello := greetings.Hello("Burak")
	fmt.Println(hello)
}