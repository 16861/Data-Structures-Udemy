package main

import (
	"fmt"

	stack_array "./StackArray"
)

func main() {
	fmt.Println("Data structures udemy")

	sa := stack_array.New(10)

	sa.Push(10)

	fmt.Println(sa.IsEmpty())
	fmt.Println(sa.Pop())
	fmt.Println(sa.IsEmpty())

}
