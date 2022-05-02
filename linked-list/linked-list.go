package linkedlist

type LinkedList struct {
	head *Node
	size int
}

func (ll *LinkedList) GetSize() int {
	return ll.size
}

func (ll *LinkedList) Init(data string) int {
	return ll.AddFirst(data)
}

func (ll *LinkedList) AddFirst(data string) int {
	node := &Node{
		data: data,
		next: nil,
	}

	node.next = ll.head
	ll.head = node

	ll.size++

	return ll.size
}

func (ll *LinkedList) AddLast(data string) int {
	if ll.head == nil {
		return ll.AddFirst(data)
	}

	node := &Node{
		data: data,
		next: nil,
	}

	currentNode := ll.head

	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = node

	ll.size++

	return ll.size
}

/**
If the position is equals to the length of the ll, should we:
	1) Replace the element: ex [1, 2, 3] => 4 => [1, 2, 4]
	2) Push the element: ex [1, 2, 3] => 4 => [1, 2, 4, 3]

If the position is greater than the length (Out of Bounds), we:
	1) Throw an error and omit the data
	2) Insert the data into the last position available
*/
func (ll *LinkedList) AddAt(position int, data string) int {
	if position == 0 {
		return ll.AddFirst(data)
	}

	if position > ll.size {
		return ll.AddLast(data)
	}

	node := &Node{
		data: data,
		next: nil,
	}

	currentNode := ll.head
	currentIndex := 0

	for currentIndex != position-1 {
		currentIndex++
		currentNode = currentNode.next
	}

	temp := currentNode.next
	currentNode.next = node
	node.next = temp

	ll.size++

	return ll.size
}

func (ll *LinkedList) Get() []string {
	linkedList := []string{}
	currentNode := ll.head

	for currentNode != nil {
		linkedList = append(linkedList, currentNode.data)
		currentNode = currentNode.next
	}

	return linkedList
}
