package main

import (
	"fmt"
	"time"
)

func main() {
	var n int = 1000000
	var slice1 []int = []int{}
	var slice2 []int = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(slice1, n))
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(slice2, n))
	//  >> Total time without preallocation: 14.5927ms
	//	   Total time with preallocation: 3.1458ms
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()

	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}