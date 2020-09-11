package singly_linked

//单向链表

type HeroNode struct {
	No   int
	Name string
	Next *HeroNode
}

type HeroLinked struct {
	First *HeroNode
}

func New() *HeroLinked {
	return &HeroLinked{
		First: &HeroNode{},
	}
}

//有序插入
func (h *HeroLinked) Add(no int, name string) {
	node := &HeroNode{
		No:   no,
		Name: name,
	}

	first := h.First
	for {
		if first.Next == nil {
			break
		}

		nextFirst := first.Next
		if no == nextFirst.No {
			node.Next = nextFirst.Next
			first.Next = node
			return
		}
		if no > first.No && no < nextFirst.No {
			first.Next = node
			node.Next = nextFirst
			return
		}

		first = first.Next
	}
	first.Next = node
}

func (h *HeroLinked) GetNameByNo(no int) string {
	node := h.First
	for {
		if node.Next == nil {
			break
		}
		node = node.Next
		if node.No == no {
			return node.Name
		}

	}
	return ""
}

func (h *HeroLinked) RmByName(name string) bool {
	flag := false
	node := h.First
	for {
		if node.Next == nil {
			break
		}
		nextFirst := node.Next
		if nextFirst.Name == name {
			node.Next = nextFirst.Next
			flag = true
			break
		}
		node = node.Next
	}
	return flag
}

func (h *HeroLinked) List() {
	node := h.First
	if node.Next == nil {
		return
	}
	node = node.Next
	for {
		println("No:", node.No, "Name:", node.Name)
		if node.Next == nil {
			break
		}
		node = node.Next
	}
}

func (h *HeroLinked) Size() int {
	num := 0
	node := h.First

	for {
		if node.Next == nil {
			break
		}
		node = node.Next
		num++
	}

	return num
}

//获取链表倒数第N个元素
func (h *HeroLinked) GetLast(n int) *HeroNode {
	target := h.Size() - n //目标数字
	if target < 0 {
		return nil
	}
	num := 0
	node := h.First
	for {
		if node.Next == nil {
			break
		}
		node = node.Next
		if num == target {
			return node
		}
		num++
	}
	return nil
}

//反转链表
func (h *HeroLinked) Reverse() {
	new := &HeroLinked{
		First: &HeroNode{},
	}
	newNode := new.First
	oldNode := h.First
	if oldNode.Next == nil {
		return
	}
	for {
		oldNext := oldNode.Next
		if oldNext == nil {
			break
		}
		node := &HeroNode{
			Name: oldNext.Name,
			No:   oldNext.No,
		}
		if newNode.Next == nil {
			newNode.Next = node
		} else {
			node.Next = newNode.Next
			newNode.Next = node
		}
		oldNode = oldNext
	}
	h.First = new.First
}

//逆序打印链表
func (h *HeroLinked) ReversePrint() {
	//判断是否是空链表
	if h.First.Next == nil {
		return
	}
	stack := NewStack(0)

	node := h.First
	for {
		next := node.Next
		if next == nil {
			break
		}
		stack.Push(next)
		node = next
	}

	for {
		nodeData := stack.Pop()
		if nodeData == nil {
			break
		}
		node := nodeData.(*HeroNode)
		println("No:", node.No, "Name:", node.Name)
	}

}

//合并两个链表并且有序
func (h *HeroLinked) PutTo(hl *HeroLinked) {
	node := hl.First
	if node.Next == nil {
		return
	}

	for {
		next := node.Next
		if next == nil {
			break
		}
		h.Add(next.No, next.Name)
		node = next
	}
}
