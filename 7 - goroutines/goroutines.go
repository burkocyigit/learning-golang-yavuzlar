package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup is simply a counter
var wg = sync.WaitGroup{}
// Mutual Exclusion
var m = sync.Mutex{}
// Read-Write Mutual Exclusion
var rwm = sync.RWMutex{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	
	wg.Wait()
	
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are %v", results)
}

func dbCall(i int) {
	// Simulate DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	save(dbData[i])
	log()
	/*m.Lock()
	results = append(results, dbData[i])
	m.Unlock()*/
	// Done decrements the WaitGroup counter by one.
	wg.Done()
}

func save(result string) {
	rwm.Lock()
	results = append(results, result)
	rwm.Unlock()
}

func log() {
	rwm.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	rwm.RUnlock()
}