package main

import (
	"fmt"
	"container/list"
)

func main() {
	fmt.Println("Go Linked List:")
	myList := list.New()
	myList.PushBack(1)
	myList.PushFront(2)
	element := myList.PushFront(3)
	myList.InsertBefore(4, element)
	fmt.Printf("Length: %v\n", myList.Len())
	fmt.Println("List: ")
	printList(myList)

	for element := myList.Front(); element != nil; element = element.Next() {
		if element.Value == 4 {
			myList.Remove(element)
		}
	}

	fmt.Println("List element with value 4 removed")
	printList(myList)
}

func printList(myList *list.List) {
	for element := myList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}
