package queue

import "fmt"

func ExampleQueue_Pop() {
	q:=Queue{1}
	q.Push(1)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	//Output:
	//1
	//false
}