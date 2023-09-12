package linkedlist

type SinglyLinkedList struct {
	head *SinglyNode
	len  int
}

type SinglyNode struct {
	value int
	next  *SinglyNode
}

func (l *SinglyLinkedList) Insert(val int) {
	n := new(SinglyNode)
	n.value = val
	if l.len == 0 {
		l.head = n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = n
			l.len++
			return
		}
		ptr = ptr.next
	}
}

func (l *SinglyLinkedList) GetAt(pos int) *SinglyNode {
	ptr := l.head
	if pos < 0 {
		return ptr
	}
	if pos > (l.len - 1) {
		return nil
	}
	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	return ptr
}

func (l *SinglyLinkedList) InsertAt(pos int, value int) {
	newNode := SinglyNode{}
	newNode.value = value
	if pos < 0 {
		return
	}
	if pos == 0 {
		l.head = &newNode
		l.len++
		return
	}
	if pos > l.len {
		return
	}
	n := l.GetAt(pos)
	newNode.next = n
	prevNode := l.GetAt(pos - 1)
	prevNode.next = &newNode
	l.len++
}
