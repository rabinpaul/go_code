package main

import (
	"fmt"
)

func main () {
	var origString [3]string
	var origInt [3]int
	
	for i, _ := range origString {
		origString[i] = "RAB"
	}
	
	fmt.Println("Original Array for string:", origString)
	
	for i, _ := range origInt {
		origInt[i] = i
	}
	
	fmt.Println("Original Array for Int:", origInt)
	
	a := [5]string{"A", "B", "C", "D"}
	fmt.Println("Original Array:", a)
	
	b := a
	fmt.Println("Original Array:", b)
	
	b[2] = "E"
	
	fmt.Println("Modified Array:", b)
	fmt.Println("Modified Original Array:", a)
	
	c := &a
	
	a[1] = "RABIN"
	
	fmt.Println("Modified Array1:", *c)
	fmt.Println("Modified Original Array1:", a)
}