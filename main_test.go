package main

import (
	"testing"
)

// TestStack tests the generic Stack implementation
func TestStack(t *testing.T) {
	// Test integer stack
	t.Run("Integer Stack", func(t *testing.T) {
		stack := &Stack[int]{}

		// Test empty stack
		_, ok := stack.Pop()
		if ok {
			t.Error("Expected false when popping from empty stack")
		}

		// Test push and pop
		stack.Push(1)
		stack.Push(2)

		val, ok := stack.Pop()
		if !ok || val != 2 {
			t.Errorf("Expected 2, got %d", val)
		}

		val, ok = stack.Pop()
		if !ok || val != 1 {
			t.Errorf("Expected 1, got %d", val)
		}
	})

	// Test string stack
	t.Run("String Stack", func(t *testing.T) {
		stack := &Stack[string]{}
		stack.Push("hello")
		stack.Push("world")

		val, ok := stack.Pop()
		if !ok || val != "world" {
			t.Errorf("Expected 'world', got %s", val)
		}
	})
}

// TestMap tests the generic Map function
func TestMap(t *testing.T) {
	t.Run("Integer to Integer Map", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		doubled := Map(numbers, func(x int) int { return x * 2 })

		expected := []int{2, 4, 6}
		if len(doubled) != len(expected) {
			t.Errorf("Expected length %d, got %d", len(expected), len(doubled))
		}

		for i, v := range doubled {
			if v != expected[i] {
				t.Errorf("At index %d: expected %d, got %d", i, expected[i], v)
			}
		}
	})

	t.Run("Integer to String Map", func(t *testing.T) {
		numbers := []int{1, 2}
		strings := Map(numbers, func(x int) string { return string(rune('A' - 1 + x)) })

		expected := []string{"A", "B"}
		if len(strings) != len(expected) {
			t.Errorf("Expected length %d, got %d", len(expected), len(strings))
		}

		for i, v := range strings {
			if v != expected[i] {
				t.Errorf("At index %d: expected %s, got %s", i, expected[i], v)
			}
		}
	})
}

// TestFilter tests the generic Filter function
func TestFilter(t *testing.T) {
	t.Run("Filter Even Numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		even := Filter(numbers, func(x int) bool { return x%2 == 0 })

		expected := []int{2, 4, 6}
		if len(even) != len(expected) {
			t.Errorf("Expected length %d, got %d", len(expected), len(even))
		}

		for i, v := range even {
			if v != expected[i] {
				t.Errorf("At index %d: expected %d, got %d", i, expected[i], v)
			}
		}
	})

	t.Run("Filter Strings by Length", func(t *testing.T) {
		words := []string{"a", "ab", "abc", "abcd"}
		longWords := Filter(words, func(s string) bool { return len(s) > 2 })

		expected := []string{"abc", "abcd"}
		if len(longWords) != len(expected) {
			t.Errorf("Expected length %d, got %d", len(expected), len(longWords))
		}

		for i, v := range longWords {
			if v != expected[i] {
				t.Errorf("At index %d: expected %s, got %s", i, expected[i], v)
			}
		}
	})
}