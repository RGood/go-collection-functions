package set

type Set[K comparable] struct {
	data map[K]struct{}
}

func NewSet[K comparable]() *Set[K] {
	return &Set[K]{data: map[K]struct{}{}}
}

func (s *Set[K]) Add(value K) {
	s.data[value] = struct{}{}
}

func (s *Set[K]) Remove(value K) {
	delete(s.data, value)
}

func (s *Set[K]) Contains(value K) bool {
	_, ok := s.data[value]

	return ok
}

func (s *Set[K]) Copy() *Set[K] {
	copiedSet := NewSet[K]()
	for e, _ := range s.data {
		copiedSet.Add(e)
	}

	return copiedSet
}

type LinkedList[K comparable] struct {
	head *Node[K]
	tail *Node[K]
	size int
}

func NewLinkedList[K comparable]() *LinkedList[K] {
	return &LinkedList[K]{size: 0}
}

func (ll *LinkedList[K]) Size() int {
	return ll.size
}

func (ll *LinkedList[K]) Head() *Node[K] {
	return ll.head
}

func (ll *LinkedList[K]) Tail() *Node[K] {
	return ll.tail
}

func (ll *LinkedList[K]) Append(value K) *Node[K] {
	node := &Node[K]{Value: value}

	if ll.size == 0 {
		ll.head = node
		ll.tail = node
	} else {
		ll.tail.Next = node
		node.Prev = ll.tail
		ll.tail = node
	}

	ll.size += 1

	return node
}

func (ll *LinkedList[K]) Remove(node *Node[K]) {
	prev := node.Prev
	next := node.Next

	if prev != nil {
		prev.Next = next
	}

	if next != nil {
		next.Prev = prev
	}
}

type Node[K comparable] struct {
	Value K
	Next  *Node[K]
	Prev  *Node[K]
}

type OrderedSet[K comparable] struct {
	data     map[K]*Node[K]
	dataList *LinkedList[K]
}

func NewOrderedSet[K comparable]() *OrderedSet[K] {
	return &OrderedSet[K]{
		data:     map[K]*Node[K]{},
		dataList: NewLinkedList[K](),
	}
}

func (os *OrderedSet[K]) Add(value K) {
	node := os.dataList.Append(value)
	os.data[value] = node
}

func (os *OrderedSet[K]) Remove(value K) {
	node, ok := os.data[value]
	if ok {
		os.dataList.Remove(node)
	}

	delete(os.data, value)
}

func (os *OrderedSet[K]) Contains(value K) bool {
	_, ok := os.data[value]
	return ok
}

func (os *OrderedSet[K]) ForEach(callback func(int, K)) {
	index := 0
	cur := os.dataList.head
	for cur != nil {
		callback(index, cur.Value)
		index++
		cur = cur.Next
	}
}
