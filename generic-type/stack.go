package main

import (
	"log"
)

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(e T) {
	s.data = append(s.data, e)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var empty T
		return empty, false
	}

	idx := len(s.data) - 1
	e := s.data[idx]
	s.data = s.data[:idx]
	return e, true
}

func main() {
	testIntStack()
	testStringStack()
}

func testIntStack() {
	stack := new(Stack[int])
	AssertEqual(true, stack.IsEmpty())

	stack.Push(1)
	AssertEqual(false, stack.IsEmpty())

	stack.Push(2)
	AssertEqual(false, stack.IsEmpty())

	val, ok := stack.Pop()
	AssertEqual(true, ok)
	AssertEqual(2, val)

	val, ok = stack.Pop()
	AssertEqual(true, ok)
	AssertEqual(1, val)

	val, ok = stack.Pop()
	AssertEqual(false, ok)
}

func testStringStack() {
	stack := new(Stack[string])
	AssertEqual(true, stack.IsEmpty())

	stack.Push("A")
	AssertEqual(false, stack.IsEmpty())

	stack.Push("B")
	AssertEqual(false, stack.IsEmpty())

	val, ok := stack.Pop()
	AssertEqual(true, ok)
	AssertEqual("B", val)

	val, ok = stack.Pop()
	AssertEqual(true, ok)
	AssertEqual("A", val)

	val, ok = stack.Pop()
	AssertEqual(false, ok)
}


func AssertEqual[T comparable](want, actual T) {
	if want != actual {
		log.Fatalf("FAILED: want %+v, actual %+v", want, actual)
	} else {
		log.Printf("PASSED")
	}
}

func AssertNotEqual[T comparable](want, actual T) {
	if want == actual {
		log.Fatalf("FAILED: NOT want %+v, actual %+v", want, actual)
	} else {
		log.Printf("PASSED")
	}
}
