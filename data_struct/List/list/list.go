package list

type Node struct {
	Data interface{}
	Next *Node
}

type List struct {
	size uint32
	head *Node
	tail *Node
}

// 构造方法，将所有变量置空
func (list *List) init() {
	(*list).size = 0
	(*list).head = nil
	(*list).tail = nil
}

func (list *List) Append(node *Node) bool {
	// 若 node 为空，直接返回
	if nil == node {
		return false
	}

	// 若 List 没有节点--头插
	// 若 List 存在节点--尾插
	if 0 == (*list).size {
		(*list).head = node
	} else {
		oldTail := (*list).tail
		(*oldTail).Next = node
	}

	(*list).tail = node
	(*list).size++

	return true
}

func (list *List) Insert(index uint32, node *Node) bool {
	if nil == node || index >= (*list).size || 0 == (*list).size {
		return false
	}

	// 如果 index 为0--头插
	if 0 == index {
		(*node).Next = (*list).head
		(*list).head = node
	} else {
		preNode := (*list).head
		for i := 1; uint32(i) < index; i++ {
			preNode = (*preNode).Next
		}

		(*node).Next = (*preNode).Next
		(*preNode).Next = node
	}

	(*list).size++
	return true
}

func (list *List) Remove(index uint32) bool {
	if index >= (*list).size {
		return false
	}

	if 0 == index { // 头删
		(*list).head = (*list).head.Next
		if 1 == (*list).size {
			(*list).tail = nil
		}
	} else {
		preNode := (*list).head
		for i := 1; uint32(i) < index; i++ {
			preNode = (*preNode).Next
		}
		(*preNode).Next = (*preNode).Next.Next

		if index == (*list).size- 2 {
			(*list).tail = preNode
		}
	}

	(*list).size--

	return true
}

func (list *List) IndexOf(index uint32) *Node {
	if index >= (*list).size {
		return nil
	}

	node := (*list).head
	for i := 0; uint32(i) < index; i++ {
		node = (*node).Next
	}

	return node
}

func (list *List) Size() uint32 {
	return (*list).size
}

func (list *List) Head() *Node {
	return (*list).head
}

func (list *List) Tail() *Node {
	return (*list).tail
}