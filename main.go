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

	fmt.Println("*********************")

	sau := stack_array.NewUnlim()
	fmt.Println(sau.IsEmpty())

	sau.Push(10)
	sau.Push(20)
	sau.Push(30)
	sau.Push(40)

	fmt.Println(sau.IsEmpty())
	fmt.Println(sau.Pop())
	fmt.Println(sau.IsEmpty())
	fmt.Println(sau.Len())
	fmt.Println(sau.Pop())
	fmt.Println(sau.Len())
}
