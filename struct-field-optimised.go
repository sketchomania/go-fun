package main

import (
	"fmt"
	"unsafe"
)

// resource: https://golangprojectstructure.com/how-to-make-go-structs-more-efficient/

// City represents a city where a person can reside.
type City uint8

// Constants representing cities.
const (
	NewYork City = iota
	London
	Paris
	Mumbai
)

// Person1 uses 16 bytes
type Person1 struct {
	currentResidence City
	uniqueID         int64
	passportNumber   int16
}

// Person2 uses 16 bytes
type Person2 struct {
	uniqueID         int64
	passportNumber   int16
	currentResidence City
}

// Person3 uses 24 bytes
type Person3 struct {
	uniqueID         int64
	currentResidence City
	passportNumber   int16
}

// OptimisedStruct is a function that demonstrates the memory usage of a Person struct.
func OptimisedStruct() {
	m1 := Person1{
		currentResidence: London,
		uniqueID:         9248511308,
		passportNumber:   10564,
	}
	m2 := Person2{
		currentResidence: London,
		uniqueID:         9248511308,
		passportNumber:   10564,
	}
	m3 := Person3{
		currentResidence: London,
		uniqueID:         9248511308,
		passportNumber:   10564,
	}

	fmt.Printf(
		"My Person struct uses %d bytes.\n",
		unsafe.Sizeof(m1),
	)
	fmt.Printf(
		"My Person struct uses %d bytes.\n",
		unsafe.Sizeof(m2),
	)
	fmt.Printf(
		"My Person struct uses %d bytes.\n",
		unsafe.Sizeof(m3),
	)
}
