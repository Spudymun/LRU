package lrucache

const (
	// ErrInvalidIndex occurs when the specified index is out of range
	ErrInvalidIndex = "invalid positional index"
)

// DoublyLinkedList is the DoublyLinkedList data structure
type DoublyLinkedList struct {
	len  int
	root Node
}

// Node represents a node in the linked list
type Node struct {
	prev  *Node
	next  *Node
	Value interface{}
	Key   string
}

// New creates a new LinkedList
func NewList() *DoublyLinkedList {
	l := new(DoublyLinkedList)
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0

	return l
}

// Length returns the length of the Linked List
func (l *DoublyLinkedList) Length() int {
	return l.len
}

// Unshift puts a at the front of the LinkedList
func (l *DoublyLinkedList) Unshift(key string, val interface{}) *Node {
	// Before end of its funcion, we must increase our list by one
	defer func() { l.len++ }()

	newNode := &Node{
		Key:   key,
		Value: val,
		next:  l.root.next,
		prev:  &l.root,
	}

	l.root.next = newNode
	newNode.next.prev = newNode

	return newNode
}

// RemoveTail is going to remove the tail Node from the LinkedList
func (l *DoublyLinkedList) RemoveTail() *Node {
	return l.Remove(l.root.prev)
}

// Remove is going to remove the a Node from the LinkedList
func (l *DoublyLinkedList) Remove(node *Node) *Node {
	defer func() {
		l.len--
	}()

	tail := l.root.prev
	l.isolate(tail)
	return tail
}

func (l *DoublyLinkedList) isolate(node *Node) {
	// node.next - root, root.prev assign to prev current tail
	node.next.prev = node.prev
	// node.prev - new tail, next of new tail assign to next of current tail
	node.prev.next = node.next
	node.next, node.prev = nil, nil
}

// Movefront is going to move the passed in Node to the front of the Linked List
func (l *DoublyLinkedList) MoveFront(node *Node) {
	// Current head node
	currentFront := l.root.next

	// Set used node on top by assign its prev to root and set roots next node to its node
	node.prev = &l.root
	l.root.next = node

	// Set prev its node for current top node on that momment
	currentFront.prev = node
	// Set for our new top node next current top node on that moment
	node.next = currentFront
	// It is all, our node on top now
}

// Head returns the head of the LinkedList
func (l *DoublyLinkedList) Head() *Node {
	return l.root.next
}

// Tail returns the head of the LinkedList
func (l *DoublyLinkedList) Tail() *Node {
	return l.root.prev
}
