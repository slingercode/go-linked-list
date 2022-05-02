package main

import (
	"fmt"
	"strings"

	linkedlist "github.com/slingercode/go-linked-list/linked-list"
)

func main() {
	fmt.Println("Linked list on go")

	ll := linkedlist.LinkedList{}

	ll.Init("First Node")
	ll.AddFirst("Second Node")
	ll.AddFirst("Third Node")
	ll.AddLast("Fourth Node")

	fmt.Println("Size", ll.GetSize())

	list := "[" + strings.Join(ll.Get(), ", ") + "]"

	fmt.Println(list)

	ll.AddAt(0, "New First Value")
	ll.AddAt(8, "New Last Value")

	list = "[" + strings.Join(ll.Get(), ", ") + "]"

	fmt.Println("Size", ll.GetSize())
	fmt.Println(list)

	ll.AddAt(1, "1")

	list = "[" + strings.Join(ll.Get(), ", ") + "]"

	fmt.Println("Size", ll.GetSize())
	fmt.Println(list)

	ll.AddAt(2, "2")

	list = "[" + strings.Join(ll.Get(), ", ") + "]"

	fmt.Println("Size", ll.GetSize())
	fmt.Println(list)

	ll.AddAt(3, "3")

	list = "[" + strings.Join(ll.Get(), ", ") + "]"

	fmt.Println("Size", ll.GetSize())
	fmt.Println(list)
}
