package doubly_linked

import "fmt"

//双向链表
type Node struct {
	pre  *Node
	no   int
	name string
	next *Node
}

type Linked struct {
	First *Node
}

func NewLinked() *Linked {
	return &Linked{
		First: &Node{},
	}
}

func (l *Linked) Add(no int, name string) {
	node := l.First
	newNode := &Node{no: no, name: name}

	for {
		next := node.next
		if next == nil {
			next = newNode
			break
		}
		node = next
	}
}

func (l *Linked) Del(no int) {
	node := l.First
	for {
		next := node.next
		if next == nil {
			break
		}
		if next.no == no {
			if next.next == nil {
				node.next = nil
			} else {
				node.next = next.next
			}
		}
		node = next
	}
}

func (l *Linked) AddBySort(no int, name string) {
	node := l.First
	newNode := &Node{no: no, name: name}
	if node.next == nil {
		node.next = newNode
		return
	}
	for {
		next := node.next
		if next == nil {
			node.next = newNode
			newNode.pre = node
			break
		} else {
			if node.no < no && no < next.no {
				if !(node.no == 0 && node.name == "") {
					newNode.pre = node
				}
				newNode.next = next
				node.next.pre = newNode
				node.next = newNode
				break
			}
		}
		node = next
	}
}

func (l *Linked) List() {
	node := l.First
	if node.next == nil {
		return
	}

	for {
		next := node.next
		if next == nil {
			break
		}
		fmt.Println("no:", next.no, "\tname:", next.name)
		node = next
	}
}
