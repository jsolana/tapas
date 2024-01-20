package tapas

import (
	"fmt"
	"testing"
)

func TestPushFront(t *testing.T) {
	l := new(List[int])
	if l.Len() != 0 {
		t.Errorf(" Wrong Len: expected 0, got: %d", l.Len())
	}

	if l.Front() != nil {
		t.Errorf(" Wrong Front: expected nil, got: %d", l.Front().Value)
	}

	if l.Back() != nil {
		t.Errorf(" Wrong Back: expected nil, got: %d", l.Back().Value)
	}

	l.PushFront(1)
	fmt.Printf("Nodes: %v\n", l.GetAll())

	if l.Len() != 1 {
		t.Errorf(" Wrong Len: expected 1, got: %d", l.Len())
	}
	if l.Front().Value != 1 {
		t.Errorf(" Wrong Front: expected 1, got: %d", l.Front().Value)
	}

	if l.Back().Value != 1 {
		t.Errorf(" Wrong Back: expected 1, got: %d", l.Back().Value)
	}
	l.PushFront(2)
	fmt.Printf("Nodes: %v\n", l.GetAll())

	if l.Len() != 2 {
		t.Errorf(" Wrong Len: expected 2, got: %d", l.Len())
	}
	if l.Front().Value != 2 {
		t.Errorf(" Wrong Front: expected 2, got: %d", l.Front().Value)
	}

	if l.Back().Value != 1 {
		t.Errorf(" Wrong Back: expected 1, got: %d", l.Back().Value)
	}
	l.PushFront(3)
	fmt.Printf("Nodes: %v\n", l.GetAll())

	if l.Len() != 3 {
		t.Errorf(" Wrong Len: expected 3, got: %d", l.Len())
	}

	if l.Front().Value != 3 {
		t.Errorf(" Wrong Front: expected 3, got: %d", l.Front().Value)
	}

	if l.Back().Value != 1 {
		t.Errorf(" Wrong Back: expected 1, got: %d", l.Back().Value)
	}

	l.Clear()

	if l.Len() != 0 {
		t.Errorf(" Wrong Len: expected 0, got: %d", l.Len())
	}

	if l.Front() != nil {
		t.Errorf(" Wrong Front: expected nil, got: %d", l.Front().Value)
	}

	if l.Back() != nil {
		t.Errorf(" Wrong Back: expected nil, got: %d", l.Back().Value)
	}

	l.PushFront(1)
	fmt.Printf("Nodes: %v\n", l.GetAll())

	if l.Len() != 1 {
		t.Errorf(" Wrong Len: expected 1, got: %d", l.Len())
	}
	if l.Front().Value != 1 {
		t.Errorf(" Wrong Front: expected 1, got: %d", l.Front().Value)
	}

	if l.Back().Value != 1 {
		t.Errorf(" Wrong Back: expected 1, got: %d", l.Back().Value)
	}
	l.PushFront(2)
	fmt.Printf("Nodes: %v\n", l.GetAll())

	if l.Len() != 2 {
		t.Errorf(" Wrong Len: expected 2, got: %d", l.Len())
	}
	if l.Front().Value != 2 {
		t.Errorf(" Wrong Front: expected 2, got: %d", l.Front().Value)
	}

	if l.Back().Value != 1 {
		t.Errorf(" Wrong Back: expected 1, got: %d", l.Back().Value)
	}
	l.PushFront(3)
	fmt.Printf("Nodes: %v\n", l.GetAll())

	if l.Len() != 3 {
		t.Errorf(" Wrong Len: expected 3, got: %d", l.Len())
	}

	if l.Front().Value != 3 {
		t.Errorf(" Wrong Front: expected 3, got: %d", l.Front().Value)
	}

	if l.Back().Value != 1 {
		t.Errorf(" Wrong Back: expected 1, got: %d", l.Back().Value)
	}

}

func TestPushBack(t *testing.T) {
	l := new(List[int])
	if l.Len() != 0 {
		t.Errorf(" Wrong Len: expected 0, got: %d", l.Len())
	}

	if l.Front() != nil {
		t.Errorf(" Wrong Front: expected nil, got: %d", l.Front().Value)
	}

	if l.Back() != nil {
		t.Errorf(" Wrong Back: expected nil, got: %d", l.Back().Value)
	}

	l.PushBack(1)
	fmt.Printf("Nodes: %v\n", l.GetAll())

	if l.Len() != 1 {
		t.Errorf(" Wrong Len: expected 1, got: %d", l.Len())
	}
	if l.Front().Value != 1 {
		t.Errorf(" Wrong Front: expected 1, got: %d", l.Front().Value)
	}

	if l.Back().Value != 1 {
		t.Errorf(" Wrong Back: expected 1, got: %d", l.Back().Value)
	}
	l.PushBack(2)
	fmt.Printf("Nodes: %v\n", l.GetAll())

	if l.Len() != 2 {
		t.Errorf(" Wrong Len: expected 2, got: %d", l.Len())
	}
	if l.Front().Value != 1 {
		t.Errorf(" Wrong Front: expected 1, got: %d", l.Front().Value)
	}

	if l.Back().Value != 2 {
		t.Errorf(" Wrong Back: expected 2, got: %d", l.Back().Value)
	}
	l.PushBack(3)
	fmt.Printf("Nodes: %v\n", l.GetAll())

	if l.Len() != 3 {
		t.Errorf(" Wrong Len: expected 3, got: %d", l.Len())
	}

	if l.Front().Value != 1 {
		t.Errorf(" Wrong Front: expected 1, got: %d", l.Front().Value)
	}

	if l.Back().Value != 3 {
		t.Errorf(" Wrong Back: expected 3, got: %d", l.Back().Value)
	}

	l.Clear()

	if l.Len() != 0 {
		t.Errorf(" Wrong Len: expected 0, got: %d", l.Len())
	}

	if l.Front() != nil {
		t.Errorf(" Wrong Front: expected nil, got: %d", l.Front().Value)
	}

	if l.Back() != nil {
		t.Errorf(" Wrong Back: expected nil, got: %d", l.Back().Value)
	}

	l.PushBack(1)
	fmt.Printf("Nodes: %v\n", l.GetAll())

	if l.Len() != 1 {
		t.Errorf(" Wrong Len: expected 1, got: %d", l.Len())
	}
	if l.Front().Value != 1 {
		t.Errorf(" Wrong Front: expected 1, got: %d", l.Front().Value)
	}

	if l.Back().Value != 1 {
		t.Errorf(" Wrong Back: expected 1, got: %d", l.Back().Value)
	}
	l.PushBack(2)
	fmt.Printf("Nodes: %v\n", l.GetAll())

	if l.Len() != 2 {
		t.Errorf(" Wrong Len: expected 2, got: %d", l.Len())
	}
	if l.Front().Value != 1 {
		t.Errorf(" Wrong Front: expected 1, got: %d", l.Front().Value)
	}

	if l.Back().Value != 2 {
		t.Errorf(" Wrong Back: expected 2, got: %d", l.Back().Value)
	}
	l.PushBack(3)
	fmt.Printf("Nodes: %v\n", l.GetAll())

	if l.Len() != 3 {
		t.Errorf(" Wrong Len: expected 3, got: %d", l.Len())
	}

	if l.Front().Value != 1 {
		t.Errorf(" Wrong Front: expected 1, got: %d", l.Front().Value)
	}

	if l.Back().Value != 3 {
		t.Errorf(" Wrong Back: expected 3, got: %d", l.Back().Value)
	}

}
