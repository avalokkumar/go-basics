package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {

	names := []string{"Clayman", "Fury", "Krazzy", "Unicorn"}
	var wg sync.WaitGroup
	wg.Add(len(names))
	for _, name := range names{
		//Creates goroutines
		//Executes each function call on a new goroutine
		go printName(name, &wg)
	}
	wg.Wait()
}

func printName(name string, wg *sync.WaitGroup) {
	result := 0.0
	//time consuming task to keep cpu busy
	for i:=0; i<100000000; i++ {
		result +=math.Pi * math.Sin(float64(len(name)))
	}
	fmt.Println("Name: ", name)
	wg.Done()
}
