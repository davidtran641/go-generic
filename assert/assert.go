package main

import (
	"log"
)

func main() {
	AssertEqual(1,1)
	AssertEqual("Apple", "Apple")
	AssertNotEqual(1,2)
}

func AssertEqual[T comparable](want, actual T) {
	if want != actual {
		log.Fatalf("FAILED: want %+v, actual %+v", want, actual)
	} else {
		log.Printf("PASSED: want %+v, actual %+v", want, actual)
	}
}

func AssertNotEqual[T comparable](want, actual T) {
	if want == actual {
		log.Fatalf("FAILED: NOT want %+v, actual %+v", want, actual)
	} else {
		log.Printf("PASSED: NOT want %+v, actual %+v", want, actual)
	}
}
