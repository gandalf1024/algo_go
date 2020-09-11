package singly_linked

//栈
type stack struct {
	arr []interface{}
}

func NewStack(len int) *stack {
	return &stack{
		arr: make([]interface{}, len),
	}
}

func (s *stack) Push(data interface{}) {
	s.arr = append(s.arr, data)
}

func (s *stack) Pop() (data interface{}) {
	l := len(s.arr)
	if l == 0 {
		return
	}
	data = s.arr[l-1]
	s.arr = s.arr[:l-1]
	return
}
