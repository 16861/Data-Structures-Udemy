package main

import (
	"testing"

	stack_array "github.com/16861/Data-Structures-Udemy/StackArray"
)

func TestPopPush(t *testing.T) {
	sa := stack_array.New(10)
	sa.Push(10)
	sa.Push(20)
	sa.Push(30)

	x := sa.Pop()
	if x != 30 {
		t.Errorf("first pop exected %d but got %d", 30, x)
	}

	x = sa.Pop()
	if x != 20 {
		t.Errorf("first pop exected %d but got %d", 20, x)
	}

	x = sa.Pop()
	if x != 10 {
		t.Errorf("first pop exected %d but got %d", 10, x)
	}
}

func TestIsEmpty(t *testing.T) {
	sa := stack_array.New(10)

	if !sa.IsEmpty() {
		t.Errorf("array must be empty afted it's created")
	}

	if sa.Pop() != -1 {
		t.Errorf("after popping from empty array get unexpected number")
	}

	if !sa.IsEmpty() {
		t.Errorf("array must be empty afted it's created and invoking pop()")
	}

	sa.Push(10)

	if sa.IsEmpty() {
		t.Errorf("array is not empty after push!")
	}

	sa.Pop()

	if !sa.IsEmpty() {
		t.Errorf("arrey must be empty after popping last element!")
	}
}
