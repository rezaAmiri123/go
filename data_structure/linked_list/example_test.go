package list_test

import (
	"container/list"
	"fmt"
	"testing"
)

func TestLinked(t *testing.T) {
	// Create a new list and put some numbers in it.
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	_ = l.PushBack(234)
	_ = l.PushBack(65)
	_ = l.PushFront(54)
	_ = l.PushFront(43)
	_ = l.PushFront(65)
	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}

//  $ go test -v data_structure/linked_list/example_test.go
// === RUN   TestLinked
// 65
// 43
// 54
// 1
// 2
// 3
// 4
// 234
// 65
// --- PASS: TestLinked (0.00s)
// PASS
// ok      command-line-arguments  0.001s
