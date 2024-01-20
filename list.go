package tapas

// List represents a doubly linked list.
type List[T any] struct {
	head *Node[T]
	len  uint
}

// Node class.
type Node[T any] struct {
	next, prev *Node[T]
	Value      T
}

// Pending possinble operations.
// - Insert: Insert an element at a specific position in the list.
// - InsertBefore inserts a new element e with value v before the mark element and returns e.
// - InsertAfter inserts a new element e with value v after the mark element and returns e.
// - Remove removes e from l if e is an element of list l.
// - MoveToFront moves element e to the front of list l.
// - MoveToBack moves element e to the back of list l.
// - MoveBefore moves element e to its new position before mark.
// - MoveAfter moves element e to its new position after mark.
// - PushBackList inserts a copy of another list at the back of list l.
// - PushFrontList inserts a copy of another list at the front of list l.
// - Pop: Extract and return the element at a given position in the list.
// - Index: Return the position of the first occurrence of an element in the list.
// - Count: Count the number of times an element appears in the list.
// - Sort: Sort the elements of the list.
// - Reverse: Reverse the order of elements in the list.
// - Extend (or Append): Add the elements of another list to the current list.

// Len returns the number of elements of the list.
func (l *List[T]) Len() uint {
	return l.len
}

// Front returns the first element of the list or nil.
func (l *List[T]) Front() *Node[T] {
	if l.len == 0 {
		return nil
	}
	return l.head
}

// Back returns the last element of the list or nil.
func (l *List[T]) Back() *Node[T] {
	if l.len == 0 {
		return nil
	}
	return l.head.prev
}

// PushFront inserts a new value at the front of list.
func (l *List[T]) PushFront(v T) {
	n := &Node[T]{Value: v}
	n.next = l.head
	n.prev = n

	if l.head != nil {
		n.prev = l.head.prev
	}

	l.head = n
	l.len++
}

// PushBack (aka Append) inserts a new value at the back of list
func (l *List[T]) PushBack(v T) {
	if l.head == nil {
		l.PushFront(v)
	} else {
		n := &Node[T]{Value: v}
		n.prev = l.Back()
		l.Back().next = n
		l.head.prev = n
		l.len++
	}
}

// GetAll return an array of T with all values of the list.
func (l *List[T]) GetAll() []T {
	var nodes []T
	for n := l.head; n != nil; n = n.next {
		nodes = append(nodes, n.Value)
	}
	return nodes
}

// Clear the list
func (l *List[T]) Clear() {
	l.head = nil
	l.len = 0
}
