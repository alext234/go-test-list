package main

import (
	"fmt"
	"container/list"
)


// test some functionality of go double linked-list
func main() {

	// create a list
	l := list.New()
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	
	// add to front
	l.PushFront(1)
	
	// print front
	fmt.Println("front = ", l.Front().Value.(int))
	
	// add to back
	l.PushBack(5)
	
	// print back
	fmt.Println("back = ", l.Back().Value.(int))

	
	// remove front
	l.Remove(l.Front())
	
	// print front
	fmt.Println("front = ", l.Front().Value.(int))

	
	// remove back
	l.Remove(l.Back())
	
	
	// print bottom
	fmt.Println("front = ", l.Back().Value.(int))
	
	// print size 
	fmt.Println("len = ", l.Len())
}

