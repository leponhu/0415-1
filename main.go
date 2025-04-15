package main

import (
	"fmt"
)

// Stack is a generic stack implementation that can store any type
type Stack[T any] struct {
	items []T
}

// Push adds an item to the top of the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// Map applies a function to each element in a slice and returns a new slice
func Map[T, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

// Filter returns a new slice containing only the elements that satisfy the predicate
func Filter[T any](slice []T, pred func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if pred(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	// Example 1: Generic Stack
	intStack := &Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	if val, ok := intStack.Pop(); ok {
		fmt.Printf("Popped from int stack: %d\n", val)
	}

	strStack := &Stack[string]{}
	strStack.Push("hello")
	strStack.Push("world")

	if val, ok := strStack.Pop(); ok {
		fmt.Printf("Popped from string stack: %s\n", val)
	}

	// Example 2: Generic Map function
	numbers := []int{1, 2, 3, 4, 5}
	doubled := Map(numbers, func(x int) int { return x * 2 })
	fmt.Printf("Original numbers: %v\n", numbers)
	fmt.Printf("Doubled numbers: %v\n", doubled)

	// Convert numbers to strings
	numberStrings := Map(numbers, func(x int) string { return fmt.Sprintf("num-%d", x) })
	fmt.Printf("Number strings: %v\n", numberStrings)

	// Example 3: Generic Filter function
	even := Filter(numbers, func(x int) bool { return x%2 == 0 })
	fmt.Printf("Even numbers: %v\n", even)
}