package cycle

import "fmt"

//双向环形链表
type Node struct {
	pre  *Node
	no   int
	name string
	next *Node
}

func (n *Node) SetPre(node *Node) {
	node.pre = n.pre
	n.pre.next = node
	node.next = n
}

func (n *Node) SetNext(node *Node) {
	node.next = n.next
	n.next = node
	node.pre = n
}

type Linked struct {
	First *Node
}

func New() *Linked {
	return &Linked{First: &Node{}}
}

func (l *Linked) Add(no int, name string) {
	node := l.First
	addNode := &Node{no: no, name: name}
	if node.next == nil {
		addNode.next = l.First
		addNode.pre = l.First
		node.next = addNode
		return
	}
	for {
		next := node.next
		if next == l.First {
			break
		}

		if next.pre.no < no && no < next.no {
			next.SetPre(addNode)
			break
		}

		if no > next.no && next.next == l.First {
			next.SetNext(addNode)
			break
		}

		node = next
	}
}

func (l *Linked) Del(no int) {

}

func (l *Linked) List() {
	node := l.First
	for {
		next := node.next
		if next == l.First {
			break
		}
		fmt.Println("no:", next.no, "\tname:", next.name)
		node = next

	}
}
